package alert

// 开闭原则


type Alerter struct {
	Handlers []AlertHandler
}

func (a *Alerter) Check(apiStatInfo ApiStatInfo) error {
	for _, handler := range a.Handlers {
		err := handler.Check(apiStatInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Alerter) AddHandler(handler AlertHandler) {
	a.Handlers = append(a.Handlers, handler)
}






