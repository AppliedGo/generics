// genericsWithInterfaces demonstrates how interfaces can be used to achieve generic-like behavior.
package main

import "errors"

// Item is a 'generic' item type
type Item interface{}

// Container is our 'generic' container type
type Container interface {
	Put(Item)
	Get() Item
}

// StringContainer is a 'string instance' of Container
type StringContainer struct {
	s []string
}

// Put adds an Item to the container, converting it into a string.
func (c *StringContainer) Put(i Item) error {
	s, ok = i.(string)
	if !ok {
		return errors.New("Wrong type. Expected: string, got: %s" + reflect.Type(i))
	}
}

func (c *StringContainer) Get() Item {
	return s
}

func main() {
	c := &StringContainer{}

}
