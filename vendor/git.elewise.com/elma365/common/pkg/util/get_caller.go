package util

import (
	"path"
	"runtime"
	"strings"
)

// Caller описание вызывающего
//
// Название файла и функции или тип.метод
type Caller struct {
	File, Function string
}

// GetCaller возвращает первого вызывающего за пределами указанного файла
//
// По сути берётся стектрейс вызова и в нём берётся первая запись, которая не содержит указанный файл (файл указывается
// без расширения .go), из этой записи возвращается имя функции (метода). Более подробно смотри в тестах.
func GetCaller(exclude ...string) Caller {
	excludes := make([]string, len(exclude))
	for i, ex := range exclude {
		excludes[i] = ex + ".go"
	}
	pc := make([]uintptr, 10)
	n := runtime.Callers(2, pc)
	pc = pc[:n]
	frames := runtime.CallersFrames(pc)
	for {
		frame, ok := frames.Next()
		if !ok {
			return Caller{"", "unknown"}
		}
		name := frame.Function
		if len(frame.File) > 0 {
			hasSuffix := false
			for _, suffix := range excludes {
				hasSuffix = strings.HasSuffix(frame.File, suffix)
				if hasSuffix {
					break
				}
			}
			if hasSuffix {
				continue
			}
			file := strings.TrimSuffix(path.Base(frame.File), ".go")
			name = path.Base(name)
			chunks := strings.SplitN(name, ".", 2)
			if len(chunks) != 2 {
				return Caller{file, name}
			}

			return Caller{file, chunks[1]}
		}
	}
}
