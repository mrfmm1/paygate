module github.com/moov-io/paygate

go 1.13

replace go4.org v0.0.0-20190430205326-94abd6928b1d => go4.org v0.0.0-20190313082347-94abd6928b1d

require (
	github.com/Shopify/sarama v1.26.4
	github.com/antihax/optional v1.0.0
	github.com/go-kit/kit v0.10.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/goftp/file-driver v0.0.0-20180502053751-5d604a0fc0c9
	github.com/goftp/server v0.0.0-20190712054601-1149070ae46b
	github.com/gorilla/mux v1.7.4
	github.com/jaegertracing/jaeger-lib v2.2.0+incompatible
	github.com/jlaffaye/ftp v0.0.0-20200602180915-5563613968bf
	github.com/lopezator/migrator v0.3.0
	github.com/mattn/go-sqlite3 v2.0.6+incompatible
	github.com/moov-io/ach v1.4.0-rc3
	github.com/moov-io/base v0.11.0
	github.com/moov-io/customers v0.4.0-rc5.0.20200605165738-057f20d4d580
	github.com/opentracing/opentracing-go v1.1.0
	github.com/ory/dockertest/v3 v3.6.0
	github.com/ory/mail/v3 v3.0.0
	github.com/pkg/sftp v1.11.0
	github.com/prometheus/client_golang v1.7.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/uber/jaeger-client-go v2.24.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	gocloud.dev v0.20.0
	gocloud.dev/pubsub/kafkapubsub v0.19.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/text v0.3.3
	gopkg.in/yaml.v2 v2.3.0
)
