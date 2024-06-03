package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"job-service/internal/core"
	"time"
)

type JobRepository struct {
	collection *mongo.Collection
}

func NewJobRepository(collection *mongo.Collection) *JobRepository {
	return &JobRepository{collection: collection}
}

func (repository *JobRepository) GetAllJob(ctx context.Context, page int64) ([]*core.Job, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	jobChannel := make(chan []*core.Job)
	errorChan := make(chan error)

	go func() {
		errorChan <- repository.retrieveJob(ctxTimeout, page, jobChannel)
		close(errorChan)
	}()

	var opened bool
	jobs := make([]*core.Job, 0)

	select {
	case <-ctxTimeout.Done():
		fmt.Println("Processing timeout in Mongo")

		break
	case jobs, opened = <-jobChannel:
		if !opened {
			fmt.Println("chanel has been closed")
			break
		}

		fmt.Printf("get from db: %v\n", jobs)
		return jobs, nil
	case err := <-errorChan:
		fmt.Printf("Processing error in Mongo: %v\n", err)

		return nil, err
	}

	return jobs, nil
}

func (repository *JobRepository) retrieveJob(ctx context.Context, page int64, channel chan<- []*core.Job) (err error) {
	coll := repository.collection

	pageSize := int64(3)
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * pageSize)
	findOptions.SetLimit(pageSize)

	cursor, err := coll.Find(
		ctx,
		bson.D{},
		findOptions,
	)

	var results []*core.Job

	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	channel <- results
	close(channel)

	return nil
}
