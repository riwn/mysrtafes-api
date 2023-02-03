package challenge

type Repository interface {
	ChallengeCreate(*Challenge) (*Challenge, error)
	ChallengeRead(ID) (*Challenge, error)
	ChallengeUpdate(*Challenge) (*Challenge, error)
	ChallengeDelete(ID) error
}

type Server interface {
	Create(*Challenge) (*Challenge, error)
	Read(ID) (*Challenge, error)
	Update(*Challenge) (*Challenge, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(c *Challenge) (*Challenge, error) {
	// TODO: Validate
	return s.repository.ChallengeCreate(c)
}

func (s *server) Read(id ID) (*Challenge, error) {
	// TODO: Validate
	return s.repository.ChallengeRead(id)
}

func (s *server) Update(c *Challenge) (*Challenge, error) {
	// TODO: Validate
	return s.repository.ChallengeUpdate(c)
}

func (s *server) Delete(id ID) error {
	// TODO: Validate
	return s.repository.ChallengeDelete(id)
}
