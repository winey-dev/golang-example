package object

import "fmt"

// Option

type Option func(*object)

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
		opt(o)
	}
	return o, nil
}

func Mode(mode string) Option {
	return Option(func(obj *object) {
		obj.mode = mode
	})
}

func WithFile(file string) Option {
	return Option(func(obj *object) {
		obj.file = file
	})
}

func (o object) View() {
	fmt.Println(o.mode)
	fmt.Println(o.file)
}
