package authorized_service

import (
	"github.com/skyzhouzj/skyCloud/configs"
	"github.com/skyzhouzj/skyCloud/internal/api/repo"
	"github.com/skyzhouzj/skyCloud/internal/api/repo/authorized_repo"
	"github.com/skyzhouzj/skyCloud/pkg/cache"
	"github.com/skyzhouzj/skyCloud/pkg/core"

	"gorm.io/gorm"
)

func (s *service) Delete(ctx core.Context, id int32) (err error) {
	// 先查询 id 是否存在
	authorizedInfo, err := authorized_repo.NewQueryBuilder().
		WhereIsDeleted(repo.EqualPredicate, -1).
		WhereId(repo.EqualPredicate, id).
		First(s.db.GetDb().WithContext(ctx.RequestContext()))

	if err == gorm.ErrRecordNotFound {
		return nil
	}

	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": ctx.UserName(),
	}

	qb := authorized_repo.NewQueryBuilder()
	qb.WhereId(repo.EqualPredicate, id)
	err = qb.Updates(s.db.GetDb().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	s.cache.Del(configs.Get().SkyCloud.RedisKeyPrefixSignature+authorizedInfo.BusinessKey, cache.WithTrace(ctx.Trace()))
	return
}
