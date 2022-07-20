package i18n

import (
	"bufio"
	"bytes"
	json "encoding/json"
	"text/template"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Максимальный размер строки выгружаемого ресурса. Берется как максимальный размер сканнера при считывании ресурсов,
// минус небольшой запас на символы, которые добавляются в строку помимо ключа и значения
const maxResourseItemSize = bufio.MaxScanTokenSize - 100

// шаблон перевода
// nolint: gochecknoglobals // шаблон должен быть глобальным
var poTemplate = template.Must(template.New("localizer").Parse(`msgctxt "{{ .Context }}"
msgid {{ .Value }}
msgstr ""

`))

// ExtractPO формирует для локализуемой строки перевод
// nolint: interfacer // нужен именно тип "POContext", не надо тут никаких "fmt.Stringer"
func ExtractPO(value interface{}, poCtxt *EntityPOContext) ([]byte, error) {
	res := &bytes.Buffer{}

	s, ok := value.(string)
	if !ok || s == "" {
		return res.Bytes(), nil
	}

	v, _ := json.Marshal(s)
	if int32(len(v)) > maxResourseItemSize {
		zap.L().Warn("Some large resourses was skipped")
		return res.Bytes(), nil
	}
	s = string(v)
	if s == "" {
		return res.Bytes(), nil
	}

	err := poTemplate.Execute(res,
		struct {
			Context string
			Value   string
		}{
			Context: poCtxt.String(),
			Value:   s,
		},
	)
	return res.Bytes(), errors.WithStack(err)
}
