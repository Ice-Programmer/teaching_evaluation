package utils

const (
	DefaultPageNum  = 1
	DefaultPageSize = 10
	MaxPageSize     = 100
)

func SetPageDefault(pageSize, pageNum int32) (int32, int32) {
	if pageNum <= 0 {
		pageNum = DefaultPageNum
	}
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	return pageSize, pageNum
}
