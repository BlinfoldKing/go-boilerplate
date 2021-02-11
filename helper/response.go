package helper

import (
	"fmt"
	"reflect"

	"github.com/kataras/iris/v12"

	"bytes"
	"encoding/json"
	"sort"
)

// Content Reponse content implemented using map slice
type Content struct {
	Key   string
	Value interface{}
}

// ContentMap Reponse content implemented using map slice
type ContentMap []Content

// CreateContentMap create new content map
func CreateContentMap(m map[string]interface{}, selections ...string) (newMap ContentMap) {
	if len(selections) > 0 {
		for _, s := range selections {
			newMap = newMap.Add(Content{s, m[s]})
		}
	} else {
		for key, val := range m {
			newMap = newMap.Add(Content{key, val})
		}
	}

	return
}

// Add new entry
func (ms ContentMap) Add(content Content) ContentMap {
	newMap := append(ms, content)
	return newMap
}

// Len get len
func (ms ContentMap) Len() int { return len(ms) }

// Swap swap content
func (ms ContentMap) Swap(i, j int) { ms[i], ms[j] = ms[j], ms[i] }

// Less custom comparator
func (ms ContentMap) Less(i, j int) bool {
	priority := map[string]int{
		"status":     0,
		"id":         0,
		"key":        0,
		"message":    1,
		"prev_url":   2,
		"next_url":   3,
		"count":      4,
		"data":       5,
		"created_at": 1000,
		"updated_at": 1001,
		"deleted_at": 1002,
	}

	var p1, p2 = 999, 999

	k1 := reflect.ValueOf(ms[i].Value).Kind()
	k2 := reflect.ValueOf(ms[j].Value).Kind()
	if val, ok := priority[ms[i].Key]; ok {
		p1 = val
	} else if k1 == reflect.Slice || k1 == reflect.Map {
		p1 = 1003
	}

	if val, ok := priority[ms[j].Key]; ok {
		p2 = val
	} else if k2 == reflect.Slice || k2 == reflect.Map {
		p2 = 1003
	}

	return p1 < p2
}

// MarshalJSON toJSON
func (ms ContentMap) MarshalJSON() ([]byte, error) {
	sort.Sort(ms)

	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for i, mi := range ms {
		b, err := json.Marshal(&mi.Value)
		if err != nil {
			return nil, err
		}

		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprintf("%v", mi.Key)))
		buf.Write(b)
		if i < len(ms)-1 {
			buf.Write([]byte{','})
		}
	}

	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

// Response represent http reponse
type Response struct {
	ok      bool
	content ContentMap
	context iris.Context
}

// CreateResponse create new response
func CreateResponse(ctx iris.Context) Response {
	response := Response{}

	response.context = ctx
	response.content = ContentMap{}

	return response
}

// Ok http 200
func (response Response) Ok() Response {
	response.content = response.content.Add(Content{"status", 200})
	return response.WithMessage("ok")
}

// WithData set data
func (response Response) WithData(data interface{}) Response {
	response.content = response.content.Add(Content{"data", data})
	return response
}

// WithField set custom field
func (response Response) WithField(field string, data interface{}) Response {
	response.content = response.content.Add(Content{field, data})
	return response
}

// WithStatus set status
func (response Response) WithStatus(status int) Response {
	response.content = response.content.Add(Content{"status", status})
	return response
}

// WithMessage set message
func (response Response) WithMessage(message string) Response {
	response.content = response.content.Add(Content{"message", message})
	return response
}

// JSON send response as JSON
func (response Response) JSON() {
	Logger.
		WithField("content", response).
		Debug()

	response.context.JSON(response.content)
}
