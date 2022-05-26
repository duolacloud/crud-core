package types

import (
	"encoding/base64"
	"fmt"
	"io"
	"strings"

	"github.com/vmihailenco/msgpack/v5"
)

type CursorDirection string

const (
	CursorDirectionAfter  CursorDirection = "AFTER"
	CursorDirectionBefore CursorDirection = "BEFORE"
)

type CursorQuery struct {
	Filter    map[string]any  `json:"filter"`    // 筛选条件
	Cursor    string          `json:"cursor"`    // 游标值
	Direction CursorDirection `json:"direction"` // 查询方向 Before：游标前 After：游标后
	Limit     int64           `json:"limit"`
	Fields    []string        `json:"fields"`
	Sort      []string        `json:"sort"`
}

type CursorExtra struct {
	HasPrevious bool   `json:"has_previous"` // 是否有更多数据
	HasNext     bool   `json:"has_next"`     // 是否有更多数据
	EndCursor   string `json:"end_cursor"`   // 结果集中的起始游标值
	StartCursor string `json:"start_cursor"` // 结果集中的结束游标值
}

type Cursor struct {
	// ID    any `msgpack:"i"`
	Value []any `msgpack:"v"`
}

func (c *Cursor) Unmarshal(s string) error {
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

func (c *Cursor) Marshal(w io.Writer) error {
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	return msgpack.NewEncoder(wc).Encode(c)
}
