package database

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "job-service/internal/config"
)

func SetupMongoDatabase(ctx context.Context, config *config.MongoConfig) (*mongo.Database, error) {
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Database.URI))

    if err != nil {
        return nil, fmt.Errorf("connect to mongo db: %w", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        return nil, fmt.Errorf("ping mongo db: %w", err)
    }

    return client.Database(config.Database.Name), nil
}
