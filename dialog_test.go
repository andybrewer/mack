/*
** Mack: Alert Test
** Test desktop dialog boxes
*/

package mack

import (
  "testing"
)

func TestBuildDialog(t *testing.T) {
  stringAssertTests := []StringAssert{
    StringAssert{
      actual: buildDialog("text", []string{}),
      expected: "display dialog \"text\"",
    },
    StringAssert{
      actual: buildDialog("text", []string{"title"}),
      expected: "display dialog \"text\" with title \"title\"",
    },
    StringAssert{
      actual: buildDialog("text", []string{"title", "answer"}),
      expected: "display dialog \"text\" with title \"title\" default answer \"answer\"",
    },
    StringAssert{
      actual: buildDialog("text", []string{"title", "answer", "5"}),
      expected: "display dialog \"text\" with title \"title\" default answer \"answer\" giving up after 5",
    },
    StringAssert{
      actual: buildDialog("text", []string{"", "", "5"}),
      expected: "display dialog \"text\" giving up after 5",
    },
  }

  runStringAssertTests("buildDialog", stringAssertTests, t)
}

func TestBuildDialogBox(t *testing.T) {
  stringAssertTests := []StringAssert{
    StringAssert{
      actual: buildDialogBox(DialogOptions{
        Text: "text",
      }),
      expected: "display dialog \"text\"",
    },
    StringAssert{
      actual: buildDialogBox(DialogOptions{
        Text: "text",
        Title: "title",
      }),
      expected: "display dialog \"text\" with title \"title\"",
    },
    StringAssert{
      actual: buildDialogBox(DialogOptions{
        Text: "text",
        Title: "title",
        Answer: "answer",
      }),
      expected: "display dialog \"text\" with title \"title\" default answer \"answer\"",
    },
    StringAssert{
      actual: buildDialogBox(DialogOptions{
        Text: "text",
        Title: "title",
        Answer: "answer",
        HiddenAnswer: true,
      }),
      expected: "display dialog \"text\" with title \"title\" default answer \"answer\" with hidden answer",
    },
    StringAssert{
      actual: buildDialogBox(DialogOptions{
        Text: "text",
        Title: "title",
        Answer: "answer",
        HiddenAnswer: false,
        Icon: "my-icon.icns",
      }),
      expected: "display dialog \"text\" with title \"title\" default answer \"answer\" with icon \"my-icon.icns\"",
    },
    StringAssert{
      actual: buildDialogBox(DialogOptions{
        Text: "text",
        Title: "title",
        Answer: "answer",
        HiddenAnswer: false,
        Icon: "0",
      }),
      expected: "display dialog \"text\" with title \"title\" default answer \"answer\" with icon 0",
    },
    StringAssert{
      actual: buildDialogBox(DialogOptions{
        Text: "text",
        Title: "title",
        Answer: "answer",
        HiddenAnswer: false,
        Icon: "0",
        Duration: 5,
      }),
      expected: "display dialog \"text\" with title \"title\" default answer \"answer\" with icon 0 giving up after 5",
    },
    StringAssert{
      actual: buildDialogBox(DialogOptions{
        Text: "text",
        Title: "title",
        Answer: "answer",
        HiddenAnswer: false,
        Icon: "0",
        Duration: 5,
        Buttons: "Yes, No, Don't Know",
      }),
      expected: "display dialog \"text\" with title \"title\" default answer \"answer\" with icon 0 giving up after 5 " +
                "buttons {\"Yes\",\"No\",\"Don't Know\"}",
    },
    StringAssert{
      actual: buildDialogBox(DialogOptions{
        Text: "text",
        Title: "title",
        Answer: "answer",
        HiddenAnswer: false,
        Icon: "0",
        Duration: 5,
        Buttons: "Yes, No, Don't Know",
        DefaultButton: "Don't Know",
      }),
      expected: "display dialog \"text\" with title \"title\" default answer \"answer\" with icon 0 giving up after 5 " +
                "buttons {\"Yes\",\"No\",\"Don't Know\"} default button \"Don't Know\"",
    },
    StringAssert{
      actual: buildDialogBox(DialogOptions{
        Text: "text",
        Title: "title",
        Answer: "answer",
        HiddenAnswer: false,
        Icon: "0",
        Duration: 5,
        Buttons: "Yes, No, Don't Know, One Too Many",
        DefaultButton: "Don't Know",
      }),
      expected: "display dialog \"text\" with title \"title\" default answer \"answer\" with icon 0 giving up after 5 " +
                "buttons {\"Yes\",\"No\",\"Don't Know\"} default button \"Don't Know\"",
    },
  }

  runStringAssertTests("buildDialogBox", stringAssertTests, t)
}
