package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBPORT     string
	DBUSER     string
	DBPASSWORD string
	DBHOST     string
	DBNAME     string
	DBSSLMODE  string
}
type KafkaConfig struct {
	Brokers []string
	Topics  KafkaTopics
}

type KafkaTopics struct {
	OrderCreated KafkaTopicConfig
}

type KafkaTopicConfig struct {
	Name              string
	Partitions        int
	ReplicationFactor int
}

type Config struct {
	AppPort   string
	DBConfig  DBConfig
	KafkaConf KafkaConfig
}

func Load() *Config {
	_ = godotenv.Load()
	dbcfg := DBConfig{
		DBPORT:     getEnv("DB_PORT", ""),
		DBHOST:     getEnv("DB_HOST", ""),
		DBNAME:     getEnv("DB_NAME", ""),
		DBPASSWORD: getEnv("DB_PASSWORD", ""),
		DBUSER:     getEnv("DB_USER", ""),
		DBSSLMODE:  getEnv("DB_SSLMODE", ""),
	}
	kafkaCfg := loadKafkaConfig()
	cfg := &Config{
		AppPort:   getEnv("APP_PORT", "3000"),
		DBConfig:  dbcfg,
		KafkaConf: kafkaCfg,
	}

	return cfg
}
func loadKafkaConfig() KafkaConfig {
	partitions, _ := strconv.Atoi(getEnv("KAFKA_ORDER_CREATED_PARTITIONS", "1"))
	replication, _ := strconv.Atoi(getEnv("KAFKA_ORDER_CREATED_REPLICATION", "1"))

	return KafkaConfig{
		Brokers: strings.Split(getEnv("KAFKA_BROKERS", ""), ","),
		Topics: KafkaTopics{
			OrderCreated: KafkaTopicConfig{
				Name:              getEnv("KAFKA_ORDER_CREATED_TOPIC", "order.created"),
				Partitions:        partitions,
				ReplicationFactor: replication,
			},
		},
	}
}

func (cfg *Config) CreateDBUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBConfig.DBUSER,
		cfg.DBConfig.DBUSER,
		cfg.DBConfig.DBHOST,
		cfg.DBConfig.DBPORT,
		cfg.DBConfig.DBNAME,
		cfg.DBConfig.DBSSLMODE,
	)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
