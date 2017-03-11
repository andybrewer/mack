package mack

import "testing"

func TestBuildList(t *testing.T) {
	input := ListOptions{
		Items:         []string{"one", "two", "three"},
		Title:         "box title",
		Message:       "box message",
		OkButton:      "ok button",
		CancelButton:  "cancel button",
		DefaultItems:  []string{"one", "two"},
		AllowMultiple: true,
		AllowEmpty:    true,
	}

	expected := `choose from list {"one","two","three"} ` +
		`with title "box title" ` +
		`with prompt "box message" ` +
		`OK button name "ok button" ` +
		`cancel button name "cancel button" ` +
		`default items {"one","two"} ` +
		`multiple selections allowed true ` +
		`empty selection allowed true`

	result := buildList(input)
	if result != expected {
		t.Errorf("expected=\n%q actual=\n%q", expected, result)
	}
}
