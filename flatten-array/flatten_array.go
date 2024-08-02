package flatten

func Flatten(nested interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range nested.([]interface{}) {
		if v == nil {
			continue
		}
		if _, ok := v.([]interface{}); ok {
			result = append(result, Flatten(v)...)
		} else {
			result = append(result, v)
		}
	}
	return result
}
