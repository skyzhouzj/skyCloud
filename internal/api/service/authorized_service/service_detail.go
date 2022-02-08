package authorized_service

import (
	"github.com/skyzhouzj/skyCloud/internal/api/repo"
	"github.com/skyzhouzj/skyCloud/internal/api/repo/authorized_repo"
	"github.com/skyzhouzj/skyCloud/pkg/core"
)

func (s *service) Detail(ctx core.Context, id int32) (info *authorized_repo.Authorized, err error) {
	qb := authorized_repo.NewQueryBuilder()
	qb.WhereIsDeleted(repo.EqualPredicate, -1)
	qb.WhereId(repo.EqualPredicate, id)

	info, err = qb.First(s.db.GetDb().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
