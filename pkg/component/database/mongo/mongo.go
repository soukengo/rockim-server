package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Client struct {
	config *Config
	client *mongo.Client
}

func NewClient(conf *Config) *Client {
	c := new(Client)
	c.config = conf
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(c.config.Address)
	if len(c.config.Username) > 0 {
		credential := options.Credential{AuthSource: c.config.AuthSource, Username: c.config.Username, Password: c.config.Password}
		clientOptions = clientOptions.SetAuth(credential)
	}
	timeout := 10 * time.Second
	clientOptions.ConnectTimeout = &timeout
	m := newMonitor()
	m.WithLog()
	m.WithTrace()
	clientOptions.Monitor = m.Entry()
	// 连接到MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	// 检查连接
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	c.client = client
	return c
}

func (c *Client) Close() error {
	return c.client.Disconnect(context.Background())
}

func (c *Client) Database(database string, opts ...*options.DatabaseOptions) *mongo.Database {
	return c.client.Database(database, opts...)
}
func (c *Client) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return c.Database(c.config.Database).Collection(name, opts...)
}

func (c *Client) Connect(ctx context.Context) error {
	return c.client.Connect(ctx)
}
func IsErrNoDocuments(source error) bool {
	return errors.Is(source, mongo.ErrNoDocuments)
}
