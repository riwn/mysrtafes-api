package result

type Repository interface {
	ResultCreate(*Result) (*Result, error)
	ResultRead(ID) (*Result, error)
	ResultUpdate(*Result) (*Result, error)
	ResultDelete(ID) error
}

type Server interface {
	Create(*Result) (*Result, error)
	Read(ID) (*Result, error)
	Update(*Result) (*Result, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(l *Result) (*Result, error) {
	// TODO: Validate
	return s.repository.ResultCreate(l)
}

func (s *server) Read(id ID) (*Result, error) {
	// TODO: Validate
	return s.repository.ResultRead(id)
}

func (s *server) Update(l *Result) (*Result, error) {
	// TODO: Validate
	return s.repository.ResultUpdate(l)
}

func (s *server) Delete(id ID) error {
	// TODO: Validate
	return s.repository.ResultDelete(id)
}
