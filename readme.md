### telegram bot 

### 非常重要

```api
goctl 暂时没有实现map[string]interface{} 在.api 不要使用map参数类型
```

### 环境安装

```
  golang 1.17.8版本
  安装包下载地址
  http://mirrors.nju.edu.cn/golang/
```

```
  goctl 版本1.3.5
  安装命令  GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@v1.3.5
```

```
  goctl-go-compact
  安装命令  GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-go-compact@latest
``` 

```
  goctl-swagger
  git clone https://github.com/zeromicro/goctl-swagger.git
  修改 go.mod 与 goctl 版本一致
  github.com/zeromicro/go-zero/tools/goctl v1.3.5
  go build -o goctl-swagger main.go
  mv goctl-swagger $GOPATH/bin
```

```
  项目tpl配置
  goctl template clean 清除
  cd $HOME/.goctl
  goctl template init 初始化``
```

### 项目构建

```api
1. 在 app/ 下面微服务目录 例如 coin目录执行，生成api目录结构
goctl api go -api ./api/*.api -dir ./api
```

### model生成

```mode
goctl model mysql ddl -src ./model/files.sql -dir ./model
```

### 文档生成

```
goctl api plugin -plugin goctl-swagger="swagger -filename verify.json -basepath /api" -api ./app/verify/api/verify.api -dir ./doc

```

### 文档迁移

``` cp ./doc/verify.json ../cointiger-golang-doc/doc 
    提交 更改
```

