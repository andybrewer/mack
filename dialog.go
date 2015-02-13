/*
** Mack: Alert
** Create a desktop dialog box
*/

package mack

import (
  "strconv"
  "strings"
)

// Dialog triggers a desktop dialog box. Either an error is returned, or the string output from the user interaction.
//  mack.Dialog("Dialog text")                                    // Display a dialog box
//  mack.Dialog("Dialog text", "My Title")                        // Display a dialog box with the title "My Title"
//  mack.Dialog("Dialog text", "My Title", "default text")        // Display a dialog box with "default text" in the input field
//  mack.Dialog("Dialog text", "My Title", "default text", "5")   // Display a dialog box that will disappear after 5 seconds
//  mack.Dialog("Dialog text", "", "", "10")                      // Display a dialog box that will disappear after 10 seconds
//  response, err := mack.Dialog("My dialog")                     // Capture the Response for the dialog box
//
// Parameters:
//
//  text string       // Required - The content of the dialog box
//  title string      // Optional - The title of the dialog box, displayed in emphasis
//  answer string     // Optional - The default text in the input field
//  duration string   // Optional - The number of seconds to wait for a user response
func Dialog(text string, options ...string) (Response, error) {
  return runWithButtons(buildDialog(text, options))
}

// Parse the dialog options and build the command
func buildDialog(text string, options []string) string {
  text = wrapInQuotes(text)

  var title, answer, duration string
  if len(options) > 0 && options[0] != "" {
    title = "with title " + wrapInQuotes(options[0])
  }
  if len(options) > 1 && options[1] != "" {
    answer = "default answer " + wrapInQuotes(options[1])
  }
  if len(options) > 2 && options[2] != "" {
    duration = "giving up after " + options[2]
  }

  return build("display dialog", text, title, answer, duration)
}

// DialogBox triggers a desktop dialog box with the option for custom buttons. Either an error is returned, or the string output from the user interaction.
//  dialog := mack.DialogOptions{
//    Text:           "Dialog text",          // Required
//    Title:          "Dialog title",         // Optional
//    Answer:         "Default answer",       // Optional
//    Duration:       5,                      // Optional
//    HiddenAnswer:   true,                   // Optional - If true, turns the input text to bullets
//    Icon:           "stop",                 // Optional - "stop", "note", "caution" or location of .icns file
//    Buttons:        "Yes, No, Don't Know",  // Optional - Comma separated list, max of 3
//    DefaultButton:  "Don't Know",           // Optional - Ignored if no ButtonList
//  }
//  response, err := mack.DialogBox(dialog)   // Display a dialog with the DialogBox settings, returns an error and Response
func DialogBox(dialog DialogOptions) (Response, error) {
  return runWithButtons(buildDialogBox(dialog))
}

// Parse the DialogBox options and build the command
func buildDialogBox(dialog DialogOptions) string {
  var title, answer, hiddenAnswer, icon, duration, buttons, defaultButton string
  text := wrapInQuotes(dialog.Text)

  if dialog.Title != "" {
    title = "with title " + wrapInQuotes(dialog.Title)
  }
  if dialog.Answer != "" {
    answer = "default answer " + wrapInQuotes(dialog.Answer)
  }
  if dialog.HiddenAnswer {
    hiddenAnswer = "with hidden answer"
  }
  if dialog.Icon != "" {
    if strings.Index(dialog.Icon, ".icns") > 0 {
      // found a filepath to an icon
      icon = "with icon " + wrapInQuotes(dialog.Icon)
    } else {
      // using a system icon
      icon = "with icon " + dialog.Icon
    }
  }
  if dialog.Duration > 0 {
    duration = "giving up after " + strconv.Itoa(dialog.Duration)
  }
  if dialog.Buttons != "" {
    buttons = makeButtonList(dialog.Buttons)

    if dialog.DefaultButton != "" {
      defaultButton = "default button " + wrapInQuotes(dialog.DefaultButton)
    }
  }

  return build("display dialog", text, title, answer, hiddenAnswer, icon, duration, buttons, defaultButton)
}

// DialogOptions are used to generate a DialogBox
type DialogOptions struct {
  Text string           // The content of the dialog box
  Title string          // The title of the dialog box, displayed in emphasis
  Answer string         // The default text in the input field
  HiddenAnswer bool     // If true, converts the answer text to bullets (like a password field)
  Icon string           // The path to a .icns file, or one of the following: "stop", "note", "caution"
  Duration int          // The number of seconds to wait for a user response

  // Buttons
  Buttons string        // The list of up to 3 buttons. Must be commas separated, ex. "Yes, No, Don't Know"
  DefaultButton string  // The default selected button from the button list, ex. "Don't Know"
}
