package util

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Copy 拷贝
func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// WriteStringToFileMethod 通过 io.WriteString 写入文件
func WriteStringToFileMethod(fileName string, writeInfo string) {
	_ = IfNoFileToCreate(fileName)
	f, err := os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
	defer f.Close()
	if err != nil {
		log.Printf("打开文件失败:%+v", err)
		return
	}
	// 将文件写进去
	if _, err = io.WriteString(f, writeInfo); err != nil {
		log.Printf("WriteStringToFileMethod2 写入文件失败:%+v", err)
		return
	}
}

// IsExists 判断所给路径文件/文件夹是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil && !os.IsExist(err) {
		return false
	}
	return true
}

// IfNoFileToCreate 文件不存在就创建文件
func IfNoFileToCreate(fileName string) (file *os.File) {
	var f *os.File
	var err error
	if !IsExists(fileName) {
		f, err = os.Create(fileName)
		if err != nil {
			return
		}
		defer f.Close()
	}
	return f
}