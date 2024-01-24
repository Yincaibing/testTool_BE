package config

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"gitlab.iglooinsure.com/axinan/backend/turbo/new-executor/component/promotion"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gitlab.iglooinsure.com/axinan/backend/common/axservice_client"
	"gitlab.iglooinsure.com/axinan/backend/common/config"
	"gitlab.iglooinsure.com/axinan/backend/common/gip_platform_pkg/pkg/grpc_service"
	"gitlab.iglooinsure.com/axinan/backend/common/log"
	"gitlab.iglooinsure.com/axinan/backend/common/services"
	"gitlab.iglooinsure.com/axinan/backend/common/trace"
	"gitlab.iglooinsure.com/axinan/backend/common/transaction_event"
	"gitlab.iglooinsure.com/axinan/backend/turbo/new-executor/component/id"
	"gitlab.iglooinsure.com/axinan/backend/turbo/new-executor/component/pubsub"
	"gitlab.iglooinsure.com/axinan/backend/turbo/new-executor/pkg/events"
	"gorm.io/gorm"
)

type Config struct {
	Host            string `json:"host"`
	config.Config   `json:",inline" mapstructure:",squash"`
	ENV             string                   `json:"env"`
	Redis           *config.RedisConfig      `json:"redis" yaml:"Redis"`
	Publisher       *events.Config           `json:"publisher" yaml:"Publisher"`
	Subscriber      *events.Config           `json:"subscriber"`
	InnerPublisher  *pubsub.PublisherConfig  `json:"InnerPublisher" yaml:"InnerPublisher"`
	InnerSubscriber *pubsub.SubscriberConfig `json:"InnerSubscriber" yaml:"InnerSubscriber"`
	SnowIDNode      *id.Config               `json:"SnowIDNode" yaml:"SnowIDNode"`
	// call back
	PolicyExpire    string            `json:"PolicyExpire" yaml:"PolicyExpire"`
	PolicyProtect   string            `json:"PolicyProtect" yaml:"PolicyProtect"`
	QuotationExpire string            `json:"QuotationExpire" yaml:"QuotationExpire"`
	PromotionConfig *promotion.Config `json:"PromotionConfig" yaml:"PromotionConfig"`
}

func (c *Config) FetchExpireURL(id uint64) string {
	return fmt.Sprintf(c.PolicyExpire, id)
}

func (c *Config) FetchProtectURL(id uint64) string {
	return fmt.Sprintf(c.PolicyProtect, id)
}

func (c *Config) FetchQuotationExpireURL(id uint64) string {
	return fmt.Sprintf(c.QuotationExpire, id)
}

func GetInnerPublisherConfig(c *Config) *pubsub.PublisherConfig {
	return c.InnerPublisher
}

func GetInnerSubscriberConfig(c *Config) *pubsub.SubscriberConfig {
	return c.InnerSubscriber
}

func GetPublisherConfig(c *Config) *events.Config {
	return c.Publisher
}

func GetSnowIDNodeConfig(c *Config) *id.Config {
	return c.SnowIDNode
}

func GetPromotionConfig(c *Config) *promotion.Config {
	if c == nil {
		return nil
	}
	return c.PromotionConfig
}

func LoadConfigWithFile(fileName string) (*Config, error) {
	var cfg = &Config{}
	c := viper.New()
	c.AddConfigPath(fileName)
	c.SetConfigType("yaml")
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	err = c.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	err = c.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, cfg.Init()
}

func (c *Config) Init() error {
	log.Initialize(c.AppConfig.Name, false)

	if c.TraceConfig != nil {
		err := trace.InitializeTracer(trace.ServiceName(c.TraceConfig.ServiceName),
			trace.ReportAddr(c.TraceConfig.ReportAddr))
		if err != nil {
			return err
		}
	}
	if len(c.AxService) > 0 {
		_ = services.Initialize(c.AxService) // Can't be err
		axservice_client.Initialize()
	}

	return nil
}

func InitMysql(c *Config) (*gorm.DB, error) {
	if c.Mysql == nil {
		return nil, nil
	}
	return c.Mysql.Build()
}

func InitRedis(c *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		DB:   c.Redis.Db,
	})
}

func GetKafkaConfig(c *Config) *transaction_event.KafkaConfig {
	return c.SmKafka
}

func InitGrpcService(c *Config) (*grpc_service.GrpcService, error) {
	if c.GRpcService == nil {
		return nil, nil
	}
	grpcService, err := grpc_service.NewGrpcService(c.GRpcService)
	if err != nil {
		return nil, err
	}
	return grpcService, nil
}
