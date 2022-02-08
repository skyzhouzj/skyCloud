package authorized_service

import (
	"github.com/skyzhouzj/skyCloud/internal/api/repo"
	"github.com/skyzhouzj/skyCloud/internal/api/repo/authorized_repo"
	"github.com/skyzhouzj/skyCloud/pkg/core"
)

func (s *service) List(ctx core.Context, searchData *SearchData) (listData []*authorized_repo.Authorized, err error) {

	qb := authorized_repo.NewQueryBuilder()
	qb = qb.WhereIsDeleted(repo.EqualPredicate, -1)

	if searchData.BusinessKey != "" {
		qb.WhereBusinessKey(repo.EqualPredicate, searchData.BusinessKey)
	}

	if searchData.BusinessSecret != "" {
		qb.WhereBusinessSecret(repo.EqualPredicate, searchData.BusinessSecret)
	}

	if searchData.BusinessDeveloper != "" {
		qb.WhereBusinessDeveloper(repo.EqualPredicate, searchData.BusinessDeveloper)
	}

	listData, err = qb.
		OrderById(false).
		QueryAll(s.db.GetDb().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
