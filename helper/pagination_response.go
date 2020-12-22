package helper

import (
	"encoding/json"
	"fmt"
	"go-boilerplate/entity"
	"net/url"

	"github.com/kataras/iris/v12"
)

// CreatePaginationResponse create pagination helper
func CreatePaginationResponse(ctx iris.Context, request entity.Pagination, list interface{}, count int) Response {
	response := CreateResponse(ctx)

	next := ctx.GetCurrentRoute().Path()
	prev := ctx.GetCurrentRoute().Path()

	data := make(map[string]interface{})
	data["count"] = count
	data["list"] = list

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

		if prevPag.Offset != nil && *prevPag.Offset >= 0 {
			data["prev_url"] = prev
		}

		if nextPag.Offset != nil && *nextPag.Offset < count {
			data["next_url"] = next
		}

	case entity.CursorPagination:
		nextPag := request.(entity.CursorPagination)
		prevPag := request.(entity.CursorPagination)

		listJSON, _ := json.Marshal(list)
		var list []map[string]interface{}
		json.Unmarshal(listJSON, &list)

		firstID, lastID := "", ""
		nextSeek, prevSeek := "next", "prev"

		if len(list) > 0 {
			firstID = list[0]["id"].(string)
			lastID = list[len(list)-1]["id"].(string)
		}

		nextPag.ID = lastID
		nextPag.Seek = &nextSeek
		prevPag.ID = firstID
		prevPag.Seek = &prevSeek

		nextQuery, err := json.Marshal(nextPag)
		if err != nil {
			return CreateErrorResponse(ctx, err).InternalServer()
		}
		next = fmt.Sprintf("%s?type=cursor&query=%s", next, url.QueryEscape(string(nextQuery)))

		prevQuery, err := json.Marshal(prevPag)
		if err != nil {
			return CreateErrorResponse(ctx, err).InternalServer()
		}
		prev = fmt.Sprintf("%s?type=cursor&query=%s", prev, url.QueryEscape(string(prevQuery)))

		data["next_url"] = next
		data["prev_url"] = prev
		// nextPag.ID = &data[0]["id"]
	}

	return response.Ok().WithData(data)
}
