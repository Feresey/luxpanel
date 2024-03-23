// DO NOT EDIT. This file was auto-generated

package {{.PackageName}}

import (
	"fmt"
	"regexp"
	"time"
	"strconv"
)

var {{.Regexp.Name}} = regexp.MustCompile(`{{.Regexp.Value}}`)

type {{.TypeName}} struct {
	{{- range $i, $field := .Fields}}
	{{$field.Name}} {{$field.Type.Name}}
	{{- end}}
}

func (c *{{.TypeName}}) Unmarhsal(src string, now time.Time) (err error) {
	res := {{.Regexp.Name}}.FindStringSubmatch(src)
	if len(res) != {{.Regexp.TotalMatches}} {
		return fmt.Errorf("%w: %d", ErrWrongLineFormat, len(res))
	}
{{range $i, $field := .Fields}}
{{- if $field.Type.IsString}}
	c.{{$field.Name}} = res[{{$field.RegexpIndex}}]
{{- else}}
	c.{{$field.Name}}, err = parseField(res[{{$field.RegexpIndex}}], "{{$field.Name}}", {{$field.Type.ParseFunc}})
	if err != nil {
		return err
	}
{{- end}}
{{- end}}

	return nil
}

func (c *{{.TypeName}}) Type() {{.PackageName | camelcase}}LineType {
	return {{.TypeName}}LineType
}

func (c *{{.TypeName}}) Time() time.Time {
	return c.LogTime
}
