module github.com/alexadastra/hisho/hisho-core-service

go 1.17

replace (
	github.com/alexadastra/hisho/hisho-core-service/pkg/api => ./pkg/api
	github.com/alexadastra/ramme => ../../ramme
)

require (
	github.com/Masterminds/squirrel v1.5.3
	github.com/alexadastra/hisho/hisho-core-service/pkg/api v0.0.0-00010101000000-000000000000
	github.com/alexadastra/ramme v0.0.0-00010101000000-000000000000
	github.com/flowchartsman/swaggerui v0.0.0-20210303154956-0e71c297862e
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3
	github.com/jackc/pgx/v4 v4.16.1
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.48.0
	gopkg.in/yaml.v2 v2.4.0
)

require google.golang.org/protobuf v1.28.0

require (
	github.com/go-critic/go-critic v0.6.3 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/quasilyte/go-ruleguard/dsl v0.3.16 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220712132514-bdd2acd4974d // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
)
