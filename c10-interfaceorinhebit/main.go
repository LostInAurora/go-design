package main
//	接口
type I interface {

}
//	组合
type A string

const (
	b A = "dfsdf"

)
//	委托

func f(a A) {

}
func main() {
	var a A = "dfs"
	f(a)
	f("fsf")
}