###没有给root对应的权限
```
mysql> GRANT ALL PRIVILEGES ON *.* TO 'root'@'192.168.1.123' IDENTIFIED BY '' WITH GRANT OPTION; 
mysql> FLUSH PRIVILEGES;
```
###为了安全LNMP默认是禁止远程连接的，开启方法
```
//查看已有的iptables规则，以序号显示

iptables -L -n --line-numbers

//删除对应的DROP规则
iptables -D INPUT 5
```