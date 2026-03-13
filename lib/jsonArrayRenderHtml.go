package lib

import (
"encoding/json"
)

func JsonArrayRenderHtml(v interface{}) string {
res:=""

pretty, _ := json.MarshalIndent(v, "  ", "  ")
var jsd []map[string]interface{}
json.Unmarshal([]byte(pretty), &jsd)
res=HtmlJsonArray(jsd)
return res
}