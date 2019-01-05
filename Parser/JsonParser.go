package Parser

import (
	"fmt"
	"os"
)

func parse(configFilePath string) {
	file, err := os.Open(configFilePath)

	if err != nil {
		fmt.Println(err)
	}

}
