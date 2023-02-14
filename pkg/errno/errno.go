
package errno

import (
	"errors"
	"fmt"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/message"
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(int64(user.ErrCode_SuccessCode), "Success")
	ServiceErr             = NewErrNo(int64(user.ErrCode_ServiceErrCode), "Service is unable to start successfully")
	ParamErr               = NewErrNo(int64(user.ErrCode_ParamErrCode), "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(int64(user.ErrCode_UserAlreadyExistErrCode), "User already exists")
	AuthorizationFailedErr = NewErrNo(int64(user.ErrCode_AuthorizationFailedErrCode), "Authorization failed")
	MessageIsNullExistErr  = NewErrNo(int64(message.ErrCode_MessageIsNullErrCode ), "Message is Null")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
