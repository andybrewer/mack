/*
** Say
** Create a voice notification
*/

package mack

// Say triggers a voice notification that will read the text provided in a given voice.
//  mack.Say("Hi in Bruce's voice!", "Bruce")   // Have the "Bruce" voice read the text
//  mack.Say("Hi in default voice!")            // Have the system default voice read the text
//
// Parameters:
//
//  text string   // Required - What the system voice will say
//  voice string  // Optional - The name of the system voice, otherwise defaults to system preferences
//                // Voice list located at: /System/Library/Speech/Voices
//                // ex. Agnes, Albert, Alex, Alice Compact, Alva Compact, Amelie Compact, Anna Compact, BadNews, Bahh, Bells, Boing, Bruce,
//                //     Bubbles, Carmit Compact, Cellos, Damayanti Compact, Daniel Compact, Deranged, Diego Compact, Ellen Compact,
//                //     Fiona Compact, Fred, GoodNews, Hysterical, Ioana Compact, Joana Compact, Junior, Kanya Compact, Karen Compact, Kathy,
//                //     Kyoko Compact, Laura Compact, Lekha Compact, Luciana Compact, Mariska Compact, Mei-Jia Compact, Melina Compact,
//                //     Milena Compact, Moira Compact, Monica Compact, Nora Compact, Organ, Paulina Compact, Princess, Ralph, Samantha Compact,
//                //     Sara Compact, Satu Compact, Sin-ji Compact, Tarik Compact, Tessa Compact, Thomas Compact, Ting-Ting Compact, Trinoids,
//                //     Veena Compact, Vicki, Victoria, Whisper, Xander Compact, Yelda Compact, Yuna Compact, Zarvox, Zosia Compact, Zuzana Compact
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
