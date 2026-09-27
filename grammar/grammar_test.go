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
		lrItems := generateLRItems(productions)

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
		lrItems := generateLRItems(productions)

		assert.Len(t, lrItems, len(expected))
		assert.Equal(t, expected, lrItems)
	})
}
