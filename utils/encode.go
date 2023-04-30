package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/klauspost/reedsolomon"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func EncodeVideo(inputFile string, outputFile string, width int, height int, a int, b int) (err error, base64Str string) {
	tmpEncodeDir := "tmp_encode_dir" + GenerateRandomName(8)
	allStartTime := time.Now()
	// 创建临时目录
	if _, err := os.Stat(tmpEncodeDir); os.IsNotExist(err) {
		fmt.Println("Encode: 创建临时目录")
		err = os.Mkdir(tmpEncodeDir, 0755)
		if err != nil {
			return fmt.Errorf("create tmp directory failed: %w\n", err), ""
		}
	}

	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("read input file failed: %w\n", err), ""
	}

	fmt.Println("Encode: 计算哈希值")
	encodeFileHash := sha256.Sum256(data)
	hashStr := fmt.Sprintf("%x", encodeFileHash)
	fmt.Println("Encode: 源文件 Hash: " + hashStr)

	fmt.Println("Encode: 生成纠错码")
	rs, err := reedsolomon.New(a, b, reedsolomon.WithMaxGoroutines(runtime.NumCPU()))
	if err != nil {
		return fmt.Errorf("create Reed-Solomon encoder failed: %w\n", err), ""
	}
	shards, err := rs.Split(data)
	if err != nil {
		return fmt.Errorf("split data into shards failed: %w\n", err), ""
	}
	fmt.Println("Encode: 开始编码")
	err = rs.Encode(shards)
	if err != nil {
		fmt.Println("Encode: 编码出现错误，看起来a和b值不能符合编码要求，尝试使用默认的a和b值进行编码:  a = 20, b = 10")
		err := rs.Encode(shards)
		if err != nil {
			fmt.Println("Encode: 编码出现错误，看起来默认提供的a和b值不能符合要求，请确认：1. 输入的文件数据长度最好是a和b的倍数 2.a和b小于256 ，建议改变文件长度或者a和b的值，然后重新编码")
			return fmt.Errorf("Reed-Solomon encoding failed: %w\n", err), ""
		}
	}
	runtime.GC()
	rows := len(shards)
	var dataEncoded []byte
	for i := 0; i < len(shards); i++ {
		for j := 0; j < len(shards[i]); j++ {
			dataEncoded = append(dataEncoded, shards[i][j])
		}
	}

	// 计算单张图像的像素数量
	numPixelsPerImage := width * height

	// 创建像素列表和计数器
	length := len(dataEncoded) * 8
	pixels := make([]color.Color, numPixelsPerImage)
	count := 0
	i := 0

	fmt.Println("Encode: 生成图片序列")
	startTime := time.Now()
	for _, byteVal := range dataEncoded {
		binary := fmt.Sprintf("%08b", byteVal) // 将每个字节转换为8位的二进制字符串
		for _, bitRune := range binary {
			bit := int(bitRune - '0')
			if bit == 0 {
				pixels[count] = color.Color(color.RGBA{R: 0, G: 0, B: 0, A: 255}) // 0对应黑色
			} else {
				pixels[count] = color.Color(color.RGBA{R: 255, G: 255, B: 255, A: 255}) // 1对应白色
			}
			count++
			if count == numPixelsPerImage {
				img := image.NewRGBA(image.Rect(0, 0, width, height))
				for y := 0; y < height; y++ {
					for x := 0; x < width; x++ {
						pixelIndex := y*width + x
						img.Set(x, y, pixels[pixelIndex])
					}
				}
				// 将图像编码为 JPEG 格式并保存到文件
				file, err := os.Create(filepath.Join(tmpEncodeDir, fmt.Sprintf("image_%09d.jpg", i)))
				if err != nil {
					return fmt.Errorf("create file failed: %w\n", err), ""
				}
				if err := jpeg.Encode(file, img, nil); err != nil {
					file.Close() // 出现错误时关闭文件句柄
					return fmt.Errorf("jpeg encode failed: %w\n", err), ""
				}
				err = file.Close() // 手动关闭文件句柄
				if err != nil {
					return fmt.Errorf("close file failed: %w\n", err), ""
				}
				if (i+1)%1000 == 0 {
					endTime := time.Now()
					duration := endTime.Sub(startTime)
					fmt.Printf("Encode: 已生成%d帧，还有%d帧未处理，总共%d帧，每1000帧耗时%f秒\n", i+1, length/numPixelsPerImage-i, length/numPixelsPerImage+1, duration.Seconds())
					startTime = time.Now()
				}
				// 重置计数器和像素列表
				count = 0
				pixels = make([]color.Color, numPixelsPerImage)
				i++
			}
		}
	}

	if count > 0 {
		// 补全最后一张图像的所有像素为黑色
		for count < numPixelsPerImage {
			pixels[count] = color.Color(color.RGBA{R: 0, G: 0, B: 0, A: 255})
			count++
		}
		// 将当前图片保存到文件中
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				pixelIndex := y*width + x
				img.Set(x, y, pixels[pixelIndex])
			}
		}
		// 将图像编码为 JPEG 格式并保存到文件
		file, err := os.Create(filepath.Join(tmpEncodeDir, fmt.Sprintf("image_%09d.jpg", i)))
		if err != nil {
			return fmt.Errorf("create file failed: %w\n", err), ""
		}
		if err := jpeg.Encode(file, img, nil); err != nil {
			file.Close() // 出现错误时关闭文件句柄
			return fmt.Errorf("jpeg encode failed: %w\n", err), ""
		}
		err = file.Close() // 手动关闭文件句柄
		if err != nil {
			return fmt.Errorf("close file failed: %w\n", err), ""
		}
		if err != nil {
			return fmt.Errorf("JPEG encoding failed: %w\n", err), ""
		}
		fmt.Printf("Encode: 已生成共%d帧\n", i+1)
	}

	fmt.Println("Encode: 图片序列生成成功")
	fmt.Println("Encode: 转换为视频")
	args := []string{
		"-i", filepath.Join(tmpEncodeDir, "image_%09d.jpg"), "-pix_fmt", "yuv420p", "-c:v", "libx264", "-crf", "18", "-y", outputFile,
	}
	if err := RunCommand("ffmpeg", args); err != nil {
		return fmt.Errorf("FFmpeg command failed: %w\n", err), ""
	}
	fmt.Println("Encode: 完成")
	allEndTime := time.Now()
	allDuration := allEndTime.Sub(allStartTime)
	fmt.Printf("Encode: 总共耗时%f秒\n", allDuration.Seconds())
	info := map[string]interface{}{
		"a":           a,
		"b":           b,
		"width":       width,
		"height":      height,
		"length":      length,
		"rows":        rows,
		"hash_sha256": hashStr,
	}
	jsonInfo, err := json.Marshal(info)
	if err != nil {
		return fmt.Errorf("JSON Marshal failed: %w\n", err), ""
	}
	b64Info := base64.StdEncoding.EncodeToString(jsonInfo)
	fmt.Println("Encode: 保存Base64配置到config.txt")
	err = os.WriteFile("config_"+strings.Replace(inputFile, "/", "-", -1)+".txt", []byte(b64Info), 0644)
	if err != nil {
		fmt.Printf("write config file failed: %w\n", err)
	}
	fmt.Printf(`Encode: 下面是解密该文件所需要的配置(Base64直接粘贴到程序下即可使用)：
———————Base64——————
%s
———————JSON————————
%s
———————————————————
`, b64Info, string(jsonInfo))

	fmt.Println("Encode: 删除临时目录")
	tries := 1
	for {
		err = os.RemoveAll(tmpEncodeDir)
		if err != nil {
			if tries >= 5 {
				fmt.Printf("删除临时目录失败，跳过任务: %w\n", err)
				break
			}
			fmt.Printf("Encode: 删除临时目录失败，重试第%d次: %w\n", tries, err)
			tries++
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}
	fmt.Printf("Encode: 视频转换成功，建议调用解码函数验证数据在部分丢失的情况下能否恢复原文件\n\n")
	return nil, b64Info
}

