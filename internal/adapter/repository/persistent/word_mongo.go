package persistent

import (
	"context"

	"github.com/joqd/slovo/internal/adapter/repository/mapper"
	"github.com/joqd/slovo/internal/adapter/repository/model"
	"github.com/joqd/slovo/internal/core/domain"
	"github.com/joqd/slovo/internal/core/port"
	"github.com/joqd/slovo/pkg/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type wordPersistent struct {
	collection *mongo.Collection
	xlog       port.Logger
}

func NewWordRespository(mongodb mongodb.MongoDB, xlog port.Logger) port.WordPersistent {
	return &wordPersistent{
		collection: mongodb.Database.Collection("words"),
		xlog:       xlog,
	}
}

func (w *wordPersistent) GetByID(ctx context.Context, id string) (*domain.Word, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrInvalidObjectID
	}

	var wordDocument model.WordDocument

	filter := bson.D{{Key: "_id", Value: objectID}}
	if err := w.collection.FindOne(ctx, filter).Decode(&wordDocument); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrDataNotFound
		}

		w.xlog.Error("mongo get error, id=%s, err=%v", id, err)
		return nil, err
	}

	return wordDocument.ToDomain(), nil
}

func (w *wordPersistent) Create(ctx context.Context, word *domain.Word) (id string, err error) {
	wordPayload := mapper.WordToWordPayload(word)

	result, err := w.collection.InsertOne(ctx, wordPayload)
	if err != nil {
		w.xlog.Error("insert mongo error, err=%v", err)
		return "", err
	}

	oid, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		w.xlog.Error("get inserted id from mongo, err=%v", err)
		return "", err
	}

	return oid.Hex(), nil
}

func (w *wordPersistent) GetByBare(ctx context.Context, bare string) (*domain.Word, error) {
	var wordDocument model.WordDocument

	filter := bson.D{{Key: "bare", Value: bare}}
	if err := w.collection.FindOne(ctx, filter).Decode(&wordDocument); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrDataNotFound
		}

		return nil, err
	}

	return wordDocument.ToDomain(), nil
}

func (w *wordPersistent) DeleteByID(ctx context.Context, id string) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrInvalidObjectID
	}

	filter := bson.D{{Key: "_id", Value: objectID}}
	result, err := w.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return domain.ErrDataNotFound
	}

	return nil
}

func (w *wordPersistent) DeleteByBare(ctx context.Context, bare string) error {
	filter := bson.D{{Key: "bare", Value: bare}}
	result, err := w.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return domain.ErrDataNotFound
	}

	return nil
}

func (w *wordPersistent) Update(ctx context.Context, word *domain.Word) error {
	oid, err := primitive.ObjectIDFromHex(word.ID)
	if err != nil {
		w.xlog.Warn("failed to parse object_id: %s; err: %v", word.ID, err)
		return domain.ErrInvalidObjectID
	}

	wordPayload := mapper.WordToWordPayload(word)

	update := bson.M{
		"$set": bson.M{
			"bare":     wordPayload.Bare,
			"accented": wordPayload.Accented,
			"disable":  wordPayload.Disable,
			"type":     wordPayload.Type,
			"level":    wordPayload.Level,
		},
	}

	filter := bson.M{"_id": oid}

	result, err := w.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		w.xlog.Error("update mongo error, err=%v", err)
		return err
	}

	if result.MatchedCount == 0 {
		w.xlog.Warn("no document matched for update, id=%s", word.ID)
		return domain.ErrDataNotFound
	}

	return nil
}
