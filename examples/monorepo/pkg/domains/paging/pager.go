package paging

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Paging interface {
	PageNumber() int
	PageSize() int
	LimitOffset() (int, int)
	SetSearch(search string)
	GetSearch() string
	GetOrderBy() (string, bool)
	GetOrder() (string, bool)
}

type _Paging struct {
	pageNumber int
	pageSize   int
	search     string
	orderBy    string
	orderDir   string
}

func New(pageNumber int, pageSize int) Paging {
	if pageNumber < 1 {
		pageNumber = 1
	}
	if pageSize < 1 {
		pageSize = 25
	}
	return &_Paging{
		pageNumber: pageNumber,
		pageSize:   pageSize,
	}
}

func FromGinContext(
	ctx *gin.Context,
) Paging {
	// Ordering
	order := ctx.DefaultQuery("orderBy", "")
	orderDir := strings.ToLower(ctx.DefaultQuery("order", "desc"))

	// trying to prevent sql injection
	if orderDir != "asc" && orderDir != "desc" {
		orderDir = "desc"
	}

	search := ctx.DefaultQuery("search", "")

	// page numbers
	pageNumberStr := ctx.DefaultQuery("page", "1")
	pageSizeStr := ctx.DefaultQuery("size", "25")

	pageNumber, err := strconv.ParseInt(pageNumberStr, 10, 64)
	if err != nil {
		pageNumber = 1
		err = nil
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = 25
	}

	return &_Paging{
		pageNumber: int(pageNumber),
		pageSize:   int(pageSize),
		orderBy:    order,
		orderDir:   orderDir,
		search:     search,
	}
}

func (p *_Paging) PageNumber() int {
	return p.pageNumber
}

func (p *_Paging) PageSize() int {
	return p.pageSize
}

func (p *_Paging) LimitOffset() (int, int) {
	return p.pageSize, (p.pageNumber - 1) * p.pageSize
}

func (p *_Paging) SetSearch(search string) {
	p.search = search
}

func (p *_Paging) GetSearch() string {
	return p.search
}

func (p *_Paging) GetOrderBy() (string, bool) {
	return p.orderBy, p.orderBy != ""
}

func (p *_Paging) GetOrder() (string, bool) {
	return p.orderDir, p.orderDir != ""
}
