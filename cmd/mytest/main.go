package main

import (
	"fmt"
	"github.com/go-kratos/kratos/contrib/config/apollo/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
)

type bootstrap struct {
	Application struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Author  string `json:"author"`
	} `json:"application"`
	Redis struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Author  string `json:"author"`
	} `json:"redis"`
}

func main() {

	c := config.New(
		config.WithSource(
			apollo.NewSource(
				apollo.WithAppID("pay123456"),
				apollo.WithCluster("default"),
				apollo.WithEndpoint("http://81.68.181.139:8080"),
				apollo.WithNamespace("application"),
				apollo.WithEnableBackup(),
				apollo.WithSecret("45450aff807242e4b04e2618a5d5e984"),
			),
		),
	)
	var bc bootstrap
	if err := c.Load(); err != nil {
		//panic(err)
	}

	scan(c, &bc)

	value(c, "application")
	//value(c, "application.name")
	//value(c, "event.array")
	//value(c, "demo.deep")

	watch(c, "application")
	<-make(chan struct{})

}
func scan(c config.Config, bc *bootstrap) {
	err := c.Scan(bc)
	fmt.Printf("=========== scan result =============\n")
	fmt.Printf("err: %v\n", err)
	fmt.Printf("cfg: %+v\n\n", bc)
}

func value(c config.Config, key string) {
	fmt.Printf("=========== value result =============\n")
	v := c.Value(key).Load()
	fmt.Printf("key=%s, load: %+v\n\n", key, v)
}

func watch(c config.Config, key string) {
	if err := c.Watch(key, func(key string, value config.Value) {
		log.Info("config(key=%s) changed: %s\n", key, value.Load())
	}); err != nil {
		//panic(err)
	}
}
