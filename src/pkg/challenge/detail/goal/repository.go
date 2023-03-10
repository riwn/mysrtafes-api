package goal

import "mysrtafes-backend/pkg/errors"

type Repository interface {
	GoalCreate(*Goal) (*Goal, error)
	GoalRead(ID) (*Goal, error)
	GoalFind(*FindOption) ([]*Goal, error)
	GoalUpdate(*Goal) (*Goal, error)
	GoalDelete(ID) error
}

type Server interface {
	Create(*Goal) (*Goal, error)
	Read(ID) (*Goal, error)
	Find(*FindOption) ([]*Goal, error)
	Update(*Goal) (*Goal, error)
	Delete(ID) error
}

type server struct {
	repository Repository
}

func NewServer(repo Repository) Server {
	return &server{repo}
}

func (s *server) Create(g *Goal) (*Goal, error) {
	// 名前のValidate
	if !g.Name.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("name", g.Name),
				},
			),
			"Name Valid error",
		)
	}
	// DescriptionのValidate
	if !g.Description.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("description", g.Description),
				},
			),
			"Description Valid error",
		)
	}
	return s.repository.GoalCreate(g)
}

func (s *server) Read(id ID) (*Goal, error) {
	// IDのValidate
	if !id.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("id", id),
				},
			),
			"ID Valid error",
		)
	}
	return s.repository.GoalRead(id)
}

func (s *server) Find(f *FindOption) ([]*Goal, error) {
	return s.repository.GoalFind(f)
}

func (s *server) Update(g *Goal) (*Goal, error) {
	// IDのValidate
	if !g.ID.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("id", g.ID),
				},
			),
			"ID Valid error",
		)
	}
	// 名前のValidate
	if !g.Name.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("name", g.Name),
				},
			),
			"Name Valid error",
		)
	}
	// DescriptionのValidate
	if !g.Description.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("description", g.Name),
				},
			),
			"Description Valid error",
		)
	}
	return s.repository.GoalUpdate(g)
}

func (s *server) Delete(id ID) error {
	// IDのValidate
	if !id.Valid() {
		return errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("id", id),
				},
			),
			"ID Valid error",
		)
	}
	return s.repository.GoalDelete(id)
}
