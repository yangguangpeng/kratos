package apollo

import (
	"fmt"
	"github.com/go-kratos/kratos/contrib/config/apollo/v2"
	"github.com/go-kratos/kratos/v2/config"
	"helloworld/internal/conf"
)

func GetConfig(baseConfig conf.Base) conf.Bootstrap {
	c := config.New(
		config.WithSource(
			apollo.NewSource(
				apollo.WithAppID(baseConfig.GetApollo().GetAppId()),
				apollo.WithCluster(baseConfig.GetApollo().GetCluster()),
				apollo.WithEndpoint(baseConfig.GetApollo().GetEndpoint()),
				apollo.WithNamespace(baseConfig.GetApollo().GetNamespace()),
				apollo.WithEnableBackup(),
				apollo.WithSecret(baseConfig.GetApollo().GetSecret()),
				apollo.WithBackupPath(baseConfig.GetApollo().GetBackupPath()),
			),
		),
	)
	var bc conf.Bootstrap
	if err := c.Load(); err != nil {
		panic(err)
	}
	scan(c, &bc)
	return bc
}

func scan(c config.Config, bc *conf.Bootstrap) {
	err := c.Scan(bc)
	fmt.Printf("=========== scan result =============\n")
	fmt.Printf("err: %v\n", err)
	fmt.Printf("cfg: %+v\n\n", bc)
}
