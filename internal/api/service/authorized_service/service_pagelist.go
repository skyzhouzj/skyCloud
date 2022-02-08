package authorized_service

import (
	"github.com/skyzhouzj/skyCloud/internal/api/repo"
	"github.com/skyzhouzj/skyCloud/internal/api/repo/authorized_repo"
	"github.com/skyzhouzj/skyCloud/pkg/core"
)

type SearchData struct {
	Page              int    `json:"page"`               // 第几页
	PageSize          int    `json:"page_size"`          // 每页显示条数
	BusinessKey       string `json:"business_key"`       // 调用方key
	BusinessSecret    string `json:"business_secret"`    // 调用方secret
	BusinessDeveloper string `json:"business_developer"` // 调用方对接人
	Remark            string `json:"remark"`             // 备注
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (listData []*authorized_repo.Authorized, err error) {

	page := searchData.Page
	if page == 0 {
		page = 1
	}

	pageSize := searchData.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

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
		Limit(pageSize).
		Offset(offset).
		OrderById(false).
		QueryAll(s.db.GetDb().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
