package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rockim/pkg/errors"
)

var (
	DocumentNotFound = errors.NotFound("DATA_NOT_FOUND", "data not found")
)

func (c *Client) FindOne(ctx context.Context, collection string, filter any, dist any, opt *options.FindOneOptions) (err error) {
	err = c.Collection(collection).FindOne(ctx, filter, opt).Decode(dist)
	if IsErrNoDocuments(err) {
		err = DocumentNotFound
	}
	return
}
