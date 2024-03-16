package repo

import (
	"context"
	"errors"
	"log"

	"github.com/ibad69/go.fields/pkg/players"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type playerRepo struct {
	Db *mongo.Client
}

func New(conn *mongo.Client) players.Repo {
	return &playerRepo{conn}
}

func (pr *playerRepo) Create(ctx context.Context, p players.Player) (players.Player, error) {
	// Insert a new player into the players collection
	res, err := pr.Db.Database("test").Collection("players").InsertOne(ctx, p)
	if err != nil {
		return players.Player{}, err
	}

	// Get the inserted ID
	id := res.InsertedID.(primitive.ObjectID)

	// Find the inserted player
	var player players.Player
	err = pr.Db.Database("go").Collection("players").FindOne(ctx, bson.M{"_id": id}).Decode(&player)
	if err != nil {
		return players.Player{}, err
	}

	// Return the inserted player
	return player, nil
}
func (pr *playerRepo) SignUp(ctx context.Context, p players.Player) (players.Player, error) {
	var player players.Player
	err := pr.Db.Database("golang").Collection("players").FindOne(ctx, bson.M{"email": p.Email}).Decode(&p)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Document not found, handle accordingly
		} else {
			// Other error, handle accordingly
			return players.Player{}, err
		}
	}
	log.Println(&p)
	log.Println(p.Email)
	if p.Email != "" {
		return players.Player{}, errors.New("email already exists")
	}
	res, err := pr.Db.Database("golang").Collection("players").InsertOne(ctx, p)
	if err != nil {
		return players.Player{}, err
	}
	id := res.InsertedID.(primitive.ObjectID)

	err = pr.Db.Database("golang").Collection("players").FindOne(ctx, bson.M{"_id": id}).Decode(&player)
	if err != nil {
		return players.Player{}, err
	}
	return player, nil
}

func (pr *playerRepo) Get(ctx context.Context, id string) ([]players.Player, error) {
	// Convert string to ObjectID
	// objID, err := primitive.ObjectIDFromHex(id)
	// if err != nil {jjj
	// 	return players.Player{}, err
	// }

	var players []players.Player

	// Find all players

	cursor, err := pr.Db.Database("golang").Collection("players").Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode the documents into the players slice
	if err = cursor.All(ctx, &players); err != nil {
		return nil, err
	}

	// Return the players
	return players, nil
}

func (pr *playerRepo) Update(ctx context.Context, p players.Player) (players.Player, error) {
	// Convert string to ObjectID
	objID, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		return players.Player{}, err
	}

	// Create a filter for the player to update
	filter := bson.M{"_id": objID}

	// Create an update for the player's data
	update := bson.M{
		"$set": p,
	}

	// Update the player
	_, err = pr.Db.Database("test").Collection("players").UpdateOne(ctx, filter, update)
	if err != nil {
		return players.Player{}, err
	}

	// Find the updated player
	var player players.Player
	err = pr.Db.Database("test").Collection("players").FindOne(ctx, filter).Decode(&player)
	if err != nil {
		return players.Player{}, err
	}

	// Return the updated player
	return player, nil
}

func (pr *playerRepo) Delete(ctx context.Context, id string) error {
	// create a player
	return nil

}
