package lib

func HtmlAttr(data map[string]interface{}) string {
	res := ""
	for k, v := range data {
		var vv string
		switch t := v.(type) {
		case map[string]interface{}:
			vv = AttrJson(t)
		case []interface{}:
			vv = ArrayAttr(t)
		default:
			vv = t.(string)
		}

		if k != "" {
			res += `` + k + `="` + vv + `" `
		} else {
			res += ` ` + vv + ` `
		}

	}
	return res
}
