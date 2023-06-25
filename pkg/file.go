package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	
	//"time"
)
func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

func WriteToFile(filepath,msg string) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		//fmt.Println("文件打开失败", err)
		return err
	}
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	m := fmt.Sprintf("%s%s", msg, "\n")
	write.WriteString(m)
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	
	return nil
}

func CopyFile(src,dst string) error  {
	if !CheckFileExist(src) {
		fmt.Printf("源文件：%v不存在",src)
		return errors.New(fmt.Sprintf("源文件：%v不存在",src))
	}
	if CheckFileExist(dst) {
		fmt.Printf("目标文件：%v已存在",dst)
		return errors.New(fmt.Sprintf("目标文件：%v已存在",dst))
	}
	source, err := os.Open(src)
	if err != nil {
		fmt.Printf("打开源文件：%v失败",src)
		return  err
	}
	defer source.Close()
	destination, err := os.Create(dst)
	if err != nil {
		fmt.Printf("创建目标文件：%v失败",dst)
		return  err
	}
	
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		fmt.Printf("拷贝文件失败")
		return  err
	}
	return nil
}


func CopyDir(srcPath string, destPath string) error {
	//检测目录正确性
	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	
	if !srcInfo.IsDir() {
		e := errors.New("srcPath不是一个正确的目录！")
			fmt.Println(e.Error())
			return e
		}
	destInfo, err := os.Stat(destPath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
		if !destInfo.IsDir() {
			e := errors.New("destInfo不是一个正确的目录！")
			fmt.Println(e.Error())
			return e
		}
		
	err = filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			path := strings.Replace(path, "\\", "/", -1)
			destNewPath := strings.Replace(path, srcPath, destPath, -1)
			fmt.Println("复制文件:" + path + " 到 " + destNewPath)
			copyFile(path, destNewPath)
		}
		return nil
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}

func copyFile(src, dest string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()
	//分割path目录
	destSplitPathDirs := strings.Split(dest, "/")
	//检测时候存在目录
	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + "/"
			b, _ := pathExists(destSplitPath)
			if b == false {
				fmt.Println("创建目录:" + destSplitPath)
				//创建目录
				err := os.Mkdir(destSplitPath, os.ModePerm)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	dstFile, err := os.Create(dest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}
//检测文件夹路径时候存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


func GetAllFilesByDir(dir string) ([]string,error) {
	var files []string
	err := filepath.Walk(dir, func (path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Printf("获取目录%v下的文件失败",dir)
	}
	
	return files,nil
}
