package players

import (
	"context"
)

// creating a player service interface this is useful to interact with services that are outside
type Service interface {
	Get(ctx context.Context, id string) ([]Player, error)
	Create(ctx context.Context, p Player) (Player, error)
	Update(ctx context.Context, p Player) (Player, error)
	Delete(ctx context.Context, id string) error
	SignUp(ctx context.Context, p Player) (Player, error)
}

// this interface is useful in interacting directly to db only the above service will tell this interface
// to contact with db
type Repo interface {
	Get(ctx context.Context, id string) ([]Player, error)
	Create(ctx context.Context, p Player) (Player, error)
	Update(ctx context.Context, p Player) (Player, error)
	Delete(ctx context.Context, id string) error
	SignUp(ctx context.Context, p Player) (Player, error)
}

// initializing the player service
type player struct {
	repo Repo
}

// function for creating a new service for interaction
func New(repo Repo) Service {
	return &player{repo}
}

func (*player) Create(ctx context.Context, p Player) (Player, error) {
	// this will call the repo methods, repo method will finally interact with the db iteself
	// p.repo.create()
	return p, nil
}

func (p *player) Get(ctx context.Context, id string) ([]Player, error) {
	return p.repo.Get(ctx, id)
}

func (*player) Update(ctx context.Context, p Player) (Player, error) {
	return p, nil
}

func (*player) Delete(ctx context.Context, id string) error {
	return nil
}

func (p *player) SignUp(ctx context.Context, pl Player) (Player, error) {
	pla, err := p.repo.SignUp(ctx, pl)
	if err != nil {
		return Player{}, err
	}
	return pla, nil
}
