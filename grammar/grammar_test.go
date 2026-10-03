package grammar

import (
	"fmt"
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
					ProductionIdx: 0,
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
		g := NewGenerator(tc.input, Statement, true)
		lrItems := g.generateLRItems()
		assert.Len(t, lrItems, len(tc.expected))
		assert.Equal(t, tc.expected, lrItems)
	}
}

func TestGenerateState(t *testing.T) {
	t.Run("testing generate state", func(t *testing.T) {
		input := []Production{
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
					Identifier,
					OperatorSemi,
				},
			},
		}
		expected := []State{
			{
				{0, 0}, // test -> * statement
				{1, 0}, // statement -> * select_stmt
				{2, 0}, // select_stmt -> select_key ident from ident semi
			},
			{
				{0, 1}, // test -> statement *
			},
			{
				{1, 1}, // statement -> select_stmt *
			},
			{
				{2, 1}, // select_stmt -> select_key * ident from ident semi
			},
			{
				{2, 2}, // select_stmt -> select_key ident * from ident semi
			},
			{
				{2, 3}, // select_stmt -> select_key ident from * ident semi
			},
			{
				{2, 4}, // select_stmt -> select_key ident from ident * semi
			},
			{
				{2, 5}, // select_stmt -> select_key ident from ident semi *
			},
		}

		g := NewGenerator(input, Statement, false)
		states := g.GenerateStates()
		assert.Equal(t, expected, states)
	})
}

func printState(g *generator, s []State) {
	for i, s := range s {
		a := g.StatesToString(s)
		fmt.Println("state", i)
		fmt.Println(a)
	}
	fmt.Println(g.productions)
}
