package main

import "fmt"



type AlertHandler interface {
	check(apiRequestInfo ApiRequestInfo) error
}



type ErrorHandler struct {

}

func(e *ErrorHandler) check(apiRequestInfo ApiRequestInfo) error {
	fmt.Println("错误检查器")
	return nil
}

type CountHandler struct {
}

func(e *CountHandler) check(apiRequestInfo ApiRequestInfo) error {
	fmt.Println("计检查器")
	return nil
}

type Alert struct {
	AlertHandlers []AlertHandler
}

func NewAlert() *Alert {
	return &Alert{
		AlertHandlers: make([]AlertHandler, 0),
	}
}

func (a *Alert) AddHandler(handlers ...AlertHandler) {
	a.AlertHandlers = append(a.AlertHandlers, handlers...)
}

func (a *Alert) DoCheck(apiRequestInfo ApiRequestInfo) error {
	for _, v := range a.AlertHandlers {
		err := v.check(apiRequestInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

type ApiRequestInfo struct {
	Api string
	RequestCount int
	ErrorCount int
}

func main() {
	a := NewAlert()
	a.AddHandler(new(CountHandler), new(ErrorHandler))
	a.DoCheck(ApiRequestInfo{})
}