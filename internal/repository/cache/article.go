package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangyul/go-blog/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

type ArticleCache interface {
	SetDetail(ctx context.Context, uid, id int64, art domain.Article) error
	GetDetail(ctx context.Context, uid, id int64) (domain.Article, error)
	GetPubDetail(ctx context.Context, uid int64, id int64) (domain.Article, error)
	SetPubDetail(ctx context.Context, uid int64, id int64, art domain.Article) error
}

type articleCache struct {
	cmd        redis.Cmdable
	expiration time.Duration
}

func NewRedisArticleCache(cmd redis.Cmdable) ArticleCache {
	return &articleCache{
		cmd:        cmd,
		expiration: 24 * time.Hour,
	}
}

func (a *articleCache) GetPubDetail(ctx context.Context, uid int64, id int64) (domain.Article, error) {
	data, err := a.cmd.Get(ctx, a.key(uid, id)).Result()
	if err != nil {
		return domain.Article{}, err
	}
	var art domain.Article
	err = json.Unmarshal([]byte(data), &art)
	if err != nil {
		return domain.Article{}, err
	}
	return art, nil
}

func (a *articleCache) SetPubDetail(ctx context.Context, uid int64, id int64, art domain.Article) error {
	data, err := json.Marshal(art)
	if err != nil {
		return err
	}
	return a.cmd.Set(ctx, a.pubKey(uid, id), data, a.expiration).Err()
}

func (a *articleCache) SetDetail(ctx context.Context, uid, id int64, art domain.Article) error {
	data, err := json.Marshal(art)
	if err != nil {
		return err
	}
	return a.cmd.Set(ctx, a.key(uid, id), data, a.expiration).Err()
}

func (a *articleCache) GetDetail(ctx context.Context, uid, id int64) (domain.Article, error) {
	res := a.cmd.Get(ctx, a.key(uid, id))
	if res.Err() != nil {
		return domain.Article{}, res.Err()
	}
	var art domain.Article
	err := json.Unmarshal([]byte(res.Val()), &art)
	if err != nil {
		return domain.Article{}, err
	}
	return art, nil
}

func (a *articleCache) key(uid, id int64) string {
	return fmt.Sprintf("article:detail:%d:%d", uid, id)
}

func (a *articleCache) pubKey(uid, id int64) string {
	return fmt.Sprintf("article:pub_detail:%d:%d", uid, id)
}
