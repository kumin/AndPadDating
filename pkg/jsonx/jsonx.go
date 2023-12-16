package jsonx

import (
	"encoding/json"
	"log"
)

func BeautifyJson(obj any) string {
	d, err := json.MarshalIndent(obj, " ", "  ")
	if err != nil {
		log.Print(err)
		return ""
	}
	return string(d)
}
