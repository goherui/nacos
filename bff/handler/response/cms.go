package response

type CmsCreate struct {
	Code int    `form:"code" json:"code" xml:"code"  binding:"required"`
	Msg  string `form:"msg" json:"msg" xml:"msg" binding:"required"`
}
type CmsDel struct {
	Code int    `form:"code" json:"code" xml:"code"  binding:"required"`
	Msg  string `form:"msg" json:"msg" xml:"msg" binding:"required"`
}
type CmsUpdate struct {
	Code int    `form:"code" json:"code" xml:"code"  binding:"required"`
	Msg  string `form:"msg" json:"msg" xml:"msg" binding:"required"`
}
type Cms struct {
	Id           int    `form:"id" json:"id" xml:"id"  binding:"required"`
	Title        string `form:"title" json:"title" xml:"title"  binding:"required"`
	CategoryName string `form:"categoryName" json:"categoryName" xml:"categoryName" binding:"required"`
	Content      string `form:"content" json:"content" xml:"content" binding:"required"`
	Status       int8   `form:"status" json:"status" xml:"status" binding:"required"`
	ViewCount    uint   `form:"viewCount" json:"viewCount" xml:"viewCount" binding:"required"`
	Creator      uint64 `form:"creator" json:"creator" xml:"creator" binding:"required"`
}
type CmsList struct {
	List []Cms  `form:"list" json:"list" xml:"list"  binding:"required"`
	Code int    `form:"code" json:"code" xml:"code"  binding:"required"`
	Msg  string `form:"msg" json:"msg" xml:"msg" binding:"required"`
}
