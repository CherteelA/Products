package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	Id          uuid.UUID `bson:"_id,omitempty" json:"id"`
	Name        string    `json:"name" binding:"required"`
	Tags        []string  `json:"tags" binding:"required"`
	TagCount    int       `json:"tag_count" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   string    `json:"created_at"`
}

var client *mongo.Client
var productCollection *mongo.Collection
var ctx = context.TODO()

func init() {
	// ЗАГРУЖАЕМ .env ФАЙЛ
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	mongoUser := os.Getenv("MONGO_USER")
	mongoPass := os.Getenv("MONGO_PASS")
	mongoHost := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB")

	// Проверяем, что все переменные установлены
	if mongoUser == "" || mongoPass == "" || mongoHost == "" || dbName == "" {
		log.Fatal("Missing MongoDB environment variables")
	}

	var err error
	// Добавляем базу данных в connection string
	connectionString := "mongodb://" + mongoUser + ":" + mongoPass + "@" + mongoHost + "/" + dbName + "?authSource=admin"

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	productCollection = client.Database(dbName).Collection("products")
	log.Println("Connected to MongoDB!")
}

func SaveProduct(p *Product) error {
	p.Id = uuid.New()
	p.CreatedAt = time.Now().Format("2006-01-02")

	// Вставляем документ в коллекцию products
	result, err := productCollection.InsertOne(ctx, p)
	if err != nil {
		return fmt.Errorf("failed to save product: %v", err)
	}

	log.Printf("Product saved successfully! ID: %v", result.InsertedID)
	return nil
}

func GetProductByID(id uuid.UUID) (*Product, error) {
	var product Product

	filter := bson.M{"_id": id}

	err := productCollection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("product not found")
		}
		return nil, fmt.Errorf("failed to get product: %v", err)
	}

	return &product, nil
}

func GetProductByName(name string) (*Product, error) {
	var product Product

	filter := bson.M{"name": name}

	err := productCollection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("product with name '%s' not found", name)
		}
		return nil, fmt.Errorf("failed to get product by name: %v", err)
	}

	return &product, nil
}
