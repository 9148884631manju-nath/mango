package lib

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func RenderHTML(data string, c *gin.Context) string {
	res := ""
	var jsd []map[string]interface{}
	json.Unmarshal([]byte(data), &jsd)
	res=HtmlJsonArray(jsd,c)
	return res
}
