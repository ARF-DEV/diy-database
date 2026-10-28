// go struct based context free grammar implementation
package grammar

type (
	Symbol     int
	Production struct {
		Symbol Symbol
		Parts  []Symbol
	}
	LRItem struct {
		ProductionIdx int
		DotPosition   int

		parentStateIdx int
	}

	// Note: states == item sets
	State []LRItem
)

const (
	// terminal symbols
	Empty Symbol = iota
	KeywordSelect
	KeywordFrom
	Identifier
	OperatorSemi
	OperatorComma

	// non-terminal symbols
	Statement // augment grammar
	StatementSelect
	SelectList

	AugmentStart
)

// symbols expect augment grammar
var symbols []Symbol = []Symbol{
	KeywordSelect,
	KeywordFrom, Identifier, OperatorSemi,
	OperatorComma,
	StatementSelect,
	SelectList,
}

type generator struct {
	startSymbol Symbol
	productions []Production
	visited     map[LRItem]struct{}
}

func NewGenerator(productions []Production, startSymbol Symbol, ignoreSymbol bool) generator {
	g := generator{
		startSymbol: startSymbol,
		visited:     map[LRItem]struct{}{},
	}
	if ignoreSymbol {
		g.productions = productions
	} else {
		p := make([]Production, len(productions)+1)
		p[0] = g.augment()
		for i := 1; i < len(p); i++ {
			p[i] = productions[i-1]
		}
		g.productions = p
	}
	return g
}

func (g *generator) generateLRItems() []LRItem {
	lrItems := []LRItem{}
	for idx := range g.productions {
		lrItems = append(lrItems, g.lrItemsFromProduction(idx)...)
	}
	return lrItems
}

// DEPRECATED
func (g *generator) lrItemsFromProduction(idx int) []LRItem {
	lrItems := []LRItem{}
	// a queue
	queue := []LRItem{
		{ProductionIdx: idx, DotPosition: 0},
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
		production := g.productions[item.ProductionIdx]
		if len(production.Parts) > item.DotPosition {
			queue = append(queue, LRItem{ProductionIdx: item.ProductionIdx, DotPosition: item.DotPosition + 1})
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
		if p.Symbol == symbol {
			lrItems = append(lrItems, LRItem{ProductionIdx: idx, DotPosition: 0})
		}
	}
	return lrItems
}

func isNonTerminal(symbol Symbol) bool {
	nonTerminalMap := map[Symbol]struct{}{
		AugmentStart:    {},
		StatementSelect: {},
		Statement:       {},
		SelectList:      {},
	}
	_, found := nonTerminalMap[symbol]
	return found
}

func (g *generator) GenerateStates() []State {
	// start from augment
	stack := []LRItem{
		{ProductionIdx: 0, DotPosition: 0},
	}
	states := []State{}
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// all state start with one start LR item
		state := g.getState([]LRItem{item})
		states = append(states, state)
		curState := len(states) - 1
		// handle transition
		for _, s := range state {
			production := g.productions[s.ProductionIdx]
			if s.DotPosition < len(production.Parts) {
				// transition
				sCopy := s
				sCopy.DotPosition++
				sCopy.parentStateIdx = curState // record valid entry point to a state candidate
				stack = append(stack, sCopy)
			}
		}
	}
	return states
}

func (g *generator) getState(stack []LRItem) []LRItem {
	finalState := []LRItem{}
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, visited := g.visited[item]; visited {
			continue
		}
		if g.visited == nil {
			g.visited = map[LRItem]struct{}{}
		}
		g.visited[item] = struct{}{}
		production := g.productions[item.ProductionIdx]
		if item.DotPosition < len(production.Parts) {
			nextSymbol := production.Parts[item.DotPosition]
			if isNonTerminal(nextSymbol) {
				stack = append(stack, g.closure(nextSymbol)...)
			}
		}
		finalState = append(finalState, item)
	}
	return finalState
}

func (g *generator) augment() Production {
	return Production{
		Symbol: AugmentStart,
		Parts: []Symbol{
			g.startSymbol,
		},
	}
}
