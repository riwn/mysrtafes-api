package detail

type Repository interface {
	DetailCreate(*Detail) (*Detail, error)
	DetailRead(ID) (*Detail, error)
	DetailUpdate(*Detail) (*Detail, error)
	DetailDelete(ID) error
}

type Server interface {
	Create(*Detail) (*Detail, error)
	Read(ID) (*Detail, error)
	Update(*Detail) (*Detail, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(l *Detail) (*Detail, error) {
	// TODO: Validate
	return s.repository.DetailCreate(l)
}

func (s *server) Read(id ID) (*Detail, error) {
	// TODO: Validate
	return s.repository.DetailRead(id)
}

func (s *server) Update(l *Detail) (*Detail, error) {
	// TODO: Validate
	return s.repository.DetailUpdate(l)
}

func (s *server) Delete(id ID) error {
	// TODO: Validate
	return s.repository.DetailDelete(id)
}
