package database

import (
	"context"
	"github.com/soukengo/gopkg/component/database"
	"github.com/soukengo/gopkg/component/paginate"
	"github.com/soukengo/gopkg/infra/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
	mgoOpts "go.mongodb.org/mongo-driver/mongo/options"
	"rockimserver/apis/rockim/shared"
	"rockimserver/app/access/admin/module/manager/biz/options"
	"rockimserver/app/access/admin/module/manager/biz/types"
	"rockimserver/app/access/admin/module/manager/data/database/convert"
	"rockimserver/app/access/admin/module/manager/data/database/entity"
	"time"
)

type SysUserData struct {
	mgo *mongo.Client
}

func NewSysUserData(mgr *database.Manager) *SysUserData {
	return &SysUserData{mgo: mgr.Mongo()}
}

func (d *SysUserData) collection() *mgo.Collection {
	return d.mgo.Collection(entity.TableSysUser)
}

func (d *SysUserData) FindByID(ctx context.Context, id string) (user *types.SysUser, err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var record entity.SysUser
	err = d.mgo.FindOne(ctx, entity.TableSysUser, bson.M{entity.MongoFieldId: objId}, &record, mgoOpts.FindOne())
	if err != nil {
		return
	}
	return convert.UserProto(&record), nil
}

func (d *SysUserData) FindIdByAccount(ctx context.Context, account string) (id string, err error) {
	projection := bson.M{entity.MongoFieldId: 1}
	var record entity.SysUser
	err = d.mgo.FindOne(ctx, entity.TableSysUser, bson.M{"account": account}, &record, mgoOpts.FindOne().SetProjection(projection))
	if err != nil {
		return
	}
	return record.Id.Hex(), nil
}

func (d *SysUserData) Create(ctx context.Context, user *types.SysUser) (err error) {
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

func (d *SysUserData) Update(ctx context.Context, user *types.SysUser) (err error) {
	objId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return
	}
	record := convert.UserEntity(user)
	_, err = d.collection().UpdateByID(ctx, objId, bson.M{"$set": record})
	return
}

func (d *SysUserData) Delete(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = d.collection().DeleteOne(ctx, bson.M{entity.MongoFieldId: objId})
	return
}

func (d *SysUserData) Paging(ctx context.Context, req *options.SysUserPagingOptions) (list []*types.SysUser, paginated *shared.Paginated, err error) {
	query := bson.M{"is_admin": bson.M{"$ne": true}}
	cursor, p, err := d.mgo.Paginate(ctx, entity.TableSysUser, query, &paginate.Paginating{PageSize: req.Paginate.PageSize, PageNo: req.Paginate.PageNo})
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	paginated = &shared.Paginated{Total: p.Total}
	records, err := mongo.ScanCursor[entity.SysUser](ctx, cursor)
	if err != nil {
		return
	}
	list = make([]*types.SysUser, len(records))
	for i, record := range records {
		list[i] = convert.UserProto(record)
	}
	return
}

func (d *SysUserData) ListRoleId(ctx context.Context, userId string) (results []string, err error) {
	objId, err := primitive.ObjectIDFromHex(userId)
	var records []*entity.SysUserRole
	err = d.mgo.FindList(ctx, entity.TableSysUserRole, bson.M{"user_id": objId}, &records, mgoOpts.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, record.RoleId.Hex())
	}
	return
}
func (d *SysUserData) SaveRoleId(ctx context.Context, userId string, roleIds []string) (err error) {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	var records []any
	for _, roleId := range roleIds {
		var roleObjId primitive.ObjectID
		roleObjId, err = primitive.ObjectIDFromHex(roleId)
		if err != nil {
			return
		}
		record := &entity.SysUserRole{UserId: userObjId, RoleId: roleObjId, CreateTime: time.Now().UnixMilli()}
		records = append(records, record)
	}
	_, err = d.mgo.Collection(entity.TableSysUserRole).InsertMany(ctx, records)
	return
}
