package request

type CmsCreate struct {
	Title      string `form:"title" json:"title" xml:"title"  binding:"required"`
	CategoryID uint64 `form:"categoryId" json:"categoryId" xml:"categoryId" binding:"required"`
	Content    string `form:"content" json:"content" xml:"content" binding:"required"`
	Status     int8   `form:"status" json:"status" xml:"status" binding:"required"`
	ViewCount  uint   `form:"viewCount" json:"viewCount" xml:"viewCount" binding:"required"`
	Creator    uint64 `form:"creator" json:"creator" xml:"creator" binding:"required"`
}
type CmsDel struct {
	Id int `form:"id" json:"id" xml:"id"  binding:"required"`
}
type CmsUpdate struct {
	Id         int    `form:"id" json:"id" xml:"id"  binding:"required"`
	Title      string `form:"title" json:"title" xml:"title"  binding:"required"`
	CategoryID uint64 `form:"categoryId" json:"categoryId" xml:"categoryId" binding:"required"`
	Content    string `form:"content" json:"content" xml:"content" binding:"required"`
	Status     int8   `form:"status" json:"status" xml:"status" binding:"required"`
	ViewCount  uint   `form:"viewCount" json:"viewCount" xml:"viewCount" binding:"required"`
	Creator    uint64 `form:"creator" json:"creator" xml:"creator" binding:"required"`
}
type CmsList struct {
	Page int `form:"page" json:"page" xml:"page"`
	Size int `form:"size" json:"size" xml:"size"`
}
