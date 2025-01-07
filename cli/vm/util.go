package vm

import (
	"errors"
	"strings"
)

func ParseParameters(param string) ([]Data, error) {
	parts := strings.Split(param, " ")
	var params []Data
	i := 0
	for i < len(parts) {
		if parts[i] == "" {
			i += 1
			continue
		}
		if strings.HasPrefix(parts[i], "\"") {
			var str []string
			hasFoundString := false
			for i < len(parts) {
				str = append(str, parts[i])
				if strings.HasSuffix(parts[i], "\"") {
					hasFoundString = true
					break
				}
				i += 1
			}
			if !hasFoundString {
				return nil, errors.New("string ending not found")
			}
			strWithQuote := strings.Join(str, " ")
			params = append(params, NewData(DT_STRING, strWithQuote[1:len(strWithQuote)-1]))
		} else {
			params = append(params, NewData(DT_STRING, parts[i]))
		}
		i += 1
	}
	return params, nil
}