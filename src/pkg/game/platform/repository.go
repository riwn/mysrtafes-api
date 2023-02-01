package platform

type Repository interface {
	PlatformCreate(*Platform) (*Platform, error)
	PlatformRead(ID) (*Platform, error)
	PlatformUpdate(*Platform) (*Platform, error)
	PlatformDelete(ID) error
}

type Server interface {
	Create(*Platform) (*Platform, error)
	Read(ID) (*Platform, error)
	Update(*Platform) (*Platform, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(p *Platform) (*Platform, error) {
	// TODO: Validate
	return s.repository.PlatformCreate(p)
}

func (s *server) Read(id ID) (*Platform, error) {
	// TODO: Validate
	return s.repository.PlatformRead(id)
}

func (s *server) Update(p *Platform) (*Platform, error) {
	// TODO: Validate
	return s.repository.PlatformUpdate(p)
}

func (s *server) Delete(id ID) error {
	// TODO: Validate
	return s.repository.PlatformDelete(id)
}
