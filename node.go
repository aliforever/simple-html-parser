package simple_html_parser

import (
	"errors"
	"io"
	"regexp"
	"strings"
)

type node struct {
	rawHtml        string
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
	tagR := regexp.MustCompile(`<([^\s^/>]+)\s?.+?>?`).FindStringSubmatch(beginTag)
	n = &node{
		rawHtml:        html,
		start:          html[i : i+len(beginTag)],
		html:           html[i+len(beginTag):],
		tag:            strings.TrimSpace(tagR[1]),
		sameTagCounter: 1,
		pause:          false,
	}
	return
}

func (n *node) readByte(r *strings.Reader) (b byte, eof bool, err error) {
	b, err = r.ReadByte()
	if err != nil && err == io.EOF {
		eof = true
		err = nil
	}
	return
}

func (n *node) readTagByte(r *strings.Reader) (data string, isOpeningTag, isClosingTag bool, eof bool, err error) {
	var (
		b          byte
		tagBytes   = []byte(n.tag)
		slashByte  = byte('/')
		endTagByte = byte('>')
	)
	for _, tagByte := range tagBytes {
		if eof {
			return
		}
		b, eof, err = n.readByte(r)
		if err != nil {
			return
		}
		data += string(b)
		if !eof && b == slashByte {
			isClosingTag = true
			b, eof, err = n.readByte(r)
			if err != nil {
				return
			}
			data += string(b)
		}
		if !eof && b != tagByte {
			isClosingTag = false
			isOpeningTag = false
			return
		} else {
			if !isClosingTag {
				isOpeningTag = true
			}
		}
	}

	b = 0
	for b != endTagByte && !eof {
		b, eof, err = n.readByte(r)
		if err != nil {
			return
		}
		data += string(b)
	}
	if b != endTagByte {
		err = errors.New("eng tag not found")
	}
	return
}

func (n *node) traverse() (endTagIndex int, err error) {
	r := strings.NewReader(n.html)

	var startTagByte = byte('<')
	var char byte
	var isOpening, isClosing bool
	var eof bool

	for char, err = r.ReadByte(); err == nil; char, err = r.ReadByte() {
		if char == startTagByte {
			_, isOpening, isClosing, eof, err = n.readTagByte(r)
			if err != nil || eof {
				return
			}
			if isOpening {
				n.sameTagCounter++
			} else if isClosing {
				n.sameTagCounter--
			}
			n.pause = true

			if n.sameTagCounter == 0 && n.pause {
				endTagIndex = int(r.Size()) - r.Len()
				return
			}
		}
	}
	err = errors.New("invalid html")
	return
}
