/*
** Mack: Tell Test
** Test tell application
 */

package mack

import "testing"

func TestTell(t *testing.T) {
	tell := func(application string, commands ...string) error {
		_, err := Tell(application, commands...)
		return err
	}

	validCommands := []ErrorAssert{
		ErrorAssert{actual: tell("TextEdit", "activate")},
		ErrorAssert{actual: tell("TextEdit", "activate")},
		ErrorAssert{actual: tell("TextEdit", "quit")},
		ErrorAssert{actual: tell("Finder",
			"activate",
			`open (POSIX file "/Applications")`)},
	}

	runErrorAssertTests("run", validCommands, t)
}
