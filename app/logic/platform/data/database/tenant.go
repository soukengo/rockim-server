package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	mgo "go.mongodb.org/mongo-driver/mongo"
	mgoopts "go.mongodb.org/mongo-driver/mongo/options"
	"rockimserver/apis/rockim/service/platform/v1/types"
	"rockimserver/apis/rockim/shared"
	"rockimserver/app/logic/platform/biz/options"
	"rockimserver/app/logic/platform/data/database/convert"
	"rockimserver/app/logic/platform/data/database/entity"
	"rockimserver/pkg/component/database/mongo"
	"strconv"
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

func (d *TenantData) GenID(ctx context.Context) (string, error) {
	seq, err := genMongoSeq(ctx, d.mgo, entity.TableTenant, true, 1)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(seq, 10), nil
}

func (d *TenantData) FindByID(ctx context.Context, id string) (t *types.Tenant, err error) {
	var record entity.Tenant
	err = d.mgo.FindOne(ctx, entity.TableTenant, bson.M{entity.MongoFieldId: id}, &record, mgoopts.FindOne())
	if err != nil {
		return
	}
	return convert.TenantProto(&record), nil
}

func (d *TenantData) FindIdByEmail(ctx context.Context, email string) (id string, err error) {
	projection := bson.M{entity.MongoFieldId: 1}
	var record entity.Tenant
	err = d.mgo.FindOne(ctx, entity.TableTenant, bson.M{"email": email}, &record, mgoopts.FindOne().SetProjection(projection))
	if err != nil {
		return
	}
	return record.Id, nil
}

func (d *TenantData) ListByIds(ctx context.Context, ids []string) (results []*types.Tenant, err error) {
	var records []*entity.Tenant
	err = d.mgo.FindList(ctx, entity.TableTenant, bson.M{entity.MongoFieldId: bson.M{"$in": ids}}, &records, mgoopts.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, convert.TenantProto(record))
	}
	return
}

func (d *TenantData) Create(ctx context.Context, t *types.Tenant) (err error) {
	record := convert.TenantEntity(t)
	if err != nil {
		return
	}
	_, err = d.collection().InsertOne(ctx, record)
	return
}

func (d *TenantData) Update(ctx context.Context, t *types.Tenant) (err error) {
	record := convert.TenantEntity(t)
	record.Id = ""
	_, err = d.collection().UpdateByID(ctx, t.Id, bson.M{"$set": record})
	return
}

func (d *TenantData) Delete(ctx context.Context, id string) (err error) {
	_, err = d.collection().DeleteOne(ctx, bson.M{entity.MongoFieldId: id})
	return
}

func (d *TenantData) Paging(ctx context.Context, opts *options.TenantPagingOptions) (ret []*types.Tenant, paginated *shared.Paginated, err error) {
	query := bson.M{}
	cursor, p, err := d.mgo.Paginate(ctx, entity.TableTenant, query, opts.Paginate)
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
	return list, p, nil
}
