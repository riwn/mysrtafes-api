package link

type Repository interface {
	LinkCreate(*Link) (*Link, error)
	LinkRead(ID) (*Link, error)
	LinkUpdate(*Link) (*Link, error)
	LinkDelete(ID) error
}

type Server interface {
	Create(*Link) (*Link, error)
	Read(ID) (*Link, error)
	Update(*Link) (*Link, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(l *Link) (*Link, error) {
	// TODO: Validate
	return s.repository.LinkCreate(l)
}

func (s *server) Read(id ID) (*Link, error) {
	// TODO: Validate
	return s.repository.LinkRead(id)
}

func (s *server) Update(l *Link) (*Link, error) {
	// TODO: Validate
	return s.repository.LinkUpdate(l)
}

func (s *server) Delete(id ID) error {
	// TODO: Validate
	return s.repository.LinkDelete(id)
}
