package props

// List represents properties from a List.
type List struct {
	// MinimumRowsBypage is the minimum amount of rows that must fit on a page.
	// This limit is used to define whether the list will start on the current page or the next one.
	MinimumRowsBypage int
}

// MakeValid from List define default values for a list.
func (l *List) MakeValid() {
	minRows := 1
	if l.MinimumRowsBypage < minRows {
		l.MinimumRowsBypage = minRows
	}
}
