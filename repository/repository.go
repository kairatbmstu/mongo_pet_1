package repository

import (
	"context"
	"mongo_peg_1/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PageRepository handles operations related to Page model.
type PageRepository struct {
	collection *mongo.Collection
}

func NewPageRepository(database *mongo.Database) *PageRepository {
	return &PageRepository{
		collection: database.Collection("pages"),
	}
}

// Create inserts a new page into the database.
func (p *PageRepository) Create(page *model.Page) error {
	_, err := p.collection.InsertOne(context.Background(), page)
	return err
}

// Update updates an existing page in the database.
func (p *PageRepository) Update(page *model.Page) error {
	filter := bson.M{"_id": page.ID}
	update := bson.M{"$set": page}
	_, err := p.collection.UpdateOne(context.Background(), filter, update)
	return err
}

// FindByID finds a page by its ID.
func (p *PageRepository) FindByID(id string) (*model.Page, error) {
	var page model.Page
	filter := bson.M{"_id": id}
	err := p.collection.FindOne(context.Background(), filter).Decode(&page)
	if err != nil {
		return nil, err
	}
	return &page, nil
}

// FindByUrl finds a page by its URL.
func (p *PageRepository) FindByUrl(url string) (*model.Page, error) {
	var page model.Page
	filter := bson.M{"url": url}
	err := p.collection.FindOne(context.Background(), filter).Decode(&page)
	if err != nil {
		return nil, err
	}
	return &page, nil
}
