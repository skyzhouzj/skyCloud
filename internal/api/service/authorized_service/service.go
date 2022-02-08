package authorized_service

import (
	"github.com/skyzhouzj/skyCloud/internal/api/repo/authorized_repo"
	"github.com/skyzhouzj/skyCloud/pkg/cache"
	"github.com/skyzhouzj/skyCloud/pkg/core"
	"github.com/skyzhouzj/skyCloud/pkg/db"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, authorizedData *CreateAuthorizedData) (id int32, err error)
	List(ctx core.Context, searchData *SearchData) (listData []*authorized_repo.Authorized, err error)
	PageList(ctx core.Context, searchData *SearchData) (listData []*authorized_repo.Authorized, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx core.Context, id int32, used int32) (err error)
	Delete(ctx core.Context, id int32) (err error)
	Detail(ctx core.Context, id int32) (info *authorized_repo.Authorized, err error)
	DetailByKey(ctx core.Context, key string) (data *CacheAuthorizedData, err error)

	CreateAPI(ctx core.Context, authorizedAPIData *CreateAuthorizedAPIData) (id int32, err error)
	DeleteAPI(ctx core.Context, id int32) (err error)
}

type service struct {
	db    db.Repo
	cache cache.Repo
}

func New(db db.Repo, cache cache.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}
