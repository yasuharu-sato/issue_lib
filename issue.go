package issue_lib

import "fmt"

type Issue struct {
	Names []string
	param *test
}

type test struct {
	param string
}

func (i *Issue) Code() {
	var anyString []string
	for _, name := range i.Names {
		anyString = append(anyString, name)
	}
	fmt.Printf("There is an issue(%s, %s).", anyString[0], i.param.param)
}

func Code2() {
	var i Issue
	i.Code()
}
