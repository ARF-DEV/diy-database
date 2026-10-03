// parser package
package parser

import "github.com/arf-dev/diy-database/tokenizer"

type Parser struct {
	stack []tokenizer.Token
}

func (p *Parser) Push(t tokenizer.Token) error {
	// TODO: handle reduce, shift, and conflict
	p.stack = append(p.stack, t)
	return nil
}
