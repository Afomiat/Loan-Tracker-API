package repositories

import (
    "context"
    "github.com/Afomiat/Loan-Tracker-API/domain"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
    collection *mongo.Collection
}

func NewUserRepository(uri, dbName string) *UserRepository {
    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        panic(err)
    }

    collection := client.Database(dbName).Collection("users")
    return &UserRepository{collection: collection}
}

func (r *UserRepository) Create(user *domain.User) error {
    _, err := r.collection.InsertOne(context.TODO(), user)
    return err
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
    var user domain.User
    err := r.collection.FindOne(context.TODO(), map[string]string{"email": email}).Decode(&user)
    return &user, err
}

func (r *UserRepository) FindByID(id string) (*domain.User, error) {
    var user domain.User
    err := r.collection.FindOne(context.TODO(), map[string]string{"id": id}).Decode(&user)
    return &user, err
}

func (r *UserRepository) VerifyUser(email, token string) error {
    // For simplicity, assume token verification is handled elsewhere
    _, err := r.collection.UpdateOne(
        context.TODO(),
        map[string]string{"email": email},
        map[string]interface{}{"$set": map[string]bool{"is_verified": true}},
    )
    return err
}
