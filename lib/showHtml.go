package lib

import (
	"github.com/gin-gonic/gin"
)

func ShowHTML(k string, v interface{}, keys interface{}, c *gin.Context) string {
	var jv string
	var res string
	switch k {
	case "content":
		s := v.(string)
		jv = s
	case "block":
		s := v.(string)
		jv = ToHtml(s,c)
	case "inblock":
		jv = JsonArrayRenderHtml(v,keys,c)
	default:
		jv = ""
	}
	res=jv
	return res
}
