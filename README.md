# vulfocus-go

[![GitHub (pre-)release](https://img.shields.io/github/release/fofapro/vulfocus-go/all.svg)](https://github.com/fofapro/vulfocus-go/releases) [![stars](https://img.shields.io/github/stars/fofapro/vulfocus-go.svg)](https://github.com/fofapro/vulfocus-go/stargazers) [![license](https://img.shields.io/github/license/fofapro/vulfocus-go.svg)](https://github.com/fofapro/vulfocus-go/blob/master/LICENSE)

[中文文档](https://github.com/fofapro/vulfocus-go/blob/master/README_zh.md)

## Vulfocus API

[`Vulfocus API`](https://fofapro.github.io/vulfocus/#/VULFOCUSAPI) is the `RESUFul API` interface provided by [`Vulfocus`](http://vulfocus.io/) for development, allowing Developers integrate [`Vulfocus`](http://vulfocus.io) in their own projects.

## Vulfocus SDK

The `GO` version of `SDK` written based on the [`Vulfocus API`](https://fofapro.github.io/vulfocus/#/VULFOCUSAPI) makes it easy for `Golang` developers to quickly integrate [`Vulfocus`](http://vulfocus.io/)  into their projects.


## Add dependency

```bash
go get github.com/fofapro/vulfocus-go
```

## Use
|field|description|
| ---- | ---- |
|`addr`|[`Vulfocus`](http://vulfocus.io/) URL|
|`username`|User login [`Vulfocus`](http://vulfocus.io/) userbox `username`|
|`licence`|Please go to the [`personal center`](http://vulfocus.fofa.so/#/profile/index) to view `API licence`|

### Pull Images

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
	err, images := client.GetImages()
	if err != nil {
		return
	}
  fmt.Printf("get %d images", len(images))
	if len(images) == 0 {
		return
	}
}
```

### Start

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
	err, images := client.GetImages()
	if err != nil {
		return
	}
  fmt.Printf("get %d images", len(images))
	if len(images) == 0 {
		return
	}
	err, exposed := client.Start(images[0].Name)
	if err != nil {
		return
	}
	println(exposed.Host, exposed.Port)
}
```

### Stop

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
	err, images := client.GetImages()
	if err != nil {
		return
	}
  fmt.Printf("get %d images", len(images))
	if len(images) == 0 {
		return
	}
	err = client.Stop(images[0].Name)
	if err != nil {
		return
	}
}
```

### Delete

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
	err, images := client.GetImages()
	if err != nil {
		return
	}
	fmt.Printf("get %d images", len(images))
	if len(images) == 0 {
		return
	}
	err = client.Delete(images[0].Name)
	if err != nil {
		return
	}
}
```

## Update Log

2021-11-25

```
- Version release
```
