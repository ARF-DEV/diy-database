package grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateLRItems(t *testing.T) {
	t.Run("empty production should return 1 LR item (with dot position = 0)", func(t *testing.T) {
		productions := []Production{
			{
				Empty,
				[]Symbol{},
			},
		}
		expectedFirstItem := LRItem{
			ProdictionIdx: 0,
			DotPosition:   0,
		}
		generator := NewGenerator(productions)
		lrItems := generator.generateLRItems()

		assert.Len(t, lrItems, 1)
		assert.Equal(t, expectedFirstItem, lrItems[0])
	})
	t.Run("return 4 lr items on SELECT IDENT and empty production", func(t *testing.T) {
		productions := []Production{
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
		}
		expected := []LRItem{
			{0, 0},
			{1, 0},
			{1, 1},
			{1, 2},
		}
		generator := NewGenerator(productions)
		lrItems := generator.generateLRItems()

		assert.Len(t, lrItems, len(expected))
		assert.Equal(t, expected, lrItems)
	})

	t.Run("valid production with nested non-terminal rules", func(t *testing.T) {
		productions := []Production{
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
		}
		expected := []LRItem{
			{0, 0},
			{0, 1},
			{1, 0},
			{1, 1},
			{1, 2},
		}
		generator := NewGenerator(productions)
		lrItems := generator.generateLRItems()

		assert.Len(t, lrItems, len(expected))
		assert.Equal(t, expected, lrItems)
	})

	t.Run("valid LR states for simple SELECT query", func(t *testing.T) {
		productions := []Production{
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
		}
		expected := []LRItem{
			{0, 0},
			{0, 1},
			{1, 0},
			{1, 1},
			{1, 2},
			{1, 3},
			{1, 4},
		}
		generator := NewGenerator(productions)
		lrItems := generator.generateLRItems()

		assert.Len(t, lrItems, len(expected))
		assert.Equal(t, expected, lrItems)
	})
}
