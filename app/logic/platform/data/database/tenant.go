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

type TenantData struct {
	mgo *mongo.Client
}

func NewTenantData(mgo *mongo.Client) *TenantData {
	return &TenantData{mgo: mgo}
}

func (d *TenantData) collection() *mgo.Collection {
	return d.mgo.Collection(entity.TableTenant)
}

func (d *TenantData) FindByID(ctx context.Context, id string) (t *types.Tenant, err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var record entity.Tenant
	err = d.mgo.FindOne(ctx, entity.TableTenant, bson.M{entity.MongoFieldId: objId}, &record, options.FindOne())
	if err != nil {
		return
	}
	return convert.TenantProto(&record), nil
}

func (d *TenantData) FindIdByEmail(ctx context.Context, email string) (id string, err error) {
	projection := bson.M{entity.MongoFieldId: 1}
	var record entity.Tenant
	err = d.mgo.FindOne(ctx, entity.TableTenant, bson.M{"email": email}, &record, options.FindOne().SetProjection(projection))
	if err != nil {
		return
	}
	return record.Id.Hex(), nil
}

func (d *TenantData) ListByIds(ctx context.Context, ids []string) (results []*types.Tenant, err error) {
	var objIdList []primitive.ObjectID
	for _, id := range ids {
		var objId primitive.ObjectID
		objId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objIdList = append(objIdList, objId)
	}
	var records []*entity.Tenant
	err = d.mgo.FindList(ctx, entity.TableTenant, bson.M{entity.MongoFieldId: bson.M{"$in": objIdList}}, &records, options.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, convert.TenantProto(record))
	}
	return
}

func (d *TenantData) Create(ctx context.Context, t *types.Tenant) (err error) {
	objId, err := primitive.ObjectIDFromHex(t.Id)
	if err != nil {
		return
	}
	record := convert.TenantEntity(t)
	record.Id = objId
	if err != nil {
		return
	}
	_, err = d.collection().InsertOne(ctx, record)
	return
}

func (d *TenantData) Update(ctx context.Context, t *types.Tenant) (err error) {
	objId, err := primitive.ObjectIDFromHex(t.Id)
	if err != nil {
		return
	}
	record := convert.TenantEntity(t)
	_, err = d.collection().UpdateByID(ctx, objId, bson.M{"$set": record})
	return
}

func (d *TenantData) Delete(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = d.collection().DeleteOne(ctx, bson.M{entity.MongoFieldId: objId})
	return
}

func (d *TenantData) Paging(ctx context.Context, req *biz.TenantPagingRequest) (res *biz.TenantPagingResponse, err error) {
	query := bson.M{}
	cursor, p, err := d.mgo.Paginate(ctx, entity.TableTenant, query, req.Paginate)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	records, err := mongo.ScanCursor[entity.Tenant](ctx, cursor)
	if err != nil {
		return
	}
	var list = make([]*types.Tenant, len(records))
	for i, record := range records {
		list[i] = convert.TenantProto(record)
	}
	res = &biz.TenantPagingResponse{List: list, Paginate: p}
	return
}
