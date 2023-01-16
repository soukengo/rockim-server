package shared

const (
	PageSizeDefault = 20
	PageSizeMax     = 100
)

func (x *Paginating) Limit() int64 {
	if x == nil {
		return PageSizeDefault
	}
	if x.PageSize > 0 && x.PageSize <= PageSizeMax {
		return int64(x.PageSize)
	}
	return PageSizeDefault
}

func (x *Paginating) Offset() int64 {
	if x == nil {
		return 0
	}
	page := x.PageNo
	if page <= 0 {
		page = 1
	}
	return (int64(page - 1)) * x.Limit()
}
