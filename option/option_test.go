package option_test

import (
	"testing"

	"github.com/jcp19/prelude/option"
	"github.com/stretchr/testify/assert"
)

func unreachable(t *testing.T) {
	assert.FailNow(t, "A suposedly unreachable location was reached.")
}

type testMapCase struct {
	input  option.Option[int]
	f      func(int) int
	output option.Option[int]
}

var testMapCases map[string]testMapCase = map[string]testMapCase{
	"none maps to none": testMapCase{
		input: option.None[int](),
		f: func(x int) int {
			return x + 1
		},
		output: option.None[int](),
	},
	"map f of some x maps to some f of x": testMapCase{
		input: option.Some(10),
		f: func(x int) int {
			return x + 1
		},
		output: option.Some(11),
	},
}

func TestMap(t *testing.T) {
	for k, v := range testMapCases {
		assert.Equal(t, v.output, option.Map(v.input, v.f), "Failed in test \"%s\".", k)
	}
}

type testCaseCase struct {
	input      option.Option[int]
	branchNone func(*testing.T)
	branchSome func(*testing.T, int)
}

var testCaseCases map[string]testCaseCase = map[string]testCaseCase{
	"case none": testCaseCase{
		input:      option.None[int](),
		branchNone: func(*testing.T) {},
		branchSome: func(t *testing.T, _ int) { unreachable(t) },
	},
	"case some": testCaseCase{
		input:      option.Some(10),
		branchNone: func(t *testing.T) { unreachable(t) },
		branchSome: func(*testing.T, int) {},
	},
}

func TestCase(t *testing.T) {
	for _, v := range testCaseCases {
		option.Case(
			v.input,
			func() { v.branchNone(t) },
			func(x int) { v.branchSome(t, x) },
		)
	}
}
