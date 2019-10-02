package helloworld

import (
	"fmt"
	"time"
)

func yyyyMMddHHmmss(t time.Time) string {
	const layout = "2006-01-02 15:04:05"
	return t.Format(layout)
}

// Hello は、挨拶を返します.
//
// 引数 t が指定されている場合、メッセージの後ろに指定された日付を付与します. (yyyy-MM-dd HH:mm:ss)
// t が nil の場合、現在時刻を付与します.
func Hello(t *time.Time) (string, error) {
	now := time.Now()
	if t != nil {
		now = *t
	}

	return fmt.Sprintf("Hello World at %s", yyyyMMddHHmmss(now)), nil
}
