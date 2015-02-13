/*
** Mack: Alert
** Create a desktop alert
*/

package mack

import (
  "strconv"
)

// Alert triggers a desktop alert with custom buttons. Either an error is returned, or the string output from the user interaction.
//  mack.Alert("Alert")                               // Display an alert with the emhpasized text "My alert"
//  mack.Alert("Alert", "Message")                    // Display an alert with the small text "My message"
//  mack.Alert("Alert", "Message", "critical")        // Display an alert styled as "critical"
//  mack.Alert("Alert", "Message", "critical", "5")   // Display an alert that will disappear after 5 seconds
//  mack.Alert("Alert", "", "", "10")                 // Display an alert that will disappear after 10 seconds
//  response, err := mack.Alert("Alert")              // Capture the Response for the alert
//
// Parameters:
//
//  title string      // Required - The title of the alert, displayed in emphasis
//  message string    // Optional - The explanatory message, displayed in small text
//  style string      // Optional - The style of the alert: "informational" (default), "warning" or "critical"
//  duration string   // Optional - The number of seconds to wait for a user response, blank or "" will keep it visible until closed
func Alert(title string, options ...string) (Response, error) {
  return runWithButtons(buildAlert(title, options))
}

// Parse the alert options and build the command
func buildAlert(title string, options []string) string {
  title = wrapInQuotes(title)

  var message, style, duration string
  if len(options) > 0 && options[0] != "" {
    message = "message " + wrapInQuotes(options[0])
  }
  if len(options) > 1 && options[1] != "" {
    style = "as " + options[1]
  }
  if len(options) > 2 && options[2] != "" {
    duration = "giving up after " + options[2]
  }

  return build("display alert", title, message, style, duration)
}

// AlertBox triggers a desktop alert with the option for custom buttons. Either an error is returned, or the string output from the user interaction.
//  alert := mack.AlertOptions{
//    Title:          "Alert title",          // Required
//    Message:        "Alert message",        // Optional
//    Style:          "critical",             // Optional
//    Duration:       5,                      // Optional
//    Buttons:        "Yes, No, Don't Know",  // Optional - Comma separated list, max of 3
//    DefaultButton:  "Don't Know",           // Optional - Ignored if no ButtonList
//  }
//  response, err := mack.AlertBox(alert)     // Display an alert with the AlertBox settings, returns an error and Response
func AlertBox(alert AlertOptions) (Response, error) {
  return runWithButtons(buildAlertBox(alert))
}

// Parse the AlertBox options and build the command
func buildAlertBox(alert AlertOptions) string {
  var message, style, duration, buttons, defaultButton string
  title := wrapInQuotes(alert.Title)

  if alert.Message != "" {
    message = "message " + wrapInQuotes(alert.Message)
  }
  if alert.Style != "" {
    style = "as " + alert.Style
  }
  if alert.Duration > 0 {
    duration = "giving up after " + strconv.Itoa(alert.Duration)
  }
  if alert.Buttons != "" {
    buttons = makeButtonList(alert.Buttons)

    if alert.DefaultButton != "" {
      defaultButton = "default button " + wrapInQuotes(alert.DefaultButton)
    }
  }

  return build("display alert", title, message, style, duration, buttons, defaultButton)
}

// AlertOptions are used to generate an AlertBox
type AlertOptions struct {
  Title string          // The title of the alert, displayed in emphasis
  Message string        // The explanatory message, displayed in small text
  Style string          // The style of the alert: "informational" (default), "warning" or "critical"
  Duration int          // The number of seconds to wait for a user response, blank or "" will keep it visible until closed

  // Buttons
  Buttons string        // The list of up to 3 buttons. Must be commas separated, ex. "Yes, No, Don't Know"
  DefaultButton string  // The default selected button from the button list, ex. "Don't Know"
}
