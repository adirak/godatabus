package bus

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// getRoot is function to get root map
func (c *Databus) getRoot() map[string]any {
	return c.v
}

// Value is function to get value by path
func (c *Databus) value(path string) interface{} {

	paths := c.splitPath(path)
	if len(paths) == 0 {
		return c.getRoot()
	}

	// value
	var curVal interface{}
	curVal = c.getRoot()

	// loop to get
	for _, spath := range paths {
		spath, idx, err := c.getArrayIndex(spath)
		if err != nil {
			panic(err)
		}
		if idx < 0 {
			val, err := c.getValueFrom(curVal, spath)
			if err != nil {
				panic(err)
			}
			curVal = val
			continue
		} else {
			val, err := c.getValueFrom(curVal, spath)
			if err != nil {
				panic(err)
			}
			val2, err := c.getValueFromArray(val, idx)
			if err != nil {
				panic(err)
			}
			curVal = val2
			continue
		}
	}

	return curVal
}

// set is function to set value by path
func (c *Databus) set(path string, value interface{}) error {

	paths := c.splitPath(path)
	if len(paths) == 0 {
		return nil
	}

	// value
	var curVal interface{}
	curVal = c.getRoot()
	setIdx := len(paths) - 1

	// loop to get
	for i, spath := range paths {

		spath, idx, err := c.getArrayIndex(spath)
		if err != nil {
			return err
		}

		// Delete Action
		if i == setIdx {

			if idx < 0 {
				mapObj, ok := curVal.(map[string]interface{})
				if !ok {
					return errors.New("data at spath=" + spath + " is not map")
				}
				mapObj[spath] = value
				return nil
			} else {

				val, err := c.getValueFrom(curVal, spath)
				if err != nil {
					return err
				}

				if val == nil {
					val = []interface{}{}
					mapObj, ok := curVal.(map[string]interface{})
					if !ok {
						return errors.New("data at spath=" + spath + " is not map")
					}
					mapObj[spath] = val
				}

				arr, err := c.setValueToArray(val, idx, value)
				if err != nil {
					return err
				}

				mapObj, ok := curVal.(map[string]interface{})
				if !ok {
					return errors.New("data at spath=" + spath + " is not map")
				}
				mapObj[spath] = arr
				return nil
			}

		}

		// Read Action
		if idx < 0 {
			val, err := c.getValueFrom(curVal, spath)
			if err != nil {
				return err
			}
			if val == nil {
				val = map[string]interface{}{}
				mapObj, ok := curVal.(map[string]interface{})
				if !ok {
					return errors.New("data at spath=" + spath + " is not map")
				}
				mapObj[spath] = val
			}
			curVal = val
			continue
		} else {
			val, err := c.getValueFrom(curVal, spath)
			if err != nil {
				return err
			}
			if val == nil {
				val = []interface{}{}
				mapObj, ok := curVal.(map[string]interface{})
				if !ok {
					return errors.New("data at spath=" + spath + " is not map")
				}
				mapObj[spath] = val
			}
			val2, err := c.getValueFromArray(val, idx)
			if err != nil {
				msg := err.Error()
				msg = strings.ToLower(msg)
				if strings.Contains(msg, "index out of bound") {
					val2 = map[string]interface{}{}
					val, err = c.setValueToArray(val, idx, val2)
					if err != nil {
						return err
					}
					mapObj, ok := curVal.(map[string]interface{})
					if !ok {
						return errors.New("data at spath=" + spath + " is not map")
					}
					mapObj[spath] = val
				}
			}
			curVal = val2
			continue
		}
	}

	return nil

}

