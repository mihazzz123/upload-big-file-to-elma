package util

// ExtractTagsFromStruct извлекает список тэгов по имени из структуры
//
// Подробнее смотри https://godoc.org/github.com/jmoiron/sqlx/reflectx#Mapper.FieldMap
func ExtractTagsFromStruct(tagName string, i interface{}) []string {
	r := NewReflector(i)
	m := r.ExtractTags(tagName, WithoutEmpty(), WithoutMinus())
	res := make([]string, 0, len(m))
	for _, tag := range m {
		res = append(res, tag)
	}

	return res
}

// StructToMapByTags преобразует структуру в хэш-таблицу по значениям тэга
//
// Подробнее смотри https://godoc.org/github.com/jmoiron/sqlx/reflectx#Mapper.FieldMap
func StructToMapByTags(tagName string, i interface{}, skipNils bool) map[string]interface{} {
	r := NewReflector(i)

	return r.ExtractValues(tagName, skipNils, WithoutEmpty(), WithoutMinus())
}
