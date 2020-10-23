---
categories: 
- docker
tags:
- docker-install-ware
- mysql
---

# docker 安装mysql

## 单机版 
```
docker run -it -d -p 3306:3306 --name docker_mysql -e MYSQL_ROOT_PASSWORD=123456  --restart=always  mysql 
```