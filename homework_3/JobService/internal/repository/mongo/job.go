package mongo

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "job-service/internal/core"
    "time"
)

type JobRepository struct {
    collection *mongo.Collection
}

func NewJobRepository(collection *mongo.Collection) *JobRepository {
    return &JobRepository{collection: collection}
}

func (repository *JobRepository) GetAll(ctx context.Context) ([]*core.Job, error) {
    ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    jobChannel := make(chan []*core.Job)
    errChannel := make(chan error)

    go func() {
        errChannel <- repository.retrieveUsers(ctxTimeout, jobChannel)
        close(errChannel)
    }()

    var jobs []*core.Job
    select {
    case <-ctxTimeout.Done():
        fmt.Println("Processing timeout in Mongo")
        break
    case jobs = <-jobChannel:
        fmt.Println("Finished processing in Mongo")
    case err := <-errChannel:
        fmt.Println("Processing error in Mongo")
        return nil, err
    }

    return jobs, nil
}

func (repository *JobRepository) retrieveUsers(ctx context.Context, channel chan<- []*core.Job) error {
    cur, err := repository.collection.Find(ctx, bson.M{})
    if err != nil {
        return fmt.Errorf("retrieve job from Mongo: %w", err)
    }

    var res []*core.Job
    if err := cur.All(ctx, &res); err != nil {
        return fmt.Errorf("processing all jobs in Mongo: %w", err)
    }

    channel <- res
    close(channel)

    return nil
}
