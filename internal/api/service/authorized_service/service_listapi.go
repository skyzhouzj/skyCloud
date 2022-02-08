package authorized_service

import (
	"github.com/skyzhouzj/skyCloud/internal/api/repo"
	"github.com/skyzhouzj/skyCloud/internal/api/repo/authorized_api_repo"
	"github.com/skyzhouzj/skyCloud/pkg/core"
)

type SearchAPIData struct {
	BusinessKey string `json:"business_key"` // 调用方key
}

func (s *service) ListAPI(ctx core.Context, searchAPIData *SearchAPIData) (listData []*authorized_api_repo.AuthorizedApi, err error) {

	qb := authorized_api_repo.NewQueryBuilder()
	qb = qb.WhereIsDeleted(repo.EqualPredicate, -1)

	if searchAPIData.BusinessKey != "" {
		qb.WhereBusinessKey(repo.EqualPredicate, searchAPIData.BusinessKey)
	}

	listData, err = qb.
		OrderById(false).
		QueryAll(s.db.GetDb().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
