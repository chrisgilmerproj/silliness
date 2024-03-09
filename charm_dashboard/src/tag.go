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
	return t.name
}

func (t Tag) Description() string {
	switch t.section {
	case instance:
		return "instance"
	default:
		return fmt.Sprintf("Items: %d", len(t.values))
	}
}

func (t Tag) Key() string {
	return t.name
}

func (t Tag) Values() []string {
	return t.values
}
