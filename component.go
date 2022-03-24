package gosong

type Component struct {
	Name string
}

func (s *Component) Setup(name string) {
	s.Name = name
}
