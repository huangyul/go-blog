package web

import (
	"github.com/gin-gonic/gin"
	interv1 "github.com/huangyul/go-blog/api/proto/gen/inter/v1"
	"github.com/huangyul/go-blog/internal/domain"
	"github.com/huangyul/go-blog/internal/pkg/errno"
	"github.com/huangyul/go-blog/internal/pkg/ginx/validator"
	"github.com/huangyul/go-blog/internal/service"
	"golang.org/x/sync/errgroup"
	"strconv"
	"time"
)

const biz = "article"

type ArticleHandler struct {
	svc service.ArticleService
	//interSvc inteService.InteractiveService
	interSvc interv1.InteractiveServiceClient
}

func NewArticleHandler(svc service.ArticleService, interSvc interv1.InteractiveServiceClient) *ArticleHandler {
	return &ArticleHandler{
		svc:      svc,
		interSvc: interSvc,
	}
}

func (h *ArticleHandler) RegisterRoutes(g *gin.Engine) {
	ug := g.Group("/article")
	ug.POST("/edit", h.Edit)
	ug.POST("/publish", h.Publish)
	ug.GET("/withdraw/:id", h.Withdraw)
	ug.POST("/list", h.List)
	ug.GET("/detail/:id", h.Detail)
	pg := ug.Group("/pub")
	pg.GET("/detail/:id", h.PubDetail)
	pg.POST("/like/:id", h.Like)
	pg.POST("/collect", h.Collect)
	pg.GET("/history", h.History)
}

