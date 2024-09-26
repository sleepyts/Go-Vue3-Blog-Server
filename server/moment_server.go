package server

import (
	"Go-Vue3-Blog-Server/globalVar"
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/models/vo"
	"Go-Vue3-Blog-Server/utils/redis_util"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func GetMomentByPage(page int, ip string) gin.H {
	var (
		moments   []*entity.Moment
		momentsVO []vo.MomentVO
		total     int64
		wg        = new(sync.WaitGroup)
	)
	pageSize := 7
	total, _ = entity.GetMomentCount()
	if page <= 0 || (page-1)*pageSize >= int(total) {
		return respose.ErrorWithMsg("Page 参数错误")
	}
	moments, _ = entity.GetMomentByPage(page, pageSize)
	wg.Add(len(moments))
	momentsVO = make([]vo.MomentVO, len(moments))
	for index, moment := range moments {
		go func(moment *entity.Moment, index int) {
			defer wg.Done()
			var commentVO vo.MomentVO
			copier.Copy(&commentVO, moment)
			commentVO.LikeCount = uint(len(globalVar.RedisDb.SMembers(redis_util.Key1(redis_util.MOMENT_LIKE_KEY, moment.Id)).Val()))
			commentVO.IsLike = globalVar.RedisDb.SIsMember(redis_util.Key1(redis_util.MOMENT_LIKE_KEY, moment.Id), ip).Val()
			momentsVO[index] = commentVO
		}(moment, index)
	}
	wg.Wait()
	res := respose.NewPageRes(total, momentsVO)
	return respose.Sucess(res)
}

func LikeOrUnlikeMoment(id int, ip string) gin.H {
	var (
		key = redis_util.Key1(redis_util.MOMENT_LIKE_KEY, id)
	)
	if globalVar.RedisDb.SIsMember(key, ip).Val() {
		globalVar.RedisDb.SRem(key, ip)
	} else {
		globalVar.RedisDb.SAdd(key, ip)
	}
	return respose.Sucess(nil)
}
