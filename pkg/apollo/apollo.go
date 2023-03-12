package apollo

import (
	"fmt"
	"github.com/go-kratos/kratos/contrib/config/apollo/v2"
	"github.com/go-kratos/kratos/v2/config"
	"helloworld/internal/conf"
)

func GetConfig(ApolloConfig *conf.Apollo) *conf.Bootstrap {
	c := config.New(
		config.WithSource(
			apollo.NewSource(
				apollo.WithAppID(ApolloConfig.GetAppId()),
				apollo.WithCluster(ApolloConfig.GetCluster()),
				apollo.WithEndpoint(ApolloConfig.GetEndpoint()),
				apollo.WithNamespace(ApolloConfig.GetNamespace()),
				apollo.WithEnableBackup(),
				apollo.WithSecret(ApolloConfig.GetSecret()),
				apollo.WithBackupPath(ApolloConfig.GetBackupPath()),
			),
		),
	)
	var bc *conf.Bootstrap
	if err := c.Load(); err != nil {
		panic(err)
	}
	scan(c, bc)
	return bc
}

func scan(c config.Config, bc *conf.Bootstrap) {
	err := c.Scan(bc)
	fmt.Printf("=========== scan result =============\n")
	fmt.Printf("err: %v\n", err)
	fmt.Printf("cfg: %+v\n\n", bc)
}
