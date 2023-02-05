package game

import (
	"mysrtafes-backend/pkg/game/platform"
	"mysrtafes-backend/pkg/game/tag"
)

type Repository interface {
	GameCreate(*Game) (*Game, error)
	GameRead(ID) (*Game, error)
	GameFind(*FindOption) ([]*Game, error)
	GameUpdate(*Game) (*Game, error)
	GameDelete(ID) error
}

type Server interface {
	Create(*Game, []platform.ID, []tag.ID) (*Game, error)
	Read(ID) (*Game, error)
	Find(*FindOption) ([]*Game, error)
	Update(*Game, []platform.ID, []tag.ID) (*Game, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(g *Game, platformIDs []platform.ID, tagIDs []tag.ID) (*Game, error) {
	// TODO: Validate
	return s.repository.GameCreate(g)
}

func (s *server) Read(id ID) (*Game, error) {
	// TODO: Validate
	return s.repository.GameRead(id)
}

func (s *server) Find(findOption *FindOption) ([]*Game, error) {
	// TODO: Validate
	return s.repository.GameFind(findOption)
}

func (s *server) Update(g *Game, platformIDs []platform.ID, tagIDs []tag.ID) (*Game, error) {
	// TODO: Validate
	return s.repository.GameUpdate(g)
}

func (s *server) Delete(id ID) error {
	// TODO: Validate
	return s.repository.GameDelete(id)
}
