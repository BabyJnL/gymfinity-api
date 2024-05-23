package Library

import 	"strconv"

func ParseInt(param string) int {
	parsedData, err := strconv.Atoi(param);
	if err != nil {
		panic(err)
	}

	return parsedData
}