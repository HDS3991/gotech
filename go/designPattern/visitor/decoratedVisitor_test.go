package visitor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecoratedVisitor(t *testing.T) {
	info := Info{}
	var v Visitor = &info

	fn1 := func(info *Info, err error) error {
		return nil
	}
	fn2 := func(info *Info, err error) error {
		return nil
	}
	v = NewDecoratedVisitor(v, fn1, fn2)

	// 从文件中读取数据
	loadContent := func(info *Info, err error) error {
		info.Name = "Test"
		info.Desc = "some desc"
		return nil
	}

	r := v.Visit(loadContent)
	assert.Equal(t, nil, r)
}
