package Parser

import (
	"ProjectCI/Configuration"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonParser struct {
	running bool
}

func (parser *JsonParser) parse(configFilePath string) {
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
