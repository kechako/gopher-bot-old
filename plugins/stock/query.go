package stock

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

const msnUrl = "http://www.msn.com/ja-jp/money/stockdetails/fi-133.1.%d.TKS"

type stock struct {
	Company       string
	StockNumber   string
	CurrentValue  string
	Direction     string
	Change        string
	PercentChange string
}

func queryStock(number int) (*stock, error) {
	u, err := createUrl(number)
	if err != nil {
		return nil, err
	}
	root, err := getHtml(u)
	if err != nil {
		return nil, err
	}

	detail := findNode(root, "div", "class", "stockdetailsheader")

	return &stock{
		Company:       getNodeText(findNode(detail.FirstChild, "div", "class", "header-companyname")),
		StockNumber:   getNodeText(findNode(detail.FirstChild, "div", "class", "subheader-symbol")),
		CurrentValue:  getNodeText(findNode(detail.FirstChild, "span", "data-role", "currentvalue")),
		Direction:     getNodeText(findNode(detail.FirstChild, "span", "data-role", "changedir")),
		Change:        getNodeText(findNode(detail.FirstChild, "div", "data-role", "change")),
		PercentChange: getNodeText(findNode(detail.FirstChild, "div", "data-role", "percentchange")),
	}, nil
}

func createUrl(number int) (string, error) {
	u, err := url.Parse(fmt.Sprintf(msnUrl, number))
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

func getHtml(u string) (*html.Node, error) {
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return html.Parse(res.Body)
}

func findNode(n *html.Node, nodeName, attrName, attrValue string) *html.Node {
	if n == nil {
		return nil
	}

	match := false
	if n.DataAtom.String() == nodeName {
		for _, attr := range n.Attr {
			if attr.Key == attrName && attr.Val == attrValue {
				match = true
				break
			}
		}
	}
	if match {
		return n
	}

	node := findNode(n.FirstChild, nodeName, attrName, attrValue)
	if node != nil {
		return node
	}

	return findNode(n.NextSibling, nodeName, attrName, attrValue)
}

func getNodeText(n *html.Node) string {
	if n == nil {
		return ""
	}

	var buf bytes.Buffer

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)

	return buf.String()
}

func dumpNode(n *html.Node) {
	if n == nil {
		return
	}

	buf := bufio.NewWriter(os.Stdout)
	defer buf.Flush()

	var f func(*html.Node, int)
	f = func(n *html.Node, level int) {
		for i := 0; i < level; i++ {
			fmt.Fprint(buf, " ")
		}

		switch n.Type {
		case html.ErrorNode:
			fmt.Fprint(buf, "#error")
		case html.TextNode:
			fmt.Fprintf(buf, "#text %s", n.Data)
		case html.DocumentNode:
			fmt.Fprint(buf, "#doc")
		case html.ElementNode:
			fmt.Fprint(buf, n.DataAtom)
			for _, a := range n.Attr {
				fmt.Fprintf(buf, " %s=%s", a.Key, a.Val)
			}
		case html.CommentNode:
			fmt.Fprint(buf, "#comment")
		case html.DoctypeNode:
			fmt.Fprint(buf, "#doctype")
		}
		fmt.Fprintln(buf)

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, level+1)
		}
	}

	f(n, 0)
}
