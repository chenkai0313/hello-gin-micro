module hello-gin-micro

go 1.14

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/blinkbean/dingtalk v0.0.0-20201231030509-45a553a84503
	github.com/elastic/go-elasticsearch/v7 v7.11.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/golang/protobuf v1.4.3
	github.com/iGoogle-ink/gotil v1.0.16
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1 // indirect
	github.com/mojocn/base64Captcha v1.3.1
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/opentracing/opentracing-go v1.1.0
	github.com/spf13/viper v1.7.1
	github.com/streadway/amqp v1.0.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible
	go.uber.org/zap v1.16.0
	golang.org/x/image v0.0.0-20190802002840-cff245a6509b
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.20.12
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
