package services

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/yhat/scrape"
)

func StartParsing() {
	// resp, err := http.Get("http://www.ebay.com/itm/Google-Pixel-FACTORY-UNLOCKED-32-128G-5-0-4GB-RAM-Quite-Black-Very-Silver-/332297629242?var=&hash=item4d5e7b6e3a:m:muQE3ZkCrGX9d2PDUujZTJA")
	resp, err := http.Get("http://www.ebay.com/itm/New-Apple-MacBook-12-Gold-1-2GHz-8GB-RAM-512GB-Flash-3-months-warranty-/162568320671?epid=1072587435&hash=item25d9d3b69f:g:U30AAOSwRXRZOm91")

	if err != nil {
		panic(err)
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	parseSchemaObjects(root)
	parseCustomObjects(root)
}

func parseSchemaObjects(root *html.Node) {
	schemaMap := make(map[string]string)
	schemaNodes := scrape.FindAllNested(root, SchemaObjectMatcher)
	for _, sn := range schemaNodes {
		prop := scrape.Attr(sn, "itemprop")
		_, present := schemaMap[prop]

		if !present {
			schemaMap[prop] = fmt.Sprintf("%s (%s)", scrape.Text(sn), scrape.Attr(sn, "itemtype"))
		}
	}

	PrintSearchResults(schemaMap, "SCHEMA OBJECTS")
}

func parseCustomObjects(root *html.Node) {
	customMap := make(map[string]string)
	var labels []string
	var values []string

	labelNodes := scrape.FindAllNested(root, LabelsMatcher)
	for _, ln := range labelNodes {
		labels = append(labels, scrape.Text(ln))
	}

	valueNodes := scrape.FindAllNested(root, ValuesMatcher)
	for _, vn := range valueNodes {
		values = append(values, scrape.Text(vn))
	}

	for i, v := range labels {
		customMap[v] = values[i]
	}

	PrintSearchResults(customMap, "CUSTOM OBJECTS")
}

func SchemaObjectMatcher(n *html.Node) bool {
	if scrape.Attr(n, "itemprop") != "" {
		return true
	}

	return false
}

func LabelsMatcher(n *html.Node) bool {
	if n.DataAtom == atom.Td && hasClass(n, "attrLabels") {
		return true
	}

	return false
}

func ValuesMatcher(n *html.Node) bool {
	validParent := n.Parent != nil && n.Parent.DataAtom == atom.Td
	if n.DataAtom == atom.Span && validParent {
		return true
	}

	return false
}

func PrintSearchResults(results map[string]string, header string) {
	fmt.Printf("\n%s\n", header)

	for k, v := range results {
		fmt.Printf("%s : %s\n", k, v)
	}
}

func hasClass(n *html.Node, class string) bool {
	classes := strings.Fields(scrape.Attr(n, "class"))

	for _, c := range classes {
		if c == class {
			return true
		}
	}

	return false
}

func getNodeValue(n *html.Node) string {
	if n.FirstChild == nil || n.LastChild == nil {
		return ""
	}

	return scrape.Text(n.FirstChild)
}
