package problem1310

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	t.Run("readme-example", func(t *testing.T) {
		l := &ListElement{
			value: 5,
			next: &ListElement{
				value: 4,
				next: &ListElement{
					value: 3,
					next: &ListElement{
						value: 2,
						next: &ListElement{
							value: 1,
							next:  nil,
						},
					},
				},
			},
		}

		Solution(l)

		assert.Equal(t, &ListElement{
			value: 1,
			next: &ListElement{
				value: 3,
				next: &ListElement{
					value: 2,
					next: &ListElement{
						value: 5,
						next: &ListElement{
							value: 4,
							next:  nil,
						},
					},
				},
			},
		}, l)
	})
}
