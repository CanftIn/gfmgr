package common

// PageReq 公共请求参数
type PageReq struct {
	DataRange []string `p:"dataRange"` // 日期范围
	PageNum   int      `p:"pageNum"`   // 当前页码
	PageSize  int      `p:"pageSize"`  // 每页数
	OrderKey  string   `p:"orderKey"`  // 排序字段
}

type Author struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
}
