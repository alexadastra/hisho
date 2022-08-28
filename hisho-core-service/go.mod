module github.com/alexadastra/hisho/hisho-core-service

go 1.18

replace github.com/alexadastra/hisho/hisho-core-service/pkg/api => ./pkg/api

require (
	github.com/Masterminds/squirrel v1.5.3
	github.com/alexadastra/hisho/hisho-core-service/pkg/api v0.0.0-00010101000000-000000000000
	github.com/alexadastra/ramme v0.0.3
	github.com/flowchartsman/swaggerui v0.0.0-20210303154956-0e71c297862e
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	github.com/jmoiron/sqlx v1.3.5
	github.com/lib/pq v1.10.6
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
)
