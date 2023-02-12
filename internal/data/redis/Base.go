package redis

import (
	"github.com/gomodule/redigo/redis"
	"helloworld/internal/conf"
)


type Base struct {
	BaseRedis
}

func NewBase (redisConf *conf.Data_Redis) (*Base, error){
	pool, err := getPool(redisConf)
	return &Base{
		BaseRedis{RedisPool: pool},
	},err
}



type BaseRedis struct {
	TTL           int
	KeyNamePrefix string
	KeyName       string

	RedisPool *redis.Pool
}

// SetTTL 设置TTL
func (t *BaseRedis) SetTTL(seconds int) {
	t.TTL = seconds
}

func (t *BaseRedis) GetKeyName() string {
	//必须在调用.SetKeyName()后，调用才有值
	return t.KeyName
}

func (t *BaseRedis) Expire() {
	if t.TTL <= 0 {
		return
	}
	r := t.RedisPool.Get()
	defer r.Close()
	r.Do("EXPIRE", t.KeyName, t.TTL)
}

func (t *BaseRedis) IncrBy(incrAmount interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("IncrBy", t.KeyName, incrAmount)
	return
}

func (t *BaseRedis) Get() (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("GET", t.KeyName)
	return
}

func (t *BaseRedis) Set(v string) {

	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	r.Do("SET", t.KeyName, v)

}

func (t *BaseRedis) HGet(field string) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("HGET", t.KeyName, field)
	return
}

func (t *BaseRedis) HGetAll() (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("HGETAll", t.KeyName)
	return
}

func (t *BaseRedis) HSet(field string, v interface{}) {

	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	r.Do("HSet", t.KeyName, field, v)
}

func (t *BaseRedis) HIncrBy(fieldName string, v interface{}) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	r.Do("HIncrBy", t.KeyName, fieldName, v)
}

func (t *BaseRedis) Delete() {
	r := t.RedisPool.Get()
	defer r.Close()

	r.Do("del", t.KeyName)
}

func (t *BaseRedis) LPush(v string) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("lPush", t.KeyName, v)
	return
}

func (t *BaseRedis) RPush(v string) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("rPush", t.KeyName, v)
	return
}

func (t *BaseRedis) LPop(v string) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("lPop", t.KeyName, v)
	return
}

func (t *BaseRedis) RPop(v string) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("rPop", t.KeyName, v)
	return
}

func (t *BaseRedis) LSize() (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("lSize", t.KeyName)
	return
}

func (t *BaseRedis) LRange(start interface{}, end interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("lRange", t.KeyName, start, end)
	return
}

func (t *BaseRedis) LPushX(v string) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("lPushx", t.KeyName, v)
	return
}

func (t *BaseRedis) RPushX(v string) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("rPushx", t.KeyName, v)
	return
}

func (t *BaseRedis) LGet(index interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("lGet", t.KeyName, index)
	return
}

func (t *BaseRedis) LSet(index interface{}, v string) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("lSet", t.KeyName, index, v)
	return
}

func (t *BaseRedis) LRemove(v string, count interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("lRem", t.KeyName, v, count)
	return
}

func (t *BaseRedis) SAdd(v string) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("sAdd", t.KeyName, v)
	return
}

func (t *BaseRedis) SRemove(v string) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("sRem", t.KeyName, v)
	return
}

func (t *BaseRedis) SIsMember(v string) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("sIsMember", t.KeyName, v)
	return
}

func (t *BaseRedis) SSize() (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("sSize", t.KeyName)
	return
}

func (t *BaseRedis) SMembers() (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("sMembers", t.KeyName)
	return
}

func (t *BaseRedis) SPop() (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("sPop", t.KeyName)
	return
}

func (t *BaseRedis) SRandMember() (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("sRandMember", t.KeyName)
	return
}

func (t *BaseRedis) ZAdd(score interface{}, v string) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("zAdd", t.KeyName, score, v)
	return
}

func (t *BaseRedis) ZRem(v string) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("zRem", t.KeyName, v)
	return
}

func (t *BaseRedis) ZRange(start interface{}, end interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("zRange", t.KeyName, start, end)
	return
}

func (t *BaseRedis) ZRevRange(start interface{}, end interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("zRevRange", t.KeyName, start, end)
	return
}

func (t *BaseRedis) ZRangeByScore(start interface{}, end interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("zRangeByScore", t.KeyName, start, end)
	return
}

func (t *BaseRedis) ZRevRangeByScore(start interface{}, end interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("zRevRangeByScore", t.KeyName, start, end)
	return
}

func (t *BaseRedis) ZRank(v interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("zRank", t.KeyName, v)
	return
}

func (t *BaseRedis) ZRevRank(v interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("zRevRank", t.KeyName, v)
	return
}

func (t *BaseRedis) ZSize() (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("zSize", t.KeyName)
	return
}

func (t *BaseRedis) ZCount(start interface{}, end interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("zCount", t.KeyName, start, end)
	return
}

func (t *BaseRedis) ZRemRangeByScore(start interface{}, end interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("zRemRangeByScore", t.KeyName, start, end)
	return
}

func (t *BaseRedis) ZScore(v interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()

	reply, err = r.Do("zScore", t.KeyName, v)
	return
}

func (t *BaseRedis) ZIncrBy(increment interface{}, v interface{}) (reply interface{}, err error) {
	r := t.RedisPool.Get()
	defer r.Close()
	defer t.Expire()

	reply, err = r.Do("zIncrBy", t.KeyName, increment, v)
	return
}
