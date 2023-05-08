package kafka

import (
	"strings"

	"github.com/Shopify/sarama"
	"github.com/go-logr/zerologr"
	"github.com/spf13/viper"
)

var (
	version string = "2"
)

type KafkaClient struct {
	Config   *sarama.Config
	Producer sarama.AsyncProducer
	logger   zerologr.Logger
}

func NewKafkaClient(logr zerologr.Logger) *KafkaClient {
	sConf := getKafkaClientConfig()

	// Handle panic
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		logr.Info("Kafa connection failed", "Rec", r)
	// 	}
	// }()

	//brokers := []string{"127.0.0.1:29092"}
	kbrokers := viper.GetString("kafka.bootstrapServers")
	kproducers := strings.Split(kbrokers, ",")
	logr.Info("Trying to connect to kafka brokers", "brokers", kproducers)

	producer, err := sarama.NewAsyncProducer(kproducers, sConf)
	if err != nil {
		logr.Error(err, "Failed to initiate kafka producer")
		panic(err)
	}

	return &KafkaClient{
		Config:   sConf,
		Producer: producer,
		logger:   logr,
	}
}

func getKafkaClientConfig() *sarama.Config {
	config := sarama.NewConfig()
	//config.Version = "0.11.0.0"
	config.Producer.Idempotent = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.Transaction.Retry.Backoff = 1
	config.Producer.Transaction.ID = "trooper_producer"
	config.Net.MaxOpenRequests = 1
	return config
}
