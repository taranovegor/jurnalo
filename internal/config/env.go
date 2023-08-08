package config

import "os"

const (
	AmqpDsn       = "AMQP_DSN"
	TcpDsn        = "TCP_DSN"
	CollectorExpr = "COLLECTOR_EXPR"
)

func GetEnv(key string) string {
	return os.Getenv(key)
}
