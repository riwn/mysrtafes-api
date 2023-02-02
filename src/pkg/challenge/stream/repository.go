package stream

type Repository interface {
	StatusCreate(*Status) (*Status, error)
	StatusRead(ID) (*Status, error)
	StatusUpdate(*Status) (*Status, error)
	StatusDelete(ID) error
}

type Server interface {
	Create(*Status) (*Status, error)
	Read(ID) (*Status, error)
	Update(*Status) (*Status, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(l *Status) (*Status, error) {
	// TODO: Validate
	return s.repository.StatusCreate(l)
}

func (s *server) Read(id ID) (*Status, error) {
	// TODO: Validate
	return s.repository.StatusRead(id)
}

func (s *server) Update(l *Status) (*Status, error) {
	// TODO: Validate
	return s.repository.StatusUpdate(l)
}

func (s *server) Delete(id ID) error {
	// TODO: Validate
	return s.repository.StatusDelete(id)
}
