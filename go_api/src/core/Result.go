package core

type resultEnum struct {
	code int

	msg string
}

type Result struct {
	Code int

	Msg string

	Data interface{}
}

const (
	SUCCESS            = "SUCCESS"
	ERROR              = "ERROR"
	USER_HAS_EXIST     = "USER_HAS_EXIST"
	USER_HAS_ACTIVATED = "USER_HAS_ACTIVATED"
)

var enumMap = map[string]resultEnum{
	SUCCESS:            {0, "成功"},
	ERROR:              {-1, "未知错误"},
	USER_HAS_EXIST:     {-1, "用户已存在"},
	USER_HAS_ACTIVATED: {-1, "用户已激活"},
}

func (r *Result) setEnum(enum resultEnum) {
	r.Code = enum.code
	r.Msg = enum.msg
}

func (r Result) SetData(data interface{}) Result {
	r.setEnum(enumMap[SUCCESS])
	r.Data = data
	return r
}

func (r Result) SetError(errorType string) Result {

	v, has := enumMap[errorType]

	if has {
		r.setEnum(v)
	} else {
		r.setEnum(enumMap[ERROR])
	}

	return r
}
