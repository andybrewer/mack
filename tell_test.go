/*
** Mack: Tell Test
** Test tell application
*/

package mack

import "testing"

func TestTell(t *testing.T) {
	validCommands := []ErrorAssert{
		ErrorAssert{actual: Tell("TextEdit", "activate")},
		ErrorAssert{actual: Tell("TextEdit", "quit")},
		ErrorAssert{actual: Tell("Finder",
			"activate",
			`open (POSIX file "/Applications")`)},
	}

	runErrorAssertTests("run", validCommands, t)
}
