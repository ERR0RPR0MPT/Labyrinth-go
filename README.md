# Labyrinth-go

Labyrinth 用于将视频画面混淆化，从而规避一些版权和审查限制。

Labyrinth-go 是 Labyrinth 基于 Go 的实现，旨在提高 Labyrinth 的转换速度。

警告：请遵守当地法规，请勿用于非法用途，开放者不承担使用者因不正当行为而导致严重后果的任何责任。

## 安装

使用以下命令进行安装：

`go get github.com/ERR0RPR0MPT/Labyrinth-go`

程序需要调用 `ffmpeg` 来实现相关功能，因此你需要安装 FFmpeg 到你的机器中。

Linux 下安装：

`apt install ffmpeg -y`

Windows 下安装：

`在官网下载预发布的二进制文件：https://ffmpeg.org`

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

一些示例
```shell
# 生成一个随机的 laby 文件用于加解密
./Labyrinth-go generate --width 1920 --height 1080 --mode hr --name index.laby

# 使用 laby 文件加密图片
./Labyrinth-go encrypt --laby index.laby --source target.png --output encrypted.png

# 使用 laby 文件解密图片
./Labyrinth-go decrypt --laby index.laby --source encrypted.png --output decrypted.png

# 加密一个视频文件
./Labyrinth-go videoencrypt --laby index.laby --source target.mp4 --framerate 30 --routines 32

# 解密一个加密的视频文件
./Labyrinth-go videodecrypt --laby index.laby --source target_output.mp4 --framerate 30 --routines 32
```

完整用法

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

以下的两个视频仅用于测试。

### ONIMAI OP (1080P24fps)

原视频(39MB)：https://www.bilibili.com/video/BV1i44y1o7yp

加密后的视频(115MB)：https://www.bilibili.com/video/BV1ks4y127nM

上传B站二压后下载的加密后的视频(59MB)

使用本地加密源解密后的视频(83MB)：https://www.bilibili.com/video/BV1jk4y1q7iD

使用B站二压源重新解密的视频(110MB)：https://www.bilibili.com/video/BV11j411w7Ku

### Bad Apple!! (1080P60fps)

原视频：https://www.bilibili.com/video/BV1x5411o7Kn

加密后的视频：https://www.bilibili.com/video/BV1ks4y127nM

本地解密后的视频：https://www.bilibili.com/video/BV17g4y1g798

## TODO

- Bug: 加密后还原的视频会出现像素丢失/错误以及色域不正确、颜色偏灰的问题，待解决（现在的效果只能说是勉强能看，后面还有二压，丢的数据更多）
- 图片颜色位移
- 视频帧间重排
- 音频分段重排
- 音频正倒放混淆
- 音调混淆
- 加密: 文件加密到音视频、图片
- 进一步提高运行效率
