package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Log(v ...any) {
	if b, ok := v[0].([]byte); ok {
		dst := &bytes.Buffer{}
		if err := json.Indent(dst, b, "", "   "); err != nil {
			panic(err)
		}

		fmt.Println(dst.String())
	} else {
		s, _ := json.MarshalIndent(v, "", "   ")
		fmt.Println(string(s))
	}
}
