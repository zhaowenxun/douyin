
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)

type UserInfoService struct {
	ctx context.Context
}

// NewUserInfoService new CreateUserService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

/// UserInfoService query user info
func (s *UserInfoService) UserInfo(req *user.UserInfoRequest) (*user.User, error) {
	
	
	user_id := req.UserId
	
	users, err := db.QueryUserInfo(s.ctx, user_id)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.AuthorizationFailedErr
	}
	u := users[0]
	user_info := pack.User(u)
	return user_info, nil
}