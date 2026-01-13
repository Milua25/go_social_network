package store

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PaginatedFeedQuery struct {
	Limit  int      `json:"limit" validate:"gte=1,lte=20"`
	Offset int      `json:"offset" validate:"gte=0"`
	Sort   string   `json:"sort" validate:"oneof=asc desc"`
	Tags   []string `json:"tags" validate:"max=5"`
	Search string   `json:"search"`
	Since  string   `json:"since"`
	Until  string   `json:"until"`
}

func (fq PaginatedFeedQuery) Parse(req *http.Request) (PaginatedFeedQuery, error) {
	qs := req.URL.Query()
	limit := qs.Get("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return fq, err
		}
		fq.Limit = l
	}

	offset := qs.Get("offset")
	if offset != "" {
		of, err := strconv.Atoi(offset)
		if err != nil {
			return fq, err
		}
		fq.Offset = of
	}

	sort := qs.Get("sort")
	if sort != "" {
		fq.Sort = sort
	}

	tags := qs.Get("tags")
	if tags != "" {
		fq.Tags = strings.Split(tags, ",")
	}

	search := qs.Get("search")
	if search != "" {
		fq.Search = search
	}

	since := qs.Get("since")
	if since != "" {
		fq.Since = parseTime(since)
		fmt.Println(fq.Since)
	}

	until := qs.Get("until")
	if since != "" {
		fq.Until = parseTime(until)
		fmt.Println(fq.Until)
	}

	return fq, nil

}

func parseTime(s string) string {
	layouts := []string{
		time.RFC3339,
		"2006-01-02 15:04:05", // no zone
		"2006-01-02",          // date only
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			t.Format(time.RFC3339)
		}
	}
	return ""

}

// Normalize to RFC3339 (e.g., 2006-01-02T15:04:05Z07:00) to match typical Postgres timestamptz text.
