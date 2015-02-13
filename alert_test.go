/*
** Mack: Alert Test
** Test desktop alerts
*/

package mack

import (
  "testing"
)

func TestBuildAlert(t *testing.T) {
  stringAssertTests := []StringAssert{
    StringAssert{
      actual: buildAlert("My Alert", []string{}),
      expected: "display alert \"My Alert\"",
    },
    StringAssert{
      actual: buildAlert("My Alert", []string{"My message"}),
      expected: "display alert \"My Alert\" message \"My message\"",
    },
    StringAssert{
      actual: buildAlert("My Alert", []string{"My message", "critical"}),
      expected: "display alert \"My Alert\" message \"My message\" as critical",
    },
    StringAssert{
      actual: buildAlert("My Alert", []string{"My message", "critical", "5"}),
      expected: "display alert \"My Alert\" message \"My message\" as critical giving up after 5",
    },
    StringAssert{
      actual: buildAlert("My Alert", []string{"", "", "5"}),
      expected: "display alert \"My Alert\" giving up after 5",
    },
  }

  runStringAssertTests("buildAlert", stringAssertTests, t)
}

func TestBuildAlertBox(t *testing.T) {
  stringAssertTests := []StringAssert{
    StringAssert{
      actual: buildAlertBox(AlertOptions{
        Title: "My Alert",
      }),
      expected: "display alert \"My Alert\"",
    },
    StringAssert{
      actual: buildAlertBox(AlertOptions{
        Title: "My Alert",
        Message: "My message",
      }),
      expected: "display alert \"My Alert\" message \"My message\"",
    },
    StringAssert{
      actual: buildAlertBox(AlertOptions{
        Title: "My Alert",
        Message: "My message",
        Style: "warning",
      }),
      expected: "display alert \"My Alert\" message \"My message\" as warning",
    },
    StringAssert{
      actual: buildAlertBox(AlertOptions{
        Title: "My Alert",
        Message: "My message",
        Style: "warning",
        Duration: 5,
      }),
      expected: "display alert \"My Alert\" message \"My message\" as warning giving up after 5",
    },
    StringAssert{
      actual: buildAlertBox(AlertOptions{
        Title: "My Alert",
        Message: "My message",
        Style: "warning",
        Duration: 5,
        Buttons: "Yes, No, Don't Know",
      }),
      expected: "display alert \"My Alert\" message \"My message\" as warning giving up after 5 buttons {\"Yes\",\"No\",\"Don't Know\"}",
    },
    StringAssert{
      actual: buildAlertBox(AlertOptions{
        Title: "My Alert",
        Message: "My message",
        Style: "warning",
        Duration: 5,
        Buttons: "Yes, No, Don't Know",
        DefaultButton: "Don't Know",
      }),
      expected: "display alert \"My Alert\" message \"My message\" as warning giving up after 5 buttons {\"Yes\",\"No\",\"Don't Know\"} " +
                "default button \"Don't Know\"",
    },
    StringAssert{
      actual: buildAlertBox(AlertOptions{
        Title: "My Alert",
        Message: "My message",
        Style: "warning",
        Duration: 5,
        Buttons: "Yes, No, Don't Know, One Too Many",
        DefaultButton: "Don't Know",
      }),
      expected: "display alert \"My Alert\" message \"My message\" as warning giving up after 5 buttons {\"Yes\",\"No\",\"Don't Know\"} " +
                "default button \"Don't Know\"",
    },
  }

  runStringAssertTests("buildAlertBox", stringAssertTests, t)
}
