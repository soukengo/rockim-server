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

type TenantResourceData struct {
	mgo *mongo.Client
}

func NewTenantResourceData(mgo *mongo.Client) *TenantResourceData {
	return &TenantResourceData{mgo: mgo}
}

func (d *TenantResourceData) collection() *mgo.Collection {
	return d.mgo.Collection(entity.TableTenantResource)
}

func (d *TenantResourceData) FindByID(ctx context.Context, id string) (resource *types.TenantResource, err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var record entity.TenantResource
	err = d.mgo.FindOne(ctx, entity.TableTenantResource, bson.M{entity.MongoFieldId: objId}, &record, options.FindOne())
	if err != nil {
		return
	}
	return convert.TenantResourceProto(&record), nil
}
func (d *TenantResourceData) List(ctx context.Context) (results []*types.TenantResource, err error) {
	var records []*entity.TenantResource
	err = d.mgo.FindList(ctx, entity.TableTenantResource, bson.M{}, &records, options.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, convert.TenantResourceProto(record))
	}
	return
}
func (d *TenantResourceData) ListByIds(ctx context.Context, ids []string) (results []*types.TenantResource, err error) {
	var objIdList []primitive.ObjectID
	for _, id := range ids {
		var objId primitive.ObjectID
		objId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objIdList = append(objIdList, objId)
	}
	var records []*entity.TenantResource
	err = d.mgo.FindList(ctx, entity.TableTenantResource, bson.M{entity.MongoFieldId: bson.M{"$in": objIdList}}, &records, options.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, convert.TenantResourceProto(record))
	}
	return
}

func (d *TenantResourceData) Create(ctx context.Context, resource *types.TenantResource) (err error) {
	objId, err := primitive.ObjectIDFromHex(resource.Id)
	if err != nil {
		return
	}
	record := convert.TenantResourceEntity(resource)
	record.Id = objId
	if err != nil {
		return
	}
	_, err = d.collection().InsertOne(ctx, record)
	return
}

func (d *TenantResourceData) Update(ctx context.Context, resource *types.TenantResource) (err error) {
	objId, err := primitive.ObjectIDFromHex(resource.Id)
	if err != nil {
		return
	}
	record := convert.TenantResourceEntity(resource)
	_, err = d.collection().UpdateByID(ctx, objId, bson.M{"$set": record})
	return
}
func (d *TenantResourceData) Delete(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = d.collection().DeleteOne(ctx, bson.M{entity.MongoFieldId: objId})
	return
}
