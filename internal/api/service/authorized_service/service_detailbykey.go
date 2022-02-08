package authorized_service

import (
	"encoding/json"
	"github.com/skyzhouzj/skyCloud/configs"
	"github.com/skyzhouzj/skyCloud/internal/api/repo"
	"github.com/skyzhouzj/skyCloud/internal/api/repo/authorized_api_repo"
	"github.com/skyzhouzj/skyCloud/internal/api/repo/authorized_repo"
	"github.com/skyzhouzj/skyCloud/pkg/cache"
	"github.com/skyzhouzj/skyCloud/pkg/core"
	"time"
)

// CacheAuthorizedData 缓存结构
type CacheAuthorizedData struct {
	Key    string         `json:"key"`     // 调用方 key
	Secret string         `json:"secret"`  // 调用方 secret
	IsUsed int32          `json:"is_used"` // 调用方启用状态 1=启用 -1=禁用
	Apis   []cacheApiData `json:"apis"`    // 调用方授权的 Apis
}

type cacheApiData struct {
	Method string `json:"method"` // 请求方式
	Api    string `json:"api"`    // 请求地址
}

func (s *service) DetailByKey(ctx core.Context, key string) (cacheData *CacheAuthorizedData, err error) {
	// 查询缓存
	cacheKey := configs.Get().SkyCloud.RedisKeyPrefixSignature + key

	if !s.cache.Exists(cacheKey) {
		// 查询调用方信息
		authorizedInfo, err := authorized_repo.NewQueryBuilder().
			WhereIsDeleted(repo.EqualPredicate, -1).
			WhereBusinessKey(repo.EqualPredicate, key).
			First(s.db.GetDb().WithContext(ctx.RequestContext()))

		if err != nil {
			return nil, err
		}

		// 查询调用方授权 API 信息
		authorizedApiInfo, err := authorized_api_repo.NewQueryBuilder().
			WhereIsDeleted(repo.EqualPredicate, -1).
			WhereBusinessKey(repo.EqualPredicate, key).
			OrderById(false).
			QueryAll(s.db.GetDb().WithContext(ctx.RequestContext()))

		if err != nil {
			return nil, err
		}

		// 设置缓存 data
		cacheData = new(CacheAuthorizedData)
		cacheData.Key = key
		cacheData.Secret = authorizedInfo.BusinessSecret
		cacheData.IsUsed = authorizedInfo.IsUsed
		cacheData.Apis = make([]cacheApiData, len(authorizedApiInfo))

		for k, v := range authorizedApiInfo {
			data := cacheApiData{
				Method: v.Method,
				Api:    v.Api,
			}
			cacheData.Apis[k] = data
		}

		cacheDataByte, _ := json.Marshal(cacheData)

		err = s.cache.Set(cacheKey, string(cacheDataByte), time.Hour*24, cache.WithTrace(ctx.Trace()))
		if err != nil {
			return nil, err
		}

		return cacheData, nil
	}

	value, err := s.cache.Get(cacheKey, cache.WithTrace(ctx.RequestContext().Trace))
	if err != nil {
		return nil, err
	}

	cacheData = new(CacheAuthorizedData)
	err = json.Unmarshal([]byte(value), cacheData)
	if err != nil {
		return nil, err
	}

	return

}
