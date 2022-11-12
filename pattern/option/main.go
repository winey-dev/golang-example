package main

import "option/object"

func main() {
	o, _ := object.New(object.Mode("hi"), object.WithFile("/home/log"))
	o.View()
}
