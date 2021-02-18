# 安装说明

> 中文都没有搞转，就不写英文文档了，害人害己，哈哈😄。

## 依赖程序

- go 1.13+ [安装教程](https://www.runoob.com/go/go-environment.html)
- nodejs (npm,vue) [安装教程](https://www.runoob.com/nodejs/nodejs-install-setup.html)
- go-bindata [教程](https://github.com/shuLhan/go-bindata)
- make (非必须)

注：如果不明白如何安装的同学，请自行Google或百度。

## 安装

### 二进制安装

[下载二进制包](https://github.com/ihaiker/sudis/releases)

### 源码安装

#### 下载源代码

```shell script
git clone https://github.com/ihaiker/sudis.git
```

#### 编译程序（make方式）

```shell script
cd sudis
make release
```
编译完的程序在 当前文件夹bin目录下，编译完成。

#### 编译程序 

第一步：编译生成前端页面（v3版已经上传编译结果，不用再执行此部分）
```shell script
$ cd sudis/webui 
webui$ npm i -g @vue/cli #安装vue cli
webui$ npm i #安装依赖
webui$ npm run build
```
> npm安装会很慢，安装[cnpm](https://npm.taobao.org/)会很快。

第二步：使用 go-bindata 把页面文件打包。

```shell script
$ go generate generator.go #执行此步需要安装go-bindata
```

第三步：下载go依赖包
```shell script
$ go mod download
```

第四步：编译
```shell script
$ go build 
```

### 程序配置
复制`libs/config/sudis.example.yaml` 到 conf/sudis.yaml`
```yaml
#### 是否已debug模式运行
#debug: false

#### node 节点唯一ID，不指定将会获取hostname,如果多机器下hostname重复需要明确指定
#key: sudis


#### 数据存储位置，
# 1、如果用未指定数据存储位置也没有设置配置文件，数据位置为: $HOME/.sudis
# 2、如果用户设置了配置文件，单位设置数据位置，默认为: dir(conf)
#data-path: $HOME/.sudis

#### OpenApi,WebAdmin 开放地址，默认不开放
address: 127.0.0.1:5984

#### 是否禁用web管理台，默认不禁用，此参数只有配置address起效
#disable-webui: false

#### 数据存储配置
#database:
#  type: sqlite3
#  url: sudis.db


#### 是否开发管理节点，默认不开放
manager: 127.0.0.1:5983

#### 安全加密盐值，
# 1、如果设置了此值，所有节点加入管理的验证将默认使用此值
# 2、若未设置节点将使用单独的设置
#salt: whosyourdaddy


#### 托管加入地址，默认不加入任何地址
#join: []
#### or
#join:
# - 192.168.1.254:5983[,salt]
# - 192.168.1.253:5983[,salt]
## 注意 如果盐值未设置经默认使用 salt 盐值（没有报错）
## 如果全平台使用相同的盐值，就需要把salt设置。如果不是每个单独设置


#### 程序默认等待时间，例如关闭时间
#maxwait: 15s

```

## 运行程序

```she
$ sudis -f conf/sudis.yaml 
```



## 集群部署

集群部署所有的配置文件类似。

> 部署准备，两管理多托管

两台管理端 ： 192.168.0.100 ,192.168.0.101， 三台托管机器，192.168.0.102，192.168.0.103，192.168.0.104

主节点管理配置：（最简化配置）

```yaml
address: 0.0.0.0:5984
manager: 0.0.0.0:5983
salt: whosyourdaddy
```

机器配置：(最简化配置)

```yaml
salt: whosyourdaddy ## 注意盐值必须和主节点一致
join:
	- 192.168.0.100:5983
	- 192.168.0.101:5983
```

注意：托管配置，没有打开webui管理功能，如果需要打开只需配置`address`即可，主节点虽然可以多个但是并非严格意义商的HA主节点，只是有多个主节点而已，不建议使用nginx直接代理（查看日志会出现问题）



## 进入管理

### webui 管理（需要配置开启）

webui默认地址： 超级管理员默认密码为：admin:12345678

```
http://localhost:5984
```

关于通知配置：请查询 [通知配置](./NOTIFY.md)



### cli 命令行管理

程序启动参数和命令可以通过 -h 帮助方式查询例如：

```shell
$ ./bin/sudis cli -h
SUDIS V3, 一个分布式进程控制程序。

Usage:
  sudis console [command]

Aliases:
  console, cli

Available Commands:
  add         添加程序管理
  delete      删除管理的程序
  detail      查看配置信息，JSON
  join        加入主控节点
  list        查看程序列表
  modify      修改程序
  start       启动管理的程序
  status      查看运行状态
  stop        停止管理的程序
  tag         添加程序标签
  tail        查看日志

Flags:
  -h, --help               help for console
  -k, --node string        执行的节点
  -s, --sock string        连接服务端sock地址.(default: ${data-path}/sudis.sock)
  -t, --timeout duration   wait timeout (default 15s)
```

