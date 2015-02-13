/*
** Mack: Notification
** Create a desktop notification
*/

package mack

// Notify triggers a desktop notification.
//  mack.Notify("My message")                                     // Display a notification with the content "My message"
//  mack.Notify("My message", "My title")                         // Display a notification with the title "My title"
//  mack.Notify("My message", "My title", "My subtitle")          // Display a notification with the subtitle "My subtitle"
//  mack.Notify("My message", "My title", "My subtitle", "Ping")  // Display a notification with a Ping sound
//  mack.Notify("My message", "", "", "Ping")                     // Display a notification with a Ping sound and no title or subtitle
//
// Parameters:
//
//  text string      // Required - The content of the notification
//  title string     // Optional - The title of the notification
//  subtitle string  // Optional - The subtitle of the notification
//  sound string     // Optional - The sound to play when showing the notification
//                   // Sounds list located at: /System/Library/Sounds/
//                   // ex. Basso, Blow, Bottle, Frog, Funk, Glass, Hero, Morse, Ping, Pop, Purr, Sosumi, Submarine, Tink
func Notify(text string, options ...string) error {
  _, err := run(buildNotification(text, options))
  return err
}

// Parse the notify options and build the command
func buildNotification(text string, options []string) string {
  var title, subtitle, sound string
  if len(options) > 0 && options[0] != "" {
    title = "with title " + wrapInQuotes(options[0])
  }
  if len(options) > 1 && options[1] != "" {
    subtitle = "subtitle " + wrapInQuotes(options[1])
  }
  if len(options) > 2 && options[2] != "" {
    sound = "sound name " + wrapInQuotes(options[2])
  }

  text = wrapInQuotes(text)
  return build("display notification", text, title, subtitle, sound)
}
