package userstate

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/farwydi/stickfinbot/pkg/domain"
	"github.com/pkg/errors"
	"log"
	"os"
)

func NewGCPUseStateGateway(ctx context.Context) (domain.UserStatGateway, func(), error) {
	projID := os.Getenv("DATASTORE_PROJECT_ID")
	if projID == "" {
		return nil, nil, errors.New(`you need to set the environment variable "DATASTORE_PROJECT_ID"`)
	}

	dsClient, err := datastore.NewClient(ctx, projID)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "fail create client datastore")
	}

	cleanup := func() {
		if err := dsClient.Close(); err != nil {
			log.Printf("Fail close gcp datastore: %v", err)
		}
	}

	return &gcpDatastoreGateway{
		dsClient: dsClient,
	}, cleanup, nil
}

type gcpDatastoreGateway struct {
	dsClient *datastore.Client
}

func (g *gcpDatastoreGateway) GetCurrentState(ctx context.Context, userID int) (*domain.UserState, error) {
	var userState UserStat
	kind := "UserState"
	name := "state"
	taskKey := datastore.NameKey(kind, name, nil)
	err := g.dsClient.Get(ctx, taskKey, &userState)
	if err != nil {
		return nil, err
	}

	return toUserState(userState), nil
}
