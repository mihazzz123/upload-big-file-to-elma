package namespace

// ConvertMap - карта преобразований namespace
type ConvertMap map[Namespace]Namespace

// Rebase - получение нового namespace в соответствии с картой преобразования
// вернёт тот же namespace если в карте нет подходящего пути
func (m ConvertMap) Rebase(namespace Namespace) Namespace {
	for key := range m {
		if key.IsSuperNamespaceFor(namespace) {
			return m[key].MakeSubNamespace(key.CutFrom(namespace)...)
		}
	}
	return namespace
}

// ExcludeSystemNamespaces фильтрует из карты системные области видимости, используемые в качестве ключей
func (m ConvertMap) ExcludeSystemNamespaces() ConvertMap {
	filteredMap := ConvertMap{}

	for key, value := range m {
		if !key.IsSystem() {
			filteredMap[key] = value
		}
	}

	return filteredMap
}

// Merge - слияние карт. В новой карте будут ключи и значения из основной карты,
// дополененной ключами и значениями из второй. При совпадении ключей приоритет у основной карты.
// Исходные карты остаются без изменения
func (m ConvertMap) Merge(second ConvertMap) ConvertMap {
	res := make(ConvertMap, len(m))

	for key := range second {
		res[key] = second[key]
	}

	for key := range m {
		res[key] = m[key]
	}

	return res
}

// ConvertMapFromNamespaceLegacy - преобразовывает namespace в ConvertMap нужно для устаревшего механизма
// Deprecated: только для совместимости. Как всезде будет переведено на новый механизм можно удалять
func ConvertMapFromNamespaceLegacy(namespace Namespace) ConvertMap {
	return ConvertMap{
		Current: namespace,
		Parent:  namespace.GetParent(),
	}
}
