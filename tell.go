/*
** Tell
** Tells an application commands.
*/

package mack

// Tell tells an application the specified commands
//  mack.Tell("TextEdit", "activate")  // Activates TextEdit
//  mack.Tell("TextEdit", "quit")      // Quits TextEdit
//  mack.Tell("Finder",
//    "activate",
//    `open (POSIX file "/Applications")`) // Activate Finder and open the "/Applications" folder
//
// Parameters:
//
//  application string   // Required - What application the system will tell to
//  commands string      // Required - What command lines the system will tell
func Tell(application string, commands ...string) error {
	_, err := run(buildTell(application, commands...))
	return err
}

// Parse the Tell options and build the command
func buildTell(application string, commands ...string) string {
	application = wrapInQuotes(application)
	args := []string{"tell application", application, "\n"}
	for _, command := range commands {
		args = append(args, command, "\n")
	}
	args = append(args, "end", "tell")
	return build(args...)
}
