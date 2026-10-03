package grammar

import (
	"fmt"
	"strings"
)

func (g *generator) StatesToString(item State) string {
	r := ""
	for _, lr := range item {
		s := ""
		s += fmt.Sprintf("%v\n", lr)
		s += fmt.Sprintf("%v\n", g.productions[lr.ProductionIdx])
		r += s + "\n"
	}
	return r
}

func (p Production) String() string {
	s := ""
	s += "symbol: " + symbolToStr[p.Symbol] + "\n"
	parts := []string{}
	for _, p := range p.Parts {
		parts = append(parts, symbolToStr[p])
	}
	s += "parts: [" + strings.Join(parts, ",") + "]\n"
	return s
}

var symbolToStr = map[Symbol]string{
	Empty:           "Empty",
	KeywordSelect:   "Keyword_Select",
	KeywordFrom:     "Keyword_From",
	Identifier:      "Ident",
	OperatorSemi:    "Semi",
	Statement:       "Statement",
	StatementSelect: "Statement_Select",
	AugmentStart:    "test",
}
