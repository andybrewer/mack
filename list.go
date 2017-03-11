/*
** Mack: Alert
** Create a desktop list selection box
 */

package mack

import (
	"strings"
)

// List triggers a desktop list selection box.
// It accepts a title and one or more string items for the user to select from
//
// Use ListWithOpts for more control over the parameters
//
//   selected, didCancel, err := mack.List("Pick Things", "thing one", "thing two")
func List(title string, items ...string) (selected []string, didCancel bool, err error) {
	opts := ListOptions{
		Title: title,
		Items: items,
	}
	return runList(opts)
}

// ListWithOpts trigger a desktop list selection box accepting
// custom parameters.
//
//   list := mack.ListOptions{
//      Items: []string{"item one", 'item two"},
//      Title: "My List Title",
//      Message: "Pick one or more items from this list",
//      DefaultItems: []string{"item one"},
//      AllowMultiple: true,
//    }
//   selected, didCancel, err := ListWithOpts(list)
func ListWithOpts(list ListOptions) (selected []string, didCancel bool, err error) {
	return runList(list)
}

func runList(list ListOptions) (selected []string, didCancel bool, err error) {
	resp, err := run(buildList(list))
	if err != nil {
		return nil, false, err
	}
	if resp == "false" {
		return nil, true, nil
	}
	items := strings.Split(resp, ",")
	for _, item := range items {
		selected = append(selected, strings.TrimSpace(item))
	}
	return selected, false, nil
}

func buildList(list ListOptions) string {
	opts := []string{"choose from list"}

	var items []string
	for _, item := range list.Items {
		items = append(items, wrapInQuotes(item))
	}
	opts = append(opts, "{"+strings.Join(items, ",")+"}")

	if list.Title != "" {
		opts = append(opts, "with title "+wrapInQuotes(list.Title))
	}

	if list.Message != "" {
		opts = append(opts, "with prompt "+wrapInQuotes(list.Message))
	}

	if list.OkButton != "" {
		opts = append(opts, "OK button name "+wrapInQuotes(list.OkButton))
	}

	if list.CancelButton != "" {
		opts = append(opts, "cancel button name "+wrapInQuotes(list.CancelButton))
	}

	if len(list.DefaultItems) > 0 {
		opts = append(opts, "default items "+mkList(list.DefaultItems...))
	}

	if list.AllowMultiple {
		opts = append(opts, "multiple selections allowed true")
	}

	if list.AllowEmpty {
		opts = append(opts, "empty selection allowed true")
	}

	return build(opts...)
}

// ListOptions supplies parameters to the ListWithOpts function.
type ListOptions struct {
	Items         []string // The items to display
	Title         string   // The title of the dialog box
	Message       string   // A message prompt to display in the obx
	OkButton      string   // Text to display on the OK button - defaults to "OK"
	CancelButton  string   // Text to display on the Cancel button - defaults to "Cancel"
	DefaultItems  []string // Optional list of items to select by default
	AllowMultiple bool     // If true, then the user can select multiple items in the list
	AllowEmpty    bool     // If true then the user can select zero items in the list
}
