# pvc挂载不上pv
* annotations 不一致会影响pvc的挂载

# 通过coredns 访问mysql.beatflow-data.svc和mysql.beatflow-data.svc.cluster.local ，有些时候在不同命名空间下可能访问不到

# coredns debug
```
kubectl run dig --rm -it --image=docker.io/azukiapp/dig /bin/sh
/ # nslookup kubernetes.default
Server:		10.96.0.10
Address:	10.96.0.10#53

Name:	kubernetes.default.svc.cluster.local
Address: 10.96.0.1
```

# 在kubeSphere 管理的k8s集群上搭建使用Prometheus Operator helm安装prometheus 起不来
* kubeSphere 已经安装了prometheus ,并且每个节点已经装了node-exporter所以导致起不来
* 目前我们只需要安装grafana就可以了
* k8s集群中可以通过注解实现prometheus 自动发现采集器 具体配置如下: 
```


```

# headless