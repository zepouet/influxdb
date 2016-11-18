package escape

import "strings"

var (
	Codes = map[byte][]byte{
		',': []byte(`\,`),
		'"': []byte(`\"`),
		' ': []byte(`\ `),
		'=': []byte(`\=`),
	}
	escaper   = strings.NewReplacer(`,`, `\,`, `"`, `\"`, ` `, `\ `, `=`, `\=`)
	unescaper = strings.NewReplacer(`\,`, `,`, `\"`, `"`, `\ `, ` `, `\=`, `=`)
)

func UnescapeString(in string) string {
	if strings.IndexByte(in, '\\') == -1 {
		return in
	}
	return unescaper.Replace(in)
}

func String(in string) string {
	return escaper.Replace(in)
}
