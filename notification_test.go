/*
** Mack: Notification Test
** Test desktop notifications
*/

package mack

import (
  "testing"
)

func TestBuildNotification(t *testing.T) {
  stringAssertTests := []StringAssert{
    StringAssert{
      actual: buildNotification("My Content", []string{"My title", "My subtitle", "Basso"}),
      expected: "display notification \"My Content\" with title \"My title\" subtitle \"My subtitle\" sound name \"Basso\"",
    },
    StringAssert{
      actual: buildNotification("My Content", []string{}),
      expected: "display notification \"My Content\"",
    },
    StringAssert{
      actual: buildNotification("My Content", []string{"", "", "Ping"}),
      expected: "display notification \"My Content\" sound name \"Ping\"",
    },
    StringAssert{
      actual: buildNotification("My Content", []string{"My title", "", "Submarine"}),
      expected: "display notification \"My Content\" with title \"My title\" sound name \"Submarine\"",
    },
  }

  runStringAssertTests("buildNotification", stringAssertTests, t)
}
