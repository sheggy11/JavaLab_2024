package mongo

import (
	"context"
	"cv-service/internal/core"
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

func (repository *CVRepository) GetAllCVBySkills(ctx context.Context, skills []string) ([]*core.CV, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	cvChannel := make(chan []*core.CV)
	errorChan := make(chan error)

	go func() {
		errorChan <- repository.retrieveCV(ctxTimeout, skills, cvChannel)
		close(errorChan)
	}()

	var opened bool
	cvs := make([]*core.CV, 0)

	select {
	case <-ctxTimeout.Done():
		fmt.Println("Processing timeout in Mongo")

		break
	case cvs, opened = <-cvChannel:
		if !opened {
			fmt.Println("chanel has been closed")

			break
		}

		fmt.Printf("get from db: %v\n", cvs)
		return cvs, nil
	case err := <-errorChan:
		fmt.Printf("Processing error in Mongo: %v\n", err)

		return nil, err
	}

	return cvs, nil
}

func (repository *CVRepository) retrieveCV(ctx context.Context, skills []string, channel chan<- []*core.CV) (err error) {
	coll := repository.collection
	aFilter := createBsonA(skills)
	filter := bson.D{{"skills", bson.D{{"$all", aFilter}}}}
	cursor, err := coll.Find(
		ctx,
		filter,
	)

	var results []*core.CV

	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	channel <- results
	close(channel)

	return nil
}

func createBsonA(strings []string) bson.A {
	res := make(bson.A, 0)
	for i := 0; i < len(strings); i++ {
		res = append(res, strings[i])
	}
	return res
}
