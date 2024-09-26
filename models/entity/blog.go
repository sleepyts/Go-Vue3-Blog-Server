package entity

import (
	"Go-Vue3-Blog-Server/globalVar"
	"Go-Vue3-Blog-Server/models/vo"
	"sync"
	"time"

	"github.com/jinzhu/copier"
)

type Blog struct {
	Id          uint       `json:"id" gorm:"primary_key"`
	Title       string     `json:"title"`
	Img         string     `json:"img"`
	Content     string     `json:"content"`
	Description string     `json:"description"`
	CategoryId  uint       `json:"categoryId"`
	CreateTime  *time.Time `json:"createTime"`
	UpdateTime  *time.Time `json:"updateTime"`
}

func (Blog) TableName() string {
	return "tb_blog"
}
func GetBlogById(id uint) Blog {
	var blog Blog
	globalVar.Db.First(&blog, id)
	return blog
}
func GetRandomBlogVO() []vo.RandomBlogVO {
	var (
		randomBlogs = []Blog{}
		wg          = sync.WaitGroup{}
		result      = []vo.RandomBlogVO{}
	)
	globalVar.Db.Order("RAND()").Limit(3).Find(&randomBlogs)
	wg.Add(len(randomBlogs))
	result = make([]vo.RandomBlogVO, len(randomBlogs))
	for index, blog := range randomBlogs {
		go func(blog Blog, index int) {
			defer wg.Done()
			randomBlogVO := vo.RandomBlogVO{}
			copier.Copy(&randomBlogVO, blog)
			result[index] = randomBlogVO
		}(blog, index)
	}
	wg.Wait()
	return result
}

func GetBlogTotalCount() int64 {
	var count int64
	globalVar.Db.Model(Blog{}).Count(&count)
	return count
}
func GetBlogByPage(page, pageSize int) []vo.BlogVO {
	var (
		blogs   []Blog
		blogVOs []vo.BlogVO
		wg      *sync.WaitGroup
	)
	globalVar.Db.Order("create_time desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&blogs)

	wg = &sync.WaitGroup{}
	wg.Add(len(blogs))
	blogVOs = make([]vo.BlogVO, len(blogs))
	for index, blog := range blogs {
		go func(blog Blog, index int) {
			defer wg.Done()
			var blogVO vo.BlogVO
			copier.Copy(&blogVO, &blog)
			var commentNum int64
			globalVar.Db.Model(&Comment{}).Where("blog_id = ?", blogVO.Id).Count(&commentNum)
			blogVO.CommentNum = int(commentNum)
			globalVar.Db.Model(&Category{}).Select("name").Where("id = ?", blogVO.CategoryId).First(&blogVO.CategoryName)
			blogVOs[index] = blogVO
		}(blog, index)
	}
	wg.Wait()
	return blogVOs
}

func GetBlogByCategoryIdAndPage(categoryId, page, pageSize int) []vo.BlogVO {
	var (
		blogs   []Blog
		blogVOs []vo.BlogVO
		wg      *sync.WaitGroup
	)
	globalVar.Db.Order("create_time desc").Limit(pageSize).Offset((page-1)*pageSize).Where("category_id = ?", categoryId).Find(&blogs)

	wg = &sync.WaitGroup{}
	wg.Add(len(blogs))
	for _, blog := range blogs {
		go func(blog Blog) {
			var blogVO vo.BlogVO
			copier.Copy(&blogVO, &blog)
			var commentNum int64
			globalVar.Db.Model(&Comment{}).Where("blog_id = ?", blogVO.Id).Count(&commentNum)
			blogVO.CommentNum = int(commentNum)
			globalVar.Db.Model(&Category{}).Select("name").Where("id = ?", blogVO.CategoryId).First(&blogVO.CategoryName)
			blogVOs = append(blogVOs, blogVO)
			wg.Done()
		}(blog)
	}
	wg.Wait()
	return blogVOs
}

func GetBlogCountByCategoryId(categoryId int) int64 {
	var count int64
	globalVar.Db.Model(Blog{}).Where("category_id = ?", categoryId).Count(&count)
	return count
}

func GetRecordVO() []vo.RecordVO {
	var (
		RecordDTOs []vo.RecordDTO
		RecordVO   []vo.RecordVO
		Years      []string
	)
	globalVar.Db.Table("tb_blog").Select("DISTINCT YEAR(create_time) AS year").Order("year asc").Pluck("year", &Years)
	RecordVO = make([]vo.RecordVO, len(Years))
	for index, year := range Years {
		globalVar.Db.Table("tb_blog").Where("YEAR(create_time) = ?", year).Order("create_time desc").Find(&RecordDTOs)
		RecordVO[index].Year = year
		RecordVO[index].RecordList = RecordDTOs
	}
	return RecordVO
}