// DecodeVideo 反向函数实现：将视频转换回原始文件
func DecodeVideo(inputFile string, outputFile string, a int, b int, width int, height int, length int, rows int, hashSha256 *string) error {
	tmpDecodeDir := "tmp_decode_jpgs" + GenerateRandomName(8)
	allStartTime := time.Now()
	fmt.Println("Decode: 转换图像序列")

	err := os.Mkdir(tmpDecodeDir, 0755)
	if err != nil {
		return fmt.Errorf("MkdirtmpDecodeDir command failed(please first delete tmpDecodeDir): %w\n", err)
	}
	args := []string{
		"-i", inputFile, "-vf", fmt.Sprintf("scale=%d:%d", width, height), "-q:v", "1", filepath.Join(tmpDecodeDir, "image_%09d.jpg"),
	}
	if err := RunCommand("ffmpeg", args); err != nil {
		return fmt.Errorf("FFmpeg command failed: %w\n", err)
	}

	fmt.Println("Decode: 转换到纠错码")
	// 创建像素列表和计数器
	pixels := make([]int, 0)
	dataByte := make([]byte, 0)
	count := 0
	i := 0
	startTime := time.Now()

	files, err := os.ReadDir(tmpDecodeDir)
	if err != nil {
		return fmt.Errorf("ReadDirTmpDecodeDir failed: %w\n", err)
	}
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".jpg" {
			imagePath := filepath.Join(tmpDecodeDir, file.Name())
			imgFile, err := os.Open(imagePath)
			if err != nil {
				return fmt.Errorf("open image failed: %w\n", err)
			}
			img, _, err := image.Decode(imgFile)
			if err != nil {
				imgFile.Close()
				return fmt.Errorf("decode image failed: %w\n", err)
			}
			bounds := img.Bounds()
			width, height = bounds.Max.X, bounds.Max.Y
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					if i == length {
						break
					}
					// 获取像素的RGB值
					r, g, b, _ := img.At(x, y).RGBA()
					r >>= 8
					g >>= 8
					b >>= 8
					if r <= 127 && g <= 127 && b <= 127 {
						pixels = append(pixels, 0)
					} else {
						pixels = append(pixels, 1)
					}
					count++
					i++
					if count == 8 {
						var res int
						for _, num := range pixels {
							res = (res << 1) | num
						}
						dataByte = append(dataByte, byte(res))
						if (i+1)%1000000 == 0 {
							endTime := time.Now()
							duration := endTime.Sub(startTime)
							startTime = time.Now()
							fmt.Printf("Decode: 已处理 %d bit，还有 %d bit 未处理，总共 %d bit，每 1000000 bit 耗时 %.8f 秒\n", i+1, length-i-1, length, duration.Seconds())
						}
						// 重置计数器和像素列表
						count = 0
						pixels = pixels[:0]
					}
					if i == length {
						break
					}
				}
				if i == length {
					break
				}
			}
			err = imgFile.Close()
			if err != nil {
				return err
			}
		}
	}

	fmt.Println("Decode: 解码纠错码")
	dataDecoded := ConvertTo2DArray(dataByte, rows)
	rs, err := reedsolomon.New(a, b, reedsolomon.WithMaxGoroutines(runtime.NumCPU()))
	if err != nil {
		return fmt.Errorf("reedsolomn new failed: %w\n", err)
	}
	err = rs.Reconstruct(dataDecoded)
	if err != nil {
		return fmt.Errorf("reedsolomn reconstruct failed: %w\n", err)
	}
	dataOriginal := make([]byte, 0)
	dataDecoded2 := dataDecoded[:a]
	for i := range dataDecoded2 {
		dataOriginal = append(dataOriginal, dataDecoded2[i]...)
	}
	fileToOutput, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("reedsolomn new failed: %w\n", err)
	}
	defer fileToOutput.Close()
	_, err = fileToOutput.Write(dataOriginal)
	if err != nil {
		return err
	}

	allEndTime := time.Now()
	allDuration := allEndTime.Sub(allStartTime)
	fmt.Printf("Decode: 总共耗时%f秒\n", allDuration.Seconds())

	if hashSha256 != nil && *hashSha256 != "" {
		fmt.Println("Decode: 对比哈希是否正确")
		h := sha256.New()
		h.Write(dataOriginal)
		actualHash := hex.EncodeToString(h.Sum(nil))
		if actualHash == *hashSha256 {
			fmt.Println("Decode: 哈希比对正确")
		} else {
			fmt.Println("Decode: 哈希比对错误，生成的文件可能损坏")
		}
		fmt.Println("Decode: 保存哈希")
		err = os.WriteFile("hash_SHA256_"+strings.Replace(outputFile, "/", "-", -1)+".txt", []byte(actualHash), 0644)
		if err != nil {
			fmt.Printf("writefile failed: %w\n", err)
		}
	}

	fmt.Println("Decode: 删除临时目录")
	tries := 1
	for {
		err = os.RemoveAll(tmpDecodeDir)
		if err != nil {
			if tries >= 5 {
				fmt.Printf("删除临时目录失败，跳过任务: %w\n", err)
				break
			}
			fmt.Printf("Encode: 删除临时目录失败，重试第%d次: %w\n", tries, err)
			tries++
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}
	fmt.Println("")
	return nil
}

// DecodeVideoByBase64 通过 base64 字符串解码视频
func DecodeVideoByBase64(encodeFile string, decodeFile string, base64Str string) error {
	fmt.Println("Decode: 读取到配置")
	jsonInfoStr, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return fmt.Errorf("base64 DecodeString failed: %w\n", err)
	}
	var jsonInfo struct {
		A          int    `json:"a"`
		B          int    `json:"b"`
		Width      int    `json:"width"`
		Height     int    `json:"height"`
		Length     int    `json:"length"`
		Rows       int    `json:"rows"`
		HashSha256 string `json:"hash_sha256"`
	}
	err = json.Unmarshal(jsonInfoStr, &jsonInfo)
	if err != nil {
		return fmt.Errorf("json Unmarshal failed: %w\n", err)
	}
	err = DecodeVideo(encodeFile, decodeFile, jsonInfo.A, jsonInfo.B, jsonInfo.Width, jsonInfo.Height, jsonInfo.Length, jsonInfo.Rows, &jsonInfo.HashSha256)
	if err != nil {
		return fmt.Errorf("DecodeVideo failed: %w\n", err)
	}
	return nil
}
