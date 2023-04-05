package utils

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func GenerateRandomLaby(width int, height int, mode string) [][]int {
	arr1 := make([]int, width)
	arr2 := make([]int, height)
	for i := 0; i < width; i++ {
		arr1[i] = i
	}
	for i := 0; i < height; i++ {
		arr2[i] = i
	}
	if mode == "r" || mode == "vr" { // 随机排列
		rand.Shuffle(len(arr1), func(i, j int) {
			arr1[i], arr1[j] = arr1[j], arr1[i]
		})
	}
	if mode == "r" || mode == "hr" { // 随机排行
		rand.Shuffle(len(arr2), func(i, j int) {
			arr2[i], arr2[j] = arr2[j], arr2[i]
		})
	}
	return [][]int{arr1, arr2}
}

func LabyToFile(s string, name string) bool {
	err := ioutil.WriteFile(name, []byte(s), 0644)
	if err != nil {
		return false
	}
	return true
}

func LabyToStr(arr [][]int) string {
	var sb strings.Builder
	for i := 0; i < len(arr[0]); i++ {
		sb.WriteString(strconv.Itoa(arr[0][i]))
		if i != len(arr[0])-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString("|")
		}
	}
	for i := 0; i < len(arr[1]); i++ {
		sb.WriteString(strconv.Itoa(arr[1][i]))
		if i != len(arr[1])-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

func LabyFileToList(name string) [][]int {
	arr := [][]int{{}, {}}
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("[WARN] 无对应的列表文件")
		os.Exit(1)
	}
	str := string(data)
	s1 := strings.Split(str, "|")
	arr1 := strings.Split(s1[0], ",")
	arr2 := strings.Split(s1[1], ",")
	for _, v := range arr1 {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("[WARN] 无法将字符串转换为整数")
			os.Exit(1)
		}
		arr[0] = append(arr[0], n)
	}
	for _, v := range arr2 {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("[WARN] 无法将字符串转换为整数")
			os.Exit(1)
		}
		arr[1] = append(arr[1], n)
	}
	return arr
}
