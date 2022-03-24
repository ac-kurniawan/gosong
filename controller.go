package gosong

type Controller struct {
	Name string
}

func (s *Controller) Setup(name string) {
	s.Name = name
}
