/*
** Mack: Say Test
** Test voice notifications
*/

package mack

import (
  "testing"
)

func TestBuildSay(t *testing.T) {
  stringAssertTests := []StringAssert{
    StringAssert{
      actual: buildSay("Hello from Zarvox!", []string{"Zarvox"}),
      expected: "say \"Hello from Zarvox!\" using \"Zarvox\"",
    },
    StringAssert{
      actual: buildSay("Hello from default!", []string{}),
      expected: "say \"Hello from default!\"",
    },
    StringAssert{
      actual: buildSay("Hello with Alex and no Bruce!", []string{"Alex", "Bruce"}),
      expected: "say \"Hello with Alex and no Bruce!\" using \"Alex\"",
    },
  }

  runStringAssertTests("buildSay", stringAssertTests, t)
}
