package log

func logParamsToZeroParams(key map[ExtraKey]interface{}) map[string]interface{} {
	params := map[string]interface{}{}

	for k, v := range key {
		params[string(k)] = v
	}
	return params
}
