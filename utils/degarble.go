package utils

import (
	"fmt"
	"image"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func RestoreOriginalImage(lst [][]int, origin, restore string) bool {
	fmt.Printf("开始解密：%s\n", restore)

	start := time.Now()

	originalImageFile, err := os.Open(origin)
	if err != nil {
		fmt.Println("[WARN] 无对应的列表文件")
		os.Exit(1)
	}
	defer originalImageFile.Close()

	originalImage, _, err := image.Decode(originalImageFile)
	if err != nil {
		fmt.Println("[WARN] 无法解码原始图像")
		os.Exit(1)
	}

	width := originalImage.Bounds().Size().X
	height := originalImage.Bounds().Size().Y

	if len(lst[1]) != height || len(lst[0]) != width {
		fmt.Printf("[WARN] 不适合的尺寸：%dx%d != %dx%d\n", len(lst), len(lst[0]), height, width)
		os.Exit(1)
	}

	newImage := image.NewRGBA(originalImage.Bounds())

	var wg sync.WaitGroup
	for i := 0; i < width; i++ { // 并发处理每一列像素
		wg.Add(1)
		go func(column int) {
			defer wg.Done()

			for row := 0; row < height; row++ {
				pixel := originalImage.At(column, row)
				newImage.Set(lst[0][column], lst[1][row], pixel)
			}
		}(i)
	}

	wg.Wait()

	outputFile, err := os.Create(restore)
	if err != nil {
		fmt.Println("[WARN] 无法创建输出文件")
		os.Exit(1)
	}
	defer outputFile.Close()

	if err := png.Encode(outputFile, newImage); err != nil {
		fmt.Println("[WARN] 无法编码新图像")
		os.Exit(1)
	}

	elapsed := time.Since(start)
	fmt.Printf("生成图片花费的时间为：%s\n", elapsed)
	return true
}

func VideoDegarble(laby string, source string, framerate int, routines int) {
	outputDir := ""
	sourceDir := ""
	labyArr := LabyFileToList(laby)
	sourceName := strings.TrimSuffix(filepath.Base(source), filepath.Ext(source))

	if !strings.Contains(source, "/") && !strings.Contains(source, "\\") {
		outputDir = filepath.Join(".", sourceName)
		sourceDir = filepath.Join(".", filepath.Base(source))
	} else {
		outputDir = filepath.Join(filepath.Dir(source), sourceName)
		sourceDir = source
	}

	if _, err := os.Stat(outputDir); !os.IsNotExist(err) {
		os.RemoveAll(outputDir)
	}

	if err := os.Mkdir(outputDir, 0755); err != nil {
		fmt.Println("创建输出目录失败")
		return
	}

	outputDirRestore := filepath.Join(outputDir, "restore_source")
	outputDirOutput := filepath.Join(outputDir, "restore_output")

	if _, err := os.Stat(outputDirRestore); os.IsNotExist(err) {
		os.Mkdir(outputDirRestore, 0755)
	}

	if _, err := os.Stat(outputDirOutput); os.IsNotExist(err) {
		os.Mkdir(outputDirOutput, 0755)
	}

	args := []string{
		"-i", fmt.Sprintf("%s", sourceDir), "-vn", "-acodec", "libmp3lame", "-b:a", "320k", "-ar", "44100", "-ac", "2", "-q:a", "0", fmt.Sprintf("%s/%s_output.mp3", outputDir, sourceName),
	}
	fmt.Println("开始提取音乐")

	if err := RunCommand("ffmpeg", args); err != nil {
		fmt.Println("提取失败")
		return
	}

	fmt.Println("提取成功")

	args = []string{
		"-i", fmt.Sprintf("%s", sourceDir), fmt.Sprintf("%s/%%d.png", outputDirRestore),
	}
	fmt.Println("开始转换图片序列")

	if err := RunCommand("ffmpeg", args); err != nil {
		fmt.Println("转换失败")
		return
	}
	fmt.Println("转换成功")

	fmt.Println("开始解密图片序列")
	files, _ := os.ReadDir(outputDirRestore)

	var wg sync.WaitGroup
	ch := make(chan struct{}, routines)

	for _, file := range files {
		ch <- struct{}{}
		wg.Add(1)

		go func(file fs.DirEntry) {
			defer wg.Done()
			source := filepath.Join(outputDirRestore, file.Name())
			output := filepath.Join(outputDirOutput, file.Name())
			RestoreOriginalImage(labyArr, source, output)
			<-ch
		}(file)
	}
	wg.Wait()

	args = []string{
		"-r", fmt.Sprintf("%d", framerate), "-i", fmt.Sprintf("%s/%%d.png", outputDirOutput), "-i", fmt.Sprintf("%s/%s_output.mp3", outputDir, sourceName), "-vcodec", "libx264", "-pix_fmt", "yuv420p", "-c:a", "copy", fmt.Sprintf("%s_output.mp4", sourceName), "-crf", "15", "-preset", "slow", "-movflags", "+faststart", "-ac", "2", "-ar", "44100", "-b:a", "320k",
	}
	fmt.Println("开始合成视频")

	if err := RunCommand("ffmpeg", args); err != nil {
		fmt.Println("合成失败")
		return
	}

	fmt.Println("合成成功")

	fmt.Println("Success")
	return
}

func Degarble(laby string, source string, output string) {
	labyArr := LabyFileToList(laby)
	RestoreOriginalImage(labyArr, source, output)
	fmt.Println("Success")
}
