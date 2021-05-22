package parser

import (
	"regexp"
	"bufio"
	"bytes"
	"strconv"
	"strings"
	"github.com/czerwonk/bird_exporter/protocol"
)

type adhocContext struct {
	line	string
	tables	[]*protocol.Adhoc
}

type adhocRegex struct {
	matched *regexp.Regexp
	total *regexp.Regexp
	name *regexp.Regexp
}

func init() {
	regex = &adhocRegex{
		matched:	regexp.MustCompile("^(?P<match>\\d+)\\sof"),
		total:		regexp.MustCompile("of\\s+(?P<match>\\d+)\\sroutes"),
		name:		regexp.MustCompile("in\\stable\\st_\\d+_(?P<match>\\w+)$"),
	}
}

var regex *adhocRegex

func ParseAdhoc(data []byte, community *protocol.LargeCommunity) []*protocol.Adhoc {
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)

	c := &adhocContext{
		tables: make([]*protocol.Adhoc, 0),
	}

	for scanner.Scan() {
		c.line = strings.Trim(scanner.Text(), " ")
		parseTable(c, community)
	}

	return c.tables
}

func parseTable(c *adhocContext, community *protocol.LargeCommunity) {
	matched := regex.matched.FindStringSubmatch(c.line)
	if matched == nil {
		return
	}
	m, _ := strconv.Atoi(matched[1])

	total := regex.total.FindStringSubmatch(c.line)
	if total == nil {
		return
	}
	t, _ := strconv.Atoi(total[1])

	name := regex.name.FindStringSubmatch(c.line)
	if name == nil {
		return
	}

	a := &protocol.Adhoc{
		Matched: m,
		Total: t,
		Name: name[1],
		Community: community,
	}

	c.tables = append(c.tables, a)
}
