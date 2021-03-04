package simple_html_parser

type Parser struct {
	HTML string
}

func NewParser(html string) *Parser {
	return &Parser{HTML: html}
}

func (p *Parser) ExtractTag(tagBegin string) (tag string, err error) {
	var n *node
	n, err = newNode(p.HTML, tagBegin)
	if err != nil {
		return
	}
	var endTagIndex int
	endTagIndex, err = n.traverse()
	if err != nil {
		return
	}
	tag = n.start + n.html[:endTagIndex]
	return
}

func (p *Parser) ExtractTags(tagBegin string) (tags []string, err error) {
	for true {
		var n *node
		n, err = newNode(p.HTML, tagBegin)
		if err != nil {
			if len(tags) > 0 {
				err = nil
			}
			return
		}
		var endTagIndex int
		endTagIndex, err = n.traverse()
		if err != nil {
			return
		}
		tags = append(tags, n.start+n.html[:endTagIndex])
		p.HTML = n.html[endTagIndex:]
	}
	return
}
