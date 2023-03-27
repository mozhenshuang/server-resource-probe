# 基于Go的服务器资源探针程序
## 介绍
这是使用Go语言开发的基于`gopsutil`获取服务器硬件资源信息的探针程序，支持两种通信方式：
- HTTP Get
- WebSocket

同时支持在Windows和Linux平台上运行，只需要运行响应平台的可执行文件即可。
## 使用
### 初始化并编译
```bash
git clone http://10.0.2.203/MoZhenShuang/server-resource-probe.git
cd server-resource-probe
go mod tidy && go build
```
> **注意**：如果是在Windows上编译成Linux的可执行文件，初始化项目之后请输入以下编译设置命令：
```bash
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build
```

### 启动默认配置
#### Windows默认启动
在程序所在目录启动命令提示符窗口，并且输入
```bash
server-resource-probe.exe
```

#### Linux默认启动
```bash
//给可执行文件赋予可执行权限
chmod +777 server-resource-probe
//运行程序
./server-resource-probe
```
> **注意**：程序默认启动将会监听两个端口，webapi：8082；websocket：8083
 
### 自定义启动
程序支持使用`-m`参数指定启动相应的服务，程序支持的启动模式如下：
- wa 以webapi服务的方式启动
- ws 以websocket服务的方式启动
- all 以webapi以及websocket服务的方式启动

同时程序支持通过`-ap`和`-sp`两个参数分别指定webapi和websocket服务监听的端口。
程序支持的命令如下：
```bash
-h 查看帮助信息，显示支持的命令
-v 查看程序版本信息
-V 查看程序版本信息
-m 指定启动相应的服务
-ap 指定webapi服务监听的端口
-sp 指定websocket服务监听的端口
```
## 各服务接口说明
### WebApi服务接口
webapi服务启动之后，其地址为：`http://IP:port`
webapi服务提供如下路由：
- [GET]`/host`：获取主机所有硬件资源信息，包括：CPU、内存、硬盘、网络IO
- [GET]`/cpu`：获取主机CPU使用情况
- [GET]`/mem`：获取主机内存使用情况
- [GET]`/disk`：获取主机硬盘各分区使用情况
- [GET]`/net`：获取主机网络情况
### WebSocket服务接口
websocket服务启动之后，其地址为：`ws://IP:port/`
客户端通过websocket接口成功连接服务端之后，程序会每秒返回主机资源信息。
## 特别说明
- 由于gopsutil框架的实现方式，当获取CPU的使用情况时，需要等两秒钟才能够统计出CPU的使用情况，因此当通过webapi服务获取所有主机硬件资源使用情况时，需要等2秒才会有响应信息，若单独请求某一项硬件资源使用情况时，通常响应时间会在几毫秒。