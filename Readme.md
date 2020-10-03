# Maven-Clean

清空maven local repository 目录下的脏文件



## 设计思路

使用`mvn help:evaluate -Dexpression=settings.localRepository` 命令获取maven repository的地址

然后`递归`清理目录下的文件 *.lastUpdated 文件 和 _remote.repositories 文件


## build windows

```
CGO_ENABLED=0;GOOS=windows;GOARCH=amd64
```


## Todo 

现在扫描文件是单线程的，后续改成多线程 ...

## Use

link [release](https://github.com/BoomManPro/maven-clean/releases/tag/1.0)

download release and run