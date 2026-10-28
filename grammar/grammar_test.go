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
				{ProductionIdx: 0, DotPosition: 0},
				{ProductionIdx: 1, DotPosition: 0},
				{ProductionIdx: 1, DotPosition: 1},
				{ProductionIdx: 1, DotPosition: 2},
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
				{ProductionIdx: 0, DotPosition: 0},
				{ProductionIdx: 0, DotPosition: 1},
				{ProductionIdx: 1, DotPosition: 0},
				{ProductionIdx: 1, DotPosition: 1},
				{ProductionIdx: 1, DotPosition: 2},
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
				{ProductionIdx: 0, DotPosition: 0},
				{ProductionIdx: 0, DotPosition: 1},
				{ProductionIdx: 1, DotPosition: 0},
				{ProductionIdx: 1, DotPosition: 1},
				{ProductionIdx: 1, DotPosition: 2},
				{ProductionIdx: 1, DotPosition: 3},
				{ProductionIdx: 1, DotPosition: 4},
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
					SelectList,
					KeywordFrom,
					Identifier,
					OperatorSemi,
				},
			},
			{
				Symbol: SelectList,
				Parts: []Symbol{
					SelectList,
					OperatorComma,
					Identifier,
				},
			},
		}
		expected := []State{
			{
				{ProductionIdx: 0, DotPosition: 0}, // augment -> * statement
				{ProductionIdx: 1, DotPosition: 0}, // statement -> * select_stmt
				{ProductionIdx: 2, DotPosition: 0}, // select_stmt -> * select_key ident from ident semi
			},
			{
				{ProductionIdx: 2, DotPosition: 1, parentStateIdx: 0}, // select_stmt -> select_key * select_list from ident semi
				{ProductionIdx: 3, DotPosition: 0, parentStateIdx: 0}, // select_list -> * select_list op_comma ident
			},
			{
				{ProductionIdx: 3, DotPosition: 1, parentStateIdx: 1}, // select_list -> select_list * op_comma ident
			},
			{
				{ProductionIdx: 3, DotPosition: 2, parentStateIdx: 2}, // select_list -> select_list op_comma * ident
			},
			{
				{ProductionIdx: 3, DotPosition: 3, parentStateIdx: 3}, // select_list -> select_list op_comma ident *
			},
			{
				{ProductionIdx: 2, DotPosition: 2, parentStateIdx: 1}, // select_stmt -> select_key select_list * from ident semi
			},
			{
				{ProductionIdx: 2, DotPosition: 3, parentStateIdx: 5}, // select_stmt -> select_key select_list from * ident semi
			},
			{
				{ProductionIdx: 2, DotPosition: 4, parentStateIdx: 6}, // select_stmt -> select_key select_list from ident * semi
			},
			{
				{ProductionIdx: 2, DotPosition: 5, parentStateIdx: 7}, // select_stmt -> select_key select_list from ident semi *
			},
			{
				{ProductionIdx: 1, DotPosition: 1, parentStateIdx: 0}, // statement -> select_stmt *
			},
			{
				{ProductionIdx: 0, DotPosition: 1, parentStateIdx: 0}, // augment -> statement *
			},
		}

		g := NewGenerator(input, Statement, false)
		states := g.GenerateStates()
		printState(&g, states)
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
