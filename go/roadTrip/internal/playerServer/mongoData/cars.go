package mongoData

import (
	context "context"
	"errors"
	. "github.com/brickshot/roadtrip/internal/playerServer"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

var carsColl *mongo.Collection

// init
func InitCars() {
	carsColl = database.Collection("cars")
}

func ShutdownCars() {
	carsColl = nil
}

// CreateCar stores a car in the datastore with the given owner and location. The id and plate will be assigned.
func (d MongoProvider) CreateCar(c Car, owner Character, location Location) (Car, error) {
	// find an unused plate number
	var plate string
	for i := 0; i < 10; i++ {
		p1, _ := gonanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 3)
		p2, _ := gonanoid.Generate("0123456789", 3)
		plate = p1 + "-" + p2
		_, err := d.GetCarByPlate(plate)
		if err == nil {
			continue
		}
		break
	}
	if plate == "" {
		return Car{}, errors.New("Could not find un-used plate in 10 attempts.")
	}

	c.Id = uuid.NewString()
	c.Plate = plate
	c.OwnerId = owner.Id
	c.Location = &location
	c.VelocityMph = 0
	c.OdometerMiles = 0

	var ctx context.Context
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	m, err := bson.Marshal(c)
	if err != nil {
		return Car{}, err
	}

	result, err := carsColl.InsertOne(ctx, m)
	if err != nil {
		return Car{}, err
	}

	newCar := Car{}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{"_id", result.InsertedID}}
	err = carsColl.FindOne(ctx, filter).Decode(&newCar)
	if err != nil {
		return Car{}, err
	}
	return newCar, nil
}

// GetCar returns the car referenced by id
func (d MongoProvider) GetCar(id string) (Car, error) {
	car := Car{}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := carsColl.FindOne(ctx, bson.D{{"id", id}}).Decode(&car)
	if err == mongo.ErrNoDocuments {
		return Car{}, errors.New("not found")
	} else if err != nil {
		log.Fatal(err)
	}
	return car, nil
}

// GetCarByPlate returns the car referenced by Plate
func (d MongoProvider) GetCarByPlate(plate string) (Car, error) {
	return Car{}, nil
}

// GetCharacters returns the characters in a car.
func (d MongoProvider) GetCharacters(id string) ([]Character, error) {
	return []Character{}, nil
}

func (d MongoProvider) GetCars() ([]Car, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := carsColl.Find(ctx, bson.D{})
	if err != nil {
		return []Car{}, err
	}
	defer cur.Close(context.Background())

	var results []Car
	if err = cur.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}
	return results, nil
}

// UpdateCar updates a car in the datastore.
func (d MongoProvider) UpdateCar(c Car) error {
	var ctx context.Context
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	m, err := bson.Marshal(c)
	if err != nil {
		return err
	}
	_, err = carsColl.ReplaceOne(
		ctx,
		bson.M{"id": c.Id},
		m)
	if err != nil {
		return err
	}
	return nil
}
