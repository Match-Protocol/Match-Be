package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"gorm.io/gorm/clause"
	"log"
	config2 "match-be/config"
	"match-be/pkg/verify"
	"match-be/repository/cache"
	"match-be/services"

	"match-be/pkg/middleware"
	"match-be/pkg/randomJ"
	"net/http"
	"time"

	rpcS "github.com/gagliardetto/solana-go/rpc"
	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// todo: 现在写的简陋版不优雅,调用起来笨重,后续再改掉整个架构

var (
	DB         *gorm.DB
	Config0    config2.Toml
	BaseClient *ethclient.Client
	Redis0     *redis.Client
	// business description
	mailLogin = "mail-login" // match-login

)

func initDB() {
	// 替换成你的 MySQL 连接信息
	if _, err := toml.DecodeFile("config.toml", &Config0); err != nil {
		panic(err)
	}
	//dsn := "root:password@tcp(127.0.0.1:3306)/user_auth?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config0.Mysql.User, Config0.Mysql.Password, Config0.Mysql.Host, Config0.Mysql.Port, Config0.Mysql.Db)
	fmt.Println(dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	// 自动迁移
	//if err := DB.AutoMigrate(&User{}); err != nil {
	//	panic("数据库迁移失败: " + err.Error())
	//}
	fmt.Println("DB already connected")
}

func main() {
	initDB()
	initRedis()
	//initRpc()
	contractAddress := common.HexToAddress(Config0.Rpc.P)
	fmt.Println(contractAddress)
	time.Sleep(time.Second)
	//go task.StartPollingContractMethod(BaseClient, contractAddress, 15*time.Second)
	//go task.Settlement(BaseClient, contractAddress, 80*time.Second)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(middleware.CorsMiddleware())
	//r.POST("/register", register)
	r.POST("/loginV0", loginV0)
	r.POST("/ses", ses)
	r.POST("/emailLogin", emailLogin)
	r.POST("/walletConnect", walletConnect)
	r.POST("/ata", tokenAmount)

	r.Run(":8099")
}

func initRpc() {
	client, err := ethclient.Dial(Config0.Rpc.Url)
	BaseClient = client
	if err != nil {
		log.Fatal(err)
	}
}

// 注册接口

// 登录接口

//  启动定时任务

// User 模型
type User struct {
	gorm.Model
	UserName string `gorm:"unique" json:"user_name"`
	Address  string `json:"address"`
}

type UserV1s struct {
	gorm.Model
	UserName  string         ` json:"user_name"`
	Mail      string         `gorm:"unique" json:"mail"`
	Address1  string         `json:"address1"`
	Address2  sql.NullString `json:"address2"`
	Address3  sql.NullString `json:"address3"`
	AddrCount sql.NullInt64  `json:"address_count"`
}
type Atas struct {
	gorm.Model
	Addr string `json:"addr"`
	Ata  string `json:"ata"`
}

func loginV0(c *gin.Context) {
	var input struct {
		Address string `json:"address" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "必须提供 address 字段"})
		return
	}
	// 查询数据库是否存在该 address
	var user User
	if err := DB.Where("address = ?", input.Address).First(&user).Error; err == nil {
		// 如果找到，返回用户信息
		c.JSON(http.StatusOK, gin.H{
			"msg":      "登录成功",
			"username": user.UserName,
			"address":  user.Address,
		})
		return
	}
	// 如果没找到，生成随机英文名并插入数据库
	username := randomJ.GenerateRandomEnglishName()
	newUser := User{
		UserName: username,
		Address:  input.Address,
	}
	if err := DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":      "新用户已创建",
		"username": username,
		"address":  input.Address,
	})
}

func initRedis() {
	Redis0 = redis.NewClient(&redis.Options{
		Addr:     Config0.Redis.Host + ":" + Config0.Redis.Port,
		Password: Config0.Redis.Password,
		DB:       Config0.Redis.Db,
	})
	err := Redis0.Ping(context.Background()).Err()
	if err != nil {
		panic("redis connection fail")
	}
	fmt.Println("Redis already connected")
}

// sol version
func ses(c *gin.Context) {
	var smsReq struct {
		Email string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&smsReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "require email address"})
		return
	}
	valid := verify.Email(smsReq.Email)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email address is invalid"})
		return
	}

	codeCache := cache.NewCodeCache(Redis0)
	code := randomJ.Number(6)
	if len(code) == 0 {
		c.Error(errors.New("Internal Server issue"))
		return
	}

	// 业务字段
	err := codeCache.Set(context.Background(), "mp-reqLogin", smsReq.Email, code)
	// err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification code,redis logic error"})
	//	return
	//}
	fmt.Println("err!!!: ", err)
	switch err {
	case nil:

	case cache.ErrCodeSendTooMany:
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": cache.ErrCodeSendTooMany.Error(),
		})
		return

	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to send verification code,redis logic error",
		})
		return
	}

	ok, err := services.Send(Config0.Sms.Key, Config0.Sms.Secret, code, smsReq.Email)
	if ok {
		c.JSON(http.StatusOK, gin.H{"msg": "verification code sent successfully",
			"code": code})
		return
	}

	if err != nil {
		er := errors.New(fmt.Sprintf("failed to send verification code,AWS err: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "during test stage,only support internal address",
			"error": er.Error()})
		return
	}

}

func emailLogin(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required"`
		Code  string `json:"code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "require email address and code"})
		return
	}
	valid := verify.Email(input.Email)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email address is invalid"})
		return
	}
	codeCache := cache.NewCodeCache(Redis0)
	if ok, err := codeCache.Verify(context.Background(), "mp-reqLogin", input.Email, input.Code); !ok {
		if err == cache.ErrCodeVerifyTooMany {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "too many verification attempts,wait a bit"})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid code"})
			return
		}
		//c.JSON()
	}

	user := &UserV1s{
		UserName:  randomJ.GenerateRandomEnglishNameV2(),
		Mail:      input.Email,
		Address1:  "",
		Address2:  sql.NullString{Valid: false},
		Address3:  sql.NullString{Valid: false},
		AddrCount: sql.NullInt64{Valid: false},
	}
	result := DB.Where("mail = ?", input.Email).First(&user)

	//result := DB.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "mail"}}, // 基于 mail 字段冲突判断
	//	DoUpdates: clause.AssignmentColumns([]string{"updated_at"}), // 更新 address1 字段
	//}).Create(&user)
	if result.Error != nil {
		if err := DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":      "login successful",
		"username": user.UserName,
		"address":  user.Address1,
	})
	return
}

/*
1.没有邮箱登录,地址+空邮箱写入
2.已经登录了邮箱,授权连接地址
*/
func walletConnect(c *gin.Context) {
	var input struct {
		Address string `json:"address" binding:"required"`
		Email   string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if input.Email != "" && !verify.Email(input.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid email address"})
		return
	}

	newUser := &UserV1s{
		Mail:     input.Email,
		Address1: input.Address,
	}

	// 如果 mail 是唯一的，就尝试插入或更新
	clause := clause.OnConflict{
		Columns:   []clause.Column{{Name: "mail"}},                // 冲突字段为 mail
		DoUpdates: clause.AssignmentColumns([]string{"address1"}), // 只更新 address1
	}

	dbResult := DB.Clauses(clause).Create(newUser)
	if dbResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库操作失败", "detail": dbResult.Error.Error()})
		return
	}
	// 返回成功信息
	c.JSON(http.StatusOK, gin.H{
		"msg":     "wallet address binding success",
		"email":   newUser.Mail,
		"address": newUser.Address1,
	})
	return
}

func tokenAmount(c *gin.Context) {
	var input struct {
		Address string `json:"address" binding:"required"`
	}
	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	var ata Atas

	res := DB.Where("addr =? ", input.Address).First(&ata)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": res.Error.Error()})
		return

	}
	if len(ata.Addr) > 0 && len(ata.Ata) > 0 {

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "can't find account info"})
		return
	}

	client := rpcS.New(Config0.Rpc.Sol)
	ataAddress := solana.MustPublicKeyFromBase58(ata.Ata)

	// 获取账户信息
	accountInfo, err := client.GetAccountInfoWithOpts(
		context.TODO(),
		ataAddress,
		&rpcS.GetAccountInfoOpts{
			Encoding: "base64", // SPL Token 是二进制格式，需指定编码
		},
	)
	if err != nil {
		log.Printf("failed to get account info: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// 解码为 SPL Token Account 类型
	// 解码为 SPL Token Account 类型
	data := accountInfo.Value.Data.GetBinary()
	var tokenAccount token.Account
	err = bin.NewBinDecoder(data).Decode(&tokenAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "success",
			"amount": tokenAccount.Amount / 1e9,
		})
		// 输出代币余额（注意是 uint64 格式，通常需要结合 decimals 才能换算成人类可读格式）
		fmt.Printf("Token amount: %d\n", tokenAccount.Amount)
		return
	}

}
