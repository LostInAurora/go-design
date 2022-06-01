package main

import "fmt"

type ApiRequest struct {

}

type Filter interface {
	doFilter(apiRequest ApiRequest) error
}

type RateLimitFilter struct {
}

func (f *RateLimitFilter) doFilter(apiRequest ApiRequest) error {
	fmt.Println("开始进行限流检查")
	return nil
}

type FilterApplication struct {
	FilterList []Filter
}

func (fa *FilterApplication) AddFilter(filter Filter) {
	fa.FilterList = append(fa.FilterList, filter)
}
func (fa *FilterApplication) HandleRequest(apiRequest ApiRequest) error {
	for _, v := range fa.FilterList {
		err := v.doFilter(apiRequest)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	fa := new(FilterApplication)
	fa.AddFilter(new(RateLimitFilter))
	fa.HandleRequest(ApiRequest{})
}