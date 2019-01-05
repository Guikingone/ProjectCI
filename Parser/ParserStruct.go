package Parser

import "time"

type Parser struct {
	filename string
	parseDate time.Time
	cachedFile bool
}

func (parser *Parser) parse(filePath string) {

}