package platesbetweencandles

import (
	"reflect"
	"testing"
)

type input struct {
	s       string
	queries [][]int
	want    []int
}

func TestPaltesBetweenCandles(t *testing.T) {

	t.Run("example test", func(t *testing.T) {

		inputs := []input{
			{s: "**|**|***|", queries: [][]int{{2, 5}, {5, 9}}, want: []int{2, 3}},
			{s: "***|**|*****|**||**|*", queries: [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}, want: []int{9, 0, 0, 0, 0}},
		}

		for _, test := range inputs {
			got := PlatesBetweenCandles(test.s, test.queries)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, but want %v", got, test.want)

			}
		}
	})
}
