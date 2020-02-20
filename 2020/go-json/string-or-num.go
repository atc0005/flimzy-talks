package main

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

// StringOrNum is a string that may be unmarshaled from either a JSON string
// or a JSON number.
type StringOrNum string

func (s *StringOrNum) UnmarshalJSON(d []byte) error {
	var v string
	if d[0] == '"' {
		if err := json.Unmarshal(d, &v); err != nil {
			return err
		}
	} else {
		v = string(d)
	}
	*s = StringOrNum(v)
	return nil
}

func main() {
	data := []byte(`[123,"321"]`)
	x := make([]StringOrNum, 0)
	if err := json.Unmarshal(data, &x); err != nil {
		panic(err)
	}
	spew.Dump(x)
}
