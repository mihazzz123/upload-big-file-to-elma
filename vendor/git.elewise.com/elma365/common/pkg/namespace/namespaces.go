package namespace

import (
	"errors"
	"sort"
)

// Namespaces - список namespace
type Namespaces []Namespace

const (
	// SuperNamespaceNotFoundError - отсутствие надNamespace
	SuperNamespaceNotFoundError = "super namespace not found"
)

// HasSuperFor - признак того, что в списке есть надNamespace для данного
func (n Namespaces) HasSuperFor(ns Namespace) bool {
	for i := range n {
		if n[i].IsSuperNamespaceFor(ns) {
			return true
		}
	}
	return false
}

// GetHighSuperFor - вернуть самый верхний (с самым коротким путём) namespace из списка,
// который является надNamespace для данного
func (n Namespaces) GetHighSuperFor(ns Namespace) (Namespace, error) {
	superNamespaces := n.GetAllSuperFor(ns)
	if len(superNamespaces) == 0 {
		return "", errors.New(SuperNamespaceNotFoundError)
	}
	superNamespaces = superNamespaces.SortByLevel()
	return superNamespaces[0], nil
}

// SortByLevel - отсортировать по уровню
func (n Namespaces) SortByLevel() Namespaces {
	res := SortByLevel(n)
	sort.Sort(SortByLevel(n))
	return Namespaces(res)
}

// SortByLevel - релизация сортировки по уровню Namespace
type SortByLevel Namespaces

// Len - реализация sort.Interface
func (n SortByLevel) Len() int {
	return len(n)
}

// Less - реализация sort.Interface
func (n SortByLevel) Less(i, j int) bool {
	return n[i].GetLevel() < n[j].GetLevel()
}

// Swap - реализация sort.Interface
func (n SortByLevel) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// GetAllSuperFor - возвращает все namespace из списка которые являются надNamespace для данного
func (n Namespaces) GetAllSuperFor(ns Namespace) Namespaces {
	res := Namespaces{}
	for i := range n {
		if n[i].IsSuperNamespaceFor(ns) {
			res = append(res, n[i])
		}
	}
	return res
}

// Contains проверяет, содержится ли пространство имен в списке.
func (n Namespaces) Contains(namespace Namespace) bool {
	for _, ns := range n {
		if ns == namespace {
			return true
		}
	}
	return false
}

// GetTopLevels возвращает верхний уровень переданных пространств имен (без повторений).
func (n Namespaces) GetTopLevels() Namespaces {
	var topLevels Namespaces

	for _, ns := range n {
		topLevel := ns.GetTopLevel()
		if topLevel == "" || topLevels.Contains(topLevel) {
			continue
		}

		topLevels = append(topLevels, topLevel)
	}

	return topLevels
}
