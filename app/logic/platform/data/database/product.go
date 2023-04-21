package database

import (
	"context"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/infra/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rockimserver/apis/rockim/service/platform/v1/types"
	"rockimserver/app/logic/platform/data/database/convert"
	"rockimserver/app/logic/platform/data/database/entity"
	"strconv"
)

type ProductData struct {
	mgo *mongo.Client
}

func NewProductData(mgr *database.Manager) *ProductData {
	return &ProductData{mgo: mgr.Mongo()}
}

func (d *ProductData) collection() *mgo.Collection {
	return d.mgo.Collection(entity.TableProduct)
}

func (d *ProductData) GenID(ctx context.Context) (string, error) {
	seq, err := genMongoSeq(ctx, d.mgo, entity.TableProduct, true, 1)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(seq, 10), nil
}

func (d *ProductData) FindByID(ctx context.Context, id string) (t *types.Product, err error) {
	var record entity.Product
	err = d.mgo.FindOne(ctx, entity.TableProduct, bson.M{entity.MongoFieldId: id}, &record, options.FindOne())
	if err != nil {
		return
	}
	return convert.ProductProto(&record), nil
}

func (d *ProductData) ListByIds(ctx context.Context, ids []string) (results []*types.Product, err error) {
	var records []*entity.Product
	err = d.mgo.FindList(ctx, entity.TableProduct, bson.M{entity.MongoFieldId: bson.M{"$in": ids}}, &records, options.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, convert.ProductProto(record))
	}
	return
}

func (d *ProductData) Create(ctx context.Context, t *types.Product) (err error) {
	record := convert.ProductEntity(t)
	_, err = d.collection().InsertOne(ctx, record)
	return
}

func (d *ProductData) Update(ctx context.Context, t *types.Product) (err error) {
	id := t.Id
	record := convert.ProductEntity(t)
	record.Id = ""
	_, err = d.collection().UpdateByID(ctx, id, bson.M{"$set": record})
	return
}

func (d *ProductData) Delete(ctx context.Context, id string) (err error) {
	_, err = d.collection().DeleteOne(ctx, bson.M{entity.MongoFieldId: id})
	return
}

func (d *ProductData) ListByTenant(ctx context.Context, tenantId string) (res []*types.Product, err error) {
	query := bson.M{"tenant_id": tenantId}
	var records []*entity.Product
	err = d.mgo.FindList(ctx, entity.TableProduct, query, &records)
	if err != nil {
		return
	}
	var list = make([]*types.Product, len(records))
	for i, record := range records {
		list[i] = convert.ProductProto(record)
	}
	res = list
	return
}
