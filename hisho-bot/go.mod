module github.com/alexadastra/hisho/hisho-bot

go 1.17

replace (
	github.com/alexadastra/hisho/hisho-bot/pkg/api => ./pkg/api
	github.com/alexadastra/ramme => ../../ramme
)

require (
	github.com/alexadastra/hisho/hisho-bot/pkg/api v0.0.0-00010101000000-000000000000
	github.com/alexadastra/ramme v0.0.0-00010101000000-000000000000
	github.com/flowchartsman/swaggerui v0.0.0-20210303154956-0e71c297862e
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.47.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220706132729-d86698d07c53 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
)
