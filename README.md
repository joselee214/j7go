# go-framwork

```
.
├── Gopkg.lock									dep文件
├── Gopkg.toml									dep文件
├── cmd										    命令行启动相关
│   └── cmd.go
├── components									框架组件
│   ├── config.go								框架配置相关
│   ├── db.go									db相关, 提供全局变量components.M用于操作数据库
│   ├── engine.go								框架主引擎
│   ├── grpc.go									grpc相关
│   ├── log.go									log相关, 提供全局变量components.L用于打印log
│   ├── redis.go								redis相关, 提供全局变量components.R用于操作redis
│   └── register.go								服务注册相关
├── conf										配置文件, 项目模板里有此目录, 实际项目需要加到忽略文件里, 创建默认的conf_default目录存放配置
│   ├── default
│   │   └── config.yml
│   ├── dev
│   │   └── config.yml
│   └── prod
│       └── config.yml
├── log											log目录, 需要加到忽略文件里
│   ├── demo.log
│   └── error.log
├── main.go										项目入口文件
├── models										db models目录
│   └── const.go								models常量配置
├── modules										moddules文件目录, 实现各种业务
│   ├── greeter									demo, 实际项目删除即可
│   │   ├── const.go
│   │   └── order.go
│   └── modules.go								用于moddules注册
├── services									services目录
│   └── orderService.go
├── start.sh									本demo启动脚本, 可忽略
└── templates									xo生成db models模板
    ├── mysql.enum.go.tpl
    ├── mysql.foreignkey.go.tpl
    ├── mysql.index.go.tpl
    ├── mysql.proc.go.tpl
    ├── mysql.query.go.tpl
    ├── mysql.querytype.go.tpl
    ├── mysql.type.go.tpl
    ├── xo_db.go.tpl
    └── xo_package.go.tpl
```



### XO
xo mysql://root:awerli123@127.0.0.1/aypcddg -o generate_files/xo --template-path /Users/joselee/codes/j7go/templates
然后自行选择需要的

### protobuf
protoc --go_out=plugins=grpc:./generate_files/proto ./protobuf/ddgadmin.proto
import  j7f/proto/common
然后自行处理 common 信息 GetHeader / GetStatus 等方法 

protoc --php_out=./class/grpc --proto_path=/Users/joselee/codes/j7go/protobuf /Users/joselee/codes/j7go/protobuf/ddgadmin.proto


ab -n 1000 -c 100 -p ./conf/json_statistics.txt -T application/x-www-form-urlencoded "http://127.0.0.1:7004/c"


RUNTIME_ENV=default go run main.go

### GRPC 调试  
grpcui -plaintext 127.0.0.1:54613




