---
categories:
- k8s
  tags:
- k8s
---

# minikube 
minikube 可以搭建一个在单节点运行的k8s集群  

使用minikube 构建k8s本地化开发环境


<!--more-->

## Minikube安装
首先我们需要下载Minikube的二进制安装包并安装:
```
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```
## Minikube启动
需要使用非root 启动
```
minikube start
```

## 安装kubectl
查看kubectl的版本号，第一次使用会直接安装kubectl
```
minikube kubectl version
```
想直接使用kubectl命令的话，可以将其复制到/bin目录下去：  

```
# 查找kubectl命令的位置
find / -name kubectl
# 找到之后复制到/bin目录下
cp /mydata/docker/volumes/minikube/_data/lib/minikube/binaries/v1.20.0/kubectl /usr/local/bin/
# 直接使用kubectl命令
kubectl version

```

## 构建本地镜像
使用与Minikube VM相同的Docker主机构建镜像，以使镜像自动存在。为此，请确保使用Minikube Docker守护进程：  
```
eval $(minikube docker-env)
```
> 注意：如果不在使用Minikube主机时，可以通过运行eval $(minikube docker-env -u)来撤消此更改。

使用Minikube Docker守护进程build Docker镜像：  
```
docker build -t hello-node:v1 .
```

## Deployment
```
kubectl run hello-node --image=hello-node:v1 --port=8080
```

这里有个坑 ,我打tag的时候打成的latest，一直报找不到镜像，原因是如果省略imagePullPolicy 镜像tag为 :latest 策略为always ，否则 策略为 IfNotPresent




