package visitor

// k8s 中 kubectl 的使用方式

type DecoratedVisitor struct {
	visitor    Visitor
	decorators []VisitorFunc
}

func NewDecoratedVisitor(v Visitor, fn ...VisitorFunc) Visitor {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{
		visitor:    v,
		decorators: fn,
	}
}

func (d DecoratedVisitor) Visit(fn VisitorFunc) error {
	return d.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}

		if err := fn(info, nil); err != nil {
			return err
		}
		for i := range d.decorators {
			if err := d.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}
