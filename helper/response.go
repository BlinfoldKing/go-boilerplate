package helper

import (
	"fmt"
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
type contentMap []Content

func (ms contentMap) Add(content Content) contentMap {
	newMap := append(ms, content)
	return newMap
}
func (ms contentMap) Len() int      { return len(ms) }
func (ms contentMap) Swap(i, j int) { ms[i], ms[j] = ms[j], ms[i] }
func (ms contentMap) Less(i, j int) bool {
	priority := map[string]int{
		"status":   0,
		"message":  1,
		"prev_url": 2,
		"next_url": 3,
		"data":     4,
	}

	var p1, p2 = 999, 999
	if val, ok := priority[ms[i].Key]; ok {
		p1 = val
	}

	if val, ok := priority[ms[j].Key]; ok {
		p2 = val
	}

	return p1 < p2
}

func (ms contentMap) MarshalJSON() ([]byte, error) {
	Logger.Info(ms)
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
	content contentMap
	context iris.Context
}

// CreateResponse create new response
func CreateResponse(ctx iris.Context) Response {
	response := Response{}

	response.context = ctx
	response.content = contentMap{}

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
