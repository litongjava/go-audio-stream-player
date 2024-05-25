package main

import (
  "flag"
  "fmt"
  "github.com/cloudwego/hertz/pkg/common/hlog"
  "github.com/hajimehoshi/oto"
  "io"
  "net/http"
  "os"
)

func main() {
  length := len(os.Args)
  hlog.Info("len:", length)
  for i := 0; i < length; i++ {
    hlog.Infof("%d %s\n", i, os.Args[i])
  }
  // 定义命令行参数
  format := flag.String("f", "s16le", "audio format")
  sampleRate := flag.Int("ar", 16000, "sample rate")
  channels := flag.Int("ac", 1, "number of channels")
  flag.Parse()

  // 检查剩余的参数，应该是 URL
  if flag.NArg() != 1 {
    fmt.Printf("Usage: %s -f s16le -ar <sample_rate> -ac <channels> <url>\n", os.Args[0])
    return
  }
  url := flag.Arg(0)

  // 验证音频格式
  if *format != "s16le" {
    fmt.Printf("Unsupported audio format: %s\n", *format)
    return
  }

  // 初始化 Oto 上下文
  context, err := oto.NewContext(*sampleRate, *channels, 2, 8192)
  if err != nil {
    fmt.Printf("Failed to create Oto context: %v\n", err)
    return
  }
  defer context.Close()

  // 创建播放器
  player := context.NewPlayer()
  defer player.Close()

  // 发送 HTTP 请求获取音频流
  resp, err := http.Get(url)
  if err != nil {
    fmt.Printf("Failed to perform HTTP request: %v\n", err)
    return
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    fmt.Printf("HTTP request failed with status: %s\n", resp.Status)
    return
  }

  // 读取并播放音频流
  buffer := make([]byte, 1024)
  for {
    n, err := resp.Body.Read(buffer)
    if n > 0 {
      _, err := player.Write(buffer[:n])
      if err != nil {
        fmt.Printf("Failed to write to player: %v\n", err)
        return
      }
    }
    if err == io.EOF {
      break
    }
    if err != nil {
      fmt.Printf("Error reading response body: %v\n", err)
      return
    }
  }
}
