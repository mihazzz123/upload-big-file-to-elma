package generator

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"io"
	"reflect"
	"strings"

	"git.elewise.com/elma365/easylocalizer/bootstrap"
)

// Generator генератор Go файлов локализации.
type Generator struct {
	out *bytes.Buffer

	pkgName     string
	pkgPath     string
	hashString  string
	structInfos []StructInfo

	imports map[string]string

	visitedFields []string
}

// StructInfo содержит информацию о структуре.
type StructInfo struct {
	Type              reflect.Type
	Name              string
	LocalizationAlias string
}

const (
	tagJSON           = "json"
	tagLocalizer      = "localizer"
	tagLocalizerID    = "id"
	tagLocalizerValue = "value"
	tagLocalizerSkip  = "-"
)

// New возвращает новый экземпляр генератора.
func New(fileName string) *Generator {
	gen := &Generator{
		imports: map[string]string{
			bootstrap.PkgGoText:     "",
			bootstrap.PkgCommonI18n: "",
			bootstrap.PkgErrors:     "",
		},
	}
	hash := fnv.New32()
	hash.Write([]byte(fileName))
	gen.hashString = fmt.Sprintf("%x", hash.Sum32())

	return gen
}

// SetPkg устанавливает название и выходной путь до пакета.
func (g *Generator) SetPkg(name, path string) {
	g.pkgName = name
	g.pkgPath = path
}

// AddStruct добавляет структуру для генерации
func (g *Generator) AddStruct(name, alias string, obj interface{}) {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	g.structInfos = append(g.structInfos, StructInfo{
		Type:              t,
		Name:              name,
		LocalizationAlias: alias,
	})
}

// Run запускает генератор и выводит сгенерированный файл с кодом в out.
func (g *Generator) Run(out io.Writer) error {
	g.out = &bytes.Buffer{}

	for _, si := range g.structInfos {
		if err := g.writeExtractPO(si); err != nil {
			return err
		}
		if err := g.writeApplyTranslation(si); err != nil {
			return err
		}
	}

	g.writeHeader(out)
	_, err := out.Write(g.out.Bytes())

	return err
}

func (g *Generator) writeExtractPO(info StructInfo) error {
	if info.Type.Kind() != reflect.Struct {
		return fmt.Errorf("cannot generate localizer for %v, not a struct type", info.Type)
	}

	g.visitedFields = make([]string, 0)

	fmt.Fprintln(g.out)
	fmt.Fprintln(g.out, "func (entity ", info.Name, ") ExtractPO(poCtxt *i18n.EntityPOContext) ([]byte, error) {")
	g.imports["bytes"] = ""
	fmt.Fprintln(g.out, "  res := &bytes.Buffer{}")
	fmt.Fprintln(g.out)
	g.writeEntityContext(info)
	fmt.Fprintln(g.out)
	if err := g.writeEntityFieldsExtractPO(info); err != nil {
		return err
	}
	fmt.Fprintln(g.out, "  return res.Bytes(), nil")
	fmt.Fprintln(g.out, "}")

	return nil
}

func (g *Generator) writeEntityContext(info StructInfo) {
	fmt.Fprintln(g.out, "  // Контекст сущности")
	fmt.Fprintln(g.out, "  entityCtxt := &i18n.EntityPOContext {")
	fmt.Fprintf(g.out, "    Name: %q,\n", info.LocalizationAlias)

	idFields := g.getLocalizerIDFields(info)
	if len(idFields) > 0 {
		formatStr := make([]string, len(idFields))
		argsStr := make([]string, len(idFields))
		for i, idField := range idFields {
			switch idField.Type.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				formatStr[i] = "%d"
			case reflect.Float32, reflect.Float64:
				formatStr[i] = "%f"
			case reflect.Bool:
				formatStr[i] = "%t"
			default:
				formatStr[i] = "%s"
			}
			argsStr[i] = fmt.Sprintf("entity.%s", idField.Name)
		}
		g.imports["fmt"] = ""
		fmt.Fprintf(g.out, "    ID: fmt.Sprintf(%q, %s),\n", strings.Join(formatStr, "."), strings.Join(argsStr, ", "))
	}
	fmt.Fprintln(g.out, "    Parent: poCtxt,")
	fmt.Fprintln(g.out, "  }")
	fmt.Fprintln(g.out, "  _ = entityCtxt")
}

