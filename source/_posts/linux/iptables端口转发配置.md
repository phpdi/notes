# 配置访问192.168.91.16端口8090转发到目标地址192.168.91.10端口80

```
sudo iptables -t nat -A PREROUTING -d 192.168.91.16 -p tcp --dport 8090 -j DNAT --to-destination 192.168.91.10:80
sudo iptables -t nat -A POSTROUTING -d 192.168.91.10 -p tcp --dport 80 -j SNAT --to-source 192.168.91.16

# 
sudo iptables -t  nat --list --line-numbers
```

```
sudo iptables -t nat -A PREROUTING -d 172.31.164.128 -p tcp --dport 32306 -j DNAT --to-destination 192.168.49.2:32306
sudo iptables -t nat -A POSTROUTING -p tcp -d 192.168.49.2 --dport 32306 -j SNAT --to-source 172.31.164.128

```
