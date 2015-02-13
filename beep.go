/*
** Mack: Beep
** Create a beep notification
*/

package mack

import (
  "strconv"
)

// Beep triggers a given number of system beeps.
//  mack.Beep(1)  // Beeps once
//  mack.Beep(3)  // Beeps 3 times
//
// Parameters:
//
//  times int  // Required - The number of beeps to play
func Beep(times int) error {
  _, err := run(buildBeep(times))
  return err
}

// Parse the beep options and build the command
func buildBeep(times int) string {
  return build("beep", strconv.Itoa(times))
}
