package goal

type Repository interface {
	GoalCreate(*Goal) (*Goal, error)
	GoalRead(ID) (*Goal, error)
	GoalUpdate(*Goal) (*Goal, error)
	GoalDelete(ID) error
}

type Server interface {
	Create(*Goal) (*Goal, error)
	Read(ID) (*Goal, error)
	Update(*Goal) (*Goal, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(l *Goal) (*Goal, error) {
	// TODO: Validate
	return s.repository.GoalCreate(l)
}

func (s *server) Read(id ID) (*Goal, error) {
	// TODO: Validate
	return s.repository.GoalRead(id)
}

func (s *server) Update(l *Goal) (*Goal, error) {
	// TODO: Validate
	return s.repository.GoalUpdate(l)
}

func (s *server) Delete(id ID) error {
	// TODO: Validate
	return s.repository.GoalDelete(id)
}
