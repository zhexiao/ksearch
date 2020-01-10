# ksearch
按页数对网站进行关键字搜索，避免在有些网站搜索数据的时候需要登录或相应的权限。

[![Build Status](https://travis-ci.org/zhexiao/ksearch.svg?branch=master)](https://travis-ci.org/zhexiao/ksearch)
[![codecov](https://codecov.io/gh/zhexiao/ksearch/branch/master/graph/badge.svg)](https://codecov.io/gh/zhexiao/ksearch)
![go](https://img.shields.io/badge/go->%3D1.13-blue)

## 使用方式

#### 第一步：下载windows程序
ksearch.exe(https://github.com/zhexiao/ksearch/releases/download/v2/ksearch.exe)

#### 第二步：双击程序，打开程序
![image](https://raw.githubusercontent.com/zhexiao/ksearch/master/_example/3.png)

#### 第三步：浏览器输入地址 localhost:18888
![image](https://raw.githubusercontent.com/zhexiao/ksearch/master/_example/4.png)

#### 第四步：填入搜索数据
![image](https://raw.githubusercontent.com/zhexiao/ksearch/master/_example/1.png)

#### 第五步：开始搜索
![image](https://raw.githubusercontent.com/zhexiao/ksearch/master/_example/2.png)

# 源码编译
#### 需求
1. Go >= 1.13

#### 开启go mod

```
根据实际情况填写路径
$ go env -w GOPATH=/goproj
$ go env -w GOBIN=/goproj/bin
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct 
```

#### 拉库
```
$ go mod download
```

#### 运行
```
$ go run main.go
```

