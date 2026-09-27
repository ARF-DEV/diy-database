// go struct based context free grammar implementation
package grammar

type (
	Symbol     int
	Production struct {
		Name  Symbol
		Parts []Symbol
	}
	LRItem struct {
		ProdictionIdx int
		DotPosition   int
	}
)

const (
	// terminal symbols
	Empty Symbol = iota
	KeywordSelect
	KeywordFrom
	Identifier

	// non-terminal symbols
	StatementSelect
)

func generateLRItems(productions []Production) []LRItem {
	lrItems := []LRItem{}
	for idx, p := range productions {
		// append the start dot ex: on SELECT IDENT -> * SELECT IDENT
		lrItems = append(lrItems, LRItem{idx, 0})
		for pos := range p.Parts {
			lrItems = append(lrItems, LRItem{idx, pos + 1})
		}
	}
	return lrItems
}
