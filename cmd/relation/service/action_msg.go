package service

import (
	"context"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/relation/dal/db"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/message/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/relation"
)
type ActionMsgService struct {
	ctx context.Context
}

func NewActionMsgService(ctx context.Context) *ActionMsgService {
	return &ActionMsgService{ctx: ctx}
}

// Create Message
func (s *ActionMsgService) MGetActionMsg(req *relation.MessageActionRequest) error {
	MessageModel := &db.Message{
		ToUserId:   req.ToUserId,
		FromUserId:  req.FromUserId,
		Content: req.Content,
	}
	return db.CreateMessage(s.ctx, []*db.Message{MessageModel})
}