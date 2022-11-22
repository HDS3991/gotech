package builder

/*
	参考：https://refactoringguru.cn/design-patterns/go
	生成器模式可以帮助你分步骤创建复杂对象，使得用相同的创建过程生成不同的产品成为可能。
*/

type IBuilder interface {
	methodProperty(str string) *Builder
	methodCausality(str string) *Builder
	methodAttribute(str string)	 *Builder
}

type Builder struct {
	property string
	causality string
	attribute string
	err error
}

type Res struct {
	Property string `json:"property"`
	Causality string `json:"causality"`
	Attribute string `json:"attribute"`
}

func NewBuilder() IBuilder {
	return &Builder{}
}

func (b *Builder) methodProperty(str string) *Builder{
	if b.err != nil {
		return b
	}
	b.property = str
	return b
}

func (b *Builder) methodCausality(str string) *Builder {
	if b.err != nil {
		return b
	}
	b.causality = str
	return b
}

func (b *Builder) methodAttribute(str string) *Builder {
	if b.err != nil {
		return b
	}
	b.attribute = str
	return b
}

func  (b *Builder) Build() (*Res, error) {
	if b.err != nil {
		return nil, b.err
	}
	return &Res{
		Property: b.property,
		Causality: b.causality,
		Attribute: b.attribute,
	},nil
}