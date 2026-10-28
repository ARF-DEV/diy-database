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
		s += fmt.Sprintf("%v\n", lr.String(g.productions))
		r += s
	}
	r += fmt.Sprintln("======================")
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

func (item LRItem) String(productions []Production) string {
	production := productions[item.ProductionIdx]
	strs := make([]string, len(production.Parts)+1)
	j := 0
	for i := range len(strs) {
		if i == item.DotPosition {
			strs[i] = "*"
		} else {
			part := production.Parts[j]
			strs[i] = symbolToStr[part]
			j++
		}
	}
	return strings.Join(strs, " ")
}

var symbolToStr = map[Symbol]string{
	Empty:           "Empty",
	KeywordSelect:   "Keyword_Select",
	KeywordFrom:     "Keyword_From",
	Identifier:      "Ident",
	SelectList:      "Select_List",
	OperatorSemi:    "Semi",
	OperatorComma:   "Comma",
	Statement:       "Statement",
	StatementSelect: "Statement_Select",
	AugmentStart:    "Augment_Start",
}