func (g *Generator) getLocalizerIDFields(info StructInfo) (idFields []reflect.StructField) {
	t := info.Type
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if tagValue, ok := field.Tag.Lookup(tagLocalizer); ok {
			switch tagValue {
			case tagLocalizerID:
				// g.visitedFields = append(g.visitedFields, fmt.Sprintf("%s.%s", info.Name, field.Name))
				idFields = append(idFields, field)
			default:
				// ничего не делаем
			}
		}
	}

	return idFields
}

func (g *Generator) writeEntityFieldsExtractPO(info StructInfo) error {
	for i := 0; i < info.Type.NumField(); i++ {
		field := info.Type.Field(i)
		fieldFullName := fmt.Sprintf("%s.%s", info.Name, field.Name)

		isVisited := false
		for _, visited := range g.visitedFields {
			if visited == fieldFullName {
				isVisited = true
				break
			}
		}
		if isVisited {
			continue
		}

		if tagValue, ok := field.Tag.Lookup(tagLocalizer); ok {
			switch tagValue {
			case tagLocalizerValue:
				g.visitedFields = append(g.visitedFields, fieldFullName)
				// Допускается отмечать тегом `localizer:"value"` только строковые поля.
				// Если отмечено поле другого типа - кидаем ошибку.
				if field.Type.Kind() == reflect.String {
					if err := g.writeStringFieldExtractPO(field); err != nil {
						return err
					}
					continue
				} else {
					return fmt.Errorf("cannot generate localizer for %v, not a string type", field)
				}
			case tagLocalizerSkip:
				g.visitedFields = append(g.visitedFields, fieldFullName)
				continue // Поле не локализуется, пропускаем
			default:
				// Ничего не делаем
			}
		}

		fieldType := field.Type
		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}
		switch fieldType.Kind() {
		case reflect.Struct:
			g.visitedFields = append(g.visitedFields, fieldFullName)
			if err := g.writeStructFieldExtractPO(field); err != nil {
				return err
			}
		case reflect.Map:
			g.visitedFields = append(g.visitedFields, fieldFullName)
			if err := g.writeMapFieldExtractPO(field); err != nil {
				return err
			}
		case reflect.Slice, reflect.Array:
			g.visitedFields = append(g.visitedFields, fieldFullName)
			if err := g.writeArrayFieldExtractPO(field); err != nil {
				return err
			}
		default:
			continue // Ничего не делаем
		}
	}

	return nil
}

func (g *Generator) writeStringFieldExtractPO(field reflect.StructField) error {
	fieldName, ok := field.Tag.Lookup(tagJSON)
	if ok {
		fieldName = strings.Split(fieldName, ",")[0]
	} else {
		fieldName = field.Name
	}

	fmt.Fprintln(g.out)
	// Заворачиваем в отдельную область видимости чтобы не было пересечений по названиям переменных
	fmt.Fprintln(g.out, "  {")
	fmt.Fprintf(g.out, "    // Генерация перевода для поля %q\n", field.Name)
	fmt.Fprintln(g.out, "    fieldCtxt := i18n.EntityPOContext {")
	fmt.Fprintf(g.out, "      Name: %q,\n", fieldName)
	fmt.Fprintln(g.out, "      Parent: entityCtxt,")
	fmt.Fprintln(g.out, "    }")
	fmt.Fprintf(g.out, "    b, err := i18n.ExtractPO(entity.%s, &fieldCtxt)\n", field.Name)
	fmt.Fprintln(g.out, "    if err != nil {return res.Bytes(), errors.WithStack(err)}")
	fmt.Fprintln(g.out, "    _, err = res.Write(b)")
	fmt.Fprintln(g.out, "    if err != nil {return res.Bytes(), errors.WithStack(err)}")
	fmt.Fprintln(g.out, "  }")

	return nil
}

