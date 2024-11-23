# RF-Relay
一个中继无线信号的非严谨实验
A non-rigorous experiment to relay wireless signals
## 简介
RF-Relay 使用 HackRF 设备进行中继无线信号实验的概念验证，无法保证通信效率。
由`client-rx` 和 `server-tx`两个程序主题构成。其中`client-rx` 接收无线信号通过 TCP 发送到 `server-tx`，`server-tx` 接收数据通过 HackRF 发送出去。
### 注意事项
- 需确保 `client-rx` 和 `server-tx` 运行在可互通的网络环境，即 `server-tx` 的 IP 和端口可被 `client-rx` 访问，先运行 `server-tx` 后运行 `server-tx` 。
- 运行程序时需要使用 `sudo`权限。
- 可以通过修改`client-rx` 和 `server-tx`中的hackrf命令频率参数在不同频率工作。
## 硬件需求
需要以下硬件设备：
1. **两台PC或虚拟机**：一台用于运行 `client-rx`，另一台用于运行 `server-tx`，需要在同一个网络（网络互通即可）中。
2. **两台 HackRF One 设备**：每台PC或虚拟机各连接一台，用于接收和发送无线信号。
## 安装步骤
### 1. 安装基础软件环境
```bash
golang默认已经安装
apt update
apt install hackrf
apt install git
git clone https://github.com/quark9981/RF-Relay.git
```
分别构建 `client-rx` 和 `server-tx`：
```bash
cd client-rx
go build -o client-rx main.go
cd ..
cd server-tx
go build -o server-tx main.go
```
## 使用
### 启动 `server-tx`
在 `server-tx` 所在目录下，运行以下命令：
```bash
sudo ./server-tx -i <server-ip> -p <server-port>
```
### 启动 `client-rx`
在 `client-rx` 所在目录下，运行以下命令：
```bash
sudo ./client-rx -i <server-ip> -p <server-port>
```
需保证两台PC在同一个网络中。

