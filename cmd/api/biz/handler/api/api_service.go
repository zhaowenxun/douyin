// Code generated by hertz generator.

package api

import (
	"context"

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
		Token:    v,
		ToUserId: req.to_user_id,
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
	resp := new(api.MessageActionResponse)
	err = rpc.MessageAction(context.Background(), &message.MessageActionRequest{
		Token:  v,
		ToUserId: req.to_user_id,
		ActionType: req.action_type,
		Content: req.Content,
		
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
