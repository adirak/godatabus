package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/adirak/godatabus/bus"
)

// Read json map function
func JsonMap(file string) map[string]any {

	jsonFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	bValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	mapObj := map[string]any{}
	json.Unmarshal(bValue, &mapObj)

	return mapObj

}

func TestDatabus(t *testing.T) {

	file := "./sample/test_databus.json"
	data := JsonMap(file)

	root := bus.NewBusWithMap(&data)

	str := root.String("id")
	num := root.Int("id")
	strL := root.String("list[2]")

	root.Set("list[5]", "buffalo")
	list := root.Value("list")

	root.Set("myList[0].a", "aaa")
	root.Set("myList[0].b", "bbb")
	root.Set("myList[0].c", "ccc")
	root.Set("myList[0].d", "ddd")
	myList := root.Value("myList")

	root.Set("'x.y.z'", "XYZ")
	xyz := root.String("'x.y.z'")

	t.Log("str:", str)
	t.Log("num:", num)
	t.Log("strL:", strL)
	t.Log("list:", list)
	t.Log("myList:", myList)
	t.Log("xyz:", xyz)
}
