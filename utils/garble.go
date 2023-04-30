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

func GenerateRandomImage(lst [][]int, origin, output string) bool {
	fmt.Printf("开始加密：%s\n", output)
	start := time.Now()

	file, err := os.Open(origin)
	if err != nil {
		fmt.Println("[WARN] 无对应的列表文件")
		os.Exit(1)
	}
	defer file.Close()

	originalImage, _, err := image.Decode(file)
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
				indexX := lst[0][column]
				indexY := lst[1][row]
				newImage.Set(column, row, originalImage.At(indexX, indexY))
			}
		}(i)
	}

	wg.Wait()

	outputFile, err := os.Create(output)
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

func VideoGarble(laby, source string, framerate int, routines int) {

}

func VideoGarble1(laby, source string, framerate int, routines int) {
	outputDir := ""
	sourceDir := ""
	labyArr := LabyFileToList(laby)
	sourceName := strings.Split(filepath.Base(source), ".")[0]
	if !strings.Contains(source, "/") || !strings.Contains(source, "\\") {
		outputDir = filepath.Join(os.Getenv("PWD"), strings.Split(filepath.Base(source), ".")[0])
		sourceDir = source
	} else {
		outputDir = filepath.Join(filepath.Dir(source), strings.Split(filepath.Base(source), ".")[0])
		sourceDir = source
	}
	if _, err := os.Stat(outputDir); err == nil {
		os.RemoveAll(outputDir)
	}
	os.Mkdir(outputDir, os.ModePerm)

	outputDirSource := filepath.Join(outputDir, "source")
	outputDirOutput := filepath.Join(outputDir, "output")
	os.Mkdir(outputDirSource, os.ModePerm)
	os.Mkdir(outputDirOutput, os.ModePerm)

	fmt.Println("输出路径：", outputDir)
	fmt.Println("资源路径：", sourceDir)
	fmt.Println("文件名：", source)

	args := []string{
		"-i", fmt.Sprintf("%s", sourceDir), "-vn", "-acodec", "libmp3lame", "-b:a", "320k", "-ar", "44100", "-ac", "2", "-q:a", "0", fmt.Sprintf("%s/%s_output.mp3", outputDir, sourceName),
	}
	fmt.Printf("开始提取音乐\n")
	if err := RunCommand("ffmpeg", args); err != nil {
		fmt.Println("提取失败")
		return
	}
	fmt.Println("提取成功")

	args = []string{
		"-i", fmt.Sprintf("%s", sourceDir), fmt.Sprintf("%s/%%d.png", outputDirSource),
	}
	fmt.Printf("开始转换图片序列\n")
	if err := RunCommand("ffmpeg", args); err != nil {
		fmt.Println("转换失败")
		return
	}
	fmt.Println("转换成功")

	fmt.Println("开始加密图片序列")
	files, _ := os.ReadDir(outputDirSource)

	var wg sync.WaitGroup
	ch := make(chan struct{}, routines)

	for _, file := range files {
		ch <- struct{}{}
		wg.Add(1)

		go func(file fs.DirEntry) {
			defer wg.Done()
			source := filepath.Join(outputDirSource, file.Name())
			output := filepath.Join(outputDirOutput, file.Name())
			GenerateRandomImage(labyArr, source, output)
			<-ch
		}(file)
	}
	wg.Wait()

	args = []string{
		"-r", fmt.Sprintf("%d", framerate), "-i", fmt.Sprintf("%s/%%d.png", outputDirOutput), "-i", fmt.Sprintf("%s/%s_output.mp3", outputDir, sourceName), "-vcodec", "libx264", "-pix_fmt", "yuv420p", "-c:a", "copy", fmt.Sprintf("%s_output.mp4", sourceName), "-crf", "15", "-preset", "slow", "-movflags", "+faststart", "-ac", "2", "-ar", "44100", "-b:a", "320k",
	}
	fmt.Printf("开始合成视频\n")
	if err := RunCommand("ffmpeg", args); err != nil {
		fmt.Println("转换失败")
		return
	}
	fmt.Println("转换成功")

	fmt.Println("Success")
	return
}

func Garble(laby, source, output string) {
	labyArr := LabyFileToList(laby)
	GenerateRandomImage(labyArr, source, output)
	fmt.Println("Success")
}