func (g *Generator) writeStructFieldExtractPO(field reflect.StructField) error {
	fieldName, ok := field.Tag.Lookup(tagJSON)
	if ok {
		fieldName = strings.Split(fieldName, ",")[0]
	} else {
		fieldName = field.Name
	}
	accessor := fmt.Sprintf("entity.%s", field.Name)

	fmt.Fprintln(g.out)
	return g.writeStructFieldExtractPOWithAlias(
		accessor,
		field.Name,
		fieldName,
		"",
		field.Type.Kind() == reflect.Ptr,
		true,
		"  ",
	)
}

func (g *Generator) writeStructFieldExtractPOWithAlias(
	accessor, name, alias, ID string,
	isPtr bool, needComment bool,
	initSpace string,
) error {
	// Заворачиваем в отдельную область видимости чтобы не было пересечений по названиям переменных
	if isPtr {
		fmt.Fprintf(g.out, initSpace+"if %s != nil {\n", accessor)
	} else {
		fmt.Fprintln(g.out, initSpace+"{")
	}
	if needComment {
		fmt.Fprintf(g.out, initSpace+"  // Генерация перевода для поля %q\n", name)
	}
	fmt.Fprintln(g.out, initSpace+"  fieldCtxt := i18n.EntityPOContext {")
	fmt.Fprintf(g.out, initSpace+"    Name: %q,\n", alias)
	if ID != "" {
		fmt.Fprintf(g.out, initSpace+"    ID: %s,\n", ID)
	}
	fmt.Fprintln(g.out, initSpace+"    Parent: entityCtxt,")
	fmt.Fprintln(g.out, initSpace+"  }")
	fmt.Fprintln(g.out, initSpace+"  var v i18n.LocalizableEntity")
	fmt.Fprintln(g.out, initSpace+"  var ok bool")
	fmt.Fprintf(g.out, initSpace+"  v, ok = interface{}(%s).(i18n.LocalizableEntity)\n", accessor)
	fmt.Fprintln(g.out, initSpace+"  if !ok {")
	fmt.Fprintf(g.out, initSpace+"    v, ok = interface{}(&%s).(i18n.LocalizableEntity)\n", accessor)
	fmt.Fprintln(g.out, initSpace+"  }")
	fmt.Fprintln(g.out, initSpace+"  if ok {")
	fmt.Fprintln(g.out, initSpace+"    b, err := v.ExtractPO(&fieldCtxt)")
	fmt.Fprintln(g.out, initSpace+"    if err != nil {return res.Bytes(), errors.WithStack(err)}")
	fmt.Fprintln(g.out, initSpace+"    _, err = res.Write(b)")
	fmt.Fprintln(g.out, initSpace+"    if err != nil {return res.Bytes(), errors.WithStack(err)}")
	fmt.Fprintln(g.out, initSpace+"  }")
	fmt.Fprintln(g.out, initSpace+"}")

	return nil
}

func (g *Generator) writeMapFieldExtractPO(field reflect.StructField) error {
	elementType := field.Type.Elem().Kind()
	switch elementType {
	case reflect.Struct:
		fieldName, ok := field.Tag.Lookup(tagJSON)
		if ok {
			fieldName = strings.Split(fieldName, ",")[0]
		} else {
			fieldName = field.Name
		}
		fmt.Fprintln(g.out)
		// Заворачиваем в отдельную область видимости чтобы не было пересечений по названиям переменных
		fmt.Fprintln(g.out, "  {")
		fmt.Fprintf(g.out, "    // Генерация перевода для поля %q\n", field.Name)
		fmt.Fprintf(g.out, "    for key, value := range entity.%s {\n", field.Name)
		if err := g.writeStructFieldExtractPOWithAlias(
			"value",
			"",
			fieldName,
			"key",
			field.Type.Kind() == reflect.Ptr,
			false,
			"      ",
		); err != nil {
			return err
		}
		fmt.Fprintln(g.out, "    }")
		fmt.Fprintln(g.out, "  }")
	default:
		return nil
	}
	return nil
}

