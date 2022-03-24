package gosong

type Repository struct {
	Name string
}

func (s *Repository) Setup(name string) {
	s.Name = name
}
