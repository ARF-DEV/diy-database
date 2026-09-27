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
	OperatorSemi

	// non-terminal symbols
	Statement // base statement // a start non terminal
	StatementSelect
)

type generator struct {
	productions []Production
	visited     map[LRItem]struct{}
}

func NewGenerator(productions []Production) generator {
	return generator{
		productions: productions,
		visited:     map[LRItem]struct{}{},
	}
}

func (g *generator) generateLRItems() []LRItem {
	lrItems := []LRItem{}
	for idx := range g.productions {
		lrItems = append(lrItems, g.lrItemsFromProduction(idx)...)
	}
	return lrItems
}

func (g *generator) lrItemsFromProduction(idx int) []LRItem {
	lrItems := []LRItem{}
	// a queue
	queue := []LRItem{
		{idx, 0},
	}
	if g.visited == nil {
		g.visited = map[LRItem]struct{}{}
	}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if _, isVisited := g.visited[item]; isVisited {
			continue
		}

		g.visited[item] = struct{}{}
		lrItems = append(lrItems, item)
		production := g.productions[item.ProdictionIdx]
		if len(production.Parts) > item.DotPosition {
			queue = append(queue, LRItem{item.ProdictionIdx, item.DotPosition + 1})
		}
		// do closure operation if non terminal
		if len(production.Parts) > 0 && item.DotPosition < len(production.Parts) {
			nextSymbol := production.Parts[item.DotPosition]
			if isNonTerminal(nextSymbol) {
				closureLRItems := g.closure(nextSymbol)
				queue = append(queue, closureLRItems...)
			}
		}
	}
	return lrItems
}

func (g *generator) closure(symbol Symbol) []LRItem {
	lrItems := []LRItem{}
	for idx, p := range g.productions {
		if p.Name == symbol {
			lrItems = append(lrItems, LRItem{idx, 0})
		}
	}
	return lrItems
}

func isNonTerminal(symbol Symbol) bool {
	nonTerminalMap := map[Symbol]struct{}{
		StatementSelect: {},
		Statement:       {},
	}
	_, found := nonTerminalMap[symbol]
	return found
}
