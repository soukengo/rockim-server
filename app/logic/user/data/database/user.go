package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rockim/app/logic/user/data/database/entity"
	"rockim/pkg/component/database/mongo"
	"strconv"
)

type UserData struct {
	mgo *mongo.Client
}

func NewUserData(mgo *mongo.Client) *UserData {
	return &UserData{mgo: mgo}
}

func (d *UserData) GenID(ctx context.Context) (string, error) {
	seq, err := genMongoSeq(ctx, d.mgo, entity.TableImUser, true, 1)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(seq, 10), nil
}

func (d *UserData) FindByAccount(ctx context.Context, appId string, account string) (string, error) {
	c := d.mgo.Collection(entity.TableImUser)
	c.FindOne(ctx, bson.M{"appId": appId, "account": account}, options.FindOne())
	return "", nil
}
