package lib

import (
"encoding/json"
"github.com/gin-gonic/gin"
)

func JsonArrayRenderHtml(v interface{},keys interface{}, c *gin.Context) string {
res:=""


	pretty, _ := json.MarshalIndent(v, "  ", "  ")
if keys!=nil{
	keysPretty, _ := json.MarshalIndent(keys, "  ", "  ")
	var keysJsd [][]string
	json.Unmarshal([]byte(keysPretty), &keysJsd)
	//res+=string(keysPretty)
	res=Keys(keysJsd,pretty,c)
	pretty = []byte(res)
}else {

}

var jsd []map[string]interface{}
json.Unmarshal([]byte(pretty), &jsd)
res=HtmlJsonArray(jsd,c)
return res
}


