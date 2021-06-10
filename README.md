# go-api-demo
基于go 构建 restful api

# 使用插件
日志管理： [github.com/lexkong/log]
参数配置： [github.com/spf13/viper]
数据库ORM： [github.com/jinzhu/gorm]
jwt：[github.com/dgrijalva/jwt-go]
web 框架：[github.com/gin-gonic/gin]
uuid 生成：[github.com/satori/go.uuid]
参数验证：[github.com/go-playground/validator]

# 参考
错误码定制
[http://open.weibo.com/wiki/Error_code]

# 命令行
``make`` 生成GO二进制
``make gotool`` 执行 gofmt -w . 和 go tool vet . 格式化代码和源码检查
``make clean`` 做一些清理工作
``make ca`` 生成证书
``make help`` 打印帮助信息