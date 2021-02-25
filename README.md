# hello-gin-micro
基于gin封装的api框架


### 使用
```
go get github.com/chenkai0313/hello-gin-micro
```

#### 使用事例
go build api.go

./main

#### 说明
- 系统基于gin封装的api框架
- 设计模式仿照php-laravel结构
- moudles 模块化设计
- 系统设计目的对接[hello-micro](https://github.com/chenkai0313/hello-micro)微服务框架
-  默认封装了jaeger (做了改动，配合使用gin使用)，初次使用，可以注释掉
- 封装了自定义的validator
- 生产/测试环境 选择基于环境变量"env"
- 内涵盖了常用的的方法(unitls目录下)
- 封装了图片验证码
- 封装了基于gorm/v2
- 封装了rabbitMq(包含了延迟队列方法封装)
- 封装了redis
- 封装了es
- 封装了小程序的相关方法（lib目录下）
- 基于makefile 的protobuf
- 服务发现使用的是etcd
- 更多内容自行挖掘
- 全部个人封装,仅供学习参考