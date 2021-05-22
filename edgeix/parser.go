package edgeix

import (
	"regexp"
	"bytes"
	"bufio"
	"fmt"
	"strconv"
	"errors"
	"strings"
)


type tableRegex struct {
	matching *regexp.Regexp
	total *regexp.Regexp
	name *regexp.Regexp
}

type tableContext struct {
	line string
	matching int
	total	int
	name	string

}

func init() {
	table = &tableRegex{
		matching:	regexp.MustCompile("^(?P<match>\\d+)\\sof"),
		total:		regexp.MustCompile("of\\s+(?P<match>\\d+)\\sroutes"),
		name:		regexp.MustCompile("in\\stable\\s(?P<match>\\w+)$"),
	}
}

var table *tableRegex

func ParseTables(data []byte) []*Table {
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)

	response := []*Table{}

	for scanner.Scan() {
		context := &tableContext{}
		context.line = strings.Trim(scanner.Text(), " ")
		err := parseLineForMatching(context)
		if err == nil {
			parsed := &Table{
				MatchingRoutes: context.matching,
				TotalRoutes: context.total,
				TableName: context.name,
			}
			response = append(response, parsed)
			fmt.Printf("%+v\n", parsed)
		}
	}

	return response
}

func parseLineForMatching(context *tableContext) (error) {
	matching := table.matching.FindStringSubmatch(context.line)
	if matching == nil {
		fmt.Println("no matching..")
		return errors.New("No matching")
	}
	//fmt.Printf("Matching: %v", matching)
	context.matching, _ = strconv.Atoi(matching[1])
	total := table.total.FindStringSubmatch(context.line)
	if total == nil {
		return errors.New("No total")
	}
	//fmt.Printf("Total: %v", total)
	context.total, _ = strconv.Atoi(total[1])
	name := table.name.FindStringSubmatch(context.line)
	if name == nil {
		return errors.New("No name")
	}
	//fmt.Printf("Name: %v", name)
	context.name = name[1]

	return nil
}
