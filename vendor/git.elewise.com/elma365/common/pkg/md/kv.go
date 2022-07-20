package md

import (
	"context"
	"strings"

	"github.com/opentracing/opentracing-go"
)

type kvContextKey struct{}

type kvRecord struct {
	key   string
	value string
	next  *kvRecord
}

// AddKV добавить строковую пару ключ-значение в контекст
//
// Ключи приводятся к нижнему регистру. Ключу соответствует список значений,
// поэтому при добавлении нескольких значений по одному ключу они все будут
// сохранены в контексте.
func AddKV(ctx context.Context, key, value string) context.Context {
	key = strings.ToLower(key)
	head, _ := ctx.Value(kvContextKey{}).(*kvRecord)
	head = &kvRecord{key, value, head}
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span.LogKV(key, value)
	}
	return context.WithValue(ctx, kvContextKey{}, head)
}

// ListV извлечь массив значений по ключу
func ListV(ctx context.Context, key string) (res []string) {
	key = strings.ToLower(key)
	for head, _ := ctx.Value(kvContextKey{}).(*kvRecord); head != nil; head = head.next {
		if head.key == key {
			res = append(res, head.value)
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// HasKV проверить наличие значения в списке по ключу
func HasKV(ctx context.Context, key, value string) bool {
	key = strings.ToLower(key)
	for head, _ := ctx.Value(kvContextKey{}).(*kvRecord); head != nil; head = head.next {
		if head.key == key && head.value == value {
			return true
		}
	}
	return false
}

// PeakV получить первое значение из списка по ключу
func PeakV(ctx context.Context, key string) (string, bool) {
	key = strings.ToLower(key)
	var res *string
	for head, _ := ctx.Value(kvContextKey{}).(*kvRecord); head != nil; head = head.next {
		if head.key == key {
			res = &head.value
		}
	}
	if res == nil {
		return "", false
	}
	return *res, true
}

// RangeKV обойти список всех пар в контексте
//
// Если visitor возвращает false, итерирование прекращается.
func RangeKV(ctx context.Context, visitor func(key string, value string) bool) {
	head, _ := ctx.Value(kvContextKey{}).(*kvRecord)
	for head != nil && visitor(head.key, head.value) {
		head = head.next
	}
}

// InjectKV записать пары в заголовки для последующего извлечения методом ExtractKV
func InjectKV(ctx context.Context, prefix string, headers OutgoingHeaders) {
	for head, _ := ctx.Value(kvContextKey{}).(*kvRecord); head != nil; head = head.next {
		headers.Add(prefix+head.key, head.value)
	}
}

// ExtractKV извлечь пары из заголовков, записанные методом InjectKV
func ExtractKV(ctx context.Context, prefix string, headers IncomingHeaders) context.Context {
	head, _ := ctx.Value(kvContextKey{}).(*kvRecord)
	headers.Range(func(key string, values []string) bool {
		if !strings.HasPrefix(key, prefix) {
			return true
		}
		key = strings.TrimPrefix(key, prefix)
		key = strings.ToLower(key)
		for i := len(values) - 1; i >= 0; i-- {
			head = &kvRecord{key, values[i], head}
		}
		return true
	})
	return context.WithValue(ctx, kvContextKey{}, head)
}
