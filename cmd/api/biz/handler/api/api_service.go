// Code generated by hertz generator.

package api

import (
	"context"
	// "fmt"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/model/api"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/mw"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/rpc"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/message"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// LoginUser .
// @router /douyin/user/login/ [POST]
func LoginUser(ctx context.Context, c *app.RequestContext) {
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// RegisterUser .
// @router /douyin/user/register/ [POST]
func RegisterUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RegisterUserRequest
	err = c.BindAndValidate(&req)
	
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	err = rpc.RegisterUser(context.Background(), &user.RegisterUserRequest{
		Username: req.Username,
		Password: req.Password,

	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// UserInfo .
// @router /douyin/user/ [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserInfoRequest
	err = c.BindAndValidate(&req)
	
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	user_info, err := rpc.UserInfo(context.Background(), &user.UserInfoRequest{
		UserId: v.(*api.User).UserID,
		// Token: v.(*api.User).Token,

	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, user_info)
}

// MessageChat .
// @router /douyin/message/chat/ [GET]
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	messages, err := rpc.MessageChat(context.Background(), &message.MessageChatRequest{
		FromUserId:    v.(*api.User).UserID,
		ToUserId: req.ToUserID,
		// Offset:    req.Offset,
		// Limit:     req.Limit,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, utils.H{
		// consts.Total: total,
		consts.Messages: messages,
	})
}

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	
	err = rpc.MessageAction(context.Background(), &message.MessageActionRequest{
		FromUserId:  v.(*api.User).UserID,
		ToUserId: req.ToUserID,
		ActionType: req.ActionType,
		Content: req.Content,
		
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	
	err = rpc.PublishAction(context.Background(), &user.PublishActionRequest{
		UserId:  v.(*api.User).UserID,
		Data: req.Data,
		Title: req.Title,
		
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	// var err error
	// var req api.PublishListRequest
	// err = c.BindAndValidate(&req)
	// if err != nil {
	// 	c.String(consts.StatusBadRequest, err.Error())
	// 	return
	// }

	// resp := new(api.PublishListResponse)

	// c.JSON(consts.StatusOK, resp)
}
