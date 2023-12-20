# Port-Releaser

Port-Releaser 是一个跨平台的命令行工具，旨在识别并可选择性地终止占用特定端口的进程。它支持 Windows、macOS 和 Linux，是开发者和系统管理员的多功能工具。

## 特性

- **跨平台支持**：在 Windows、macOS 和 Linux 上运行。
- **易于使用**：通过简单的命令快速找到使用特定端口的进程。
- **安全终止进程**：识别进程后提供安全终止的选项。

## 安装

要安装 Port-Releaser，克隆此仓库并使用 Go 构建工具：

```bash
git clone https://github.com/feiandxs/port-releaser.git
cd port-releaser
go build
```

## 使用方法

使用 Port-Releaser，只需运行以下命令并加上所需的端口号：

```bash
./port-releaser <端口号>
```

例如，要找出使用端口 8080 的进程：

```bash
./port-releaser 8080
```

该工具将列出使用指定端口的所有进程，并提示您是否希望终止它们。

## 贡献

欢迎并感谢对 Port-Releaser 的贡献。如果您有改进意见或发现了错误，请提交 pull request 或开设 issue。

## 许可证

算了吧，没有一行代码是我写的，都是 ChatGPT 生成的。
