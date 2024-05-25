# go-audio-stream-player

`go-audio-stream-player` 是一个用 Go 语言编写的音频流播放器，能够通过 HTTP 实时接收和播放 PCM
音频流。该项目使用 `github.com/hajimehoshi/oto` 库来实现音频播放。

## 特性

- 支持 `s16le` 音频格式
- 实时接收和播放 PCM 音频流
- 支持自定义采样率和声道数

## 安装

1. 安装 Go 语言环境（如果尚未安装），请参考 [Go 官方文档](https://golang.org/doc/install)。
2. 获取项目依赖：

```bash
go get github.com/hajimehoshi/oto
```

## 编译

克隆项目并编译：

```bash
git clone https://github.com/litongjava/go-audio-stream-player.git
cd go-audio-stream-player
go build
go install
```

## 使用

编译完成后，可以使用以下命令运行播放器：

```bash
./go-audio-stream-player -f s16le -ar <sample_rate> -ac <channels> <url>
```

参数说明：

- `-f`：音频格式，当前只支持 `s16le`
- `-ar`：采样率，例如 `16000`
- `-ac`：声道数，例如 `1`
- `<url>`：音频流的 URL，例如 `http://localhost/tts`

示例：

```bash
./go-audio-stream-player -f s16le -ar 16000 -ac 1 http://localhost/tts
```

## 许可证

此项目使用 [MIT 许可证](LICENSE)。