package object

import "fmt"

// Option

type Option interface {
	set(*object)
}

type option func(*object)

func (o option) set(obj *object) {
	o(obj)
}

// Object
type ObjectAPI interface {
	View()
}

type object struct {
	mode string
	file string
}

func New(opts ...Option) (ObjectAPI, error) {
	o := new(object)

	for _, opt := range opts {
		opt.set(o)
	}
	return o, nil
}

func Mode(mode string) Option {
	return option(func(obj *object) {
		obj.mode = mode
	})
}

func WithFile(file string) Option {
	return option(func(obj *object) {
		obj.file = file
	})
}

func (o object) View() {
	fmt.Println(o.mode)
	fmt.Println(o.file)
}
