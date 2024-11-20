package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type IntArray []int

func (a *IntArray) Scan(src interface{}) error {
	if src == nil {
		*a = nil
		return nil
	}

	var srcStr string
	switch v := src.(type) {
	case string:
		srcStr = v
	case []byte:
		srcStr = string(v)
	default:
		return fmt.Errorf("Não foi possivel scanear o tipo %T em IntArray: %v", src, src)
	}

	srcStr = strings.TrimPrefix(srcStr, "{")
	srcStr = strings.TrimSuffix(srcStr, "}")
	if srcStr == "" {
		*a = []int{}
		return nil
	}

	elements := strings.Split(srcStr, ",")
	result := make([]int, len(elements))
	for i, elem := range elements {
		val, err := strconv.Atoi(strings.TrimSpace(elem))
		if err != nil {
			return fmt.Errorf("Erro ao converter array para inteiro: %v", err)
		}
		result[i] = val
	}

	*a = result
	return nil
}
