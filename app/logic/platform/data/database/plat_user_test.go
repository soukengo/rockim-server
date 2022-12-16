package database

import (
	"context"
	"github.com/stretchr/testify/suite"
	sharedTypes "rockim/api/rockim/shared/types"
	"rockim/app/logic/platform/biz"
	"rockim/pkg/component/database/mongo"
	"testing"
)

type PlatUserDataSuite struct {
	suite.Suite
	target *PlatUserData
}

func (s *PlatUserDataSuite) SetupSuite() {
	mgo := mongo.NewClient(&mongo.Config{
		Address:    "mongodb://127.0.0.1:27017",
		Username:   "rockim",
		Password:   "rockim2022",
		Database:   "rockim",
		AuthSource: "admin",
	})
	s.target = NewPlatUserData(mgo)
}

func (s *PlatUserDataSuite) Test_Paging() {
	ctx := context.TODO()
	res, err := s.target.Paging(ctx, &biz.PlatUserPagingRequest{
		Paginate: &sharedTypes.Paginating{
			PageNo:   1,
			PageSize: 10,
		},
	})
	if err != nil {
		s.Fail("Paging error: %v", err)
	}
	s.T().Logf("res: %v", res)
}

func Test_PlatUserData(t *testing.T) {
	suite.Run(t, new(PlatUserDataSuite))
}
