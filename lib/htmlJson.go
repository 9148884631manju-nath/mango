package lib

import (
	"github.com/gin-gonic/gin"
)

func HtmlJson(data interface{}, c *gin.Context) string {
	res := ""
	res = RenderHTML(data.(string), c)
	return res
}
