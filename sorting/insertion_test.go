package sorting_test

import (
	"cmp"
	"reflect"
	"testing"

	"github.com/luisnogit/algosingo.git/sorting"
)

type Test_table[T cmp.Ordered] struct {
	name   string
	input  []T
	output []T
}

func TestSortEmpty(t *testing.T) {
	test := []interface{}{
		Test_table[int]{
			name:   "Empty test case",
			input:  []int{},
			output: []int{},
		},
		Test_table[int]{
			name:   "worst case",
			input:  []int{5, 4, 3, 2, 1},
			output: []int{1, 2, 3, 4, 5},
		},
		Test_table[string]{
			name:   "string case",
			input:  []string{"banana", "apple", "cherry", "date"},
			output: []string{"apple", "banana", "cherry", "date"},
		},
	}
	for _, v := range test {
		switch typetest := v.(type) {
		case Test_table[int]:
			typetest.input = sorting.Insertion(typetest.input)
			if reflect.DeepEqual(typetest.input, typetest.output) == false {
				t.Errorf("Test Faild: %+v", v)
			}
		case Test_table[float64]:
			typetest.input = sorting.Insertion(typetest.input)
			if reflect.DeepEqual(typetest.input, typetest.output) == false {
				t.Errorf("Test Faild: %+v", v)
			}

		case Test_table[string]:
			typetest.input = sorting.Insertion(typetest.input)
			if reflect.DeepEqual(typetest.input, typetest.output) == false {
				t.Errorf("Test Faild: %+v", v)
			}

		}
	}
}
