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

type PlatUserData struct {
	mgo *mongo.Client
}

func NewPlatUserData(mgo *mongo.Client) *PlatUserData {
	return &PlatUserData{mgo: mgo}
}

func (d *PlatUserData) collection() *mgo.Collection {
	return d.mgo.Collection(entity.TablePlatUser)
}

func (d *PlatUserData) FindByID(ctx context.Context, id string) (user *types.PlatUser, err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var record entity.PlatUser
	err = d.mgo.FindOne(ctx, entity.TablePlatUser, bson.M{entity.MongoFieldId: objId}, &record, options.FindOne())
	if err != nil {
		return
	}
	return convert.UserProto(&record), nil
}

func (d *PlatUserData) FindByAccount(ctx context.Context, account string) (id string, err error) {
	projection := bson.M{entity.MongoFieldId: 1}
	var record entity.PlatUser
	err = d.mgo.FindOne(ctx, entity.TablePlatUser, bson.M{"account": account}, &record, options.FindOne().SetProjection(projection))
	if err != nil {
		return
	}
	return record.Id.Hex(), nil
}

func (d *PlatUserData) Create(ctx context.Context, user *types.PlatUser) (err error) {
	objId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return
	}
	record := convert.UserEntity(user)
	record.Id = objId
	if err != nil {
		return
	}
	_, err = d.collection().InsertOne(ctx, record)
	return
}

func (d *PlatUserData) Update(ctx context.Context, user *types.PlatUser) (err error) {
	objId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return
	}
	record := convert.UserEntity(user)
	_, err = d.collection().UpdateByID(ctx, objId, bson.M{"$set": record})
	return
}

func (d *PlatUserData) Delete(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = d.collection().DeleteOne(ctx, bson.M{entity.MongoFieldId: objId})
	return
}

func (d *PlatUserData) Paging(ctx context.Context, req *biz.PlatUserPagingRequest) (res *biz.PlatUserPagingResponse, err error) {
	query := bson.M{}
	cursor, p, err := d.mgo.Paginate(ctx, entity.TablePlatUser, query, req.Paginate)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	//var records []*entity.PlatUser
	records, err := mongo.ScanCursor[entity.PlatUser](ctx, cursor)
	//err = cursor.All(ctx, &records)
	if err != nil {
		return
	}
	var list = make([]*types.PlatUser, len(records))
	for i, record := range records {
		list[i] = convert.UserProto(record)
	}
	res = &biz.PlatUserPagingResponse{List: list, Paginate: p}
	return
}

func (d *PlatUserData) ListRoleId(ctx context.Context, userIds []string) (results []string, err error) {
	var objIdList []primitive.ObjectID
	for _, id := range userIds {
		var objId primitive.ObjectID
		objId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objIdList = append(objIdList, objId)
	}
	var records []*entity.PlatUserRole
	err = d.mgo.FindList(ctx, entity.TablePlatUserRole, bson.M{"user_id": bson.M{"$in": objIdList}}, &records, options.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, record.RoleId)
	}
	return
}
