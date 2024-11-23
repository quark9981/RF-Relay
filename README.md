# RF-Relay
一个中继无线信号的有趣实验An interesting experiment in relaying wireless signals
## 简介
RF-Relay 使用 HackRF 设备进行中继无线信号实验的概念验证，无法保证通信效率。
由`client-rx` 和 `server-tx`两个程序主题构成：。其中`client-rx` 接收无线信号通过 TCP 发送到 `server-tx`，`server-tx` 接收数据通过 HackRF 发送出去。
### 注意事项

- 需确保 `client-rx` 和 `server-tx` 运行在可互通的网络环境，即 `server-tx` 的 IP 和端口可被 `client-rx` 访问，先运行 `server-tx` 后运行 `server-tx` 。
- 运行程序时需要使用 `sudo`权限。
- 可以通过修改`client-rx` 和 `server-tx`中的hackrf命令频率参数在不同频率工作。

## 实验原理
### `server-tx`

1. `server-tx` 程序启动后，首先解析命令行参数获取监听的 IP 和端口。
2. 创建一个命名管道 `/tmp/pipetx.ipc`，用于存储从客户端接收到的数据。
3. 启动一个 Goroutine 执行 `hackrf_transfer` 命令，将命名管道中的数据通过 HackRF 设备发送出去。
4. 主线程打开命名管道并监听指定的 IP 和端口。
5. 接受客户端连接，从 TCP 连接读取数据，并写入命名管道。

### `client-rx`

1. `client-rx` 程序启动后，首先解析命令行参数获取服务器的 IP 和端口。
2. 创建一个命名管道 `/tmp/piperx.ipc`，用于存储从 HackRF 接收到的数据。
3. 启动一个 Goroutine 执行 `hackrf_transfer` 命令，从 HackRF 设备接收无线信号，并将数据写入命名管道。
4. 主线程打开命名管道并连接到服务器。
5. 从命名管道读取数据，并通过 TCP 发送到服务器。

`client-rx` 接收到的无线信号通过 TCP 传输到 `server-tx`，再通过 HackRF 设备重新发送，实现信号的中继。

## 硬件需求

需要以下硬件设备：

1. **两台PC或虚拟机**：一台用于运行 `client-rx`，另一台用于运行 `server-tx`，需要在同一个网络（网络互通即可）中。
2. **两台 HackRF One 设备**：每台PC或虚拟机各连接一台，用于接收和发送无线信号。

## 安装步骤

### 1. 安装基础软件环境

在每台PC或虚拟机上，安装下列基础软件：

1. **Go**：需要安装 Go 1.18 或更高版本。


2. **HackRF Tools**：可通过以下命令安装：

   ```bash
   sudo apt update
   sudo apt install hackrf
   ```

3. **Git**：用于克隆代码仓库

   ```bash
   sudo apt install git
   ```

### 2. 构建与使用

克隆代码仓库：

```bash
git clone https://github.com/quark9981/RF-Relay.git
```

分别构建 `client-rx` 和 `server-tx`：

#### 构建 `client-rx`

```bash
go mod init client-rx
cd client-rx
go build -o client-rx main.go
```

#### 构建 `server-tx`

```bash
go mod init server-tx
cd server-tx
go build -o server-tx main.go
```

### 3. 验证 HackRF 安装

验证各主机或虚拟机 HackRF 设备是否就绪。

1. **检查 HackRF 工具是否已安装**

   ```bash
   hackrf_info
   ```

  若环境就绪，上述命令将分别打印 HackRF 设备的信息。

2. **检查 HackRF 设备连接**

   ```bash
   lsusb
   ```
  若环境就绪，上述命令将输出 HackRF 设备接口信息。


### 4. 启动 RF-Relay

分别构建 `client-rx` 和 `server-tx`：

### 构建 `client-rx`

```bash
cd client-rx
go build -o client-rx main.go
```

### 构建 `server-tx`

```bash
cd server-tx
go build -o server-tx main.go
```

## 使用

### 启动 `server-tx`

在 `server-tx` 所在目录下，运行以下命令：

```bash
sudo ./server-tx -i <server-ip> -p <server-port>
```

例如：

```bash
sudo ./server-tx -i 192.168.1.100 -p 8080
```

### 启动 `client-rx`

在 `client-rx` 所在目录下，运行以下命令：

```bash
sudo ./client-rx -i <server-ip> -p <server-port>
```

例如：

```bash
sudo ./client-rx -i 192.168.1.100 -p 8080
```

需保证两台PC在同一个网络中。

### 5. 测试

如果一切正常，`client-rx` 和 `server-tx` 应该会显示OK输出，说明 HackRF 设备已启动并正在发送和接收数据。

- `client-rx` 应该显示类似如下的输出：

  ```
  RX OK
  ```

- `server-tx` 应该显示类似如下的输出：

  ```
  TX OK
  ```
