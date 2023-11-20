package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

var path string
var width1, width2, width3, height1, height2, height3 int

func main() {
	fmt.Println("文件夹路径为: ")
	fmt.Scanln(&path)
	fmt.Println("当图片宽度大于高度时，请设置宽度、高度: ")
	fmt.Scan(&width1)
	fmt.Scan(&height1)
	fmt.Println("当图片宽度小于高度时，请设置宽度、高度: ")
	fmt.Scan(&width2)
	fmt.Scan(&height2)
	fmt.Println("当图片宽度等于高度时，请设置宽度、高度: ")
	fmt.Scan(&width3)
	fmt.Scan(&height3)
	dir, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dir.Close()

	// 读取文件夹中的所有文件
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 遍历文件并打印它们的名称
	for _, file := range files {
		// builder 拼接字符串
		var builder strings.Builder
		builder.WriteString(path)
		builder.WriteString("/")
		builder.WriteString(file.Name())

		alter(builder.String(), width1, height1, width2, height2)
	}
}

func alter(pic string, width1, height1, width2, height2 int) {
	// 打开原始图片文件
	file, err := os.Open(pic)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 解码图片
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// 获取原始图片的尺寸
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	// 根据原始尺寸和设定的大小，强制改变图片大小
	var resizedImg image.Image
	if width > height {
		resizedImg = resize.Resize(uint(width1), uint(height1), img, resize.Lanczos3)
	} else if width < height {
		resizedImg = resize.Resize(uint(width2), uint(height2), img, resize.Lanczos3)
	} else {
		resizedImg = resize.Resize(uint(width3), uint(height3), img, resize.Lanczos3)
	}

	// 创建新图片文件
	newFile, err := os.Create(pic)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// 保存新图片
	err = jpeg.Encode(newFile, resizedImg, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("图片大小已改变并保存为", pic)
}
