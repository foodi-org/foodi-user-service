# foodi-user-service

the service of foodi user center

### proto

便于文件拆分避免单文件过于庞大而缺少可读性，该目录用于统一存放`proto`文件，自行按照业务需求拆分。

使用`goctl`服务分组模式生成service，故执行生成命令时请带上`-m`参数，由于`goctl`规则限制，请勿在`proto`文件中使用`import`导入。

生成`service`的命令示例:

```shell
rpc protoc foodi-article-service.proto --go_out=../ --go-grpc_out=../  --zrpc_out=.. -m=
```

### gormctl

生成model文件工具，使用`gorm`映射。方便习惯写struct生成`model`的同学。

使用方式：

* 在`/modelgen`目录下定义`model`。(可自行定义package,在run.go中注册即可)
* 执行`/gormctl/run.go`文件内的生成方法。

### client

存放`service client`，方便`proxy`导入。

// todo 解决导入默认分支问题
