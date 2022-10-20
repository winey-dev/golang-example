package argument

type Arguments struct {
	Loc       string
	Read      bool
	ReadQuery bool
	ReadFile  string
	Write     bool
	W_field   bool
	W_tag     bool
}

func (arg *Arguments) ValidCheck() bool {

	return true
}
