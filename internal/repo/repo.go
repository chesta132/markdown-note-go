package repo

import (
	"fmt"
	"strings"
)

type Where struct {
	N   string // Field to filter | REQUIRED
	V   any    // Value of N 			| REQUIRED
	Ind string // Where indent 		| e.g. ("OR", "NOT") 											| Default: "AND"
	Op  string // Where operator  | e.g. ("=", ">", "<", "LIKE", "IN", etc) | Default: "="
}

type Query struct {
	Q string
	V []any
}

func BuildWhere(where []Where) Query {
	query := ""
	value := []any{}
	for i, w := range where {
		indent := ""
		operator := "="

		if i != 0 {
			if w.Ind != "" {
				indent = fmt.Sprintf(" %v ", w.Ind)
			} else {
				indent = " AND "
			}
		}

		if w.Op != "" {
			operator = w.Op
		}

		if strings.ToUpper(w.Op) == "IN" {
			placeholders := strings.Repeat("?,", len(w.V.([]any)))
			placeholders = placeholders[:len(placeholders)-1]
			query += fmt.Sprintf("%v%v IN (%v)", indent, w.N, placeholders)
			value = append(value, w.V.([]any)...)
			continue
		}

		query = query + fmt.Sprintf("%v%v %v ?", indent, w.N, operator)
		value = append(value, w.V)
	}
	return Query{query, value}
}
