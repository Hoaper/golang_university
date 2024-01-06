package repositories

import (
	"context"
	"github.com/Hoaper/golang_university/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName          = "mongo_university"
	usersCollection = "users"
)

type UserRepository struct {
	Client *mongo.Client
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{Client: client}
}
func (r *UserRepository) CreateUser(user *models.User) error {
	collection := r.Client.Database(dbName).Collection(usersCollection)
	_, err := collection.InsertOne(context.Background(), models.User{Role: "student", Login: user.Login, Password: user.Password})
	return err
}

func (r *UserRepository) GetUserByLogin(login string) (*models.User, error) {
	collection := r.Client.Database(dbName).Collection(usersCollection)

	filter := bson.D{{"login", login}}
	user := &models.User{}

	err := collection.FindOne(context.Background(), filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	collection := r.Client.Database(dbName).Collection(usersCollection)
	filter := bson.D{{"login", user.Login}}

	update := bson.D{{"$set", bson.D{
		{"password", user.Password},
		{"role", user.Role},
		{"issuances", user.Issuances},
	}}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}
