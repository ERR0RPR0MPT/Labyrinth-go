# Labyrinth-go

Labyrinth 允许数据在文件与储存冗余数据的帧序列中转换，以及使视频画面混淆化，从而使文件/视频成功规避一些版权和审查限制。

警告：请遵守当地法规，请勿用于非法用途，开放者不承担使用者因不正当行为而导致严重后果的任何责任。

## 原理

Labyrinth 使用两种模式对数据进行处理，一种为编码，另一种为混淆。

### 编码(Recommend)

Labyrinth 使用 Reed-Solomon 算法对文件进行纠错编码，使得部分损坏的文件仍然可以被解码。

编码过程将文件分为 a 份原始数据切片和 b 份冗余数据切片，其中 a : b == 原始数据切片数量 : 冗余数据切片数量。

将原始数据切片和冗余数据切片转换为视频帧，实现以视频的形式存储文件。

通过解码可以从视频帧中还原出原始数据切片和冗余数据切片，从而还原出原始文件。

### 混淆

Labyrinth 使用简单的混淆算法专门对视频及图像进行混淆处理。

这样的处理使得视频帧在一定程度上可以绕过审查，因为不明所以的人可能无法看出这是什么，但是在一定程度上仅通过人的判断就会暴露大部分明文数据，所以不推荐使用混淆。

编码在进行解码后可以得到原始文件，但由于处理过程存在大量的有损压缩导致的数据丢失，因此混淆几乎无法还原出原始文件，但反混淆得到的画面与原始视频基本上是一致的。

## 依赖

- Go 1.19

## 安装

下载源码：

`git clone https://github.com/ERR0RPR0MPT/Labyrinth-go.git`

程序需要调用 `ffmpeg` 来实现相关功能，因此你需要安装 `ffmpeg` 到你的机器中。

Linux 下安装：

`apt install ffmpeg -y`

Windows 下安装：

`在官网下载预发布的二进制文件：https://ffmpeg.org`

## 用法

运行 Labyrinth-go 时需要指定以下任意一个命令：

### 使用编码：

- encode：将文件转换为冗余视频的形式
- decode：将冗余视频的形式还原为视频
- decodeb64：通过 Base64 配置将冗余视频的形式还原为视频

一些示例
```shell
# 编码冗余视频
./Labyrinth-go encode --input input.mp4 --output encoded.mp4 --width 64 --height 36 --a 100 --b 150

# 解码冗余视频（使用 encode 自动输出的 Base64 配置）
./Labyrinth-go decodeb64 --input encoded.mp4 --output decoded.mp4 --base64 "..."

# 解码冗余视频
./Labyrinth-go decode --input encoded.mp4 --output decoded.mp4 --width 64 --height 36 --a 100 --b 150 --length 1000 --rows 10 --hash "..."
```

其中 a : b == 原始数据切片数量 : 冗余数据切片数量

注意：最好使 a, b 是文件数据长度的倍数，否则可能会导致解码失败（原因在于无法解决数据长度能否被整除，待解决）

### 使用混淆：

- generate：生成用于混淆的 laby 文件
- garble：混淆图像文件
- degarble：反混淆图像文件
- videogarble：混淆视频文件
- videodegarble：反混淆视频文件

目前更加推荐使用编码的方式进行转换，因为编码的方式转换速度更快，转换后的视频文件体积更小。 混淆的视频更易被审查拦截，因为存在明文信息的直接泄露。

注意：在进行混淆时，必须使用和图片宽高度匹配的 laby 文件，只有使用匹配的 laby 文件进行解密才能得到近似于加密之前的图像/视频数据

