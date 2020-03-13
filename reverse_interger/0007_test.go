package reverse_interger

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_reverse(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 123,
			},
			a: ans{
				one: 321,
			},
		},
		question{
			p: para{
				one: -123,
			},
			a: ans{
				one: -321,
			},
		},
		question{
			p: para{
				one: 1234567899,
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: -1234567899,
			},
			a: ans{
				one: 0,
			},
		},
	}
	for _, q := range qs {
		a, p := q.a, q.p
		fmt.Println(reverse(p.one))
		ast.Equal(a.one, reverse(p.one), "输入:%v", p)
	}
}
