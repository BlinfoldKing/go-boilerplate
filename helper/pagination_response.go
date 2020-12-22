package helper

import (
	"encoding/json"
	"fmt"
	"go-boilerplate/entity"
	"net/url"

	"github.com/kataras/iris/v12"
)

// CreatePaginationResponse create pagination helper
func CreatePaginationResponse(ctx iris.Context, request entity.Pagination, list interface{}) Response {
	response := CreateResponse(ctx)

	next := ctx.GetCurrentRoute().Path()
	prev := ctx.GetCurrentRoute().Path()

	switch request.(type) {
	case entity.OffsetPagination:
		nextPag := request.(entity.OffsetPagination)
		prevPag := request.(entity.OffsetPagination)

		limit := 10
		if nextPag.Limit != nil {
			limit = *nextPag.Limit
		}

		if nextPag.Offset != nil {
			nextOffset := *nextPag.Offset + limit
			prevOffset := *prevPag.Offset - limit
			nextPag.Offset = &nextOffset
			prevPag.Offset = &prevOffset
		} else {
			nextPag.Offset = &limit
			prevPag.Offset = &limit
		}

		nextQuery, err := json.Marshal(nextPag)
		if err != nil {
			return CreateErrorResponse(ctx, err).InternalServer()
		}
		next = fmt.Sprintf("%s?type=offset&query=%s", next, url.QueryEscape(string(nextQuery)))

		prevQuery, err := json.Marshal(prevPag)
		if err != nil {
			return CreateErrorResponse(ctx, err).InternalServer()
		}
		prev = fmt.Sprintf("%s?type=offset&query=%s", prev, url.QueryEscape(string(prevQuery)))

		fmt.Println(string(nextQuery))
		fmt.Println(string(prevQuery))
	case entity.CursorPagination:
		next = fmt.Sprintf("%s?type=cursor", next)
	}

	data := make(map[string]interface{})
	data["list"] = list
	data["next_url"] = next
	data["prev_url"] = prev

	return response.Ok().WithData(data)
}
