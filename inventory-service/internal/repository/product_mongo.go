package repository

import (
	"context"
	"inventory-service/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productMongoRepo struct {
	collection *mongo.Collection
}

func NewProductMongoRepo(col *mongo.Collection) ProductRepository {
	return &productMongoRepo{collection: col}
}

func (r *productMongoRepo) Create(ctx context.Context, product *entity.Product) error {
	product.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, product)
	return err
}

func (r *productMongoRepo) GetAll(ctx context.Context, filter map[string]interface{}, limit, offset int64) ([]entity.Product, error) {
	mongoFilter := bson.M{}

	if category, ok := filter["category"]; ok {
		mongoFilter["category"] = category
	}
	if minPrice, ok := filter["minPrice"]; ok {
		mongoFilter["price"] = bson.M{"$gte": minPrice}
	}
	if maxPrice, ok := filter["maxPrice"]; ok {
		if p, ok := mongoFilter["price"].(bson.M); ok {
			p["$lte"] = maxPrice
		} else {
			mongoFilter["price"] = bson.M{"$lte": maxPrice}
		}
	}

	opts := options.Find().SetLimit(limit).SetSkip(offset)

	cursor, err := r.collection.Find(ctx, mongoFilter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []entity.Product
	for cursor.Next(ctx) {
		var p entity.Product
		if err := cursor.Decode(&p); err == nil {
			products = append(products, p)
		}
	}

	return products, nil
}

func (r *productMongoRepo) GetByID(ctx context.Context, id primitive.ObjectID) (*entity.Product, error) {
	var product entity.Product
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	return &product, err
}

func (r *productMongoRepo) Update(ctx context.Context, id primitive.ObjectID, product *entity.Product) error {
	update := bson.M{
		"$set": bson.M{
			"name":     product.Name,
			"category": product.Category,
			"stock":    product.Stock,
			"price":    product.Price,
		},
	}
	_, err := r.collection.UpdateByID(ctx, id, update)
	return err
}

func (r *productMongoRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
