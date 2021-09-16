package mongoData

import (
	"context"
	"errors"
	. "github.com/brickshot/roadtrip/internal/playerServer"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

var charactersColl *mongo.Collection

// init
func InitCharacters() {
	charactersColl = database.Collection("characters")
}

func ShutdownCharacters() {
	charactersColl = nil
}

// CreateCharacter
func (d MongoProvider) CreateCharacter(name string) (Character, error) {
	id := uuid.NewString()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := charactersColl.InsertOne(ctx, bson.D{{"id", id}, {"name", name}, {"deleted", false}})
	if err != nil {
		return Character{}, err
	}
	return Character{
		Id:   id,
		Name: name,
		Car:  nil,
	}, nil
}

// GetCharacter returns the Character from the database with its Car populated
func (d MongoProvider) GetCharacter(id string) (Character, error) {
	char := Character{}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := charactersColl.FindOne(ctx, bson.D{{"id", id}, {"deleted", false}}).Decode(&char)
	if err == mongo.ErrNoDocuments {
		return Character{}, errors.New("not found")
	} else if err != nil {
		log.Fatal(err)
	}

	// Character's car is a singleton so we always return it with the character
	car := Car{}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = carsColl.FindOne(ctx, bson.D{{"owner_id", id}}).Decode(&car)
	if err == mongo.ErrNoDocuments {
		return Character{}, errors.New("missing car")
	} else if err != nil {
		log.Fatal(err)
	}
	char.Car = &car

	return char, nil
}

// DeleteCharacter
func (d MongoProvider) DeleteCharacter(id string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := charactersColl.UpdateOne(ctx, bson.D{{"id", id}}, bson.D{{"deleted", true}})
	return err
}

// SetCar
func (d MongoProvider) SetCar(charId string, carUUID string) (Character, error) {
	return Character{}, nil
}
