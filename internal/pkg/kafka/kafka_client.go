package kafka

import (
	"account-consumer-service/internal/config"
	"errors"
	"time"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type KafkaClient struct {
	SchemaRegistry *SchemaRegistry
	Client         sarama.Client
	GroupClient    sarama.ConsumerGroup
}

func NewKafkaClient(cfg *config.KafkaConfig) (*KafkaClient, error) {
	kafkaConfig, err := generateSaramaConfig(cfg)
	if err != nil {
		zap.S().Fatalf("Error creating kafka configuration %v", err)
		return nil, err
	}

	sr, err := NewSchemaRegistry(cfg.SchemaRegistryHost, cfg.SchemaRegistryUser, cfg.SchemaRegistryPassword)
	if err != nil {
		zap.S().Fatal(err)
	}

	kafkaClient, err := sarama.NewClient(cfg.Hosts, kafkaConfig)
	if err != nil {
		zap.S().Fatalf("Error creating kafka client: %v", err)
		return nil, err
	}

	groupClient, err := sarama.NewConsumerGroupFromClient(cfg.ConsumerGroup, kafkaClient)
	if err != nil {
		zap.S().Fatalf("Error creating consumer group groupClient: %v", err)
	}

	return &KafkaClient{sr, kafkaClient, groupClient}, nil
}

func generateSaramaConfig(cfg *config.KafkaConfig) (*sarama.Config, error) {
	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.ClientID = cfg.ClientId
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Retry.Max = 5
	kafkaConfig.Producer.Return.Successes = cfg.EnableEvents
	kafkaConfig.Producer.Return.Errors = cfg.EnableEvents
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	kafkaConfig.Consumer.Offsets.AutoCommit.Enable = false

	kafkaConfig.Net.ReadTimeout = 60 * time.Second
	kafkaConfig.Net.WriteTimeout = 60 * time.Second
	kafkaConfig.Net.DialTimeout = 60 * time.Second
	kafkaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	kafkaConfig.Consumer.Return.Errors = true
	kafkaConfig.Consumer.Group.Session.Timeout = 30 * time.Second
	kafkaConfig.Consumer.Group.Rebalance.Timeout = 20 * time.Second
	kafkaConfig.Consumer.MaxProcessingTime = 30 * time.Second
	kafkaConfig.Consumer.Group.Heartbeat.Interval = 8 * time.Second

	kafkaConfig.Version = sarama.V3_1_0_0

	if cfg.UseAuthentication {
		kafkaConfig.Net.SASL.Mechanism = sarama.SASLMechanism(cfg.SaslMechanism)
		kafkaConfig.Net.SASL.User = cfg.User
		kafkaConfig.Net.SASL.Password = cfg.Password
		kafkaConfig.Net.TLS.Enable = cfg.EnableTLS
		kafkaConfig.Net.SASL.Enable = cfg.UseAuthentication
		if err := setAuthentication(kafkaConfig); err != nil {
			return nil, err
		}
	}
	return kafkaConfig, nil
}

func setAuthentication(conf *sarama.Config) error {
	switch conf.Net.SASL.Mechanism {
	case sarama.SASLTypeSCRAMSHA512:
		scram512Fn := func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
		conf.Net.SASL.SCRAMClientGeneratorFunc = scram512Fn
	case sarama.SASLTypeSCRAMSHA256:
		scram256Fn := func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA256} }
		conf.Net.SASL.SCRAMClientGeneratorFunc = scram256Fn
	case sarama.SASLTypePlaintext:
	default:
		return errors.New("invalid sasl mechanism")
	}
	return nil
}
