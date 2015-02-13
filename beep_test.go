/*
** Mack: Beep Test
** Test beep notifications
*/

package mack

import (
  "testing"
)

func TestBuildBeep(t *testing.T) {
  stringAssertTests := []StringAssert{
    StringAssert{
      actual: buildBeep(1),
      expected: "beep 1",
    },
    StringAssert{
      actual: buildBeep(7),
      expected: "beep 7",
    },
  }

  runStringAssertTests("buildBeep", stringAssertTests, t)
}
