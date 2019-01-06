package Parser

import (
	"ProjectCI/Configuration"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Parser struct {
	filename string
	parseDate time.Time
	cachedFile bool
}

func parseJSON(configFilePath string) {
	file, err := os.Open(configFilePath)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The configuration file can be opened.")

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var config Configuration.Configuration

	json.Unmarshal(byteValue, &config)

}