package lib

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"fmt"
)

func ToHtml(file string, c *gin.Context) string {
	data := ReadData(file)
	res := ""
	
	if data != "File Not Found" {
		var jsd map[string]interface{}
		err := json.Unmarshal([]byte(data), &jsd)
		if syntaxErr, ok := err.(*json.SyntaxError); ok {
			line, col := findLineAndCol(data, syntaxErr.Offset)
			res = fmt.Sprintf("Error at Line %d, column %d: %v", line, col, err)
		} else {
			res = RenderHTML(data, c)
		}
	} else {
		res = data + " " + file
	}
	return res
}
