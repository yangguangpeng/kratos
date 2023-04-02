package main

import (
	"fmt"
	"github.com/go-kratos/kratos/contrib/config/apollo/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/conf"
)

//type bootstrap struct {
//	Application struct {
//		Name    string `json:"name"`
//		Version string `json:"version"`
//		Author  string `json:"author"`
//	} `json:"application"`
//	Test struct {
//		Address string `json:"address"`
//		Tls     struct {
//			Enable   bool   `json:"enable"`
//			CertFile string `json:"cert_file"`
//			KeyFile  string `json:"key_file"`
//		} `json:"tls"`
//	} `json:"test"`
//}

func main() {

	c := config.New(
		config.WithSource(
			apollo.NewSource(
				apollo.WithAppID("ms"),
				apollo.WithCluster("dev"),
				apollo.WithEndpoint("http://192.168.1.5:8080"),
				apollo.WithNamespace("application,redis,test"),
				//apollo.WithEnableBackup(),
				apollo.WithSecret("bbfbc03b62534917aa41d20337db2f9b"),
			),
		),
	)
	var bc conf.Bootstrap
	if err := c.Load(); err != nil {
		//panic(err)
	}

	scan(c, &bc)

	value(c, "application")
	value(c, "application.name")
	//value(c, "event.array")
	//value(c, "demo.deep")

	watch(c, "application")
	<-make(chan struct{})

}
func scan(c config.Config, bc *conf.Bootstrap) {
	err := c.Scan(bc)
	//bc.Test.Address = `sdfsdfsdf`
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
