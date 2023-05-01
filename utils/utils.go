package utils

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// 随机字符集
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateRandomName 生成随机名称函数
func GenerateRandomName(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetNextMultiple(length int) int {
	a := 1
	for i := 2; i <= length; i++ {
		if length%i == 0 {
			a = i
			break
		}
	}
	return a
}

func ConvertTo2DArray(arr []byte, rows int) [][]byte {
	cols := len(arr) / rows
	result := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]byte, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = arr[i*cols+j]
		}
	}
	return result
}

func RunCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败")
		return err
	}
	cmd.Dir = dir

	if s, err := cmd.Output(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Println("提取失败")
			fmt.Println(s)
			fmt.Printf("Error: exit status %d\n", exitError.ExitCode())
			fmt.Println(string(exitError.Stderr))
			return err
		}
	}
	return nil
}

func Generate(width int, height int, mode string, name string) {
	arr := GenerateRandomLaby(width, height, mode)
	LabyToFile(LabyToStr(arr), name)
	labyArr := LabyFileToList(name)
	if fmt.Sprintf("%v", arr) == fmt.Sprintf("%v", labyArr) {
		fmt.Println("Success")
	} else {
		fmt.Println("Failed")
	}
}