func (g *Generator) writeArrayFieldExtractPO(field reflect.StructField) error {
	elementType := field.Type.Elem().Kind()
	switch elementType {
	case reflect.Struct:
		fieldName, ok := field.Tag.Lookup(tagJSON)
		if ok {
			fieldName = strings.Split(fieldName, ",")[0]
		} else {
			fieldName = field.Name
		}
		fmt.Fprintln(g.out)
		// Заворачиваем в отдельную область видимости чтобы не было пересечений по названиям переменных
		fmt.Fprintln(g.out, "  {")
		fmt.Fprintf(g.out, "    // Генерация перевода для поля %q\n", field.Name)
		fmt.Fprintf(g.out, "    for _, value := range entity.%s {\n", field.Name)
		if err := g.writeStructFieldExtractPOWithAlias(
			"value",
			"",
			fieldName,
			"",
			field.Type.Kind() == reflect.Ptr,
			false,
			"      ",
		); err != nil {
			return err
		}
		fmt.Fprintln(g.out, "    }")
		fmt.Fprintln(g.out, "  }")
	default:
		return nil
	}
	return nil
}

func (g *Generator) writeApplyTranslation(info StructInfo) error {
	if info.Type.Kind() != reflect.Struct {
		return fmt.Errorf("cannot generate localizer for %v, not a struct type", info.Type)
	}

	g.visitedFields = make([]string, 0)

	fmt.Fprintln(g.out)
	fmt.Fprintln(g.out, "func (entity *", info.Name, ") ApplyTranslation(poCtxt *i18n.EntityPOContext, translator gotext.Translator) {")
	g.writeEntityContext(info)
	fmt.Fprintln(g.out)
	if err := g.writeEntityFieldsApplyTranslation(info); err != nil {
		return err
	}
	fmt.Fprintln(g.out, "}")

	return nil
}

func (g *Generator) writeEntityFieldsApplyTranslation(info StructInfo) error {
	for i := 0; i < info.Type.NumField(); i++ {
		field := info.Type.Field(i)
		fieldFullName := fmt.Sprintf("%s.%s", info.Name, field.Name)

		isVisited := false
		for _, visited := range g.visitedFields {
			if visited == fieldFullName {
				isVisited = true
				break
			}
		}
		if isVisited {
			continue
		}

		if tagValue, ok := field.Tag.Lookup(tagLocalizer); ok {
			switch tagValue {
			case tagLocalizerValue:
				g.visitedFields = append(g.visitedFields, fieldFullName)
				// Допускается отмечать тегом `localizer:"value"` только строковые поля.
				// Если отмечено поле другого типа - кидаем ошибку.
				if field.Type.Kind() == reflect.String {
					if err := g.writeStringFieldApplyTranslation(field); err != nil {
						return err
					}
					continue
				} else {
					return fmt.Errorf("cannot generate localizer for %v, not a string type", field)
				}
			case tagLocalizerSkip:
				g.visitedFields = append(g.visitedFields, fieldFullName)
				continue // Поле не локализуется, пропускаем
			default:
				// Ничего не делаем
			}
		}
		fieldType := field.Type
		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}
		switch fieldType.Kind() {
		case reflect.Struct:
			g.visitedFields = append(g.visitedFields, fieldFullName)
			if err := g.writeStructFieldApplyTranslation(field); err != nil {
				return err
			}
		case reflect.Map:
			g.visitedFields = append(g.visitedFields, fieldFullName)
			if err := g.writeMapFieldApplyTranslation(field); err != nil {
				return err
			}
		case reflect.Slice, reflect.Array:
			g.visitedFields = append(g.visitedFields, fieldFullName)
			if err := g.writeArrayFieldApplyTranslation(field); err != nil {
				return err
			}
		default:
			continue // Ничего не делаем
		}
	}
	return nil
}

