package simple_html_parser

type Parser struct {
	HTML string
}

func NewParser(html string) *Parser {
	return &Parser{HTML: html}
}

func (p *Parser) ExtractTag(tagBegin string) (tag string, err error) {
	n := newNode(p.HTML, tagBegin)
	return n.traverse()
}
