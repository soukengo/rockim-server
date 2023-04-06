package database

import (
	"context"
	"github.com/soukengo/gopkg/component/database/mongo"
	"github.com/soukengo/gopkg/component/paginate"
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

type SysRoleData struct {
	mgo *mongo.Client
}

func NewSysRoleData(mgo *mongo.Client) *SysRoleData {
	return &SysRoleData{mgo: mgo}
}

func (d *SysRoleData) collection() *mgo.Collection {
	return d.mgo.Collection(entity.TableSysRole)
}

func (d *SysRoleData) FindByID(ctx context.Context, id string) (resource *types.SysRole, err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var record entity.SysRole
	err = d.mgo.FindOne(ctx, entity.TableSysRole, bson.M{entity.MongoFieldId: objId}, &record, mgoOpts.FindOne())
	if err != nil {
		return
	}
	return convert.RoleProto(&record), nil
}
func (d *SysRoleData) ListByIds(ctx context.Context, ids []string) (results []*types.SysRole, err error) {
	var objIdList []primitive.ObjectID
	for _, id := range ids {
		var objId primitive.ObjectID
		objId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objIdList = append(objIdList, objId)
	}
	var records []*entity.SysRole
	err = d.mgo.FindList(ctx, entity.TableSysRole, bson.M{entity.MongoFieldId: bson.M{"$in": objIdList}}, &records, mgoOpts.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, convert.RoleProto(record))
	}
	return
}

func (d *SysRoleData) Create(ctx context.Context, resource *types.SysRole) (err error) {
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

func (d *SysRoleData) Update(ctx context.Context, resource *types.SysRole) (err error) {
	objId, err := primitive.ObjectIDFromHex(resource.Id)
	if err != nil {
		return
	}
	record := convert.RoleEntity(resource)
	_, err = d.collection().UpdateByID(ctx, objId, bson.M{"$set": record})
	return
}

func (d *SysRoleData) Delete(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = d.collection().DeleteOne(ctx, bson.M{entity.MongoFieldId: objId})
	return
}

func (d *SysRoleData) Paging(ctx context.Context, req *options.SysRolePagingOptions) (list []*types.SysRole, paginated *shared.Paginated, err error) {
	query := bson.M{}
	cursor, p, err := d.mgo.Paginate(ctx, entity.TableSysRole, query, &paginate.Paginating{PageSize: req.Paginate.PageSize, PageNo: req.Paginate.PageNo})
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	paginated = &shared.Paginated{Total: p.Total}
	records, err := mongo.ScanCursor[entity.SysRole](ctx, cursor)
	if err != nil {
		return
	}
	list = make([]*types.SysRole, len(records))
	for i, record := range records {
		list[i] = convert.RoleProto(record)
	}
	return
}

func (d *SysRoleData) ListResourceId(ctx context.Context, roleIds []string) (results []string, err error) {
	var objIdList []primitive.ObjectID
	for _, id := range roleIds {
		var objId primitive.ObjectID
		objId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objIdList = append(objIdList, objId)
	}
	var records []*entity.SysRoleResource
	err = d.mgo.FindList(ctx, entity.TableSysRoleResource, bson.M{"role_id": bson.M{"$in": objIdList}}, &records, mgoOpts.Find())
	if err != nil {
		return
	}
	for _, record := range records {
		results = append(results, record.ResourceId.Hex())
	}
	return
}

func (d *SysRoleData) SaveResourceId(ctx context.Context, roleId string, resourceIds []string) (err error) {
	roleObjId, err := primitive.ObjectIDFromHex(roleId)
	var records []any
	for _, resourceId := range resourceIds {
		var resourceObjId primitive.ObjectID
		resourceObjId, err = primitive.ObjectIDFromHex(resourceId)
		if err != nil {
			return
		}
		record := &entity.SysRoleResource{RoleId: roleObjId, ResourceId: resourceObjId, CreateTime: time.Now().UnixMilli()}
		records = append(records, record)
	}
	_, err = d.mgo.Collection(entity.TableSysRoleResource).InsertMany(ctx, records)
	return
}
