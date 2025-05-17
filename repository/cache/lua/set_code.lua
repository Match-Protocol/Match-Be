local key = KEYS[1]   -- 第一個key是幹嘛?  就是一个独立的键名
local cntKey =key.. ":cnt"   -- 自己起的另一个健明remain verify times
--  key 和 ":cnt" 连接
-- 你准备存储的验证码
local  val = ARGV[1]

--  tonumber what   redis.call("command",object)
local ttl = tonumber(redis.call("ttl",key))
if ttl == -1 then--  key存在但是没有过期时间,这是一种可能
redis.call("del", key)
 return -2      --  给后端返回的-2 很严重,居然验证码一直有效
 elseif ttl >=590 then
 --    发送太频繁
    return -1
elseif ttl == -2  or ttl < 600 then    --  -2  表示key不存在 小于600秒内的验证码还有效
--      key不存在 或还在10分钟有效期,可以重新发送
   redis.call("set",key,val)
   redis.call("expire",key,600)
   redis.call("set",cntKey,3) --    只能验证三次
   redis.call("expire",cntKey,600)
   return 0
end