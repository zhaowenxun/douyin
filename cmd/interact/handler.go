package main

import (
	"context"
	// "fmt"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/interact/service"




	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	interact "github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/interact"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *interact.FavoriteActionRequest) (resp *interact.FavoriteActionResponse, err error) {
	resp = new(interact.FavoriteActionResponse)

	if len(req.Token) == 0 || req.VideoId == 0 || req.ActionType == 0 {
		resp = pack.BuildFavoriteBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp = pack.BuildFavoriteBaseResp(err)
		return resp, nil
	}
	resp = pack.BuildFavoriteBaseResp(errno.Success)
	return resp, nil
	return
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (resp *interact.FavoriteListResponse, err error) {
	resp = new(interact.FavoriteListResponse)

	if req.UserId == 0 {
		resp = pack.BuildFavoriteListBaseResp(errno.ParamErr)
		return resp, nil
	}

	videoList, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		resp = pack.BuildFavoriteListBaseResp(err)
		return resp, nil
	}

	resp = pack.BuildFavoriteListBaseResp(errno.Success)
	resp.VideoList = videoList
	return resp, nil
}

// CommentSrvImpl implements the last service interface defined in the IDL.
type CommentSrvImpl struct{}

func packErr1(err error) *interact.CommentActionResponse {
	msg := err.Error()
	return &interact.CommentActionResponse{StatusCode: errno.CommentError, StatusMsg: &msg}
}

func packErr2(err error) *interact.CommentListResponse {
	msg := err.Error()
	return &interact.CommentListResponse{StatusCode: errno.SuccessCode, StatusMsg: &msg,
		CommentList: []*interact.Comment{}}
}

func getUser(ctx context.Context, id int, token string) (*interact.User, error) {
	resp, err := rpc.Info(ctx, &user.UserRequest{Token: token, UserId: int64(id)})
	if err != nil {
		fmt.Println(err)
		panic(err)
		return nil, err
	}
	return &interact.User{Id: resp.Id, Name: resp.Name, FollowCount: resp.FollowCount,
		FollowerCount: resp.FollowerCount, IsFollow: resp.IsFollow}, nil
}

// TODO service
// CommentAction implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentAction(ctx context.Context, req *interact.CommentActionRequest) (resp *interact.CommentActionResponse, err error) {
	// TODO: Your code here...
	log.Info("get interact action req", *req)
	resp = new(interact.CommentActionResponse)

	//TODO check video id
	if req.ActionType == 1 {
		cmt := &db.Comment{UserId: int(req.UserId), Content: *req.CommentText,
			VideoId: int(req.VideoId), IsValid: true, CreateTime: time.Now().String()}
		if err := db.CreateComment(ctx, cmt); err != nil {
			return packErr1(err), nil
		}
		user, err := getUser(ctx, cmt.UserId, req.Token)
		if err != nil {
			return nil, err
		}
		return &interact.CommentActionResponse{StatusCode: errno.SuccessCode,
			Comment: &interact.Comment{Id: int64(cmt.ID), User: user, Content: cmt.Content, CreateDate: cmt.CreateTime}}, nil
	} else if req.ActionType == 2 {
		cmt := &db.Comment{ID: uint(*req.CommentId)}
		tmp, err := db.SelectComment(ctx, int(cmt.ID))
		if err != nil {
			return packErr1(err), nil
		}
		if tmp == nil {
			return &interact.CommentActionResponse{StatusCode: errno.CommentNotFound}, nil
		}
		if err := db.DeleteComment(ctx, cmt); err != nil {
			return packErr1(err), nil
		}
		user, err := getUser(ctx, tmp.UserId, req.Token)
		if err != nil {
			return nil, err
		}
		return &interact.CommentActionResponse{StatusCode: errno.SuccessCode,
			Comment: &interact.Comment{Id: int64(tmp.ID), User: user, Content: tmp.Content, CreateDate: tmp.CreateTime}}, nil
	} else {
		msg := "err"
		return &interact.CommentActionResponse{StatusCode: errno.ActionTypeErrCode, StatusMsg: &msg}, nil
	}
}

// CommentList implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentList(ctx context.Context, req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
	log.Info("get interact list req", *req)
	// TODO: Your code here...
	resp = new(interact.CommentListResponse)
	cmts, err := db.QueryComments(ctx, int(req.VideoId))
	if err != nil {
		return packErr2(err), nil
	}
	res := []*interact.Comment{}
	for _, c := range cmts {
		fmt.Println("get user , id = ", c.UserId)
		user, err := getUser(ctx, c.UserId, req.Token)
		if err != nil {
			return nil, err
		}
		tmp := &interact.Comment{Id: int64(c.ID), Content: c.Content,
			CreateDate: c.CreateTime, User: user}
		res = append(res, tmp)
	}
	return &interact.CommentListResponse{StatusCode: errno.SuccessCode, CommentList: res}, nil
}