package bigCache

import (
	redis "helloworld/pkg/cache/toolRedis"
	"strconv"
)

type User struct {
	Base
}

func UserInstance() *User {
	instance := &User{
		Base{
			BaseRedis: redis.BaseRedis{
				TTL:           -1,
				KeyNamePrefix: "language",
			},
		},
	}

	instance.Init()

	return instance
}

// Init 真正的keyName的获取会比较复杂
func (t *User) Init() {
	t.Base.Init()
}

func (t *User) UseSlave() *User {
	t.Base.UseSlave()
	return t
}

//SetKeyName 通过本方法，构建真实的 key
func (t *User) SetKeyName(userId int64) *User {
	t.BaseRedis.KeyName = `user_id:` + strconv.Itoa(int(userId))
	return t
}
