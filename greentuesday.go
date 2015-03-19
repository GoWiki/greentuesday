// greentuesday project greentuesday.go
package greentuesday

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Policy struct {
	Add []AttrEle
	//Remove []AttrEle
}

type AttrEle struct {
	Tag       string
	Attribute html.Attribute
}

func (p Policy) Massage(s string) string {

	r := strings.NewReader(s)

	var buff bytes.Buffer
	tokenizer := html.NewTokenizer(r)

	for {
		if tokenizer.Next() == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				return buff.String()
			}
			return ""
		}

		token := tokenizer.Token()
		switch token.Type {
		case html.StartTagToken:
			for _, add := range p.Add {
				if add.Tag == token.Data {
					token.Attr = append(token.Attr, add.Attribute)
				}
			}
			buff.WriteString(token.String())
		default:
			buff.WriteString(token.String())
		}

	}
}
