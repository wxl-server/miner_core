package nacos

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"miner_core/sal/config"
)

func Register2Nacos(ctx context.Context) naming_client.INamingClient {
	nacosConfig := config.Config.Nacos

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				NamespaceId: nacosConfig.Namespace,
				Username:    nacosConfig.Username,
				Password:    nacosConfig.Password,
			},
			ServerConfigs: []constant.ServerConfig{
				*constant.NewServerConfig(nacosConfig.Host, nacosConfig.Port, constant.WithGrpcPort(nacosConfig.GrpcPort)),
			},
		},
	)
	if err != nil {
		logger.CtxErrorf(ctx, "[Init] nacos init failed, err = %v", err)
		panic(err)
	}
	return cli
}
