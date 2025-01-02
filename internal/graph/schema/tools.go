package schema

import "github.com/nmarsollier/resourcesgo/internal/graph/model"

func mapReqValues(data []*model.KeyValueInput) map[string]string {
	result := make(map[string]string)
	for _, kv := range data {
		result[kv.Key] = kv.Value
	}
	return result
}
func mapValues(values map[string]string) []*model.KeyValue {
	var result []*model.KeyValue
	for k, v := range values {
		result = append(result, &model.KeyValue{
			Key:   k,
			Value: v,
		})
	}

	return result
}
