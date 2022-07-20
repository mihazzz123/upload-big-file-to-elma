package patch

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

const patchTagName = "patch"
const patchJSONTagName = "json"

// Extracter - интерфейс объекта, который сам создает экстракты
type Extracter interface {
	GetExtracts(target string) (Extracts, error)
}

// CreateExtracts - пытается создать экстракты анализируя структуру
//
// Если значение реализует интерфейс patch.Extracter, то он имеет приоритет
func CreateExtracts(val interface{}, target string) (Extracts, error) {
	e, ok := val.(Extracter)
	if ok {
		return e.GetExtracts(target)
	}

	s := reflect.ValueOf(val)

	return fromReflectValue(s, target)
}

func fromReflectValue(s reflect.Value, target string) (Extracts, error) {
	if !s.IsValid() {
		return nil, nil
	}
	e, ok := s.Interface().(Extracter)
	if ok {
		return e.GetExtracts(target)
	}

	switch s.Kind() {
	case reflect.Map:
		return fromMap(s, target)
	case reflect.Slice:
		return fromSlice(s, target)
	case reflect.Struct:
		return fromStruct(s, target)
	case reflect.Ptr:
		return fromReflectValue(s.Elem(), target)
	case reflect.String:
		return Extracts{NewExtract()}, nil
	default:
		return nil, nil
	}
}

func fromMap(s reflect.Value, target string) (Extracts, error) {
	keys := s.MapKeys()
	if len(keys) == 0 {
		return nil, nil
	}
	if keys[0].Kind() != reflect.String {
		return nil, nil
	}
	extracts := Extracts{}
	for i := range keys {
		val := s.MapIndex(keys[i])
		subExtracts, err := fromReflectValue(val, target)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		prefixedSubExtracts := subExtracts.Prefix(keys[i].String())
		extracts = append(extracts, prefixedSubExtracts...)
	}
	return extracts, nil
}

func fromSlice(s reflect.Value, target string) (Extracts, error) {
	extracts := Extracts{}
	for i := 0; i < s.Len(); i++ {
		subExtracts, err := fromReflectValue(s.Index(i), target)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		prefixedSubExtracts := subExtracts.Prefix(i)
		extracts = append(extracts, prefixedSubExtracts...)
	}
	return extracts, nil
}

func fromStruct(s reflect.Value, target string) (Extracts, error) {
	sType := s.Type()
	extracts := Extracts{}
	for i := 0; i < s.NumField(); i++ {
		subExtracts, err := fromStructField(s, sType.Field(i), target)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		extracts = append(extracts, subExtracts...)
	}
	return extracts, nil
}

func fromStructField(s reflect.Value, field reflect.StructField, target string) (Extracts, error) {
	fieldName := field.Name
	patchTags := field.Tag.Get(patchTagName)
	jsonTags := field.Tag.Get(patchJSONTagName)
	subVal := s.FieldByName(fieldName)

	// Анонимные поля (родительские структуы) без тэгов patch и json обрабатываются на этом же уровне
	if field.Anonymous && patchTags == "" && jsonTags == "" {
		subExtracts, err := fromReflectValue(subVal, target)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return subExtracts, nil
	}

	if !tagContainsValue(patchTags, target) {
		return Extracts{}, nil
	}
	patchName, ok := getPatchName(field)
	if !ok {
		return Extracts{}, nil
	}

	switch subVal.Kind() {
	case reflect.String:
		return fromString(subVal, patchName), nil

	default:
		subExtracts, err := fromReflectValue(subVal, target)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return subExtracts.Prefix(patchName), nil
	}
}

func fromString(val reflect.Value, patchName string) Extracts {
	// В случае отсутствия значения у филда не добавляем экстракты.
	// Тк у некоторых филдов используется тег `omitempty`, бывают случаи,
	// что поля с нулевыми значениями не вносятся в JSON, но присутствуют в экстрактах.
	// Это вызывает ошибку при проверке полноты экстрактов в методе patch.CheckExtracts
	if val.String() == "" {
		return nil
	}
	return Extracts{NewExtract(patchName)}
}

func tagContainsValue(tags, value string) bool {
	if strings.TrimSpace(tags) == "" {
		return false
	}
	tagList := strings.Split(tags, ",")
	for i := range tagList {
		if strings.TrimSpace(tagList[i]) == strings.TrimSpace(value) {
			return true
		}
	}
	return false
}

func getPatchName(f reflect.StructField) (string, bool) {
	tag := f.Tag
	jsonTag := tag.Get(patchJSONTagName)
	jsonTagList := strings.Split(jsonTag, ",")
	jsonTagPatchName := strings.TrimSpace(jsonTagList[0])
	if jsonTagPatchName != "" {
		if jsonTagPatchName == "-" {
			return "", false
		}
		return jsonTagPatchName, true
	}
	return f.Name, true
}
