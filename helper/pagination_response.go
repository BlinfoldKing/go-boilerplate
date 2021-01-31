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
	t := ctx.URLParam("type")

	next := ctx.GetCurrentRoute().Path()
	prev := ctx.GetCurrentRoute().Path()

	data := make(map[string]interface{})
	data["count"] = count

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

		if t == "offset" {
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
		} else {
			baseQuery := "?type=inline_offset"
			group := ctx.URLParamDefault("group", "")
			if group != "" {
				baseQuery = fmt.Sprintf("%s&group=%s", baseQuery, group)
			}
			sort := ctx.URLParamDefault("sort", "")
			if sort != "" {
				baseQuery = fmt.Sprintf("%s&sort=%s", baseQuery, sort)
			}
			filter := ctx.URLParamDefault("filter", "")
			if filter != "" {
				baseQuery = fmt.Sprintf("%s&filter=%s", baseQuery, filter)
			}
			distinct := ctx.URLParamDefault("distinct", "")
			if distinct != "" {
				baseQuery = fmt.Sprintf("%s&distinct=%s", baseQuery, distinct)
			}
			selection := ctx.URLParamDefault("select", "")
			if selection != "" {
				baseQuery = fmt.Sprintf("%s&select=%s", baseQuery, selection)
			}

			skip := ctx.URLParamIntDefault("skip", 0)
			take := ctx.URLParamIntDefault("take", 10)

			if prevPag.Offset != nil && *prevPag.Offset >= 0 {
				data["prev_url"] = url.QueryEscape(fmt.Sprintf("%s%s&skip=%d&take=%d", prev, baseQuery, skip-take, take))
			}

			if nextPag.Offset != nil && *nextPag.Offset < count {
				data["next_url"] = url.QueryEscape(fmt.Sprintf("%s%s&skip=%d&take=%d", next, baseQuery, skip+take, take))
			}
		}

		req := request.(entity.OffsetPagination)

		if req.GroupBy != nil {
			newList := make([]map[string]interface{}, 0)
			b, _ := json.Marshal(list)
			json.Unmarshal(b, &newList)
			groupedList := formatGroup(req.Query, *req.GroupBy, newList)
			data["data"] = groupedList
		} else {
			newList := make([]map[string]interface{}, 0)
			b, _ := json.Marshal(list)
			json.Unmarshal(b, &newList)

			result := make([]interface{}, 0)
			for _, item := range newList {
				if req.Selection != nil {
					result = append(result, CreateContentMap(item, *req.Selection...))
				} else {
					result = append(result, CreateContentMap(item))
				}
			}

			data["data"] = result
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

		req := request.(entity.CursorPagination)

		if req.GroupBy != nil {
			newList := make([]map[string]interface{}, 0)
			b, _ := json.Marshal(list)
			json.Unmarshal(b, &newList)
			groupedList := formatGroup(req.Query, *req.GroupBy, newList)
			data["data"] = groupedList
		} else {
			newList := make([]map[string]interface{}, 0)
			b, _ := json.Marshal(list)
			json.Unmarshal(b, &newList)

			result := make([]interface{}, 0)
			for _, item := range newList {
				if req.Selection != nil {
					result = append(result, CreateContentMap(item, *req.Selection...))
				} else {
					result = append(result, CreateContentMap(item))
				}
			}

			data["data"] = result
		}
	}

	for k := range data {
		response = response.WithField(k, data[k])
	}

	return response.Ok()
}

func formatGroup(req entity.Query, groups []entity.PaginationGroup, data []map[string]interface{}) interface{} {
	if len(groups) > 0 {
		head := groups[0]
		tail := groups[1:]

		selector := head.Selector
		var currValue interface{} = ""

		groupItems := make([][]map[string]interface{}, 0)
		tempGroup := make([]map[string]interface{}, 0)
		for _, item := range data {
			if item[selector] == currValue {
				tempGroup = append(tempGroup, item)
			} else {
				groupItems = append(groupItems, tempGroup)
				tempGroup = []map[string]interface{}{
					item,
				}
				currValue = item[selector]
			}
		}
		groupItems = append(groupItems, tempGroup)

		groupItems = groupItems[1:]

		result := make([]map[string]interface{}, 0)
		for _, items := range groupItems {
			group := make(map[string]interface{})
			group["key"] = items[0][selector]
			group["count"] = len(items)
			group["summary"] = []interface{}{len(items)}
			if head.IsExpanded != nil && *head.IsExpanded {
				group["items"] = formatGroup(req, tail, items)
			} else {
				group["items"] = nil
			}
			result = append(result, group)
		}

		return result
		// items := formatGroup(tail, data)
	}

	result := make([]interface{}, 0)

	for _, item := range data {
		if req.Selection != nil {
			result = append(result, CreateContentMap(item, *req.Selection...))
		} else {
			result = append(result, CreateContentMap(item))
		}
	}
	return result
}
