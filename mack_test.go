/*
** Mack Test
** Test the wrapper for Mac's Notification Center
*/

package mack

import (
  "errors"
  "strconv"
  "testing"
)

// Test executing the commands
func TestRun(t *testing.T) {
  // Invalid commands
  badAlert := AlertOptions{
    Buttons: "Yes, No, Don't Know",
    DefaultButton: "Non-existant-button",
  }
  badAlert2 := AlertOptions{
    Style: "none",
  }
  badDialog := DialogOptions{
    Icon: "invalid",
  }

  _, err1 := AlertBox(badAlert)
  _, err2 := AlertBox(badAlert2)
  _, err3 := DialogBox(badDialog)

  commands := []ErrorAssert{
    ErrorAssert{
      actual: Say("hi!", "non-existant-voice"),
      expected: errors.New("exit status 1: 0:36: execution error: Voice wasn’t found. (-244) " +
                           "(say \"hi!\" using \"non-existant-voice\")"),
    },
    ErrorAssert{
      actual: err1,
      expected: errors.New("exit status 1: 0:87: execution error: Specified button does not exist. (-50) " +
                           "(display alert \"\" buttons {\"Yes\",\"No\",\"Don't Know\"} default button \"Non-existant-button\")"),
    },
    ErrorAssert{
      actual: err2,
      expected: errors.New("exit status 1: 20:24: execution error: The variable none is not defined. (-2753) " +
                           "(display alert \"\" as none)"),
    },
    ErrorAssert{
      actual: err3,
      expected: errors.New("exit status 1: 28:35: execution error: The variable invalid is not defined. (-2753) " +
                           "(display dialog \"\" with icon invalid)"),
    },
  }

  // Valid commands, comment out for active development
  validCommands := []ErrorAssert{
    ErrorAssert{actual: Say("hi!")},
    ErrorAssert{actual: Say("hi!", "Agnes")},
    ErrorAssert{actual: Beep(1)},
    ErrorAssert{actual: Beep(2)},
  }
  commands = append(commands, validCommands...)

  runErrorAssertTests("run", commands, t)
}

// Test wrapping a string in quotes
func TestWrapInQuotes(t *testing.T) {
  stringAssertTests := []StringAssert{
    StringAssert{
      actual: wrapInQuotes("Hello world!"),
      expected: "\"Hello world!\"",
    },
  }

  runStringAssertTests("wrapInQuotes", stringAssertTests, t)
}

// Test building a command
func TestBuild(t *testing.T) {
  stringAssertTests := []StringAssert{
    StringAssert{
      actual: build("say","\"Hello world!\"", "using \"Agnes\""),
      expected: "say \"Hello world!\" using \"Agnes\"",
    },
  }

  runStringAssertTests("build", stringAssertTests, t)
}

// Test parsing a button list
func TestMakeButtonList(t *testing.T) {
  stringAssertTests := []StringAssert{
    StringAssert{
      actual: makeButtonList("One, Two, Three"),
      expected: "buttons {\"One\",\"Two\",\"Three\"}",
    },
    StringAssert{
      actual: makeButtonList("One"),
      expected: "buttons {\"One\"}",
    },
    StringAssert{
      actual: makeButtonList("One, Two, Three, Four"),
      expected: "buttons {\"One\",\"Two\",\"Three\"}",
    },
    StringAssert{
      actual: makeButtonList("One,Two"),
      expected: "buttons {\"One\",\"Two\"}",
    },
  }

  runStringAssertTests("makeButtonList", stringAssertTests, t)
}

// Test parsing an alert or dialog response
func TestParseResponse(t *testing.T) {
  response1 := "button returned:OK"
  response2 := "button returned:My button, gave up:false"
  response3 := "button returned:My button 2, gave up:false"
  response4 := "button returned:, gave up:true"
  response5 := "button returned:OK, text returned:my text, gave up:false"
  response6 := "button returned:button, text returned:this is the text returned: blah, gave up:false"

  stringAssertTests := []StringAssert{
    StringAssert{
      actual: parseResponse(response1, []string{"OK","Cancel"}).Clicked,
      expected: "OK",
    },
    StringAssert{
      actual: parseResponse(response1, []string{"OK","Cancel"}).Clicked,
      expected: "OK",
    },
    StringAssert{
      actual: strconv.FormatBool(parseResponse(response1, []string{"OK","Cancel"}).GaveUp),
      expected: "false",
    },
    StringAssert{
      actual: parseResponse(response2, []string{"My button","My button2"}).Clicked,
      expected: "My button",
    },
    StringAssert{
      actual: strconv.FormatBool(parseResponse(response2, []string{"My button","My button 2"}).GaveUp),
      expected: "false",
    },
    StringAssert{
      actual: parseResponse(response3, []string{"My button","My button 2"}).Clicked,
      expected: "My button 2",
    },
    StringAssert{
      actual: strconv.FormatBool(parseResponse(response3, []string{"My button","My button 2"}).GaveUp),
      expected: "false",
    },
    StringAssert{
      actual: parseResponse(response4, []string{"My button","My button 2"}).Clicked,
      expected: "",
    },
    StringAssert{
      actual: strconv.FormatBool(parseResponse(response4, []string{"My button","My button 2"}).GaveUp),
      expected: "true",
    },
    StringAssert{
      actual: parseResponse(response5, []string{"OK","Cancel"}).Clicked,
      expected: "OK",
    },
    StringAssert{
      actual: strconv.FormatBool(parseResponse(response5, []string{"OK","Cancel"}).GaveUp),
      expected: "false",
    },
    StringAssert{
      actual: parseResponse(response5, []string{"OK","Cancel"}).Text,
      expected: "my text",
    },
    StringAssert{
      actual: parseResponse(response6, []string{"button","button 2"}).Clicked,
      expected: "button",
    },
    StringAssert{
      actual: strconv.FormatBool(parseResponse(response6, []string{"button","button 2"}).GaveUp),
      expected: "false",
    },
    StringAssert{
      actual: parseResponse(response6, []string{"button","button 2"}).Text,
      expected: "this is the text returned: blah",
    },
  }

  runStringAssertTests("parseResponse", stringAssertTests, t)
}

/*
** Test Helpers
*/

// Compare two string values
func runStringAssertTests(process string, tests []StringAssert, t *testing.T) {
  for _, test := range tests {
    if test.actual != test.expected {
      fail(process, test.expected, test.actual, t)
    }
  }
}

// Compare two error values
func runErrorAssertTests(process string, tests []ErrorAssert, t *testing.T) {
  for _, test := range tests {
    // Give every result an error so we can compare the error strings
    if test.actual == nil {
      test.actual = errors.New("nil")
    }
    if test.expected == nil {
      test.expected = errors.New("nil")
    }
    if test.actual.Error() != test.expected.Error() {
      fail("run", test.expected.Error(), test.actual.Error(), t)
    }
  }
}

// Global fail message format
func fail(process string, expected string, got string, t *testing.T) {
  t.Error("\n!! " + process + " failed !!\nExpected: " + expected + "\nReceived: " + got + "\n")
}

type StringAssert struct {
  actual string
  expected string
}

type ErrorAssert struct {
  actual error
  expected error
}
