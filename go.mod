module github.com/vulcanize/eth-statediff-fill-service

go 1.16

require (
	github.com/ethereum/go-ethereum v1.10.17
	github.com/jmoiron/sqlx v1.3.5
	github.com/joho/godotenv v1.4.0
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.19.0
	github.com/prometheus/client_golang v1.12.2
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.11.0
	github.com/vulcanize/ipld-eth-server/v4 v4.0.1-alpha
)

replace github.com/ethereum/go-ethereum v1.10.17 => github.com/vulcanize/go-ethereum v1.10.17-statediff-4.0.1-alpha
