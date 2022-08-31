package mongo

import (
	"context"

	"github.com/reearth/reearth/server/internal/infrastructure/mongo/migration"
	"github.com/reearth/reearth/server/internal/usecase/repo"
	"github.com/reearth/reearth/server/pkg/scene"
	"github.com/reearth/reearth/server/pkg/user"
	"github.com/reearth/reearthx/mongox"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(ctx context.Context, mc *mongo.Client, databaseName string) (*repo.Container, error) {
	if databaseName == "" {
		databaseName = "reearth"
	}

	db := mc.Database(databaseName)
	lock, err := NewLock(db.Collection("locks"))
	if err != nil {
		return nil, err
	}

	client := mongox.NewClientWithDatabase(db)
	c := &repo.Container{
		Asset:          NewAsset(client),
		AuthRequest:    NewAuthRequest(client),
		Config:         NewConfig(db.Collection("config"), lock),
		DatasetSchema:  NewDatasetSchema(client),
		Dataset:        NewDataset(client),
		Layer:          NewLayer(client),
		Plugin:         NewPlugin(client),
		Project:        NewProject(client),
		PropertySchema: NewPropertySchema(client),
		Property:       NewProperty(client),
		Scene:          NewScene(client),
		Tag:            NewTag(client),
		Workspace:      NewWorkspace(client),
		User:           NewUser(client),
		SceneLock:      NewSceneLock(client),
		Transaction:    mongox.NewTransaction(client),
		Policy:         NewPolicy(client),
		Lock:           lock,
	}

	// migration
	m := migration.Client{Client: client, Config: c.Config}
	if err := m.Migrate(ctx); err != nil {
		return nil, err
	}

	return c, nil
}

func applyWorkspaceFilter(filter interface{}, ids user.WorkspaceIDList) interface{} {
	if ids == nil {
		return filter
	}
	return mongox.And(filter, "team", bson.M{"$in": ids.Strings()})
}

func applySceneFilter(filter interface{}, ids scene.IDList) interface{} {
	if ids == nil {
		return filter
	}
	return mongox.And(filter, "scene", bson.M{"$in": ids.Strings()})
}

func applyOptionalSceneFilter(filter interface{}, ids scene.IDList) interface{} {
	if ids == nil {
		return filter
	}
	return mongox.And(filter, "", bson.M{"$or": []bson.M{
		{"scene": bson.M{"$in": ids.Strings()}},
		{"scene": nil},
		{"scene": ""},
	}})
}