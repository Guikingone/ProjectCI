package Parser

import "time"

type JsonParser struct {
	filename string
	parseDate time.Time
	cachedFile bool
}

func (parser *JsonParser) parse(filePath string) {

}