type ArticleEditReq struct {
	ID      int64  `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (h *ArticleHandler) Edit(ctx *gin.Context) {

	var req ArticleEditReq
	if err := ctx.ShouldBind(&req); err != nil {
		WriteErrno(ctx, errno.ErrBadRequest.SetMessage(validator.Translate(err)))
		return
	}
	userId := ctx.MustGet("user_id").(int64)
	id, err := h.svc.Save(ctx, domain.Article{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
		Author: domain.Author{
			ID: userId,
		},
	})
	if err != nil {
		WriteErrno(ctx, errno.ErrInternalServer.SetMessage(err.Error()))
		return
	}
	WriteSuccess(ctx, gin.H{"id": id})

}

func (h *ArticleHandler) Publish(ctx *gin.Context) {
	type Req struct {
		ID      int64  `json:"id"`
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	var req Req
	if err := ctx.ShouldBind(&req); err != nil {
		WriteErrno(ctx, errno.ErrBadRequest.SetMessage(validator.Translate(err)))
		return
	}
	userId := ctx.MustGet("user_id").(int64)
	id, err := h.svc.Publish(ctx, domain.Article{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
		Author: domain.Author{
			ID: userId,
		},
	})
	if err != nil {
		WriteErrno(ctx, errno.ErrInternalServer.SetMessage(err.Error()))
		return
	}
	WriteSuccess(ctx, gin.H{"id": id})
}

func (h *ArticleHandler) Withdraw(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		WriteErrno(ctx, errno.ErrBadRequest.SetMessage("id illegal"))
		return
	}
	userId := ctx.MustGet("user_id").(int64)
	err = h.svc.Withdraw(ctx, userId, id)
	if err != nil {
		WriteErrno(ctx, errno.ErrInternalServer.SetMessage(err.Error()))
		return
	}
	WriteSuccess(ctx, nil)
}

func (h *ArticleHandler) List(ctx *gin.Context) {
	var req Page
	if err := ctx.ShouldBind(&req); err != nil {
		WriteErrno(ctx, errno.ErrBadRequest.SetMessage(validator.Translate(err)))
		return
	}
	req.SetDefault()

	userId := ctx.MustGet("user_id").(int64)
	data, err := h.svc.List(ctx, userId, req.Page, req.PageSize)
	if err != nil {
		WriteErrno(ctx, errno.ErrInternalServer.SetMessage(err.Error()))
		return
	}
	WriteSuccess(ctx, gin.H{"data": data})
}

func (h *ArticleHandler) Detail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		WriteErrno(ctx, errno.ErrBadRequest.SetMessage("id illegal"))
		return
	}
	userId := ctx.MustGet("user_id").(int64)
	art, err := h.svc.Detail(ctx, userId, id)
	if err != nil {
		WriteErrno(ctx, errno.ErrInternalServer.SetMessage(err.Error()))
		return
	}
	WriteSuccess(ctx, gin.H{"data": art})
}

func (h *ArticleHandler) PubDetail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		WriteErrno(ctx, errno.ErrBadRequest.SetMessage("id illegal"))
		return
	}
	userId := ctx.MustGet("user_id").(int64)
	var (
		eg   errgroup.Group
		art  domain.Article
		inte *interv1.InteractiveResponse
	)
	eg.Go(func() error {
		var er error
		art, er = h.svc.PubDetail(ctx, userId, id, biz)
		return er
	})
	eg.Go(func() error {
		var er error
		inte, er = h.interSvc.Get(ctx, &interv1.GetRequest{
			Uid: userId,
			Id:  id,
			Biz: biz,
		})
		return er
	})
	err = eg.Wait()
	if err != nil {
		WriteErrno(ctx, errno.ErrInternalServer.SetMessage(err.Error()))
		return
	}
	res := pubDetail{
		ID:         art.ID,
		Title:      art.Title,
		Content:    art.Content,
		AuthorName: art.Author.Name,
		AuthorID:   art.Author.ID,
		UpdatedAt:  art.UpdatedAt.Format(time.DateOnly),
		CreatedAt:  art.CreatedAt.Format(time.DateOnly),
		ReadCnt:    int(inte.ReadCnt),
		LikeCnt:    int(inte.LikeCnt),
		CollectCnt: int(inte.CollectCnt),
		Liked:      inte.Liked,
		Collected:  inte.Collected,
	}
	//go func() {
	//	h.interSvc.IncrReadCnt(ctx, art.ID, biz)
	//}()
	WriteSuccess(ctx, gin.H{"data": res})
}

func (h *ArticleHandler) Like(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		WriteErrno(ctx, errno.ErrBadRequest.SetMessage("id illegal"))
		return
	}
	type Req struct {
		Like bool `json:"like"`
	}
	var req Req
	if err := ctx.ShouldBind(&req); err != nil {
		WriteErrno(ctx, errno.ErrBadRequest.SetMessage(validator.Translate(err)))
		return
	}
	userId := ctx.MustGet("user_id").(int64)
	var er error
	if req.Like {
		_, er = h.interSvc.Like(ctx, &interv1.LikeRequest{
			Biz:    biz,
			BizId:  id,
			UserId: userId,
		})
	} else {
		_, er = h.interSvc.CancelLike(ctx, &interv1.CancelLikeRequest{
			Biz:    biz,
			BizId:  id,
			UserId: userId,
		})
	}
	if er != nil {
		WriteErrno(ctx, errno.ErrInternalServer.SetMessage(er.Error()))
		return
	}
	WriteSuccess(ctx, nil)
}

func (h *ArticleHandler) Collect(ctx *gin.Context) {
	type Req struct {
		ID  int64 `json:"id" binding:"required"`
		CID int64 `json:"cid"`
	}
	var req Req
	if err := ctx.ShouldBind(&req); err != nil {
		WriteErrno(ctx, errno.ErrBadRequest.SetMessage(validator.Translate(err)))
		return
	}
	userId := ctx.MustGet("user_id").(int64)
	_, err := h.interSvc.Collect(ctx, &interv1.CollectRequest{
		Uid: userId,
		Biz: biz,
		Id:  req.ID,
		Cid: req.CID,
	})
	if err != nil {
		WriteErrno(ctx, errno.ErrInternalServer.SetMessage(err.Error()))
		return
	}
	WriteSuccess(ctx, nil)
}

func (h *ArticleHandler) History(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(int64)
	arts, err := h.svc.GetHistory(ctx, userId)
	if err != nil {
		WriteErrno(ctx, errno.ErrInternalServer.SetMessage(err.Error()))
		return
	}
	WriteSuccess(ctx, arts)
}

type pubDetail struct {
	ID         int64  `json:"id"`
	AuthorName string `json:"author_name"`
	AuthorID   int64  `json:"author_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	ReadCnt    int    `json:"read_cnt"`
	LikeCnt    int    `json:"like_cnt"`
	CollectCnt int    `json:"collect_cnt"`
	Liked      bool   `json:"liked"`
	Collected  bool   `json:"collected"`
}
