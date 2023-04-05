# Labyrinth-go

Labyrinth 用于将视频画面混淆化，从而规避一些版权和审查限制。

Labyrinth-go 是 Labyrinth 基于 Go 的实现，旨在提高 Labyrinth 的转换速度。

警告：请遵守当地法规，请勿用于非法用途，开放者不承担使用者因不正当行为而导致严重后果的任何责任。

## 安装

使用以下命令进行安装：

`go get github.com/ERR0RPR0MPT/Labyrinth-go`

下载源代码：

`git clone https://github.com/ERR0RPR0MPT/Labyrinth-go.git`

## 用法

运行 Labyrinth-go 时需要指定以下任意一个命令：

- generate：生成用于加解密的 laby 文件
- encrypt：加密图像文件
- decrypt：解密图像文件
- videoencrypt：加密视频文件
- videodecrypt：解密视频文件

注意：在进行加解密时，必须使用和图片宽高度匹配的 laby 文件

只有使用匹配的 laby 文件进行解密才能得到近似于加密之前的图像/视频数据

```shell
Usage: C:\Users\Weclont\AppData\Local\Temp\GoLand\___go_build_main_go.exe [command] [options]

Commands:
 generate	Generate a labyrinth image
 Options:
 --width	the width of the generated data
 --height	the height of the generated data
 --mode	the mode of the generated data (hr/vr/r)
 --name	the name of the generated data file
 encrypt	Encrypt a source image using a labyrinth file
 Options:
 --laby	the labyrinth file to use for encryption
 --source	the source file to encrypt
 --output	the output file name
 decrypt	Decrypt an encrypted image using a labyrinth file
 Options:
 --laby	the labyrinth file to use for decryption
 --source	the source file to decrypt
 --output	the output file name
 videoencrypt	Encrypt a video using a labyrinth file
 Options:
 --laby	the labyrinth file to use for video encryption
 --source	the source video file to encrypt
 --framerate	the framerate of the video
 --routines	the number of encryption routines to run in parallel
 videodecrypt	Decrypt a video using a labyrinth file
 Options:
 --laby	the labyrinth file to use for video decryption
 --source	the source video file to decrypt
 --framerate	the framerate of the video
 --routines	the number of decryption routines to run in parallel
```

## Benchmark

仅支持 CPU 运算。

使用配置为 8H16G 混淆一张 1080P 的图片时间大约在 400ms 左右

转换一个 3分11秒 的 1080P60fps 视频总用时约 10min

资源占用情况：占用 CPU 100% ，内存占用在 800M 左右

实际表现会根据机器配置上下浮动

## 效果

原视频(1080P60fps)：https://www.bilibili.com/video/BV1x5411o7Kn

加密后的视频(1080P60fps)：https://www.bilibili.com/video/BV1ks4y127nM

恢复后的视频(1080P60fps)：https://www.bilibili.com/video/BV17g4y1g798

## TODO

- Bug: 加密后还原的视频会出现部分像素丢失/错误
- 图片颜色位移
- 视频帧间重排
- 音频分段重排
- 音频正倒放混淆
- 音调混淆
- 加密: 文件加密到音视频、图片
- 进一步提高运行效率