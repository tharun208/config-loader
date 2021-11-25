package config

func merge(data map[string]interface{}, dataToAdd map[string]interface{}) map[string]interface{} {
	for key, val := range dataToAdd {
		if dataVal, ok := data[key]; ok {
			dataMap, ok := dataVal.(map[string]interface{})
			dataToAddMap, ok2 := val.(map[string]interface{})
			if ok && ok2 {
				data[key] = merge(dataMap, dataToAddMap)
			}
		}
		data[key] = val
	}
	return data
}
