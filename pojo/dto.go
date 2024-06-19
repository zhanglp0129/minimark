package pojo

// GoodsPageDTO 接收条件分页查询前端传的参数
type GoodsPageDTO struct {
	PageNum    int     `form:"pageNum"`
	PageSize   int     `form:"pageSize"`
	CategoryID *int    `form:"categoryId"`
	GoodsName  *string `form:"goodsName"`
}
