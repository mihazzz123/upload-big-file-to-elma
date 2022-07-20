package permissions

//go:generate ../../tooling/bin/enumer -json -transform=snake -type=AccessAtom -trimprefix=AccessAtom -output=access_string.go access.go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"unsafe"
)

// AccessAtom атомарное разрешение
type AccessAtom int

const (
	// READ чтение элемента
	READ AccessAtom = 1 << iota
	// CREATE создание элемента (в папке или справочнике)
	CREATE
	// UPDATE изменение элемента
	UPDATE
	// DELETE удаление элемента
	DELETE
	// ASSIGN изменение прав
	ASSIGN
	// BPMANAGE управление процессами элемента
	BPMANAGE
	// EXPORT экспорт данных
	EXPORT
	// IMPORT импорт данных
	IMPORT
)

// Access набор разрешений
type Access AccessAtom

const (
	// FullAccess Полный доступ
	FullAccess = Access(READ | CREATE | UPDATE | DELETE | ASSIGN | BPMANAGE | EXPORT | IMPORT)
	// FullFileAccess Полный доступ к файлу
	FullFileAccess = Access(READ | UPDATE | DELETE | ASSIGN | EXPORT | IMPORT)
	// FullDirAccess Полный доступ к директории
	FullDirAccess = Access(READ | CREATE | UPDATE | DELETE | ASSIGN | EXPORT | IMPORT)
	// EmptyAccess Отсутствие доступа
	EmptyAccess = Access(0)
)

// Has проверяет наличие атомарного разрешения в наборе
func (a Access) Has(t AccessAtom) bool {
	return a&Access(t) > 0
}

// MarshalJSON implements json.Marshaler
func (a Access) MarshalJSON() ([]byte, error) {
	buf := bytes.Buffer{}
	_ = buf.WriteByte('[')
	size := uint(unsafe.Sizeof(Access(0)) * 8) //nolint:gosec // необходимо для определения размера
	for i := uint(0); i < size; i++ {
		t := AccessAtom(1 << i)
		if a.Has(t) {
			_, _ = buf.WriteString(fmt.Sprintf("%q,", t))
		}
	}
	if buf.Len() > 1 {
		// Обрежем последнюю запятую
		buf.Truncate(buf.Len() - 1)
	}
	_ = buf.WriteByte(']')

	return buf.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (a *Access) UnmarshalJSON(data []byte) error {
	var atoms []AccessAtom

	if err := json.Unmarshal(data, &atoms); err != nil {
		return err
	}

	*a = 0
	for _, atom := range atoms {
		*a |= Access(atom)
	}

	return nil
}
