package bus

import (
	"sync"
)

// RootBus is root data of data bus
type Databus struct {
	mu sync.Mutex
	v  map[string]any
}

// NewBus is function to create new databus
func NewBus() *Databus {
	sm := Databus{}
	v := map[string]any{}
	sm.v = v
	return &sm
}

// NewBus is function to create new databus
func NewBusWithMap(databus *map[string]any) *Databus {

	if databus == nil {
		return NewBus()
	}
	sm := Databus{}
	v := *databus
	sm.v = v
	return &sm
}

// GetRoot is function to get root map
func (c *Databus) GetRoot() map[string]any {
	return c.getRoot()
}

// Valid is function to check valid path
func (c *Databus) Valid(path string) bool {

	// Locking
	c.mu.Lock()
	defer c.mu.Unlock()

	// if not null it's valid
	val := c.value(path)
	return val != nil
}

// Value is function to get value by path
func (c *Databus) Value(path string) any {

	// Locking
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value(path)
}

// String is function to get value and convert it's to string
func (c *Databus) String(path string) (str string) {

	// Locking
	c.mu.Lock()
	defer c.mu.Unlock()

	// Get Value
	value := c.value(path)

	// convert to string
	str, err := c.toString(value)
	if err != nil {
		panic(err)
	}

	return
}

// Int is function to get value and convert it's to int64
func (c *Databus) Int(path string) int64 {

	// Locking
	c.mu.Lock()
	defer c.mu.Unlock()

	// Get Value
	value := c.value(path)

	// convert to integer
	val, err := c.toInteger(value)
	if err != nil {
		panic(err)
	}

	return val
}

// Float is function to get value and convert it's to float64
func (c *Databus) Float(path string) float64 {

	// Locking
	c.mu.Lock()
	defer c.mu.Unlock()

	// Get Value
	value := c.value(path)

	// convert to floating
	val, err := c.toFloat64(value)
	if err != nil {
		panic(err)
	}

	return val
}

// Bool is function to get value and convert it's to boolean
func (c *Databus) Bool(path string) bool {

	// Locking
	c.mu.Lock()
	defer c.mu.Unlock()

	// Get Value
	value := c.value(path)

	// convert to boolean
	val := c.toBoolean(value)

	return val
}

// Map is function to get value and convert it's to map
func (c *Databus) Map(path string) map[string]any {

	// Locking
	c.mu.Lock()
	defer c.mu.Unlock()

	// Get Value
	value := c.value(path)

	// convert to map
	val, err := c.toMap(value)
	if err != nil {
		panic(err)
	}

	return val
}

// Array is function to get value and convert it's to array
func (c *Databus) Array(path string) []any {

	// Locking
	c.mu.Lock()
	defer c.mu.Unlock()

	// Get Value
	value := c.value(path)

	// convert to array
	val, err := c.toArray(value)
	if err != nil {
		panic(err)
	}

	return val
}

// Set is function to set value to path
func (c *Databus) Set(path string, value any) {

	// Locking
	c.mu.Lock()
	defer c.mu.Unlock()

	// call set value
	err := c.set(path, value)
	if err != nil {
		panic(err)
	}
}

// Del is function to delete value by path
func (c *Databus) Del(path string) {

	// Locking
	c.mu.Lock()
	defer c.mu.Unlock()

	// call set value
	err := c.del(path)
	if err != nil {
		panic(err)
	}
}
