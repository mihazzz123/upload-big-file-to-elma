package patch

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/pkg/errors"
)

// PathPart кусочек пути в json-документе
type PathPart interface {
	Get(*gabs.Container) *gabs.Container
	Set(*gabs.Container, interface{}) error
	String() string
}

// NewPart — конструктор
func NewPart(x interface{}) PathPart {
	switch tp := x.(type) {
	default:
		panic(errors.Errorf("invalid type (%[1]T) of path part %+[1]v", x))

	case int:

		return index(tp)

	case string:

		return key(tp)
	}
}

type index int

func (i index) Get(container *gabs.Container) *gabs.Container {
	return container.Index(int(i))
}

func (i index) Set(container *gabs.Container, value interface{}) error {
	_, err := container.SetIndex(value, int(i))

	return errors.WithStack(err)
}

func (i index) String() string {
	return fmt.Sprintf("[%d]", i)
}

type key string

func (k key) Get(container *gabs.Container) *gabs.Container {
	return container.Search(string(k))
}

func (k key) Set(container *gabs.Container, value interface{}) error {
	if container == nil {
		return errors.WithStack(errors.New(
			fmt.Sprintf(`undefined gabs container (key: "%s", value: "%s")`, k, value)),
		)
	}

	_, err := container.Set(value, string(k))

	return errors.WithStack(err)
}

func (k key) String() string {
	return fmt.Sprintf(".%s", string(k))
}

// Path путь внутри json-документа
type Path []PathPart

// NewPath построение пути по кусочкам
func NewPath(parts ...interface{}) Path {
	path := make(Path, 0, len(parts))
	for _, part := range parts {
		if partPath, ok := part.(Path); ok {
			path = append(path, partPath...)
		} else {
			path = append(path, NewPart(part))
		}
	}

	return path
}

// Варианты:
// .some-identifier   - точка плюс cтрока
// [123]              - число в скобках
// .[some-identifier] - точка плюс строка, не начинающаяся с цифры
// .                  - просто точка (идентификатор - пустая строка)
var partRE = regexp.MustCompile(`(\.[^.[]+|\[\d+\]|\.\[\D[^.[]+\]|\.)`)

// FromString преобразует строку в json-путь
func FromString(s string) (Path, error) {
	parts := partRE.FindAllString(s, -1)
	if strings.Join(parts, "") != s {
		return nil, errors.Errorf("invalid path: %q", s)
	}
	path := make(Path, 0, len(parts))
	for _, part := range parts {
		switch part[0] {
		case '.':
			path = append(path, key(part[1:]))

		case '[':
			i, err := strconv.ParseInt(part[1:len(part)-1], 10, 64)
			if err != nil {
				// Видимо, в скобках строка, а не число
				path = append(path, key(part[1:]))
				break
			}

			path = append(path, index(i))
		}
	}

	return path, nil
}

// String возвращает
func (path Path) String() string {
	builder := &strings.Builder{}

	for _, part := range path {
		_, _ = builder.WriteString(part.String())
	}

	return builder.String()
}

// Prefix path with given parts
func (path Path) Prefix(parts ...interface{}) Path {
	return append(NewPath(parts...), path...)
}

// Get value by path
func (path Path) Get(container *gabs.Container) interface{} {
	for _, part := range path {
		container = part.Get(container)
	}

	return container.Data()
}

// Set value by path
func (path Path) Set(container *gabs.Container, value interface{}) error {
	for _, part := range path[:len(path)-1] {
		container = part.Get(container)
		if container == nil {
			return nil
		}
	}

	return path[len(path)-1].Set(container, value)
}

// Search - поиск вложенного контейнера по пути
func (path Path) Search(container *gabs.Container) *gabs.Container {
	for i := range path {
		container = path[i].Get(container)
		if container == nil {
			return nil
		}
	}
	return container
}
