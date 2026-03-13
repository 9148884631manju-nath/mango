package lib

import (
	"encoding/json"
	"fmt"
)

func ToHtml(file string) string {
	data := ReadData(file)
	res := ""
	
	if data != "File Not Found" {
		var jsd map[string]interface{}
		err := json.Unmarshal([]byte(data), &jsd)
		if syntaxErr, ok := err.(*json.SyntaxError); ok {
			line, col := findLineAndCol(data, syntaxErr.Offset)
			res = fmt.Sprintf("Error at Line %d, column %d: %v", line, col, err)
		} else {
			res = RenderHTML(data)
		}
	} else {
		res = data + " " + file
	}
	return res
}