func (g *Generator) writeStringFieldApplyTranslation(field reflect.StructField) error {
	fieldName, ok := field.Tag.Lookup(tagJSON)
	if ok {
		fieldName = strings.Split(fieldName, ",")[0]
	} else {
		fieldName = field.Name
	}

	fmt.Fprintln(g.out)
	// Заворачиваем в отдельную область видимости чтобы не было пересечений по названиям переменных
	fmt.Fprintln(g.out, "  {")
	fmt.Fprintf(g.out, "if entity.%s != \"\" {\n", field.Name)
	fmt.Fprintf(g.out, "      // Применяем перевод к полю %q\n", field.Name)
	fmt.Fprintln(g.out, "      fieldCtxt := i18n.EntityPOContext {")
	fmt.Fprintf(g.out, "        Name: %q,\n", fieldName)
	fmt.Fprintln(g.out, "        Parent: entityCtxt,")
	fmt.Fprintln(g.out, "      }")
	fmt.Fprintf(g.out, "  s := translator.GetC(entity.%s, fieldCtxt.String())\n", field.Name)
	fmt.Fprintf(g.out, "  if s == \"\" || s == entity.%s {s = translator.Get(entity.%s)}\n", field.Name, field.Name)
	fmt.Fprintf(g.out, "  if s != \"\" {entity.%s = s}\n", field.Name)
	fmt.Fprintln(g.out, "    }")
	fmt.Fprintln(g.out, "  }")

	return nil
}

func (g *Generator) writeStructFieldApplyTranslation(field reflect.StructField) error {
	fieldName, ok := field.Tag.Lookup(tagJSON)
	if ok {
		fieldName = strings.Split(fieldName, ",")[0]
	} else {
		fieldName = field.Name
	}
	accessor := fmt.Sprintf("entity.%s", field.Name)

	fmt.Fprintln(g.out)
	return g.writeStructFieldApplyTranslationWithAlias(
		accessor,
		field.Name,
		fieldName,
		"",
		field.Type.Kind() == reflect.Ptr,
		true,
		true,
		"  ",
	)
}

func (g *Generator) writeStructFieldApplyTranslationWithAlias(
	accessor, name, alias, ID string,
	isPtr, needComment, genFieldContext bool,
	initSpace string,
) error {
	// Заворачиваем в отдельную область видимости чтобы не было пересечений по названиям переменных
	if isPtr {
		fmt.Fprintf(g.out, initSpace+"if %s != nil {\n", accessor)
	} else {
		fmt.Fprintln(g.out, initSpace+"{")
	}

	if needComment {
		fmt.Fprintf(g.out, initSpace+"  // Генерация перевода для поля %q\n", name)
	}
	if genFieldContext {
		fmt.Fprintln(g.out, initSpace+"  fieldCtxt := i18n.EntityPOContext {")
		fmt.Fprintf(g.out, initSpace+"    Name: %q,\n", alias)
		if ID != "" {
			fmt.Fprintf(g.out, initSpace+"    ID: %s,\n", ID)
		}
		fmt.Fprintln(g.out, initSpace+"    Parent: entityCtxt,")
		fmt.Fprintln(g.out, initSpace+"  }")
	}
	fmt.Fprintln(g.out, initSpace+"  var v i18n.LocalizableEntity")
	fmt.Fprintln(g.out, initSpace+"  var ok bool")
	fmt.Fprintf(g.out, initSpace+"  v, ok = interface{}(%s).(i18n.LocalizableEntity)\n", accessor)
	fmt.Fprintln(g.out, initSpace+"  if !ok {")
	fmt.Fprintf(g.out, initSpace+"    v, ok = interface{}(&%s).(i18n.LocalizableEntity)\n", accessor)
	fmt.Fprintln(g.out, initSpace+"  }")
	fmt.Fprintln(g.out, initSpace+"  if ok {")
	fmt.Fprintln(g.out, initSpace+"    v.ApplyTranslation(&fieldCtxt, translator)")
	fmt.Fprintln(g.out, initSpace+"  }")
	fmt.Fprintln(g.out, initSpace+"}")

	return nil
}

