package repositories

import (
	"context"
	"errors"
	"gin-mongodb-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserRepository(userCollection *mongo.Collection, ctx context.Context) UserRepository {
	return &UserRepositoryImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserRepositoryImpl) CreateUser(user *models.User) error {

	_, err := u.userCollection.InsertOne(u.ctx, user)

	return err

}

func (u *UserRepositoryImpl) GetUser(id *string) (*models.User, error) {

	var user *models.User

	query := bson.D{bson.E{Key: "id", Value: id}}

	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)

	return user, err
}

func (u *UserRepositoryImpl) UpdateUser(user *models.User) error {

	filter := bson.D{bson.E{Key: "id", Value: user.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "name", Value: user.Name},
		bson.E{Key: "email", Value: user.Email}}}}

	result, _ := u.userCollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount < 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *UserRepositoryImpl) GetAll() ([]*models.User, error) {

	var users []*models.User

	cursor, err := u.userCollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) < 1 {
		return nil, errors.New("documents not found")
	}

	return users, nil
}

func (u *UserRepositoryImpl) DeleteUser(id *string) error {

	filter := bson.D{bson.E{Key: "id", Value: id}}

	result, _ := u.userCollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount < 1 {
		return errors.New("no matched document found for delete")
	}

	return nil
}
