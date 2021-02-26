package simple_html_parser

import (
	"errors"
	"regexp"
	"strings"
)

type node struct {
	html           string
	tag            string
	start          string
	isDone         bool
	sameTagCounter int
	pause          bool
}

func newNode(html string, beginTag string) (n *node, err error) {
	i := strings.Index(html, beginTag)
	if i == -1 {
		err = errors.New("begin_tag_not_found")
		return
	}
	tagR := regexp.MustCompile(`<([^\s]+).*?>`).FindStringSubmatch(beginTag)
	n = &node{
		start:          html[i : i+len(beginTag)],
		html:           html[i+len(beginTag):],
		tag:            strings.TrimSpace(tagR[1]),
		sameTagCounter: 1,
		pause:          false,
	}
	return
}

func (n *node) traverse() (result string, err error) {
	r := strings.NewReader(n.html)
	var tagBytes = []byte(n.tag)

	var startTagByte = byte('<')
	var endTagByte = byte('>')
	var endTagSeparatorByte = byte('/')
	var char byte

Loop:
	for char, err = r.ReadByte(); err == nil; char, err = r.ReadByte() {
		result += string(char)
		isEnding := false

		if char == startTagByte {
			for _, tagByte := range tagBytes {
				char, err = r.ReadByte()
				if err != nil {
					result = ""
					err = errors.New("invalid html tag " + err.Error())
					return
				}
				result += string(char)
				if char == endTagSeparatorByte {
					isEnding = true
					char, err = r.ReadByte()
					if err != nil {
						result = ""
						err = errors.New("invalid html tag " + err.Error())
						return
					}
					result += string(char)
				}
				if char != tagByte {
					continue Loop
				}
			}

			for true {
				char, err = r.ReadByte()
				if err != nil {
					err = errors.New("invalid html tag " + err.Error())
					result = ""
					return
				}
				result += string(char)
				if char == endTagByte {
					break
				}
			}

			if !isEnding {
				n.sameTagCounter++
			} else {
				n.sameTagCounter--
			}
			n.pause = true

			if n.sameTagCounter == 0 && n.pause {
				result = n.start + result
				return
			}
		}
	}
	result = ""
	err = errors.New("invalid html")
	return
}
