package tag

type Repository interface {
	TagCreate(*Tag) (*Tag, error)
	TagRead(ID) (*Tag, error)
	TagUpdate(*Tag) (*Tag, error)
	TagDelete(ID) error
}

type Server interface {
	Create(*Tag) (*Tag, error)
	Read(ID) (*Tag, error)
	Update(*Tag) (*Tag, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(t *Tag) (*Tag, error) {
	// TODO: Validate
	return s.repository.TagCreate(t)
}

func (s *server) Read(id ID) (*Tag, error) {
	// TODO: Validate
	return s.repository.TagRead(id)
}

func (s *server) Update(t *Tag) (*Tag, error) {
	// TODO: Validate
	return s.repository.TagUpdate(t)
}

func (s *server) Delete(id ID) error {
	// TODO: Validate
	return s.repository.TagDelete(id)
}
