package link

// Link string
type Link string

// String implements fmt.Stringer interface
func (l Link) String() string {
	return string(l)
}
