package patch

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"text/template"

	"git.elewise.com/elma365/common/pkg/errs"

	"github.com/Jeffail/gabs"
	"github.com/pkg/errors"
	"github.com/vporoshok/pogo"
	"go.uber.org/zap"
)

//nolint: gochecknoglobals // шаблоны должны быть глобальными
var poTemplate = template.Must(template.New("po").Parse(`msgid ""
msgstr ""
{{- range .Items }}

msgctxt "{{ .Path }}"
msgid {{ .Value }}
msgstr {{ .Value }}
{{- end }}
`))

var ruABC = regexp.MustCompile(`[а-яёА-ЯЁ]+`)

// Максимальный размер строки выгружаемого ресурса. Берется как максимальный размер сканнера при считывании ресурсов,
// минус небольшой запас на символы, которые добавляются в строку помимо ключа и значения
const maxResourseItemSize = bufio.MaxScanTokenSize - 100

type poItem struct {
	Path  string
	Value string
}

// CheckExtracts - проверка полноты экстрактов
func CheckExtracts(blob json.RawMessage, extracts Extracts) ([]string, error) {
	container, err := gabs.ParseJSON(blob)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "parse JSON"))
	}

	replaces := make(Replaces, len(extracts))
	for i := range extracts {
		replaces[i] = extracts[i].ToReplace("")
		if replaces[i].path.Search(container) == nil {
			return nil, errors.WithStack(errors.New(fmt.Sprintf("pat %s not found", replaces[i].path)))
		}
	}

	testBlob, err := replaces.Apply(blob)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ruABC.FindAllString(string(testBlob), 10), nil
}

// ExtractPO извлекает ресурсы в po-файл
func ExtractPO(blob json.RawMessage, extracts Extracts) ([]byte, error) {
	res, _, err := extract(blob, extracts)
	return res, err
}

// CheckLargeResourses проверяет есть ли слишком большие ресурсы, которые будут пропущены при выгрузке
func CheckLargeResourses(blob json.RawMessage, extracts Extracts) (bool, error) {
	_, largeItemsSkipped, err := extract(blob, extracts)
	return largeItemsSkipped, err
}

func extract(blob json.RawMessage, extracts Extracts) (result []byte, largeItemsSkipped bool, err error) {
	resources, err := extracts.Apply(blob)
	if err != nil {
		return nil, false, err
	}
	items := make([]poItem, 0, len(resources))
	largeItemsSkipped = false
	for path, val := range resources {
		s, ok := val.(string)
		if !ok || s == "" {
			continue
		}

		checkItem, _ := json.Marshal(s)
		if int32(len(checkItem)) > maxResourseItemSize {
			largeItemsSkipped = true
			zap.L().Warn(
				fmt.Sprintf("Some large resourses of '%s' was skipped", path),
			)
			continue
		}

		value, _ := json.Marshal(s)
		items = append(items, poItem{
			Path:  path,
			Value: string(value),
		})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Path < items[j].Path
	})

	res := &bytes.Buffer{}
	err = poTemplate.Execute(res, struct{ Items []poItem }{Items: items})

	return res.Bytes(), largeItemsSkipped, errors.WithStack(err)
}

// ReplacePO заменяет ресурсы в json-документе
func ReplacePO(blob json.RawMessage, po []byte) (json.RawMessage, error) {
	replacer := poReplacer{
		po:        po,
		blob:      blob,
		container: nil,
		poFile:    nil,
		replaces:  nil,
	}
	return replacer.Replace()
}

type poReplacer struct {
	po   []byte
	blob json.RawMessage

	container *gabs.Container
	poFile    *pogo.POFile
	replaces  Replaces
}

// Replace - применить po файл
func (replacer *poReplacer) Replace() (json.RawMessage, error) {
	err := errs.WithRecover(func() {
		replacer.replace()
	})
	return replacer.container.Bytes(), err
}

func (replacer *poReplacer) replace() {
	replacer.loadPoFile()
	replacer.initContainer()
	replacer.extractReplaces()
	if len(replacer.replaces) == 0 {
		return
	}
	replacer.applyReplaces()
}

func (replacer *poReplacer) loadPoFile() {
	var err error
	buf := bytes.NewReader(replacer.po)
	replacer.poFile, err = pogo.ReadPOFile(buf)
	if err != nil {
		panic(errors.WithStack(err))
	}
}

func (replacer *poReplacer) initContainer() {
	var err error
	replacer.container, err = gabs.ParseJSON(replacer.blob)
	if err != nil {
		panic(errors.WithStack(err))
	}
}

func (replacer *poReplacer) extractReplaces() {
	for i := range replacer.poFile.Entries {
		replacer.extractReplace(&replacer.poFile.Entries[i])
	}
}

func (replacer *poReplacer) extractReplace(entity *pogo.POEntry) {
	path, ok := replacer.checkEntityAndExtractPath(entity)
	if !ok {
		return
	}
	replacer.replaces = append(replacer.replaces, NewReplace(entity.MsgStr, path))
}

// checkEntityAndExtractPath - проверка, нужно ли использовать этот ресурс и получение Path
// реализовано в одном методе только чтобы не генерировать Path 2 раза
func (replacer *poReplacer) checkEntityAndExtractPath(entity *pogo.POEntry) (Path, bool) {
	if entity.Obsolete {
		return nil, false
	}
	if entity.MsgCtxt == "" || entity.MsgStr == "" {
		return nil, false
	}

	path, err := FromString(entity.MsgCtxt)
	if err != nil {
		panic(errors.WithStack(err))
	}

	if path.Search(replacer.container) == nil {
		return nil, false
	}

	return path, true
}

func (replacer *poReplacer) applyReplaces() {
	err := replacer.replaces.ApplyToContainer(replacer.container)
	if err != nil {
		panic(errors.WithStack(err))
	}
}
