package config

import (
	"github.com/foodi-org/foodi-user-service/internal/pkg/pkgconsul"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

var (
	servConf = Config{}
)

type (
	// Config go-zero rpc 服务配置结构体
	Config struct {
		zrpc.RpcServerConf

		Mysql struct {
			DataSource string `json:"dataSource"`
		}

		Consul consul.Conf
	}

	// ServYaml etc yaml 文件配置映射结构体
	ServYaml struct {
		UnmarshalKey string     `yaml:"UnmarshalKey"`
		ConsulYaml   ConsulYaml `yaml:"Consul"`
	}

	ConsulYaml struct {
		Key   string   `yaml:"Key"`
		Meta  Meta     `yaml:"Meta"`
		Tag   []string `yaml:"Tag"`
		Host  string   `yaml:"Host"`
		Token string   `yaml:"Token"`
		TTL   int      `yaml:"TTL"`
	}

	Meta struct {
		Protocol string `yaml:"Protocol"`
	}

	// UserConsulConf consul 配置映射结构体
	UserConsulConf struct {
		// ServiceName 服务名
		ServiceName string `json:"serviceName"`

		// ListenOn 声明监听端口 0.0.0.0:xxx
		ListenOn string `json:"listenOn"`

		Mysql struct {
			Datasource string `json:"datasource"`
		} `json:"mysql"`

		Redis struct {

			// redis 服务地址 ip:port, 如果是 redis cluster 则为 ip1:port1,ip2:port2,ip3:port3...(暂不支持redis sentinel)
			Host string `json:"host"`

			// node:单节点 redis;cluster:redis 集群
			Type string `json:"type"`

			Password string `json:"password"`

			// 是否开启tls
			TLS bool `json:"tls"`
		} `json:"redis"`
	}
)

func ServConf() *Config {
	return &servConf
}

// InitServConf
//
//	@dDescription: 初始化 service 配置
//	@param path: 项目路径
//	@param filename: 使用的配置文件名称
//	@return error
func InitServConf(path string, filename string) error {
	var (
		servYaml       ServYaml
		userConsulConf UserConsulConf
	)

	// 解析yaml配置文件
	file, err := os.ReadFile(filepath.Join(path, "etc", filename))
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(file, &servYaml); err != nil {
		return err
	}

	servConf.Consul = consul.Conf{
		Host:  servYaml.ConsulYaml.Host,
		Key:   servYaml.ConsulYaml.Key,
		Token: servYaml.ConsulYaml.Token,
		Tag:   servYaml.ConsulYaml.Tag,
		TTL:   servYaml.ConsulYaml.TTL,
	}

	if err = servConf.Consul.Validate(); err != nil {
		return err
	}

	consulClient := pkgconsul.NewConsulDO()
	consulClient.SetConfig(servConf.Consul)

	// consul connect
	if err = consulClient.Connect(); err != nil {
		panic(err)
	}

	// load service config from consul with key
	if err = consulClient.LoadJsonConfig(servYaml.UnmarshalKey, &userConsulConf); err != nil {
		panic(err)
	}

	// 设置 zrpc service config
	servConf.Name = userConsulConf.ServiceName
	servConf.ListenOn = userConsulConf.ListenOn
	servConf.Mysql.DataSource = userConsulConf.Mysql.Datasource

	servConf.Redis = redis.RedisKeyConf{
		RedisConf: redis.RedisConf{
			Host: userConsulConf.Redis.Host,
			Type: userConsulConf.Redis.Type,
			Pass: userConsulConf.Redis.Password,
			Tls:  userConsulConf.Redis.TLS,
		},
		Key: userConsulConf.ServiceName,
	}
	return nil
}
