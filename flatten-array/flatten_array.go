// package flatten implements a routine to flatten a list
package flatten

import "reflect"

// Flatten takes in an array and returns a flattened array
func Flatten(data interface{}) []interface{} {
    result := []interface{}{}
	for _, item := range data.([]interface{}) {
		if item == nil {
			continue
		}
		refType := reflect.TypeOf(item)
		switch refType.Kind() {
		case reflect.Slice:
			result = append(result, Flatten(item.([]interface{}))...)
		default:
			result = append(result, item)
		}
	}
	return result
}

