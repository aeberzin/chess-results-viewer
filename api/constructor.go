package api

import (
	"reflect"
)

type constructor struct {
	indexes map[int]string
	config  map[string]string
	Items   []interface{}
	reflect reflect.Type
}

func (c *constructor) add(row []string) {
	element := reflect.New(c.reflect).Elem()
	for i, v := range row {
		if name, ok := c.indexes[i]; ok {
			element.FieldByName(name).SetString(v)
		}
	}
	c.Items = append(c.Items, element.Interface())
}

func (c *constructor) setIndexes(row []string) {
	c.indexes = make(map[int]string)
	for i, v := range row {
		if val, ok := c.config[v]; ok {
			c.indexes[i] = val
		}
	}
}

func (c *constructor) fill(table [][]string) {
	for _, v := range table[1:] {
		c.add(v)
	}
}
