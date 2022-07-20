package errs

// ConstantError может использоваться для создания обрабатываемых исключений
type ConstantError string

// Error возвращает значение как строку
//
// Implements: error
func (ce ConstantError) Error() string {
	return string(ce)
}
