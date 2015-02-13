/*
** Mack
** A wrapper for AppleScript
*/

// Mack is a Golang wrapper for AppleScript. With Mack, you can easily trigger
// OS X desktop notifications and system sounds from within your Go application.
// Mack is ideal for local workflow optimization or OS X binary applications.
//
// Repository: http://github.com/everdev/mack
package mack

import (
  "regexp"
  "errors"
  "os/exec"
  "strings"
)

// Build the AppleScript command from a set of optional parameters, return the output
func run(command string) (string, error) {
  cmd := exec.Command("osascript", "-e", command)
  output, err := cmd.CombinedOutput()
  prettyOutput := strings.Replace(string(output), "\n", "", -1)

  // Ignore errors from the user hitting the cancel button
  if err != nil && strings.Index(string(output), "User canceled.") < 0 {
    return "", errors.New(err.Error() + ": " + prettyOutput + " (" + command + ")")
  }

  return prettyOutput, nil
}

// Build the AppleScript command from a set of optional parameters, return a Response
func runWithButtons(command string) (Response, error) {
  output, err := run(command)

  // Return if the user hit the default cancel button
  if strings.Index(output, "execution error: User canceled. (-128)") > 0 {
    response := Response{
      Clicked: "Cancel",
    }
    return response, err
  }

  // Parse the buttons
  re := regexp.MustCompile("buttons {(.*)}")
  buttonMatches := re.FindStringSubmatch(command)
  var buttons []string
  if len(buttonMatches) > 1 {
    buttons = strings.Split(buttonMatches[1], ",")
  } else {
    buttons = []string{"OK","Cancel"}
  }

  return parseResponse(output, buttons), err
}

// Wrap text in quotes for proper command line formatting
func wrapInQuotes(text string) string {
  return "\"" + text + "\""
}

// Build the AppleScript command, ignoring any blank optional parameters
func build(params ...string) string {
  var validParams []string

  for _, param := range params {
    if param != "" {
      validParams = append(validParams, param)
    }
  }

  return strings.Join(validParams, " ")
}

// Parse and format the button values
func makeButtonList(buttons string) string {
  buttonList := strings.Split(buttons, ",")

  if len(buttonList) > 3 {
    buttonList = buttonList[:3]
  }

  var wrappedButtons []string
  for _, button := range buttonList {
    wrappedButtons = append(wrappedButtons, wrapInQuotes(strings.TrimSpace(button)))
  }

  return "buttons {" + strings.Join(wrappedButtons, ",") + "}"
}

// Parse a button response
func parseResponse(output string, buttons []string) Response {
  var clicked, text string
  var gaveUp bool

  // Find out if the notification gave up
  gaveUpRe := regexp.MustCompile("gave up:(true|false)")
  gaveUpMatches := gaveUpRe.FindStringSubmatch(output)
  if len(gaveUpMatches) > 1 && gaveUpMatches[1] == "true" {
    gaveUp = true
  }

  if !gaveUp {
    for _, button := range buttons {
      // Find which button was clicked
      buttonStr := "button returned:" + button
      clickedRe := regexp.MustCompile(buttonStr + ",")
      if clickedRe.MatchString(output) || output == buttonStr {
        clicked = button
        break
      }
    }

    // Don't mess around with regex, just get the text returned
    if strings.Index(output, ", text returned:") > 0 {
      output = strings.Replace(output, "button returned:" + clicked + ", ", "", 1)
      output = strings.Replace(output, ", gave up:false", "", 1)
      output = strings.Replace(output, "text returned:", "", 1)
      text = output
    }
  }

  // Find out if the user entered text

  response := Response{
    Clicked: clicked,
    GaveUp: gaveUp,
    Text: text,
  }

  return response
}

// The response format after a button click on an alert or dialog box
type Response struct {
  Clicked string  // The name of the button clicked
  GaveUp bool     // True if the user failed to respond in the duration specified
  Text string     // Only on Dialog boxes - The return value of the input field
}
