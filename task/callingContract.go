package task

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	bingdings "match-be/contract/bindings"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func StartPollingContractMethod(client *ethclient.Client, contractAddress common.Address, interval time.Duration) {
	instance, err := bingdings.NewBingdings(contractAddress, client)
	if err != nil {
		log.Println("实例化合约失败:", err)
		return
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// 创建固定不变的调用选项（提升到循环外部）
	baseOpts := &bind.CallOpts{
		Pending: false,
		From:    contractAddress,      // 如果fromAddress固定不变
		Context: context.Background(), // 基础上下文
	}

	for {
		select {
		case <-ticker.C:
			go func(opts *bind.CallOpts) { // 通过参数传递避免闭包陷阱
				// 复制基础配置以保持可扩展性
				localOpts := *opts

				// 可以在此处添加动态配置（如超时控制）
				// localOpts.Context = context.WithTimeout(opts.Context, 5*time.Second)

				games, err := instance.GetAllGames(&localOpts)
				if err != nil {
					log.Println("调用失败:", err)
					return
				}
				log.Printf("找到 %d 个游戏记录\n", len(games))
				for i, game := range games {
					log.Printf("游戏 %d:", i+1)
					log.Printf("  ID: %s", game.Id.String())
					log.Printf("  名称: %s", game.Name)
					log.Printf("  玩家: %v", game.Players) // []common.Address
					log.Printf("  开始时间: %s", game.StartTime.String())
					log.Printf("  结束时间: %s", game.EndTime.String())
					log.Printf("  是否已结算: %s", game.IsSettled.String())
					log.Printf("  玩家0余额: %s", game.Player0Balance.String())
					log.Printf("  玩家1余额: %s", game.Player1Balance.String())
				}

				log.Println("定时返回:", games)
			}(baseOpts)
		}
	}
}

// Settlement 定时查询游戏数量，并对每个游戏调用结算（Settlement）方法
func Settlement(client *ethclient.Client, contractAddress common.Address, interval time.Duration) {
	// 实例化合约
	instance, err := bingdings.NewBingdings(contractAddress, client)
	if err != nil {
		log.Println("实例化合约失败:", err)
		return
	}

	// 获取私钥（需在环境变量中设置 Private）
	privateKey := os.Getenv("Private")
	if len(privateKey) == 0 {
		log.Println("ERROR: 请设置 .env 文件中的 Private")
		return
	}

	// 为只读调用构造基础 callOpts（这里用的是从私钥派生的地址）
	baseTransactor, err := GetTransactor(client, privateKey)
	if err != nil {
		log.Println("获取基础 transactor 失败:", err)
		return
	}
	baseCallOpts := &bind.CallOpts{
		Pending: false,
		From:    baseTransactor.From,
		Context: context.Background(),
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C

		// 调用只读函数获取当前游戏数量
		count, err := instance.GetGameCount(baseCallOpts)
		if err != nil {
			log.Println("调用 GetGameCount 失败:", err)
			continue
		}
		countInt := count.Int64()
		fmt.Printf("当前游戏数量: %d\n", countInt)

		// 针对每个游戏发起一次结算交易
		for i := int64(1); i <= countInt; i++ {
			// 为每笔交易构造新的 transactor 以确保 nonce 和 gasPrice 是最新的
			transactor, err := GetTransactor(client, privateKey)
			if err != nil {
				log.Printf("获取 transactor 失败，gameID = %d: %s\n", i, err)
				continue
			}

			// 调用 Settlement 方法传入 gameId（这里假设 gameId 为 i）
			tx, err := instance.Settlement(transactor, big.NewInt(i))
			if err != nil {
				log.Printf("结算失败, gameID = %d: %s\n", i, err)
				continue
			}

			log.Printf("轮询调用结算方法成功, gameID = %d, tx Hash: %s\n", i, tx.Hash().Hex())

			// 可选：短暂等待一下再发下一笔交易，避免 nonce 问题或网络延迟
			time.Sleep(2 * time.Second)
		}
	}
}

// GetTransactor 根据私钥构造 *bind.TransactOpts，用于交易签名
func GetTransactor(client *ethclient.Client, privateKeyHex string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // 发送交易时不附带 ETH
	auth.GasLimit = uint64(3000000) // 设置足够的 Gas 限额
	// 为确保替换交易有效，额外在 gasPrice 上加 1 Gwei
	auth.GasPrice = new(big.Int).Add(gasPrice, big.NewInt(1e9))

	return auth, nil
}
