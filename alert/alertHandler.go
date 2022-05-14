package alert

import "fmt"

type AlertHandler interface {
	Check(apiStatInfo ApiStatInfo) error
}

type ErrorHandler struct {}

func (e *ErrorHandler) Check(api ApiStatInfo)  error {
	if api.ErrorCount > 10 {
		return fmt.Errorf("error count is higher than 10")
	}
	return nil
}

