package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"rockim/pkg/errors"
)

var (
	reason           = "DB_ERROR"
	DocumentNotFound = errors.NotFound("DATA_NOT_FOUND", "data not found")
)

func (c *Client) FindOne(ctx context.Context, collection string, filter any, dist any, opts ...*options.FindOneOptions) (err error) {
	err = c.Collection(collection).FindOne(ctx, filter, opts...).Decode(dist)
	if IsErrNoDocuments(err) {
		err = DocumentNotFound
	}
	return
}

func (c *Client) FindList(ctx context.Context, collection string, filter any, result any, opts ...*options.FindOptions) (err error) {
	cursor, err := c.Collection(collection).Find(ctx, filter, opts...)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	return c.scanCursor(ctx, cursor, result)
}
func (c *Client) scanCursor(ctx context.Context, cursor *mongo.Cursor, result interface{}) (err error) {
	resultv := reflect.ValueOf(result)
	if resultv.Kind() != reflect.Ptr {
		return errors.InternalServer(reason, "result argument must be a slice address")
	}

	slicev := resultv.Elem()

	if slicev.Kind() == reflect.Interface {
		slicev = slicev.Elem()
	}
	if slicev.Kind() != reflect.Slice {
		return errors.InternalServer(reason, "result argument must be a slice address")
	}

	slicev = slicev.Slice(0, slicev.Cap())
	elemt := slicev.Type().Elem()
	i := 0
	for cursor.Next(ctx) {
		if slicev.Len() == i {
			elemp := reflect.New(elemt)
			if err = cursor.Decode(elemp.Interface()); err != nil {
				break
			}
			slicev = reflect.Append(slicev, elemp.Elem())
			slicev = slicev.Slice(0, slicev.Cap())
		} else {
			if err = cursor.Decode(slicev.Index(i).Addr().Interface()); err != nil {
				break
			}
		}
		i++
	}
	resultv.Elem().Set(slicev.Slice(0, i))
	return
}
