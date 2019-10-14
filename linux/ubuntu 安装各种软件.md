### chrome
```bash
sudo wget https://repo.fdzh.org/chrome/google-chrome.list -P /etc/apt/sources.list.d/ &&
wget -q -O - https://dl.google.com/linux/linux_signing_key.pub  | sudo apt-key add - &&
sudo apt-get update &&
sudo apt-get install google-chrome-stable
```
### yarn
```bash
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
sudo apt-get update && sudo apt-get install yarn
```
### npm
```bash
sudo apt install nodejs
sudo apt install npm
#我在安装npm的过程中提示缺少npde-gyp的依赖，尝试加-f 参数来递归安装依赖，但是不行，最后还是一个一个的去安装的。

sudo npm config set registry https://registry.npm.taobao.org #设置镜像
sudo npm install npm@latest -g  #升级npm为最新版本

sudo npm install -g n #通过n模块安装指定的nodejs
sudo n stable #安装官方稳定版本
sudo n 8.* #安装指定版本的nodejs

```

### pavucontrol(解决机箱前置耳机没声音)
```bash
 sudo apt-get install pavucontrol
# 选择配置-》模拟立体声双工
# 输出设备-》模拟耳机
```

### composer 
```bash
sudo apt-get install composer 

 composer config -g repo.packagist composer https://packagist.laravel-china.org #laravel-china 社区镜像
```

### lnmp
```bash
sudo apt-get install libcurl4-gnutls-dev #ubuntu18.04下先安这个curl扩展,不然php 中curl 请求https会报502

# https://lnmp.org/auto.html 这里可以生成无人值守安装命令
wget http://soft.vpser.net/lnmp/lnmp1.5.tar.gz -cO lnmp1.5.tar.gz && tar zxf lnmp1.5.tar.gz && cd lnmp1.5 && LNMP_Auto="y" DBSelect="4" DB_Root_Password="root" InstallInnodb="y" PHPSelect="8" SelectMalloc="1" ./install.sh lnmp
```
### xx_net
```bash
cd /usr/local && https://github.com/XX-net/XX-Net.git
sudo apt-get install miredo #开启ipv6
```
### qq
```bash
cd /usr/local/ && sudo git clone https://github.com/wszqkzqk/deepin-wine-ubuntu.git && cd deepin-wine-ubuntu && ./install
sudo apt-get install -f # 如果提示依赖,执行这个
wget https://gitee.com/wszqkzqk/deepin-wine-containers-for-ubuntu/raw/master/deepin.com.qq.im_8.9.19983deepin23_i386.deb #qq 
```

### supervisor(进程管理工具)  
Supervisor是用Python开发的一套通用的进程管理程序，能将一个普通的命令行进程变为后台daemon，并监控进程状态，异常退出时能自动重启。
	
```bash
sudo apt-get install supervisor
```
	
###  java 
	
```bash
java -version 	#输出
Command 'java' not found, but can be installed with:
sudo apt install default-jre            
sudo apt install openjdk-11-jre-headless
sudo apt install openjdk-8-jre-headless 

sudo apt install default-jre  #安装默认版本的java,ubuntu 默认安装的是java10.0.2
```

### jmeter (apache压测工具)
```bash
sudo apt install jmeter
```
### 系统托盘
安装 Gnome Shell 插件：TopIcons Plus

### golang 环境安装
1.到[官网](https://golang.org/dl/)下载最新的golang包，需要翻墙
2.将包解压到/usr/local下
```bash
sudo tar -C /usr/local -xzf go1.11.4.linux-amd64.tar.gz
```
3.将go执行命令放到环境变量中
```bash
sudo vim  ~/.profile
export PATH=$PATH:/usr/local/go/bin

source ~/.profile
```
4. go version 测试是否安装成功

### Wireshark
抓包工具

sudo add-apt-repository ppa:wireshark-dev/stable 
sudo apt update
sudo apt -y install wireshark

