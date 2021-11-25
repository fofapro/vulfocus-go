# vulfocus-go

此仓库是开源漏洞集成平台 [vulfocus](https://github.com/fofapro/vulfocus) 的 Golang SDK。

## 安装

```bash
go get github.com/fofapro/vulfocus-go
```

## 使用方式

使用 `接口地址` 、`用户账号` 、`licence` 创建一个客户端，即可调用相关函数。

```go
package main

import (
	"fmt"
	vulfocus "github.com/fofapro/vulfocus-go"
)

const (
	addr     = "http://vulfocus.fofa.so"
	username = ""
	licence  = ""
)

func main() {
	client := vulfocus.NewClient(addr, username, licence)
	// 获取镜像
	err, images := client.GetImages()
	if err != nil {
		return
	}
	fmt.Printf("get %d images", len(images))
	if len(images) == 0 {
		return
	}

	// 创建
	err, exposed := client.Start(images[0].Name)
	if err != nil {
		return
	}
	println(exposed.Host, exposed.Port)

	// 停止
	err = client.Stop(images[0].Name)
	if err != nil {
		return
	}

	// 删除
	err = client.Delete(images[0].Name)
	if err != nil {
		return
	}
}
```