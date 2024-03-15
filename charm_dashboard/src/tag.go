package main

import "fmt"

type section int

func (s section) getNext() section {
	if s == resource {
		return tagKey
	}
	return s + 1
}

func (s section) getPrev() section {
	if s == tagKey {
		return resource
	}
	return s - 1
}

const (
	tagKey section = iota
	tagValue
	resource
)

/* Custom SectionItem */
type SectionItem struct {
	section section
	name    string
	values  []string
}

func (t SectionItem) FilterValue() string {
	return t.name
}

func (t SectionItem) Title() string {
	if len(t.name) == 0 {
		return "Empty Value"
	}
	return truncateString(t.name, 39)
}

func (t SectionItem) Description() string {
	switch t.section {
	case resource:
		return "resource"
	default:
		return fmt.Sprintf("Items: %d", len(t.values))
	}
}

func (t SectionItem) Key() string {
	if len(t.name) == 0 {
		return "Empty Value"
	}
	return t.name
}

func (t SectionItem) Values() []string {
	return t.values
}
