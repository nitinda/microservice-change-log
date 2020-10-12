package console

import (
	"encoding/json"
	"fmt"
	"log"
)

// ToJSON to convert data in json
func ToJSON(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
