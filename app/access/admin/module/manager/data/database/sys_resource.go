package database

import (
	"context"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/infra/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
	mgoOpts "go.mongodb.org/mongo-driver/mongo/options"
	"rockimserver/app/access/admin/module/manager/biz/types"
	"rockimserver/app/access/admin/module/manager/data/database/convert"
	"rockimserver/app/access/admin/module/manager/data/database/entity"
)

type SysResourceData struct {
	mgo *mongo.Client
}

func NewSysResourceData(mgr *database.Manager) *SysResourceData {
	return &SysResourceData{mgo: mgr.Mongo()}
}

func (d *SysResourceData) collection() *mgo.Collection {
	return d.mgo.Collection(entity.TableSysResource)
}

func (d *SysResourceData) FindByID(ctx context.Context, id string) (resource *types.SysResource, err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var record entity.SysResource
	err = d.mgo.FindOne(ctx, entity.TableSysResource, bson.M{entity.MongoFieldId: objId}, &record, mgoOpts.FindOne())
	if err != nil {
		return
	}
	return convert.ResourceProto(&record), nil
}
func (d *SysResourceData) List(ctx context.Context) (results []*types.SysResource, err error) {
	var records []*entity.SysResource
	err = d.mgo.FindList(ctx, entity.TableSysResource, bson.M{}, &records, mgoOpts.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, convert.ResourceProto(record))
	}
	return
}
func (d *SysResourceData) ListByIds(ctx context.Context, ids []string) (results []*types.SysResource, err error) {
	var objIdList []primitive.ObjectID
	for _, id := range ids {
		var objId primitive.ObjectID
		objId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objIdList = append(objIdList, objId)
	}
	var records []*entity.SysResource
	err = d.mgo.FindList(ctx, entity.TableSysResource, bson.M{entity.MongoFieldId: bson.M{"$in": objIdList}}, &records, mgoOpts.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, convert.ResourceProto(record))
	}
	return
}

func (d *SysResourceData) Create(ctx context.Context, resource *types.SysResource) (err error) {
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

func (d *SysResourceData) Update(ctx context.Context, resource *types.SysResource) (err error) {
	objId, err := primitive.ObjectIDFromHex(resource.Id)
	if err != nil {
		return
	}
	record := convert.ResourceEntity(resource)
	_, err = d.collection().UpdateByID(ctx, objId, bson.M{"$set": record})
	return
}
func (d *SysResourceData) Delete(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = d.collection().DeleteOne(ctx, bson.M{entity.MongoFieldId: objId})
	return
}
