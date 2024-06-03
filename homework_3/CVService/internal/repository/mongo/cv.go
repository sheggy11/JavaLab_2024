package mongo

import (
	"CVService/internal/core"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CVRepository struct {
	collection *mongo.Collection
}

func NewCVRepository(collection *mongo.Collection) *CVRepository {
	return &CVRepository{collection: collection}
}

func (repository *CVRepository) GetAllCV(ctx context.Context) ([]*core.CV, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	cvChannel := make(chan []*core.CV)
	var err error

	go func() {
		err = repository.retrieveAll(ctx, cvChannel)
	}()

	if err != nil {
		return nil, err
	}

	var cvs []*core.CV

	select {
	case <-ctxTimeout.Done():
		fmt.Println("Processing timeout in Mongo")
		break
	case cvs = <-cvChannel:
		fmt.Println("Finished processing in Mongo")
	}

	return cvs, nil
}

func (repository *CVRepository) retrieveAll(ctx context.Context, channel chan<- []*core.CV) (err error) {

	var cvs []*core.CV

	cursor, err := repository.collection.Find(ctx, bson.M{})

	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &cvs); err != nil {
		return err
	}

	channel <- cvs

	return nil
}

func (repository *CVRepository) GetByActivityType(ctx context.Context, activityType string) (*core.CV, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	cvChannel := make(chan *core.CV)
	var err error

	go func() {
		err = repository.retrieveCV(ctx, activityType, cvChannel)
	}()

	if err != nil {
		return nil, err
	}

	var cv *core.CV

	select {
	case <-ctxTimeout.Done():
		fmt.Println("Processing timeout in Mongo")
		break
	case cv = <-cvChannel:
		fmt.Println("Finished processing in Mongo")
	}

	return cv, nil
}

func (repository *CVRepository) retrieveCV(ctx context.Context, activityType string, channel chan<- *core.CV) (err error) {

	cv := &core.CV{}


	err = repository.collection.FindOne(ctx, bson.M{"activityType": activityType}).Decode(cv)

	if err != nil {
		return err
	}

	channel <- cv

	return nil
}
