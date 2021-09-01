package aliyunfcgoruntime

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"
)

type Handler struct {
	initialize func(ctx *FCContext) error
	invoke     func(ctx *FCContext, event []byte) ([]byte, error)
	httpInvoke func(ctx *FCContext, w http.ResponseWriter)
	prefreeze  func(ctx *FCContext) error
	prestop    func(ctx *FCContext) error
}

var handler *Handler

func initializeHandler(w http.ResponseWriter, req *http.Request) {
	fcCtx := NewFromContext(req)
	fmt.Println(fmt.Sprintf(fcInitializeLogStartPrefix, fcCtx.RequestID))
	defer func() {
		if r := recover(); r != nil {
			w.Header().Set(fcStatus, "404")
			w.WriteHeader(404)
			w.Write([]byte(fmt.Sprintf("Error: %+v\nStack: %s", r, string(debug.Stack()))))
		}
		fmt.Println(fmt.Sprintf(fcInitializeLogEndPrefix, fcCtx.RequestID))
	}()

	if handler.initialize == nil {
		panic("This function doesn't have initialize handler.")
	}

	err := handler.initialize(fcCtx)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(""))
}

func invokeHandler(w http.ResponseWriter, req *http.Request) {
	fcCtx := NewFromContext(req)
	fmt.Println(fmt.Sprintf(fcInvokeLogStartPrefix, fcCtx.RequestID))
	defer func() {
		if r := recover(); r != nil {
			w.Header().Set(fcStatus, "404")
			w.WriteHeader(404)
			w.Write([]byte(fmt.Sprintf("Error: %+v\nStack: %s", r, string(debug.Stack()))))
		}
		fmt.Println(fmt.Sprintf(fcInvokeLogEndPrefix, fcCtx.RequestID))
	}()

	event, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	if handler.invoke == nil {
		panic("This function doesn't have invoke handler.")
	}

	resp, err := handler.invoke(fcCtx, event)
	if err != nil {
		panic(err)
	}

	w.Write(resp)
}

func httpInvokeHandler(w http.ResponseWriter, req *http.Request) {
	fcCtx := NewFromContext(req)
	fmt.Println(fmt.Sprintf(fcInvokeLogStartPrefix, fcCtx.RequestID))
	defer func() {
		if r := recover(); r != nil {
			w.Header().Set(fcStatus, "404")
			w.WriteHeader(404)
			w.Write([]byte(fmt.Sprintf("Error: %+v\nStack: %s", r, string(debug.Stack()))))
		}
		fmt.Println(fmt.Sprintf(fcInvokeLogEndPrefix, fcCtx.RequestID))
	}()

	if handler.httpInvoke == nil {
		panic("This function doesn't have http invoke handler.")
	}

	handler.httpInvoke(fcCtx, w)
}

func prefreezeHandler(w http.ResponseWriter, req *http.Request) {
	fcCtx := NewFromContext(req)
	fmt.Println(fmt.Sprintf(fcPreFreezeLogStartPrefix, fcCtx.RequestID))
	defer func() {
		if r := recover(); r != nil {
			w.Header().Set(fcStatus, "404")
			w.WriteHeader(404)
			w.Write([]byte(fmt.Sprintf("Error: %+v\nStack: %s", r, string(debug.Stack()))))
		}
		fmt.Println(fmt.Sprintf(fcPreFreezeLogEndPrefix, fcCtx.RequestID))
	}()

	if handler.prefreeze == nil {
		panic("This function doesn't have prefreeze handler.")
	}

	err := handler.prefreeze(fcCtx)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(""))
}

func prestopHandler(w http.ResponseWriter, req *http.Request) {
	fcCtx := NewFromContext(req)
	fmt.Println(fmt.Sprintf(fcPreStopStartPrefix, fcCtx.RequestID))
	defer func() {
		if r := recover(); r != nil {
			w.Header().Set(fcStatus, "404")
			w.WriteHeader(404)
			w.Write([]byte(fmt.Sprintf("Error: %+v\nStack: %s", r, string(debug.Stack()))))
		}
		fmt.Println(fmt.Sprintf(fcPreStopEndPrefix, fcCtx.RequestID))
	}()

	if handler.prestop == nil {
		panic("This function doesn't have prestop handler.")
	}

	err := handler.prestop(fcCtx)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(""))
}

func handle(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			w.Header().Set(fcStatus, "404")
			w.WriteHeader(404)
			w.Write([]byte(fmt.Sprintf("Error: %+v\nStack: %s", r, string(debug.Stack()))))
		}
	}()
	controlPath := req.Header.Get(fcControlPath)
	switch controlPath {
	case "/initialize":
		initializeHandler(w, req)
	case "/invoke":
		invokeHandler(w, req)
	case "/http-invoke":
		httpInvokeHandler(w, req)
	case "/pre-freeze":
		prefreezeHandler(w, req)
	case "/pre-srop":
		prestopHandler(w, req)
	default:
		panic("Unknown control path.")
	}
}

func Start(h *Handler) {
	handler = h
	fmt.Println("FunctionCompute go runtime inited.")
	http.HandleFunc("/", handle)
	port := os.Getenv("FC_SERVER_PORT")
	if port == "" {
		port = "9000"
	}
	http.ListenAndServe(":"+port, nil)
}
