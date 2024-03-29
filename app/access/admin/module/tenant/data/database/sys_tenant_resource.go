package database

import (
	"context"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/infra/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
	mgoopts "go.mongodb.org/mongo-driver/mongo/options"
	"rockimserver/app/access/admin/module/tenant/biz/types"
	"rockimserver/app/access/admin/module/tenant/data/database/convert"
	"rockimserver/app/access/admin/module/tenant/data/database/entity"
)

type SysTenantResourceData struct {
	mgo *mongo.Client
}

func NewTenantResourceData(mgr *database.Manager) *SysTenantResourceData {
	return &SysTenantResourceData{mgo: mgr.Mongo()}
}

func (d *SysTenantResourceData) collection() *mgo.Collection {
	return d.mgo.Collection(entity.TableSysTenantResource)
}

func (d *SysTenantResourceData) FindByID(ctx context.Context, id string) (resource *types.SysTenantResource, err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var record entity.SysTenantResource
	err = d.mgo.FindOne(ctx, entity.TableSysTenantResource, bson.M{entity.MongoFieldId: objId}, &record, mgoopts.FindOne())
	if err != nil {
		return
	}
	return convert.SysTenantResourceProto(&record), nil
}

func (d *SysTenantResourceData) List(ctx context.Context) (results []*types.SysTenantResource, err error) {
	var records []*entity.SysTenantResource
	err = d.mgo.FindList(ctx, entity.TableSysTenantResource, bson.M{}, &records, mgoopts.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, convert.SysTenantResourceProto(record))
	}
	return
}
func (d *SysTenantResourceData) ListByIds(ctx context.Context, ids []string) (results []*types.SysTenantResource, err error) {
	var objIdList []primitive.ObjectID
	for _, id := range ids {
		var objId primitive.ObjectID
		objId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objIdList = append(objIdList, objId)
	}
	var records []*entity.SysTenantResource
	err = d.mgo.FindList(ctx, entity.TableSysTenantResource, bson.M{entity.MongoFieldId: bson.M{"$in": objIdList}}, &records, mgoopts.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, convert.SysTenantResourceProto(record))
	}
	return
}

func (d *SysTenantResourceData) Create(ctx context.Context, resource *types.SysTenantResource) (err error) {
	objId, err := primitive.ObjectIDFromHex(resource.Id)
	if err != nil {
		return
	}
	record := convert.SysTenantResourceEntity(resource)
	record.Id = objId
	if err != nil {
		return
	}
	_, err = d.collection().InsertOne(ctx, record)
	return
}

func (d *SysTenantResourceData) Update(ctx context.Context, resource *types.SysTenantResource) (err error) {
	objId, err := primitive.ObjectIDFromHex(resource.Id)
	if err != nil {
		return
	}
	record := convert.SysTenantResourceEntity(resource)
	_, err = d.collection().UpdateByID(ctx, objId, bson.M{"$set": record})
	return
}
func (d *SysTenantResourceData) Delete(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = d.collection().DeleteOne(ctx, bson.M{entity.MongoFieldId: objId})
	return
}