// del is function to delete value by path
func (c *Databus) del(path string) error {

	paths := c.splitPath(path)
	if len(paths) == 0 {
		return nil
	}

	// value
	var curVal interface{}
	curVal = c.getRoot()
	delIdx := len(paths) - 1

	// loop to get
	for i, spath := range paths {

		spath, idx, err := c.getArrayIndex(spath)
		if err != nil {
			return err
		}

		// Delete Action
		if i == delIdx {

			if idx < 0 {
				mapObj, ok := curVal.(map[string]interface{})
				if !ok {
					return errors.New("data at spath=" + spath + " is not map")
				}
				delete(mapObj, spath)
				return nil
			} else {

				val, err := c.getValueFrom(curVal, spath)
				if err != nil {
					return err
				}
				arr, err := c.removeValueFromArray(val, idx)
				if err != nil {
					return err
				}
				mapObj, ok := curVal.(map[string]interface{})
				if !ok {
					return errors.New("data at spath=" + spath + " is not map")
				}
				mapObj[spath] = arr
				return nil
			}

		}

		// Read Action
		if idx < 0 {
			val, err := c.getValueFrom(curVal, spath)
			if err != nil {
				return err
			}
			curVal = val
			continue
		} else {
			val, err := c.getValueFrom(curVal, spath)
			if err != nil {
				return err
			}
			val2, err := c.getValueFromArray(val, idx)
			if err != nil {
				return err
			}
			curVal = val2
			continue
		}
	}

	return nil
}

// splitPath is function to split path by type
func (c *Databus) splitPath(path string) (paths []string) {

	// init paths
	paths = []string{}
	if path == "" {
		return
	}

	var sb strings.Builder
	iMax := len(path) - 1
	oQuote := false
	for i, r := range path {

		// Check dot operator
		dot := r == '.'
		nQuote := (r == '\'' || r == '"')
		if nQuote {
			oQuote = !oQuote
		}

		// It's changed
		if dot && !oQuote {
			if sb.Len() > 0 {
				spath := sb.String()
				paths = append(paths, spath)
			}
			sb.Reset()
		}

		// Keep data
		if oQuote && !nQuote {
			sb.WriteRune(r)
		} else if !dot && !nQuote {
			sb.WriteRune(r)
		}

		// End
		if i == iMax {
			spath := sb.String()
			paths = append(paths, spath)
		}
	}

	return paths
}

// getArrayIndex is function to get array index
func (c *Databus) getArrayIndex(spath string) (string, int, error) {
	if strings.Contains(spath, "[") && strings.HasSuffix(spath, "]") {
		s := strings.Index(spath, "[") + 1
		e := len(spath) - 1
		ssp := spath[s:e]
		val, err := strconv.ParseInt(ssp, 10, 64)
		if err != nil {
			return spath, -1, err
		}
		npath := spath[:s-1]
		return npath, int(val), nil
	}
	return spath, -1, nil
}

// getValueFrom is function to get value from current data
func (c *Databus) getValueFrom(curVal interface{}, spath string) (interface{}, error) {

	if spath == "" {
		return curVal, nil
	}

	// Curvalue is map
	mapObj, ok := curVal.(map[string]interface{})
	if !ok {
		return nil, errors.New("value at name=" + spath + " is not map")
	}
	return mapObj[spath], nil
}

// getValueFromArray is function to get value from array
func (c *Databus) getValueFromArray(curVal interface{}, index int) (interface{}, error) {

	if index < 0 {
		return nil, fmt.Errorf("array index out of bound, index=%v", index)
	}

	// Curvalue is map
	arr, ok := curVal.([]interface{})
	if !ok {
		return nil, errors.New("value at is not array")
	}

	l := len(arr)
	if index >= l {
		return nil, fmt.Errorf("array index out of bound, index=%v", index)
	}
	return arr[index], nil
}

// setValueToArray is function to set value from array
func (c *Databus) setValueToArray(curVal interface{}, index int, value interface{}) (interface{}, error) {

	if index < 0 {
		return nil, fmt.Errorf("array index out of bound, index=%v", index)
	}

	// Curvalue is map
	arr, ok := curVal.([]interface{})
	if !ok {
		return nil, errors.New("value at is not array")
	}

	l := len(arr)
	if index >= l {
		for i := l - 1; i <= index; i++ {
			if i >= l {
				arr = append(arr, nil)
			}
		}
	}

	// set value by index
	arr[index] = value
	return arr, nil
}

// removeValueFromArray is function to remove array by index
func (c *Databus) removeValueFromArray(curVal interface{}, index int) (interface{}, error) {

	if index < 0 {
		return nil, fmt.Errorf("array index out of bound, index=%v", index)
	}

	// Curvalue is map
	arr, ok := curVal.([]interface{})
	if !ok {
		return nil, errors.New("value at is not array")
	}

	l := len(arr)
	if index >= l {
		return nil, fmt.Errorf("array index out of bound, index=%v", index)
	}

	arr = append(arr[:index], arr[index+1:]...)

	return arr, nil
}
