package bus

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// toString is function to convert object to string
func (c *Databus) toString(value interface{}) (string, error) {

	// Empty String
	if value == nil {
		return "", nil
	}

	str, ok := value.(string)
	if ok {
		return str, nil
	}

	// map
	mapObj, ok := value.(map[string]interface{})
	if ok {
		bData, err := json.Marshal(mapObj)
		if err != nil {
			return "", err
		}
		str = string(bData)
		return str, nil
	}

	// array
	arrObj, ok := value.([]interface{})
	if ok {
		bData, err := json.Marshal(arrObj)
		if err != nil {
			return "", err
		}
		str = string(bData)
		return str, nil
	}

	// Other type
	str = fmt.Sprintf("%v", value)
	return str, nil

}

// toInteger is function to convert number to int64
func (c *Databus) toInteger(value interface{}) (int64, error) {

	if value == nil {
		return 0, nil
	}

	switch num := value.(type) {
	case int:
		return int64(num), nil
	case int64:
		return num, nil
	case int32:
		return int64(num), nil
	case int16:
		return int64(num), nil
	case float64:
		return int64(num), nil
	case float32:
		return int64(num), nil
	default:
		str := fmt.Sprintf("%v", num)
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			offset := strings.Index(str, ".")
			nStr := str[:offset]
			i, err = strconv.ParseInt(nStr, 10, 64)
			if err != nil {
				return 0, err
			}
		}
		return i, nil
	}
}

// toFloat64 is function to convert number to float64
func (c *Databus) toFloat64(value interface{}) (float64, error) {

	if value == nil {
		return 0, nil
	}

	switch num := value.(type) {
	case int:
		return float64(num), nil
	case int64:
		return float64(num), nil
	case int32:
		return float64(num), nil
	case int16:
		return float64(num), nil
	case float64:
		return num, nil
	case float32:
		return float64(num), nil
	default:
		str := fmt.Sprintf("%v", num)
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, err
		}
		return f, nil
	}

}

// toBoolean is function to convert object to boolean
func (c *Databus) toBoolean(value interface{}) bool {

	// return false
	if value == nil {
		return false
	}

	switch num := value.(type) {
	case int:
		return num != 0
	case int64:
		return num != 0
	case int32:
		return num != 0
	case int16:
		return num != 0
	case float64:
		return num != 0
	case float32:
		return num != 0
	default:

		str := fmt.Sprintf("%v", num)
		str = strings.ToLower(str)
		str = strings.TrimSpace(str)
		if strings.Contains(str, "true") {
			return true
		} else if strings.Contains(str, "false") {
			return false
		} else if strings.EqualFold(str, "1") {
			return true
		} else if strings.EqualFold(str, "0") {
			return false
		} else if len(str) > 0 {
			return true
		}
		return false

	}

}

// toMap is function to convert object to map
func (c *Databus) toMap(value interface{}) (map[string]interface{}, error) {

	switch obj := value.(type) {

	case map[string]interface{}:

		return obj, nil

	case string:

		mapObj := map[string]interface{}{}
		err := json.Unmarshal([]byte(obj), &mapObj)
		if err != nil {
			return nil, errors.New("data is not map")
		}

		return mapObj, nil

	default:
		return nil, errors.New("data is not map")
	}

}

// toArray is function to convert object to array
func (c *Databus) toArray(value interface{}) ([]interface{}, error) {

	switch obj := value.(type) {

	case []interface{}:

		return obj, nil

	case string:

		arrayObj := []interface{}{}
		err := json.Unmarshal([]byte(obj), &arrayObj)
		if err != nil {
			return nil, errors.New("data is not array")
		}

		return arrayObj, nil

	default:
		return nil, errors.New("data is not array")
	}

}
