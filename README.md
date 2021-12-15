# golib
为`gin` 框架提供的基础功能库

## 功能
- wlog 日志

## 使用
### 安装
```shell
go get -u github.com/weirwei/golib
```

### wlog
对`logrus` 进行封装，使用更加便捷
使用前需要先进行初始化
```go
import "github.com/weirwei/golib/wlog"

wlog.InitLog(wlog.LogConfig{
    Level:     "info",
    Stdout:    true,
    FileOut:   true,
    Path:      "./logs",
    Formatter: "json",
})
```
使用时仅直接调用方法即可
```go
wlog.Infof(context, "get a info %s", m)
```

### middleware
中间件
#### AccessLog 
打印单次请求的日志
效果如下：
```json
{"clientIP":"127.0.0.1","cost":0,"level":"info","method":"POST","msg":"notice","query":"[\"a=niua\",\"b=123\"]","requestID":"2872274240","requestParam":"{\"msg\":\"niu niu niu\",\"id\":2}","response":"{\"errNo\":0,\"errMsg\":\"\",\"data\":\"niu niu niu\"}","status":200,"time":"2021-12-01 16:44:44","uri":"/post?a=niua\u0026b=123"}
```
