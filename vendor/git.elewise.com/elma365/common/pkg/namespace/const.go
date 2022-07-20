package namespace

const (
	// Separator - разделитель пространств имен
	Separator = "."
	// System - системное пространство имен
	System Namespace = "system"
	// Global - глобальное пространство имен
	Global Namespace = "global"
	// Virtual - виртуальное пространство имен
	Virtual Namespace = "virtual"
	// Current - ссылочное пространство имен. При вычислении абсолютного заменяется на текущее пространство имен
	// Deprecated: отказались от относительных namespace
	Current Namespace = "$current"
	// Parent - ссылочное пространство имен. При вычислении абсолютного заменяется на родительское пространство имен
	// Deprecated: отказались от относительных namespace
	Parent Namespace = "$parent"
)
