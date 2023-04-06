package database

import (
	"context"
	"github.com/soukengo/gopkg/component/database/mongo"
	"github.com/stretchr/testify/suite"
	"rockimserver/apis/rockim/shared"
	"rockimserver/app/access/admin/module/manager/biz/options"
	"testing"
)

type SysUserDataSuite struct {
	suite.Suite
	target *SysUserData
}

func (s *SysUserDataSuite) SetupSuite() {
	mgo := mongo.NewClient(&mongo.Config{
		Address:    "mongodb://127.0.0.1:27017",
		Username:   "rockim",
		Password:   "rockim2022",
		Database:   "rockim",
		AuthSource: "admin",
	})
	s.target = NewSysUserData(mgo)
}

func (s *SysUserDataSuite) Test_Paging() {
	ctx := context.TODO()
	list, p, err := s.target.Paging(ctx, &options.SysUserPagingOptions{
		Paginate: &shared.Paginating{
			PageNo:   1,
			PageSize: 10,
		},
	})
	if err != nil {
		s.Fail("Paging error: %v", err)
	}
	s.T().Logf("list: %v, p: %v", list, p)
}

func Test_SysUserData(t *testing.T) {
	suite.Run(t, new(SysUserDataSuite))
}
