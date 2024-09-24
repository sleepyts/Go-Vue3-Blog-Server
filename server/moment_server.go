package server

import (
	"Go-Vue3-Blog-Server/models/entity"
	"Go-Vue3-Blog-Server/models/respose"
	"Go-Vue3-Blog-Server/models/vo"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func GetMomentByPage(page int) gin.H {
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
	for _, moment := range moments {
		go func(moment *entity.Moment) {
			defer wg.Done()
			var ommentVO vo.MomentVO
			copier.Copy(&ommentVO, moment)
			// TODO 获取点赞数量
			momentsVO = append(momentsVO, ommentVO)
		}(moment)
	}
	wg.Wait()
	res := respose.NewPageRes(total, momentsVO)
	return respose.Sucess(res)
}
