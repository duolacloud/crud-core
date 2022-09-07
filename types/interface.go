package types

import (
	"fmt"
	"sort"
	"strings"
)

type ID interface {
	int | int32 | int64 | string | any
}

func FormatID(id ID) string {
	switch v := id.(type) {
	case string:
		return v
	case int, int32, int64:
		return fmt.Sprintf("%d", v)
	case map[string]any:
		// 确保排序一致
		keys := make([]string, 0, len(v))
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		values := make([]string, 0, len(keys))
		for _, k := range keys {
			values = append(values, fmt.Sprintf("%v", v[k]))
		}
		return strings.Join(values, "|")
	default:
		return fmt.Sprintf("%v", v)
	}
}
