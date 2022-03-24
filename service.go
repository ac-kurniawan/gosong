package gosong

type Service struct {
	Name string
}

func (s *Service) Setup(name string) {
	s.Name = name
}