一些示例
```shell
# 生成一个随机的 laby 文件用于加解密
./Labyrinth-go generate --width 1920 --height 1080 --mode hr --name index.laby

# 使用 laby 文件加密图片
./Labyrinth-go garble --laby index.laby --source target.png --output garbleed.png

# 使用 laby 文件解密图片
./Labyrinth-go degarble --laby index.laby --source garbleed.png --output degarbleed.png

# 加密一个视频文件
./Labyrinth-go videogarble --laby index.laby --source target.mp4 --framerate 30 --routines 32

# 解密一个加密的视频文件
./Labyrinth-go videodegarble --laby index.laby --source target_output.mp4 --framerate 30 --routines 32
```

完整用法

```shell
Usage: Labyrinth-go.exe [command] [options]

Commands:
encode  Encode a file
 Options:
 --input        the input file to encode
 --output       the output file name
 --width        the width of the encoded video
 --height       the height of the encoded video
 --a    the dataShards value of the encoded video
 --b    the parityShards value of the encoded video
decode  Decode a file
 Options:
 --input        the input file to decode
 --output       the output file name
 --a    the dataShards value of the decoded video
 --b    the parityShards value of the decoded video
 --width        the width of the decoded video
 --height       the height of the decoded video
 --length       the length of the decoded video
 --rows the rows of the decoded video
 --hash the hash of the decoded video
decodeb64       Decode a base64 string
 Options:
 --input        the input file to decode
 --output       the output file name
 --base64str    the base64 string to decode
 generate       Generate a labyrinth image
 Options:
 --width        the width of the generated data
 --height       the height of the generated data
 --mode the mode of the generated data (hr/vr/r)
 --name the name of the generated data file
 garble Garble a source image using a labyrinth file
 Options:
 --laby the labyrinth file to use for garbleion
 --source       the source file to garble
 --output       the output file name
 degarble       Degarble an garbleed image using a labyrinth file
 Options:
 --laby the labyrinth file to use for degarbleion
 --source       the source file to degarble
 --output       the output file name
 videogarble    Garble a video using a labyrinth file
 Options:
 --laby the labyrinth file to use for video garbleion
 --source       the source video file to garble
 --framerate    the framerate of the video
 --routines     the number of garbleion routines to run in parallel
 videodegarble  Degarble a video using a labyrinth file
 Options:
 --laby the labyrinth file to use for video degarbleion
 --source       the source video file to degarble
 --framerate    the framerate of the video
 --routines     the number of degarbleion routines to run in parallel
```

## 混淆的相关说明

### Benchmark

仅支持 CPU 运算。

使用配置为 8H16G 混淆一张 1080P 的图片时间大约在 400ms 左右

转换一个 3分11秒 的 1080P60fps 视频总用时约 10min

资源占用情况：占用 CPU 100% ，内存占用在 800M 左右

实际表现会根据机器配置上下浮动

### 效果

以下的两个视频仅用于测试。

#### ONIMAI OP (1080P24fps)

原视频(39MB)：https://www.bilibili.com/video/BV1i44y1o7yp

加密后的视频(115MB)：https://www.bilibili.com/video/BV1ks4y127nM

上传B站二压后下载的加密后的视频(59MB)

使用本地加密源解密后的视频(83MB)：https://www.bilibili.com/video/BV1jk4y1q7iD

使用B站二压源重新解密的视频(110MB)：https://www.bilibili.com/video/BV11j411w7Ku

#### Bad Apple!! (1080P60fps)

原视频：https://www.bilibili.com/video/BV1x5411o7Kn

加密后的视频：https://www.bilibili.com/video/BV1ks4y127nM

本地解密后的视频：https://www.bilibili.com/video/BV17g4y1g798

### TODO(混淆)

- Bug: 加密后还原的视频会出现像素丢失/错误以及色域不正确、颜色偏灰的问题，待解决（现在的效果只能说是勉强能看，后面还有二压，丢的数据更多）
- Bug: 转换所需的磁盘空间要求非常大，成本很高，解决方案是需要将视频帧转换内置到程序内部完成
- 图片颜色位移
- 视频帧间重排
- 音频分段重排
- 音频正倒放混淆
- 音调混淆
- 加密: 文件加密到音视频、图片
- 进一步提高运行效率
