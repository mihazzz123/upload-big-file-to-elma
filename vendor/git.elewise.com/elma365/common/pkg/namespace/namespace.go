package namespace

import (
	"fmt"
	"regexp"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// Namespace - тип для пространства имен
type Namespace string

// SystemCodePrefix - префикс для системных кодов, которые иначе не получается пометить как системный
const SystemCodePrefix = "_"

var extensionNamespaceRegex = regexp.MustCompile("^ext_[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")

// FromParts - получение пространства имён из массива составляющих
func FromParts(parts []string) Namespace {
	if len(parts) == 0 {
		return Global
	}
	return Namespace(strings.Join(parts, Separator))
}

// FromPartsArg - получение пространства имён из аргументов
func FromPartsArg(parts ...string) Namespace {
	return FromParts(parts)
}

// IsSubNamespaceOf - проверяет является ли текущее пространство подпространством контекстного
func (ns Namespace) IsSubNamespaceOf(context Namespace) bool {
	if ns == context {
		return true
	}
	if ns.GetRoot() != context.GetRoot() {
		return false
	}
	if context == Global {
		return true
	}
	return strings.HasPrefix(ns.String(), context.String()+Separator)
}

// GetRoot - возвращает корневое пространство имён (global или system)
func (ns Namespace) GetRoot() Namespace {
	sections := ns.GetSections()
	if Namespace(sections[0]) == System {
		return System
	}
	return Global
}

// IsSuperNamespaceFor - проверяет является ли текущее пространство имён надпространством для ns2
func (ns Namespace) IsSuperNamespaceFor(ns2 Namespace) bool {
	return ns2.IsSubNamespaceOf(ns)
}

// IsSuperNamespaceForAny проверяет, является ли "ns" надпространством над хотя бы одним пространством из "namespaces".
func (ns Namespace) IsSuperNamespaceForAny(namespaces Namespaces) bool {
	for _, checkNamespace := range namespaces {
		if ns.IsSuperNamespaceFor(checkNamespace) {
			return true
		}
	}
	return false
}

// GetLevel - уроваень вложенности пространства имён
func (ns Namespace) GetLevel() int {
	if ns == Global {
		return 0
	}
	return len(ns.GetSections())
}

// GetSections - составляющие пространства имён в виде массива строк
func (ns Namespace) GetSections() []string {
	return strings.Split(ns.String(), Separator)
}

// AsRelative - построение относительно пути в контексте пространства имён context
// Deprecated: отказываемся от относительных путей
func (ns Namespace) AsRelative(context Namespace) Namespace {
	if context == Global {
		return ns
	} else if ns == context {
		return Current
	} else if context.IsSuperNamespaceFor(ns) {
		return Namespace(strings.Replace(ns.String(), context.String(), Current.String(), 1))
	} else if parent := context.GetParent(); parent != Global && parent.IsSuperNamespaceFor(ns) {
		return Namespace(strings.Replace(ns.String(), parent.String(), Parent.String(), 1))
	} else {
		return ns
	}
}

// GetParent - возвращает родительское пространство имён
func (ns Namespace) GetParent() Namespace {
	switch ns.GetLevel() {
	case 0, 1:
		return Global
	default:
		sections := ns.GetSections()
		parent := sections[:len(sections)-1]
		return FromParts(parent)
	}
}

// AsAbsolute - получение абсолютного пространства имен по относительному
// Deprecated: отказываемся от относительных путей
func (ns Namespace) AsAbsolute(context Namespace) Namespace {
	if strings.HasPrefix(string(ns), string(Current)) {
		return Namespace(strings.Replace(string(ns), string(Current), string(context), 1))
	}
	if strings.HasPrefix(string(ns), string(Parent)) {
		path := strings.Split(string(context), Separator)
		return Namespace(strings.Replace(string(ns), string(Parent), path[0], 1))
	}
	return ns
}

// MakeCollectionID формирует id коллекции для данного раздела и кода
func (ns Namespace) MakeCollectionID(code string) uuid.UUID {
	return uuid.NewV5(uuid.NamespaceOID, fmt.Sprintf("%s:%s", ns, code))
}

// String имплементация fmt.Stringer - строковое представление ns
func (ns Namespace) String() string {
	return string(ns)
}

// IsSystemCode - является ли приложение с неймспейсом ns и кодом code системным.
// текущий признак системного приложения - код начинается с символа "_"
func (ns Namespace) IsSystemCode(code string) bool {
	return strings.HasPrefix(code, SystemCodePrefix)
}

// GetDeepestSection - возвращает наиболее глубокую часть. Например достать код приложения из полного неймспейса
func (ns Namespace) GetDeepestSection() string {
	sections := ns.GetSections()
	if len(sections) == 0 {
		return Global.String()
	}
	return sections[len(sections)-1]
}

// IsSystem - признак системности
func (ns Namespace) IsSystem() bool {
	if System.IsSuperNamespaceFor(ns) {
		return true
	}
	sections := ns.GetSections()
	for i := range sections {
		if !strings.HasPrefix(sections[i], SystemCodePrefix) {
			return false
		}
	}
	return true
}

// CutFrom - вырезает ns из начала ns2 и возвращает оставшиеся части списком
func (ns Namespace) CutFrom(ns2 Namespace) []string {
	if !ns.IsSuperNamespaceFor(ns2) {
		return ns2.GetSections()
	}
	return ns2.GetSections()[ns.GetLevel():]
}

// MakeSubNamespace - создать подNamespace добавив части из аргумента
func (ns Namespace) MakeSubNamespace(subSections ...string) Namespace {
	if len(subSections) == 0 {
		return ns
	}
	if ns == Global {
		return FromParts(subSections)
	}
	subNamespaceSections := append(ns.GetSections(), subSections...)
	return FromParts(subNamespaceSections)
}

// GetTopLevel - получить namespace верхнего уровня (код раздела)
func (ns Namespace) GetTopLevel() Namespace {
	sections := ns.GetSections()
	return Namespace(sections[0])
}

// IsExtension проверяет, является ли пространством имён расширения
func (ns Namespace) IsExtension() bool {
	return extensionNamespaceRegex.MatchString(ns.String())
}

// GetDirID возвращает код дирректории пространства имен
func (ns Namespace) GetDirID() uuid.UUID {
	if ns == Global || ns == System {
		return uuid.NewV5(uuid.NamespaceOID, string(System))
	}

	return ns.GetParent().MakeCollectionID(ns.GetTopLevel().String())
}
