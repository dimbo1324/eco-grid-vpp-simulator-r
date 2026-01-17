package main

import (
	"fmt"
	"reflect"
)

func main() {
	properties := map[string]string{
		"qwe": "awe",
	}

	fmt.Println(reflect.TypeOf(properties["qwe"]))

}
