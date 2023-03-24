package capture

import "fmt"

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 17:26
**/

type WareError struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func (w *WareError) Error() string {
	return fmt.Sprintf(`{"code":%d,"msg":"%s"}`, w.Code, w.Msg)
}

func Handle(err error) *WareError {
	if err == nil {
		return &WareError{
			Code: 0,
			Msg:  "请求成功",
		}
	}

	wareErr, ok := err.(*WareError)
	if ok {
		return wareErr
	}

	return &WareError{Code: -1, Msg: err.Error()}
}
