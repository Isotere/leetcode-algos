package unique_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MySolution(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
		k      int
	}{
		{
			input:  []int{},
			output: []int{},
			k:      0,
		},
		{
			input:  []int{1},
			output: []int{1},
			k:      1,
		},
		{
			input:  []int{1, 2},
			output: []int{1, 2},
			k:      2,
		},
		{
			input:  []int{1, 1, 2},
			output: []int{1, 2},
			k:      2,
		},
		{
			input:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			output: []int{0, 1, 2, 3, 4},
			k:      5,
		},
	}

	for _, current := range cases {
		k := removeDuplicates(current.input)

		assert.Equal(t, current.k, k)

		for i := 0; i < k; i++ {
			assert.Equal(t, current.output[i], current.input[i])
		}
	}
}
