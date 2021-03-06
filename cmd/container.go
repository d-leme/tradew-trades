package cmd

import (
	"context"

	"github.com/d-leme/tradew-trades/pkg/core"
	"github.com/d-leme/tradew-trades/pkg/trades"
	"github.com/d-leme/tradew-trades/pkg/trades/external/inventory"
	"github.com/d-leme/tradew-trades/pkg/trades/external/inventory/proto"
	"github.com/d-leme/tradew-trades/pkg/trades/mongodb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

// Container ...
type Container struct {
	Settings *core.Settings

	Authenticate *core.Authenticate

	MongoClient *mongo.Client

	InventoryServiceConnection *grpc.ClientConn
	InventoryService           inventory.Service

	TradeRepository trades.Repository
	TradeService    trades.Service
	TradeController trades.Controller
}

// NewContainer creates new instace of Container
func NewContainer(settings *core.Settings) *Container {

	container := new(Container)

	container.Settings = settings

	container.MongoClient = connectMongoDB(settings.MongoDB)

	container.Authenticate = core.NewAuthenticate(settings.JWT.Secret)

	// GRPC
	container.InventoryServiceConnection = connectGRPC(settings.InventoryService)
	container.InventoryService = proto.NewService(
		proto.NewInventoryServiceClient(container.InventoryServiceConnection),
	)

	// Trades
	container.TradeRepository = mongodb.NewRepository(container.MongoClient, settings.MongoDB.Database)
	container.TradeService = trades.NewService(container.TradeRepository, container.InventoryService)
	container.TradeController = trades.NewController(container.Authenticate, container.TradeService)

	return container
}

// Controllers maps all routes and exposes them
func (c *Container) Controllers() []core.Controller {
	return []core.Controller{
		&c.TradeController,
	}
}

// Close terminates every opened resource
func (c *Container) Close() {
	c.MongoClient.Disconnect(context.Background())
	c.InventoryServiceConnection.Close()
}

func connectMongoDB(conf *core.MongoDBConfig) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.ConnectionString))

	if err != nil {
		logrus.
			WithError(err).
			Fatal("error connecting to MongoDB")
	}

	client.Connect(context.Background())

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		logrus.
			WithError(err).
			Fatal("error pinging MongoDB")
	}

	logrus.Info("connected to mongodb")

	return client
}

func connectGRPC(srv *core.GRPCService) *grpc.ClientConn {
	conn, err := grpc.Dial(srv.URL, grpc.WithInsecure())
	if err != nil {
		logrus.
			WithError(err).
			Fatal("error pinging connecting to GRPC endpoint")
	}

	return conn

}
