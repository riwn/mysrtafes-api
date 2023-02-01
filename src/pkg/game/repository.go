package game

type Repository interface {
	GameCreate(*Game) (*Game, error)
	GameRead(ID) (*Game, error)
	GameUpdate(*Game) (*Game, error)
	GameDelete(ID) error
}

type Server interface {
	Create(*Game) (*Game, error)
	Read(ID) (*Game, error)
	Update(*Game) (*Game, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(g *Game) (*Game, error) {
	// TODO: Validate
	return s.repository.GameCreate(g)
}

func (s *server) Read(id ID) (*Game, error) {
	// TODO: Validate
	return s.repository.GameRead(id)
}

func (s *server) Update(g *Game) (*Game, error) {
	// TODO: Validate
	return s.repository.GameUpdate(g)
}

func (s *server) Delete(id ID) error {
	// TODO: Validate
	return s.repository.GameDelete(id)
}
