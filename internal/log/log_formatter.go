package log

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// TxtFormatter 文本格式
type TxtFormatter struct {
	SortKeys   []string `json:"sort_keys"`
	TimeFormat string   `json:"time_format"`
	kvFormat   string   // k:v格式，默认是`key:val`
}

func NewCustomTextFormatter(sortKeys []string, timeFormat string) *TxtFormatter {
	return &TxtFormatter{
		SortKeys:   sortKeys,
		TimeFormat: timeFormat,
		kvFormat:   `"%s":"%v"`,
	}
}

// Format 文本格式
//  1. 将entry里面的
//
// LogField,
func (t *TxtFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Level,Time,Caller,Message,Err
	timeFormat := time.RFC3339
	if t.TimeFormat != "" {
		timeFormat = t.TimeFormat
	}
	userMapData := map[string]any{
		FieldLevel: strings.ToUpper(entry.Level.String()),
		FieldTime:  entry.Time.Format(timeFormat),
		FieldMsg:   entry.Message,
	}
	// WithFields
	for k, v := range entry.Data {
		userMapData[k] = v
	}

	// logData -> slice: 直接字符串拼接
	ss := make([]string, 0, len(entry.Data)+6)
	for k, v := range userMapData {
		ss = append(ss, fmt.Sprintf(t.kvFormat, k, v))
	}

	// 无排序，字符串顺序排序
	if t.SortKeys == nil {
		sort.Strings(ss)
		return []byte(strings.Join(ss, "|") + "\n"), nil
	}

	// 有排序，需要针对userMapData，按SortKey排序字段排序，剩余的userMapData按字符串顺序排序
	ex := make(map[string]struct{})
	sortSs := make([]string, 0, len(t.SortKeys))
	for _, key := range t.SortKeys {
		if v, ok := userMapData[key]; ok {
			sortSs = append(sortSs, fmt.Sprintf(t.kvFormat, key, v))
			ex[key] = struct{}{}
		}
	}

	// 剩余map内容排序
	var sortSs2 []string
	for k, v := range userMapData {
		if _, ok := ex[k]; ok {
			continue
		}
		sortSs2 = append(sortSs2, fmt.Sprintf(t.kvFormat, k, v))
	}
	sort.Strings(sortSs2)

	// 组装
	sortSs = append(sortSs, sortSs2...)
	return []byte(strings.Join(sortSs, "|") + "\n"), nil
}
