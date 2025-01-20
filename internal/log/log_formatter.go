package log

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// WisdomTextLogFormatter 文本格式
type WisdomTextLogFormatter struct {
	SortKeys   []string `json:"sort_keys"`
	TimeFormat string   `json:"time_format"`
	kvFormat   string   // k:v格式，默认是`key:val`
}

func NewCustomTextFormatter(sortKeys []string, timeFormat string) *WisdomTextLogFormatter {
	return &WisdomTextLogFormatter{
		SortKeys:   sortKeys,
		TimeFormat: timeFormat,
		kvFormat:   `"%s":"%v"`,
	}
}

// Format 文本格式
func (t *WisdomTextLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Level,Time,Caller,Message,Err
	timeFormat := time.RFC3339
	if t.TimeFormat != "" {
		timeFormat = t.TimeFormat
	}
	kvLogData := map[string]any{
		FieldLevel: strings.ToUpper(entry.Level.String()),
		FieldTime:  entry.Time.Format(timeFormat),
		FieldMsg:   entry.Message,
	}
	// WithFields
	for k, v := range entry.Data {
		kvLogData[k] = v
	}

	// 若没有设置排序规则，采用log meta的首字符默认排序
	if t.SortKeys == nil {
		// logData -> slice: 直接字符串拼接
		ss := make([]string, 0, len(kvLogData))
		for k, v := range kvLogData {
			ss = append(ss, fmt.Sprintf(t.kvFormat, k, v))
		}
		sort.Strings(ss)
		return []byte(strings.Join(ss, "|") + "\n"), nil
	}

	// 若有排序，需要针对userMapData，按SortKey排序字段排序，剩余的userMapData按字符串顺序排序
	ex := make(map[string]struct{})
	sortedList1 := make([]string, 0, len(t.SortKeys))
	for _, key := range t.SortKeys {
		if v, ok := kvLogData[key]; ok {
			sortedList1 = append(sortedList1, fmt.Sprintf(t.kvFormat, key, v))
			ex[key] = struct{}{}
		}
	}

	// 剩余map内容按字符排序
	var sortedList2 []string
	for k, v := range kvLogData {
		if _, ok := ex[k]; ok {
			continue
		}
		sortedList2 = append(sortedList2, fmt.Sprintf(t.kvFormat, k, v))
	}
	sort.Strings(sortedList2)

	// 统一组装
	sortedList1 = append(sortedList1, sortedList2...)
	return []byte(strings.Join(sortedList1, "|") + "\n"), nil
}
