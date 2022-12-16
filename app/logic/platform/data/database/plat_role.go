package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/logic/platform/biz"
	"rockim/app/logic/platform/data/database/convert"
	"rockim/app/logic/platform/data/database/entity"
	"rockim/pkg/component/database/mongo"
)

type PlatRoleData struct {
	mgo *mongo.Client
}

func NewPlatRoleData(mgo *mongo.Client) *PlatRoleData {
	return &PlatRoleData{mgo: mgo}
}

func (d *PlatRoleData) collection() *mgo.Collection {
	return d.mgo.Collection(entity.TablePlatRole)
}

func (d *PlatRoleData) FindByID(ctx context.Context, id string) (resource *types.PlatRole, err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var record entity.PlatRole
	err = d.mgo.FindOne(ctx, entity.TablePlatRole, bson.M{entity.MongoFieldId: objId}, &record, options.FindOne())
	if err != nil {
		return
	}
	return convert.RoleProto(&record), nil
}
func (d *PlatRoleData) ListByIds(ctx context.Context, ids []string) (results []*types.PlatRole, err error) {
	var objIdList []primitive.ObjectID
	for _, id := range ids {
		var objId primitive.ObjectID
		objId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objIdList = append(objIdList, objId)
	}
	var records []*entity.PlatRole
	err = d.mgo.FindList(ctx, entity.TablePlatRole, bson.M{entity.MongoFieldId: bson.M{"$in": objIdList}}, &records, options.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, convert.RoleProto(record))
	}
	return
}

func (d *PlatRoleData) Create(ctx context.Context, resource *types.PlatRole) (err error) {
	objId, err := primitive.ObjectIDFromHex(resource.Id)
	if err != nil {
		return
	}
	record := convert.RoleEntity(resource)
	record.Id = objId
	if err != nil {
		return
	}
	_, err = d.collection().InsertOne(ctx, record)
	return
}

func (d *PlatRoleData) Update(ctx context.Context, resource *types.PlatRole) (err error) {
	objId, err := primitive.ObjectIDFromHex(resource.Id)
	if err != nil {
		return
	}
	record := convert.RoleEntity(resource)
	_, err = d.collection().UpdateByID(ctx, objId, bson.M{"$set": record})
	return
}

func (d *PlatRoleData) Delete(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = d.collection().DeleteOne(ctx, bson.M{entity.MongoFieldId: objId})
	return
}

func (d *PlatRoleData) Paging(ctx context.Context, req *biz.PlatRolePagingRequest) (res *biz.PlatRolePagingResponse, err error) {
	query := bson.M{}
	cursor, p, err := d.mgo.Paginate(ctx, entity.TablePlatRole, query, req.Paginate)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	records, err := mongo.ScanCursor[entity.PlatRole](ctx, cursor)
	if err != nil {
		return
	}
	var list = make([]*types.PlatRole, len(records))
	for i, record := range records {
		list[i] = convert.RoleProto(record)
	}
	res = &biz.PlatRolePagingResponse{List: list, Paginate: p}
	return
}

func (d *PlatRoleData) ListResourceId(ctx context.Context, roleIds []string) (results []string, err error) {
	var objIdList []primitive.ObjectID
	for _, id := range roleIds {
		var objId primitive.ObjectID
		objId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objIdList = append(objIdList, objId)
	}
	var records []*entity.PlatRoleResource
	err = d.mgo.FindList(ctx, entity.TablePlatRoleResource, bson.M{"role_id": bson.M{"$in": objIdList}}, &records, options.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, record.ResourceId)
	}
	return
}
