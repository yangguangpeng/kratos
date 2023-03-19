package bigCache

import (
	redis "helloworld/pkg/cache/toolRedis"
	"strconv"
	"strings"
)

type Language struct {
	Base
}

func LanguageInstance() *Language {
	instance := &Language{
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
func (t *Language) Init() {
	t.Base.Init()
}

//SetKeyName 通过本方法，构建真实的 key
func (t *Language) SetKeyName(name string, langAppId int, lang string) *Language {
	t.BaseRedis.KeyName = `lang_item_` + strconv.Itoa(langAppId) + `_` + strings.ToLower(lang) + `_` + name
	return t
}
