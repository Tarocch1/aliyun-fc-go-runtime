package aliyunfcgoruntime

const (
	fcRequestID       = "x-fc-request-id"
	fcAccessKeyID     = "x-fc-access-key-id"
	fcAccessKeySecret = "x-fc-access-key-secret"
	fcSecurityToken   = "x-fc-security-token"

	fcFunctionName          = "x-fc-function-name"
	fcFunctionHandler       = "x-fc-function-handler"
	fcFunctionMemory        = "x-fc-function-memory"
	fcFunctionTimeout       = "x-fc-function-timeout"
	fcFunctionInitializer   = "x-fc-function-initializer"
	fcInitializationTimeout = "x-fc-initialization-timeout"

	fcServiceName       = "x-fc-service-name"
	fcServiceLogProject = "x-fc-service-logproject"
	fcServiceLogstore   = "x-fc-service-logstore"

	fcRegion     = "x-fc-region"
	fcAccountID  = "x-fc-account-id"
	fcQualifier  = "x-fc-qualifier"
	fcVersionID  = "x-fc-version-id"
	fcRetryCount = "x-fc-retry-count"

	fcStatus      = "x-fc-status"
	fcControlPath = "x-fc-control-path"

	fcInitializeLogStartPrefix = "FC Initialize Start RequestId: %s"
	fcInitializeLogEndPrefix   = "FC Initialize End RequestId: %s"
	fcInvokeLogStartPrefix     = "FC Invoke Start RequestId: %s"
	fcInvokeLogEndPrefix       = "FC Invoke End RequestId: %s"
	fcPreFreezeLogStartPrefix  = "FC Pre Freeze Start RequestId: %s"
	fcPreFreezeLogEndPrefix    = "FC Pre Freeze End RequestId: %s"
	fcPreStopStartPrefix       = "FC Pre Stop Start RequestId: %s"
	fcPreStopEndPrefix         = "FC Pre Stop End RequestId: %s"
)
