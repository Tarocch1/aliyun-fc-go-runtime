# Aliyun fc go runtime

阿里云函数计算 golang custom runtime 框架。

## Install

```sh
go get github.com/Tarocch1/aliyun-fc-go-runtime
```

## Event Function

```go
package main

import (
	"fmt"

	gr "github.com/Tarocch1/aliyun-fc-go-runtime"
)

func initialize(ctx *gr.FCContext) error {
	fmt.Println("init golang!")
	return nil
}

func invoke(ctx *gr.FCContext, event []byte) ([]byte, error) {
	fmt.Println(fmt.Sprintf("hello golang!\ncontext = %+v", ctx))
	return event, nil
}

func main() {
	handler := &gr.Handler{
		Initialize: initialize,
		Invoke:     invoke,
	}
	gr.Start(handler)
}
```

## HTTP Function

```go
package main

import (
	"fmt"
	"net/http"

	gr "github.com/Tarocch1/aliyun-fc-go-runtime"
)

func initialize(ctx *gr.FCContext) error {
	fmt.Println("init golang!")
	return nil
}

func httpInvoke(ctx *gr.FCContext, w http.ResponseWriter) error {
	w.Write([]byte(fmt.Sprintf("hello golang!\ncontext = %+v", ctx)))
	return nil
}

func main() {
	handler := &gr.Handler{
		Initialize: initialize,
		HttpInvoke: httpInvoke,
	}
	gr.Start(handler)
}
```
