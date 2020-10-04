package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	defer timeCost()()
	path, err := getMvnLocalRepositoryPath()
	if err != nil {
		fmt.Printf("发生异常:%v", err)
	}
	fmt.Printf("获取到maven Local Repository Path -> : %s\n", path)

	fmt.Printf("清理目录下的文件 *.lastUpdated 文件 和 _remote.repositories 文件\n")

	fileList, err := getAllLastUpdateFile(path)
	if err != nil {
		fmt.Printf("发生异常:%v", err)
	}

	for i := range fileList {
		fmt.Printf("删除文件: %s\n", fileList[i])
		go os.Remove(fileList[i])
	}

	fmt.Printf("\n\n推荐依赖仓库地址: ->%s", "<url>https://maven.aliyun.com/repository/public</url>")

}

// 获取所有异常文件  *.lastUpdated 文件 和 _remote.repositories
func getAllLastUpdateFile(dir string) ([]string, error) {
	var result []string
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("%v", err)
			}
			if !info.IsDir() && (strings.Contains(path, "lastUpdated") || strings.Contains(path, "_remote.repositories")) {
				result = append(result, path)
			}
			return nil
		})
	return result, err
}

func getMvnLocalRepositoryPath() (string, error) {
	//根据命令 mvn
	//获取本地仓库地址
	//清除无效文件
	fmt.Printf("mvn help:evaluate -Dexpression=settings.localRepository | grep -v '\\[INFO\\]'\n")
	cmd := exec.Command("mvn", "help:evaluate", "-Dexpression=settings.localRepository")

	stdout, err := cmd.StdoutPipe()

	//获取输出对象，可以从该对象中读取输出结果
	if err != nil {
		return "", err
	}
	// 保证关闭输出流
	defer stdout.Close()

	// 运行命令
	if err := cmd.Start(); err != nil {
		return "", err
	}
	// 读取输出结果
	if opBytes, err := ioutil.ReadAll(stdout); err != nil {
		return "", err
	} else {
		return parserLocalRepositoryPath(string(opBytes))
	}

}

func parserLocalRepositoryPath(content string) (string, error) {
	lineList := strings.Split(content, "\n")
	for i := range lineList {
		if strings.Index(lineList[i], "[INFO]") == -1 {
			result := strings.TrimRight(lineList[i], "\r")
			return result, nil
		}
	}
	return "", errors.New(fmt.Sprintf("没有找到maven Local Repository maven command result: \n%s", content))
}

//@brief：耗时统计函数
func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("\ntime cost = %v\n", tc)
	}
}
