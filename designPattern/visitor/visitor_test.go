package visitor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type NameVisitor struct {
	visitor Visitor
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}

		err = fn(info, err)
		if err != nil {
			info.ErrList = append(info.ErrList, err)
		}
		return nil
	})
}

type DescVisitor struct {
	visitor Visitor
}

func (v DescVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}

		err = fn(info, err)
		if err != nil {
			info.ErrList = append(info.ErrList, err)
		}
		return nil
	})
}

func TestVisitor(t *testing.T) {
	info := Info{}
	var v Visitor = &info
	v = NameVisitor{v}
	v = DescVisitor{v}

	// 从文件中读取数据
	loadContent := func(info *Info, err error) error {
		info.Name = "Test"
		info.Desc = "some desc"
		return nil
	}
	// 同时激活 NameVisitor 和 DescVisitor 的  Visit 方法
	r := v.Visit(loadContent)
	assert.Equal(t, nil, r)

}
