package aliyunfcgoruntime

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Credentials struct {
	AccessKeyID     string
	AccessKeySecret string
	SecurityToken   string
}

type FunctionMeta struct {
	Name                  string
	Handler               string
	Memory                int
	Timeout               int
	Initializer           string
	InitializationTimeout int
}

type ServiceMeta struct {
	ServiceName string
	LogProject  string
	LogStore    string
	Qualifier   string
	VersionID   string
}

type FCContext struct {
	RequestID   string
	Credentials Credentials
	Function    FunctionMeta
	Service     ServiceMeta
	Region      string
	AccountID   string
	RetryCount  int
	Req         *http.Request
	Log         func(map[string]interface{})
}

func log(data map[string]interface{}) {
	j, _ := json.Marshal(data)
	if j != nil {
		fmt.Println(string(j))
	}
}

func NewFromContext(req *http.Request) *FCContext {
	mStr := req.Header.Get(fcFunctionMemory)
	m, err := strconv.Atoi(mStr)
	if err != nil {
		m = -1
	}
	tStr := req.Header.Get(fcFunctionTimeout)
	t, err := strconv.Atoi(tStr)
	if err != nil {
		t = -1
	}
	itStr := req.Header.Get(fcInitializationTimeout)
	it, err := strconv.Atoi(itStr)
	if err != nil {
		it = -1
	}
	retryStr := req.Header.Get(fcRetryCount)
	retryCount, err := strconv.Atoi(retryStr)
	if err != nil {
		retryCount = 0
	}
	ctx := &FCContext{
		RequestID: req.Header.Get(fcRequestID),
		Credentials: Credentials{
			AccessKeyID:     req.Header.Get(fcAccessKeyID),
			AccessKeySecret: req.Header.Get(fcAccessKeySecret),
			SecurityToken:   req.Header.Get(fcSecurityToken),
		},
		Function: FunctionMeta{
			Name:                  req.Header.Get(fcFunctionName),
			Handler:               req.Header.Get(fcFunctionHandler),
			Memory:                m,
			Timeout:               t,
			Initializer:           req.Header.Get(fcFunctionInitializer),
			InitializationTimeout: it,
		},
		Service: ServiceMeta{
			ServiceName: req.Header.Get(fcServiceName),
			LogProject:  req.Header.Get(fcServiceLogProject),
			LogStore:    req.Header.Get(fcServiceLogstore),
			Qualifier:   req.Header.Get(fcQualifier),
			VersionID:   req.Header.Get(fcVersionID),
		},
		Region:     req.Header.Get(fcRegion),
		AccountID:  req.Header.Get(fcAccountID),
		RetryCount: retryCount,
		Req:        req,
		Log:        log,
	}
	return ctx
}
