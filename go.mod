module miner_core

go 1.20

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/bytedance/gopkg v0.1.1
	github.com/qcq1/common v0.0.0-20250115161123-7612381b82d1
	github.com/qcq1/rpc_miner_core v0.0.0-20250115172548-fe67e732ceaa
	go.uber.org/dig v1.18.0
	gorm.io/driver/mysql v1.4.4
	gorm.io/gen v0.3.26
	gorm.io/gorm v1.25.9
	gorm.io/plugin/dbresolver v1.5.0
)

require (
	code.byted.org/aiops/apm_vendor_byted v0.0.26 // indirect
	code.byted.org/aiops/metrics_codec v0.0.21 // indirect
	code.byted.org/aiops/monitoring-common-go v0.0.4 // indirect
	code.byted.org/bytedtrace/bytedtrace-client-go v1.2.3-pre // indirect
	code.byted.org/bytedtrace/bytedtrace-common/go v0.0.13 // indirect
	code.byted.org/bytedtrace/bytedtrace-conf-provider-client-go v0.0.26 // indirect
	code.byted.org/bytedtrace/bytedtrace-gls-switch v1.3.0 // indirect
	code.byted.org/bytedtrace/interface-go v1.0.20 // indirect
	code.byted.org/bytedtrace/serializer-go v1.0.0 // indirect
	code.byted.org/gopkg/apm_vendor_interface v0.0.3 // indirect
	code.byted.org/gopkg/asynccache v0.0.0-20201112072351-d630cb60c767 // indirect
	code.byted.org/gopkg/consul v1.2.6 // indirect
	code.byted.org/gopkg/ctxvalues v0.4.0 // indirect
	code.byted.org/gopkg/debug v0.10.1 // indirect
	code.byted.org/gopkg/env v1.6.7 // indirect
	code.byted.org/gopkg/etcd_util v0.0.0-20181016075009-149305e55690 // indirect
	code.byted.org/gopkg/etcdproxy v0.1.1 // indirect
	code.byted.org/gopkg/logid v0.0.0-20241008043456-230d03adb830 // indirect
	code.byted.org/gopkg/logs v1.2.23 // indirect
	code.byted.org/gopkg/logs/v2 v2.1.54 // indirect
	code.byted.org/gopkg/metainfo v0.1.1 // indirect
	code.byted.org/gopkg/metrics v1.4.25 // indirect
	code.byted.org/gopkg/metrics/v3 v3.1.35 // indirect
	code.byted.org/gopkg/metrics/v4 v4.1.3 // indirect
	code.byted.org/gopkg/metrics_core v0.0.38 // indirect
	code.byted.org/gopkg/net2 v1.5.0 // indirect
	code.byted.org/gopkg/stats v1.2.12 // indirect
	code.byted.org/gopkg/tccclient v1.4.2 // indirect
	code.byted.org/gopkg/thrift v1.4.13 // indirect
	code.byted.org/kite/kitex v1.18.1 // indirect
	code.byted.org/kite/rpal v0.1.22 // indirect
	code.byted.org/kitex/apache_monitor v0.1.0 // indirect
	code.byted.org/lang/trace v0.0.3 // indirect
	code.byted.org/lidar/profiler v0.4.4 // indirect
	code.byted.org/lidar/profiler/kitex v0.4.6 // indirect
	code.byted.org/log_market/gosdk v0.0.0-20220328031951-809cbf0ba485 // indirect
	code.byted.org/log_market/loghelper v0.1.9 // indirect
	code.byted.org/log_market/tracelog v0.1.4 // indirect
	code.byted.org/log_market/ttlogagent_gosdk v0.0.6 // indirect
	code.byted.org/log_market/ttlogagent_gosdk/v4 v4.0.51 // indirect
	code.byted.org/middleware/fic_client v0.2.8 // indirect
	code.byted.org/middleware/gocaller v0.0.6 // indirect
	code.byted.org/security/go-spiffe-v2 v1.0.0 // indirect
	code.byted.org/security/sensitive_finder_engine v0.3.17 // indirect
	code.byted.org/security/zti-jwt-helper-golang v1.0.9 // indirect
	code.byted.org/service_mesh/shmipc v0.2.16 // indirect
	code.byted.org/trace/trace-client-go v1.3.6 // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.18 // indirect
	github.com/apache/thrift v0.16.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.13.0 // indirect
	github.com/bits-and-blooms/bloom/v3 v3.6.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/bytedance/sonic v1.12.5 // indirect
	github.com/bytedance/sonic/loader v0.2.0 // indirect
	github.com/caarlos0/env/v6 v6.2.2 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/configmanager v0.2.2 // indirect
	github.com/cloudwego/dynamicgo v0.4.7-0.20241220085612-55704ea4ca8f // indirect
	github.com/cloudwego/fastpb v0.0.5 // indirect
	github.com/cloudwego/frugal v0.2.3 // indirect
	github.com/cloudwego/gopkg v0.1.3 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/cloudwego/kitex v0.12.1 // indirect
	github.com/cloudwego/kitex/pkg/protocol/bthrift v0.0.0-20250103074449-204f2bdb87d3 // indirect
	github.com/cloudwego/localsession v0.1.1 // indirect
	github.com/cloudwego/netpoll v0.6.5 // indirect
	github.com/cloudwego/runtimex v0.1.0 // indirect
	github.com/cloudwego/thriftgo v0.3.18 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/pprof v0.0.0-20240727154555-813a5fbdbec8 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hbollon/go-edlib v1.6.0 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/jhump/protoreflect v1.8.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kitex-contrib/registry-nacos v0.1.2 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/gls v0.0.0-20220109145502-612d0167dce5 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/nacos-group/nacos-sdk-go v1.1.5 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/shirou/gopsutil/v3 v3.22.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.19.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/tidwall/gjson v1.17.3 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	github.com/zeebo/errs v1.2.2 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	go4.org/unsafe/assume-no-moving-gc v0.0.0-20230525183740-e7c30c78aeb2 // indirect
	golang.org/x/arch v0.2.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	google.golang.org/genproto v0.0.0-20240213162025-012b6fc9bca9 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240314234333-6e1732d8331c // indirect
	google.golang.org/grpc v1.62.1 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/square/go-jose.v2 v2.4.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/datatypes v1.1.1-0.20230130040222-c43177d3cf8c // indirect
	gorm.io/hints v1.1.0 // indirect
)
