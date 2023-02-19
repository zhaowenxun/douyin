// Code generated by hertz generator.

package relation

import (
	"context"
	// "fmt"
	api_relation "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/model/api_relation"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/model/api_user"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/mw"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/rpc"

	// "github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/message"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// MessageChat .
// @router /douyin/message/chat/ [GET]
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api_relation.MessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	messages, err := rpc.MessageChat(context.Background(), &message.MessageChatRequest{
		FromUserId:    v.(*api_relation.User).ID,
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
	var req api_relation.MessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	
	err = rpc.MessageAction(context.Background(), &message.MessageActionRequest{
		FromUserId:  v.(*api_relation.User).ID,
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
