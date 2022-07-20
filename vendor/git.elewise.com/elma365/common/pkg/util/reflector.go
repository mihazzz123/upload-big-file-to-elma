package util

import (
	"encoding"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/pkg/errors"
)

// Reflector рефлексия вокруг структурных тэгов
type Reflector struct {
	stype reflect.Type
	value reflect.Value
}

// NewReflector — конструктор
func NewReflector(i interface{}) Reflector {
	r := Reflector{
		stype: reflect.TypeOf(i),
		value: reflect.ValueOf(i),
	}
	if r.stype.Kind() == reflect.Ptr {
		r.stype = r.stype.Elem()
		r.value = r.value.Elem()
	}

	return r
}

// FromValue конструирует рефлектор по рефлексии объекта
func FromValue(v reflect.Value) Reflector {
	r := Reflector{
		stype: v.Type(),
		value: v,
	}
	if r.stype.Kind() == reflect.Ptr {
		r.stype = r.stype.Elem()
		r.value = r.value.Elem()
	}

	return r
}

// Value возвращает объект рефлексии
func (r Reflector) Value() interface{} {
	return r.value.Addr().Interface()
}

type extractConfig struct {
	tagName      string
	skipEmbedded bool
	skipEmpty    bool
	skipMinus    bool
}

// ExtractOption опция обработчика
type ExtractOption interface {
	Apply(extractConfig) extractConfig
}

type extractOptionFunc func(extractConfig) extractConfig

func (f extractOptionFunc) Apply(cfg extractConfig) extractConfig {
	return f(cfg)
}

// WithoutEmbedded игнорировать встроенные структуры
func WithoutEmbedded() ExtractOption {
	return extractOptionFunc(func(cfg extractConfig) extractConfig {
		cfg.skipEmbedded = true

		return cfg
	})
}

// WithoutEmpty игнорировать поля без тэга
func WithoutEmpty() ExtractOption {
	return extractOptionFunc(func(cfg extractConfig) extractConfig {
		cfg.skipEmpty = true

		return cfg
	})
}

// WithoutMinus игнорировать поля с тэгом равным "-"
func WithoutMinus() ExtractOption {
	return extractOptionFunc(func(cfg extractConfig) extractConfig {
		cfg.skipMinus = true

		return cfg
	})
}

// ExtractValues возвращает хеш-таблицу, где значению тэга сопоставлено значение поля
func (r Reflector) ExtractValues(tagName string, skipNils bool, opts ...ExtractOption) map[string]interface{} {
	tags := r.ExtractTags(tagName, opts...)
	res := make(map[string]interface{}, len(tags))
	for fieldName, tag := range tags {
		val := r.value.FieldByName(fieldName)
		if skipNils {
			if val.Kind() == reflect.Ptr && val.IsNil() || val.Kind() == reflect.Slice && val.Len() == 0 {
				continue
			}
		}
		res[tag] = val.Interface()
	}
	return res
}

// ExtractTags возвращает хеш-таблицу, где имени поля сопоставлено значение тэга
func (r Reflector) ExtractTags(tagName string, opts ...ExtractOption) map[string]string {
	cfg := extractConfig{
		tagName: tagName,
	}
	for _, opt := range opts {
		cfg = opt.Apply(cfg)
	}

	return r.extractTags(cfg, r.stype, map[string]string{})
}

func (r Reflector) extractTags(cfg extractConfig, t reflect.Type, m map[string]string) map[string]string {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag, _ := field.Tag.Lookup(cfg.tagName)
		if !r.isExtractable(cfg, field, tag) {
			continue
		}
		if field.Anonymous {
			r.extractTags(cfg, field.Type, m)
		} else {
			m[field.Name] = tag
		}
	}
	return m
}

func (r Reflector) isExtractable(cfg extractConfig, field reflect.StructField, tag string) bool {
	isExported := func() bool {
		ch, _ := utf8.DecodeRuneInString(field.Name)
		return unicode.IsUpper(ch)
	}
	isEmbeddedAndShouldBeSkipped := func() bool {
		return field.Anonymous && cfg.skipEmbedded
	}
	isEmptyTagAndShouldBeSkipped := func() bool {
		return !field.Anonymous && tag == "" && cfg.skipEmpty
	}
	isMinusTagAndShouldBeSkipped := func() bool {
		return tag == "-" && cfg.skipMinus
	}
	return isExported() &&
		!isEmbeddedAndShouldBeSkipped() &&
		!isEmptyTagAndShouldBeSkipped() &&
		!isMinusTagAndShouldBeSkipped()
}

// Apply применяет данные из переданной хеш-таблицы на структуру, вокруг которой построен рефлектор
func (r Reflector) Apply(m map[string]string) error {
	s := r.value
	for k, v := range m {
		f := s.FieldByName(k)
		if err := r.processValue(f, v); err != nil {
			ft, _ := r.stype.FieldByName(k)

			return errors.WithMessage(err, ft.Name)
		}
	}

	return nil
}

// nolint:gocyclo
func (r Reflector) processValue(value reflect.Value, source string) error {
	t := value.Type()
	if source == "" {
		value.Set(reflect.Zero(t))

		return nil
	}
	dst := reflect.New(t).Interface()
	if unmarshaler, ok := dst.(encoding.TextUnmarshaler); ok {
		if err := unmarshaler.UnmarshalText([]byte(source)); err != nil {
			return errors.WithStack(err)
		}
		value.Set(reflect.ValueOf(dst).Elem())

		return nil
	}
	if t.PkgPath() == "time" && t.Name() == "Duration" {
		d, err := time.ParseDuration(source)
		if err != nil {
			return errors.WithStack(err)
		}
		value.SetInt(int64(d))

		return nil
	}
	if t.Kind() == reflect.Slice {
		sources := strings.Split(source, ",")
		values := reflect.MakeSlice(t, len(sources), len(sources))
		for i, val := range sources {
			err := r.processValue(values.Index(i), val)
			if err != nil {
				return err
			}
		}
		value.Set(values)

		return nil
	}
	if t.Kind() == reflect.String {
		value.SetString(source)

		return nil
	}
	if _, err := fmt.Sscan(source, dst); err != nil {
		return errors.WithStack(err)
	}

	value.Set(reflect.ValueOf(dst).Elem())

	return nil
}
