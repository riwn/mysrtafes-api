package game

import (
	"mysrtafes-backend/pkg/errors"
	"mysrtafes-backend/pkg/game/platform"
	"mysrtafes-backend/pkg/game/tag"
)

type Repository interface {
	GameCreate(*Game, []platform.ID, []tag.ID) (*Game, error)
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
	// NameのValidate
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

	// PublisherのValidate
	if !g.Publisher.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("publisher", g.Publisher),
				},
			),
			"Publisher Valid error",
		)
	}

	// DeveloperのValidate
	if !g.Developer.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("developer", g.Developer),
				},
			),
			"Developer Valid error",
		)
	}

	for _, link := range g.Links {
		// link.TitleのValidate
		if !link.Title.Valid() {
			return nil, errors.NewInvalidRequest(
				errors.Layer_Domain,
				errors.NewInformation(
					errors.ID_InvalidParams,
					"",
					[]errors.InvalidParams{
						errors.NewInvalidParams("links.title", link.Title),
					},
				),
				"links.title Valid error",
			)
		}
		// link.DescriptionのValidate
		if !link.LinkDescription.Valid() {
			return nil, errors.NewInvalidRequest(
				errors.Layer_Domain,
				errors.NewInformation(
					errors.ID_InvalidParams,
					"",
					[]errors.InvalidParams{
						errors.NewInvalidParams("links.description", link.LinkDescription),
					},
				),
				"links.description Valid error",
			)
		}
	}
	return s.repository.GameCreate(g, platformIDs, tagIDs)
}

func (s *server) Read(id ID) (*Game, error) {
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
	return s.repository.GameRead(id)
}

func (s *server) Find(findOption *FindOption) ([]*Game, error) {
	return s.repository.GameFind(findOption)
}

func (s *server) Update(g *Game, platformIDs []platform.ID, tagIDs []tag.ID) (*Game, error) {
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
	// NameのValidate
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

	// PublisherのValidate
	if !g.Publisher.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("publisher", g.Publisher),
				},
			),
			"Publisher Valid error",
		)
	}

	// DeveloperのValidate
	if !g.Developer.Valid() {
		return nil, errors.NewInvalidRequest(
			errors.Layer_Domain,
			errors.NewInformation(
				errors.ID_InvalidParams,
				"",
				[]errors.InvalidParams{
					errors.NewInvalidParams("developer", g.Developer),
				},
			),
			"Developer Valid error",
		)
	}

	for _, link := range g.Links {
		// link.TitleのValidate
		if !link.Title.Valid() {
			return nil, errors.NewInvalidRequest(
				errors.Layer_Domain,
				errors.NewInformation(
					errors.ID_InvalidParams,
					"",
					[]errors.InvalidParams{
						errors.NewInvalidParams("links.title", link.Title),
					},
				),
				"links.title Valid error",
			)
		}
		// link.DescriptionのValidate
		if !link.LinkDescription.Valid() {
			return nil, errors.NewInvalidRequest(
				errors.Layer_Domain,
				errors.NewInformation(
					errors.ID_InvalidParams,
					"",
					[]errors.InvalidParams{
						errors.NewInvalidParams("links.description", link.LinkDescription),
					},
				),
				"links.description Valid error",
			)
		}
	}
	return s.repository.GameUpdate(g)
}

func (s *server) Delete(id ID) error {
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
	return s.repository.GameDelete(id)
}
