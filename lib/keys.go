package lib

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func Keys(ar [][]string, con []byte, c *gin.Context) string {
	res:= string(con)
	for i:=0; i<len(ar); i++{
		switch ar[i][0]{
		case "request":
				res=strings.ReplaceAll(res,ar[i][2],c.DefaultQuery(ar[i][1],ar[i][3]))
		default:
		}
	}
	return res
}
