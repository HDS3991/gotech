package visitor

/*
访问者模式是 Kubernetes 的 kubectl 命令中的使用到到的一个编程模式
适用于多个 Visitor 是来访问一个数据结构的不同部分
优点：
1.解耦了数据和逻辑；
2.用 pipeline 使逻辑清晰易读
*/

/*
	参考：https://coolshell.cn/articles/21263.html
*/

type Info struct {
	Name    string
	Desc    string
	ErrList []error
}

type VisitorFunc func(*Info, error) error

type Visitor interface {
	Visit(VisitorFunc) error
}

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}
