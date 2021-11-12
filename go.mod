module github.com/irisnet/irismod

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.44.0
	github.com/gin-gonic/gin v1.7.3 // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.4
	github.com/tendermint/tmlibs v0.9.0
	github.com/tidwall/gjson v1.11.0
	github.com/xeipuuv/gojsonschema v1.2.0
	google.golang.org/genproto v0.0.0-20210828152312-66f60bf46e71
	google.golang.org/grpc v1.41.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/cosmos/cosmos-sdk => /Users/zhangzhiqiang/workspace/github/dreamer-zq/cosmos-sdk

replace github.com/cosmos/cosmos-sdk/db => /Users/zhangzhiqiang/workspace/github/dreamer-zq/cosmos-sdk/db
