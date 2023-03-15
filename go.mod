module rockimserver

go 1.19

require (
	github.com/Shopify/sarama v1.23.1
	github.com/antlabs/timer v0.0.10
	github.com/bsm/sarama-cluster v2.1.15+incompatible
	github.com/emirpasic/gods v1.18.1
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/go-kratos/kratos/contrib/log/zap/v2 v2.0.0-20221009085009-468630cc4bf6
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20221026062414-3c65f16737e3
	github.com/go-kratos/kratos/contrib/registry/nacos/v2 v2.0.0-20230129033620-b242403bc125
	github.com/go-kratos/kratos/contrib/registry/zookeeper/v2 v2.0.0-20221026062414-3c65f16737e3
	github.com/go-kratos/kratos/v2 v2.5.3
	github.com/go-redis/redis/extra/rediscmd v0.2.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-resty/resty/v2 v2.7.0
	github.com/go-zookeeper/zk v1.0.3
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/websocket v1.5.0
	github.com/jinzhu/copier v0.3.5
	github.com/json-iterator/go v1.1.12
	github.com/lesismal/nbio v1.3.10
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/nacos-group/nacos-sdk-go v1.1.4
	github.com/panjf2000/gnet/v2 v2.2.5
	github.com/rs/xid v1.4.0
	github.com/samber/lo v1.37.0
	github.com/sony/sonyflake v1.1.0
	github.com/speps/go-hashids/v2 v2.0.1
	github.com/spf13/viper v1.13.0
	github.com/stretchr/testify v1.8.1
	github.com/zhenjl/cityhash v0.0.0-20131128155616-cdd6a94144ab
	go.etcd.io/etcd/client/v3 v3.5.4
	go.mongodb.org/mongo-driver v1.11.1
	go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo v0.36.4
	go.opentelemetry.io/otel v1.11.1
	go.opentelemetry.io/otel/trace v1.11.1
	go.uber.org/atomic v1.10.0
	go.uber.org/zap v1.23.0
	golang.org/x/exp v0.0.0-20221227203929-1b447090c38c
	golang.org/x/net v0.5.0
	google.golang.org/genproto v0.0.0-20220524023933-508584e28198
	google.golang.org/grpc v1.46.2
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/DataDog/zstd v1.3.6-0.20190409195224-796139022798 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.18 // indirect
	github.com/antlabs/stl v0.0.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/eapache/go-resiliency v1.3.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230111030713-bf00bc1b83b6 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/jonboulle/clockwork v0.3.0 // indirect
	github.com/klauspost/compress v1.15.14 // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/pierrec/lz4 v0.0.0-20190327172049-315a67e90e41 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.etcd.io/etcd/api/v3 v3.5.4 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.4 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/jcmturner/aescts.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/dnsutils.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/goidentity.v3 v3.0.0 // indirect
	gopkg.in/jcmturner/gokrb5.v7 v7.2.3 // indirect
	gopkg.in/jcmturner/rpc.v1 v1.1.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