func (g *Generator) writeMapFieldApplyTranslation(field reflect.StructField) error {
	elementType := field.Type.Elem().Kind()
	switch elementType {
	case reflect.Struct:
		fieldName, ok := field.Tag.Lookup(tagJSON)
		if ok {
			fieldName = strings.Split(fieldName, ",")[0]
		} else {
			fieldName = field.Name
		}
		fmt.Fprintln(g.out)
		// Заворачиваем в отдельную область видимости чтобы не было пересечений по названиям переменных
		fmt.Fprintln(g.out, "  {")
		fmt.Fprintf(g.out, "    // Генерация перевода для поля %q\n", field.Name)
		fmt.Fprintf(g.out, "    for key, value := range entity.%s {\n", field.Name)
		if err := g.writeStructFieldApplyTranslationWithAlias(
			"value",
			"",
			fieldName,
			"key",
			field.Type.Kind() == reflect.Ptr,
			false,
			true,
			"      ",
		); err != nil {
			return err
		}
		fmt.Fprintf(g.out, "      entity.%s[key] = value\n", field.Name)
		fmt.Fprintln(g.out, "    }")
		fmt.Fprintln(g.out, "  }")
	default:
		return nil
	}
	return nil
}

func (g *Generator) writeArrayFieldApplyTranslation(field reflect.StructField) error {
	elementType := field.Type.Elem().Kind()
	switch elementType {
	case reflect.Struct:
		fieldName, ok := field.Tag.Lookup(tagJSON)
		if ok {
			fieldName = strings.Split(fieldName, ",")[0]
		} else {
			fieldName = field.Name
		}
		fmt.Fprintln(g.out)
		// Заворачиваем в отдельную область видимости чтобы не было пересечений по названиям переменных
		fmt.Fprintln(g.out, "  {")
		fmt.Fprintf(g.out, "    // Генерация перевода для поля %q\n", field.Name)
		fmt.Fprintln(g.out, "    fieldCtxt := i18n.EntityPOContext {")
		fmt.Fprintf(g.out, "      Name: %q,\n", fieldName)
		fmt.Fprintln(g.out, "      Parent: entityCtxt,")
		fmt.Fprintln(g.out, "    }")
		fmt.Fprintf(g.out, "    for i, value := range entity.%s {\n", field.Name)
		if err := g.writeStructFieldApplyTranslationWithAlias(
			"value",
			"",
			fieldName,
			"",
			field.Type.Kind() == reflect.Ptr,
			false,
			false,
			"      ",
		); err != nil {
			return err
		}
		fmt.Fprintf(g.out, "      entity.%s[i] = value\n", field.Name)
		fmt.Fprintln(g.out, "    }")
		fmt.Fprintln(g.out, "  }")
	default:
		return nil
	}
	return nil
}

func (g *Generator) writeHeader(out io.Writer) {
	fmt.Fprintln(out, "// Code generated by localizer. DO NOT EDIT.")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "package ", g.pkgName)
	fmt.Fprintln(out)
	if len(g.imports) > 0 {
		fmt.Fprintln(out, "import (")
		for im, al := range g.imports {
			if al != "" {
				fmt.Fprintf(out, "  %s %q\n", al, im)
			} else {
				fmt.Fprintf(out, "  %q\n", im)
			}
		}
		fmt.Fprintln(out, ")")
	}
}
