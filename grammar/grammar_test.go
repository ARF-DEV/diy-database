package grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateLRItems(t *testing.T) {
	cases := []struct {
		name     string
		expected []LRItem
		input    []Production
	}{
		{
			name: "empty production should return 1 LR item (with dot position = 0)",
			expected: []LRItem{
				{
					ProdictionIdx: 0,
					DotPosition:   0,
				},
			},
			input: []Production{
				{
					Empty,
					[]Symbol{},
				},
			},
		},
		{
			name: "return 4 lr items on SELECT IDENT and empty production",
			expected: []LRItem{
				{0, 0},
				{1, 0},
				{1, 1},
				{1, 2},
			},
			input: []Production{
				{
					Empty,
					[]Symbol{},
				},
				{
					StatementSelect,
					[]Symbol{
						KeywordSelect,
						Identifier,
					},
				},
			},
		},
		{
			name: "valid production with nested non-terminal rules",
			expected: []LRItem{
				{0, 0},
				{0, 1},
				{1, 0},
				{1, 1},
				{1, 2},
			},
			input: []Production{
				{
					Statement,
					[]Symbol{
						StatementSelect,
					},
				},
				{
					StatementSelect,
					[]Symbol{
						KeywordSelect,
						Identifier,
					},
				},
			},
		},
		{
			name: "valid LR states for simple SELECT query",
			expected: []LRItem{
				{0, 0},
				{0, 1},
				{1, 0},
				{1, 1},
				{1, 2},
				{1, 3},
				{1, 4},
			},
			input: []Production{
				{
					Statement,
					[]Symbol{
						StatementSelect,
					},
				},
				{
					StatementSelect,
					[]Symbol{
						KeywordSelect,
						Identifier,
						KeywordFrom,
						OperatorSemi,
					},
				},
			},
		},
	}

	for _, tc := range cases {
		g := NewGenerator(tc.input)
		lrItems := g.generateLRItems()
		assert.Len(t, lrItems, len(tc.expected))
		assert.Equal(t, tc.expected, lrItems)
	}
}
