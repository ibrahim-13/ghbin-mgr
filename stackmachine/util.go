package stackmachine

import (
	"errors"
	"strconv"
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
		} else if vi, err := strconv.Atoi(parts[i]); err != nil {
			params = append(params, NewData(DT_INT, vi))
		} else if vf, err := strconv.ParseFloat(parts[i], 64); err != nil {
			params = append(params, NewData(DT_FLOAT, vf))
		} else {
			params = append(params, NewData(DT_STRING, parts[i]))
		}
		i += 1
	}
	return params, nil
}
