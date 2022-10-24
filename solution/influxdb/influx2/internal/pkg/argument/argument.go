package argument

type Arguments struct {
	Read      bool
	ReadQuery bool
	ReadFile  string
	Write     bool
	W_field   bool
	W_tag     bool
	W_stat    bool
}

func (arg *Arguments) ValidCheck() bool {

	return true
}
