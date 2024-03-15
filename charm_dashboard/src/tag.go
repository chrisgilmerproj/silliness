package main

import "fmt"

/* Custom Tag */
type Tag struct {
	section section
	name    string
	values  []string
}

func (t Tag) FilterValue() string {
	return t.name
}

func (t Tag) Title() string {
	if len(t.name) == 0 {
		return "Empty Value"
	}
	return truncateString(t.name, 39)
}

func (t Tag) Description() string {
	switch t.section {
	case resource:
		return "resource"
	default:
		return fmt.Sprintf("Items: %d", len(t.values))
	}
}

func (t Tag) Key() string {
	if len(t.name) == 0 {
		return "Empty Value"
	}
	return t.name
}

func (t Tag) Values() []string {
	return t.values
}
