package helper

import (
	"fmt"

	"github.com/kataras/iris/v12"

	"bytes"
	"encoding/json"
	"sort"
)

// Content Reponse Content implemented using map slice
type Content struct {
	Key   string
	Value interface{}
}

// ContentMap Reponse Content implemented using map slice
type ContentMap []Content

// CreateContentMap create new Content map
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
func (ms ContentMap) Add(Content Content) (newMap ContentMap) {
	switch Content.Value.(type) {
	case map[string]interface{}:
		Content.Value = CreateContentMap(Content.Value.(map[string]interface{}))
	case []interface{}:
		slice := Content.Value.([]interface{})
		value := make([]interface{}, 0)
		for _, item := range slice {
			switch item.(type) {
			case map[string]interface{}:
				v := CreateContentMap(item.(map[string]interface{}))
				value = append(value, v)
			default:
				v := item
				value = append(value, v)
			}

		}
		Content.Value = value
	}

	newMap = append(ms, Content)
	return newMap
}

// Len get len
func (ms ContentMap) Len() int { return len(ms) }

// Swap swap Content
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

	if val, ok := priority[ms[i].Key]; ok {
		p1 = val
	} else {
		switch ms[i].Value.(type) {
		case []interface{}, []map[string]interface{}:
			p1 = 1003
		}
	}

	if val, ok := priority[ms[j].Key]; ok {
		p2 = val
	} else {
		switch ms[i].Value.(type) {
		case []interface{}, []map[string]interface{}:
			p2 = 1003
		}
	}

	if p1 == p2 {
		return ms[i].Key < ms[j].Key
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
	Content ContentMap
	context iris.Context
}

// CreateResponse create new response
func CreateResponse(ctx iris.Context) Response {
	response := Response{}

	response.context = ctx
	response.Content = ContentMap{}

	return response
}

// Ok http 200
func (response Response) Ok() Response {
	response.Content = response.Content.Add(Content{"status", 200})
	return response.WithMessage("ok")
}

// WithData set data
func (response Response) WithData(data interface{}) Response {
	response.Content = response.Content.Add(Content{"data", data})
	return response
}

// WithField set custom field
func (response Response) WithField(field string, data interface{}) Response {
	response.Content = response.Content.Add(Content{field, data})
	return response
}

// WithStatus set status
func (response Response) WithStatus(status int) Response {
	response.Content = response.Content.Add(Content{"status", status})
	return response
}

// WithMessage set message
func (response Response) WithMessage(message string) Response {
	response.Content = response.Content.Add(Content{"message", message})
	return response
}

// JSON send response as JSON
func (response Response) JSON() {
	Logger.
		WithField("Content", response.Content).
		Debug()

	response.context.JSON(response.Content)
}
