package aliyunfcgoruntime

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"
)

type Handler struct {
	Initialize func(ctx *FCContext) error
	Invoke     func(ctx *FCContext, event []byte) ([]byte, error)
	HttpInvoke func(ctx *FCContext, w http.ResponseWriter) error
	Prefreeze  func(ctx *FCContext) error
	Prestop    func(ctx *FCContext) error
}

var handler *Handler

func initializeHandler(w http.ResponseWriter, req *http.Request) {
	fcCtx := NewFromContext(req)
	fmt.Println(fmt.Sprintf(fcInitializeLogStartPrefix, fcCtx.RequestID))
	defer func() {
		if r := recover(); r != nil {
			handleRec(r, w)
		}
		fmt.Println(fmt.Sprintf(fcInitializeLogEndPrefix, fcCtx.RequestID))
	}()

	if handler.Initialize == nil {
		panic("This function doesn't have initialize handler.")
	}

	err := handler.Initialize(fcCtx)
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
			handleRec(r, w)
		}
		fmt.Println(fmt.Sprintf(fcInvokeLogEndPrefix, fcCtx.RequestID))
	}()

	event, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	if handler.Invoke == nil {
		panic("This function doesn't have invoke handler.")
	}

	resp, err := handler.Invoke(fcCtx, event)
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
			handleRec(r, w)
		}
		fmt.Println(fmt.Sprintf(fcInvokeLogEndPrefix, fcCtx.RequestID))
	}()

	if handler.HttpInvoke == nil {
		panic("This function doesn't have http invoke handler.")
	}

	err := handler.HttpInvoke(fcCtx, w)
	if err != nil {
		panic(err)
	}
}

func prefreezeHandler(w http.ResponseWriter, req *http.Request) {
	fcCtx := NewFromContext(req)
	fmt.Println(fmt.Sprintf(fcPreFreezeLogStartPrefix, fcCtx.RequestID))
	defer func() {
		if r := recover(); r != nil {
			handleRec(r, w)
		}
		fmt.Println(fmt.Sprintf(fcPreFreezeLogEndPrefix, fcCtx.RequestID))
	}()

	if handler.Prefreeze == nil {
		panic("This function doesn't have prefreeze handler.")
	}

	err := handler.Prefreeze(fcCtx)
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
			handleRec(r, w)
		}
		fmt.Println(fmt.Sprintf(fcPreStopEndPrefix, fcCtx.RequestID))
	}()

	if handler.Prestop == nil {
		panic("This function doesn't have prestop handler.")
	}

	err := handler.Prestop(fcCtx)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(""))
}

func handle(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			handleRec(r, w)
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

func handleRec(r interface{}, w http.ResponseWriter) {
	errorInfo := fmt.Sprintf("Error: %+v\nStack: %s", r, string(debug.Stack()))
	fmt.Println(errorInfo)
	w.Header().Set(fcStatus, "404")
	w.WriteHeader(404)
	w.Write([]byte(errorInfo))
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
