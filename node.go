package simple_html_parser

import (
	"bytes"
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

func newNode(html string, beginTag string) *node {
	i := strings.Index(html, beginTag)
	if i == -1 {
		return nil
	}
	tagR := regexp.MustCompile(`<([^\s]+).*?>`).FindStringSubmatch(beginTag)
	n := &node{
		start:          html[i : i+len(beginTag)],
		html:           html[i+len(beginTag):],
		tag:            tagR[1],
		sameTagCounter: 1,
		pause:          false,
	}
	return n
}

func (n *node) traverse() (result string, err error) {
	r := strings.NewReader(n.html)
	var tagBytes = []byte(n.tag)
	var startTagByte = byte('<')
	var endTagSeparatorByte = byte('/')
	var char byte
	for char, err = r.ReadByte(); err == nil; char, err = r.ReadByte() {
		result += string(char)
		if char == startTagByte {
			var nextBytes = make([]byte, len(tagBytes))
			char, err = r.ReadByte()
			if err != nil {
				err = errors.New("invalid html tag " + err.Error())
				return
			}
			result += string(char)
			isEnding := false
			if char == endTagSeparatorByte {
				isEnding = true
				nextBytes = make([]byte, len(tagBytes)+1)
			} else {
				nextBytes = make([]byte, len(tagBytes)-1)
			}
			_, err = r.Read(nextBytes)
			if err != nil {
				err = errors.New("invalid html tag " + err.Error())
				return
			}
			var compareBytes []byte
			if !isEnding {
				compareBytes = append([]byte{char}, nextBytes...)
			} else {
				compareBytes = nextBytes[:len(nextBytes)-1]
			}

			if bytes.Compare(tagBytes, compareBytes) == 0 {
				if isEnding {
					n.sameTagCounter--
				} else {
					n.sameTagCounter++
				}
			}
			result += string(nextBytes)
			if n.sameTagCounter == 1 && n.pause {
				result = n.start + result
				return
			}
			n.pause = true
		}
	}
	err = errors.New("invalid html")
	return
}
