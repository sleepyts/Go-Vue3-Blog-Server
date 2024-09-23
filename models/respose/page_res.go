package respose

// PageRes 定义分页结果的结构
type PageRes[T any] struct {
	Total int64    `json:"total"` // 总记录数
	Rows  []T       `json:"rows"`  // 当前页的数据，具体类型的切片
}

// NewPageRes 创建一个新的分页结果实例
func NewPageRes[T any](total int64, rows []T) *PageRes[T] {
	return &PageRes[T]{
		Total: total,
		Rows:  rows,
	}
}
