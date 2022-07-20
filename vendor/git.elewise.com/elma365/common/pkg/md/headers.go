package md

// OutgoingHeaders интерфейс заголовков в которые можно сохранить метаданные
type OutgoingHeaders interface {
	Add(key, value string)
}

// IncomingHeaders интерфейс заголовков из которых можно извлечь метаданные
type IncomingHeaders interface {
	Values(key string) []string
	Range(func(key string, values []string) bool) bool
}
