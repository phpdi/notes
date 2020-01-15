### 安装grpc
1. [Ubuntu终端使用XX-Net代理](https://blog.csdn.net/miaoqiucheng/article/details/71244665)
2. go get -u -insecure google.golang.org/grpc 使用go get安装grpc ,不加-insecure 参数会报错  x509: certificate signed by unknown authority
3. [入门示例](https://grpc.io/docs/quickstart/go/)
下载相关依赖包
下载/google.golang.org/genproto遇到错误
```
error: RPC failed; curl 18 transfer closed with 15363920 bytes remaining to read

解决
  git config --global http.postBuffer 524288000 加大Buffer
  git clone https://github.com/google/go-genproto.git --depth 1 $GOPATH/src/google.golang.org/genproto 我是加了--depth 1  这个指令才解决的
```
### 安装 protobuf
1.  go get -u github.com/google/protobuf
2.  cd $GOPATH/src/github.com/google/protobuf
3. ./autogen.sh
4. ./configure
5. make
6. make check
7. sudo make install
8. sudo ldconfig # refresh shared library cache.


### rpc使用
protoc --go_out=plugins=grpc:. route_guide.proto 生成pb.go文件