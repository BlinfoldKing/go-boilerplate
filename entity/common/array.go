package common

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StrArr string array
type StrArr []string

// FromDB create object from db
func (arr *StrArr) FromDB(bts []byte) error {
	if len(bts) == 0 {
		return nil
	}

	str := string(bts)
	if strings.HasPrefix(str, "{") {
		str = str[1:]
	}

	if strings.HasSuffix(str, "}") {
		str = str[0 : len(str)-1]
	}

	items := strings.Split(str, ",")
	str = ""
	for _, item := range items {
		s := fmt.Sprintf(`"%s"`, item)
		if str == "" {
			str = s
		} else {
			str += ", " + s
		}
	}

	str = fmt.Sprintf(`{"data": [%s]}`, str)

	var data map[string]StrArr
	err := json.Unmarshal([]byte(str), &data)
	*arr = data["data"]
	return err
}

// ToDB create db string
func (arr *StrArr) ToDB() ([]byte, error) {
	str := ""
	for _, item := range *arr {
		if str == "" {
			str = item
		} else {
			str += ", " + item
		}
	}

	return []byte(fmt.Sprintf("{%s}", str)), nil
}
