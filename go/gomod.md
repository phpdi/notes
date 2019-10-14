###环境变量GO111MODULE
>GO111MODULE=off go命令从不使用新模块支持。使用GOPATH模式(查找vendor目录和GOPATH路径下的依赖)
>GO111MODULE=on go命令开启模块支持,此时执行go get 命令,所有依赖都会下到$GOPATH/pkg/mod ,go依赖查找也只会到$GOPATH/pkg/mod进行查找
>GO111MODULE=auto 默认值,go命令根据当前目录启用或禁用模块支持。仅当当前目录位于$GOPATH/src之外并且其本身包含go.mod文件或位于包含go.mod文件的目录下时，才启用模块支持。简而言之,就是项目包含go.mod文件时会开启模块支持

### go mod命令
```

download    下载依赖的module到本地cache
edit        编辑go.mod文件
graph       打印模块依赖图
init        在当前文件夹下初始化一个新的module, 创建go.mod文件
tidy        增加丢失的module，去掉未用的module
vendor      将依赖复制到vendor下,注意依赖需要在import 中声明后才能进行导入
verify      校验依赖
why         解释为什么需要依赖
```
>go get ./... 命令可以查找出当前项目的依赖

### 阿里云代理
export GOPROXY=http://mirrors.aliyun.com/goproxy/

### 翻墙
将对应的包替换为对应github上面的包
```
replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20181127143415-eb0de9b17e85
	golang.org/x/net => github.com/golang/net v0.0.0-20181114220301-adae6a3d119a
)
```

### 使用本地包
```
require (
	modtest v0.0.0
)

replace (
	modtest v0.0.0 => ../modtest
)
```

###打包命令
1.使用GOPATH模式进行打包
```
export GO111MODULE=off
export CGO_ENABLED=0
go build  -a -v -o app main.go
```
2.使用vendor目录下包来进行打包
```
export GO111MODULE=on
export CGO_ENABLED=0
go build -mod=vendor -a -v -o app main.go
```
