/*
** Say
** Create a voice notification
 */

package mack

// Say triggers a voice notification that will read the text provided in a given voice.
//
//	mack.Say("Hi in Fred's voice!", "Fred")   // Have the "Fred" voice read the text
//	mack.Say("Hi in default voice!")          // Have the system default voice read the text
//
// Parameters:
//
//	text string   // Required - What the system voice will say
//	voice string  // Optional - The name of the system voice, otherwise defaults to system preferences
//	              // Voice list located at: /System/Library/Speech/Voices/ (macOS 10.15+)
//	              //                   or: /System/Library/Speech/Voices/ (older macOS)
//	              // Common voices: Alex, Fred, Samantha, Victoria, Daniel, Karen, Moira, Rishi, Flo, Grandma, Grandpa
//	              // Note: Many voices are now downloaded on-demand in modern macOS versions.
//	              // Use "say -v '?'" in Terminal to see all available voices on your system.
func Say(text string, options ...string) error {
	_, err := run(buildSay(text, options))
	return err
}

// Parse the say options and build the command
func buildSay(text string, options []string) string {
	var voice string
	if len(options) > 0 && options[0] != "" {
		voice = "using " + wrapInQuotes(options[0])
	}

	text = wrapInQuotes(text)
	return build("say", text, voice)
}
