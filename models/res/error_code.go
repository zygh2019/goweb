package res

type ErrorCode int

var (
	ErrorMap = map[ErrorCode]string{
		500: "系统错误",
	}
)
