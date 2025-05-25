package paging

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"{{.moduleName}}/pkg/utils/strs"

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
	GetPageState() []byte
	StartDateTime() time.Time
	EndDateTime() time.Time
	TimeRange() (time.Time, time.Time)
	Json() string
}

type _Paging struct {
	Page_number     int       `json:"page_number"`
	Page_size       int       `json:"page_size"`
	Page_state      []byte    `json:"page_state"`
	Search          string    `json:"search"`
	Order_by        string    `json:"order_by"`
	Order_dir       string    `json:"order_dir"`
	Start_date_time time.Time `json:"start_date_time"`
	End_date_time   time.Time `json:"end_date_time"`
}

func (p *_Paging) TimeRange() (time.Time, time.Time) {
	return p.Start_date_time, p.End_date_time
}

func (p *_Paging) EndDateTime() time.Time {
	return p.End_date_time
}

func (p *_Paging) StartDateTime() time.Time {
	return p.Start_date_time
}

func (p *_Paging) GetPageState() []byte {
	return p.Page_state
}

func New(pageNumber int, pageSize int) Paging {
	if pageNumber < 1 {
		pageNumber = 1
	}
	if pageSize < 1 {
		pageSize = 25
	}
	return &_Paging{
		Page_number: pageNumber,
		Page_size:   pageSize,
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
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = 25
	}

	nextToken := ctx.DefaultQuery("nextToken", "")

	if nextToken != "" {
		nextTokenBuff, err := strs.Base64Decode(nextToken)
		if err != nil {
			nextToken = ""
		} else {
			nextToken = string(nextTokenBuff)
		}

	}

	startDateTimeStr := ctx.DefaultQuery("start_time", "")
	endDateTimeStr := ctx.DefaultQuery("end_time", "")

	startDateTime, err := time.Parse(time.RFC3339, startDateTimeStr)
	if err != nil {
		startDateTime = time.Time{}
	}

	endDateTime, err := time.Parse(time.RFC3339, endDateTimeStr)
	if err != nil {
		endDateTime = time.Time{}
	}

	return &_Paging{
		Page_number:     int(pageNumber),
		Page_size:       int(pageSize),
		Order_by:        order,
		Order_dir:       orderDir,
		Search:          search,
		Page_state:      []byte(nextToken),
		Start_date_time: startDateTime,
		End_date_time:   endDateTime,
	}
}

func (p *_Paging) PageNumber() int {
	return p.Page_number
}

func (p *_Paging) PageSize() int {
	return p.Page_size
}

func (p *_Paging) LimitOffset() (int, int) {
	return p.Page_size, (p.Page_number - 1) * p.Page_size
}

func (p *_Paging) SetSearch(search string) {
	p.Search = search
}

func (p *_Paging) GetSearch() string {
	return p.Search
}

func (p *_Paging) GetOrderBy() (string, bool) {
	return p.Order_by, p.Order_by != ""
}

func (p *_Paging) GetOrder() (string, bool) {
	return p.Order_dir, p.Order_dir != ""
}

func (p *_Paging) Json() string {
	byts, _ := json.Marshal(p)
	return string(byts)
}
