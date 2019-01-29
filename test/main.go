package main

import (
	"fmt"
	"github.com/bytedance/go-tagexpr/validator"
)

func main() {
	var vd = validator.New("vd")

	type InfoRequest struct {
		Name string `vd:"($!='Alice'||(Age)$==18) && regexp('\\w')"`
		Age  int    `vd:"$>0"`
	}
	info := &InfoRequest{Name: "Alice", Age: 18}
	fmt.Println(vd.Validate(info) == nil)
	// Output:
	// true
}
