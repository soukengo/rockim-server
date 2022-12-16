package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/logic/platform/data/database/convert"
	"rockim/app/logic/platform/data/database/entity"
	"rockim/pkg/component/database/mongo"
)

type PlatResourceData struct {
	mgo *mongo.Client
}

func NewPlatResourceData(mgo *mongo.Client) *PlatResourceData {
	return &PlatResourceData{mgo: mgo}
}

func (d *PlatResourceData) collection() *mgo.Collection {
	return d.mgo.Collection(entity.TablePlatResource)
}

func (d *PlatResourceData) FindByID(ctx context.Context, id string) (resource *types.PlatResource, err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var record entity.PlatResource
	err = d.mgo.FindOne(ctx, entity.TablePlatResource, bson.M{entity.MongoFieldId: objId}, &record, options.FindOne())
	if err != nil {
		return
	}
	return convert.ResourceProto(&record), nil
}
func (d *PlatResourceData) List(ctx context.Context) (results []*types.PlatResource, err error) {
	var records []*entity.PlatResource
	err = d.mgo.FindList(ctx, entity.TablePlatResource, bson.M{}, &records, options.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, convert.ResourceProto(record))
	}
	return
}
func (d *PlatResourceData) ListByIds(ctx context.Context, ids []string) (results []*types.PlatResource, err error) {
	var objIdList []primitive.ObjectID
	for _, id := range ids {
		var objId primitive.ObjectID
		objId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objIdList = append(objIdList, objId)
	}
	var records []*entity.PlatResource
	err = d.mgo.FindList(ctx, entity.TablePlatResource, bson.M{entity.MongoFieldId: bson.M{"$in": objIdList}}, &records, options.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, convert.ResourceProto(record))
	}
	return
}

func (d *PlatResourceData) Create(ctx context.Context, resource *types.PlatResource) (err error) {
	objId, err := primitive.ObjectIDFromHex(resource.Id)
	if err != nil {
		return
	}
	record := convert.ResourceEntity(resource)
	record.Id = objId
	if err != nil {
		return
	}
	_, err = d.collection().InsertOne(ctx, record)
	return
}

func (d *PlatResourceData) Update(ctx context.Context, resource *types.PlatResource) (err error) {
	objId, err := primitive.ObjectIDFromHex(resource.Id)
	if err != nil {
		return
	}
	record := convert.ResourceEntity(resource)
	_, err = d.collection().UpdateByID(ctx, objId, bson.M{"$set": record})
	return
}
func (d *PlatResourceData) Delete(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = d.collection().DeleteOne(ctx, bson.M{entity.MongoFieldId: objId})
	return
}
