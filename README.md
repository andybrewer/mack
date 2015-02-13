# Mack
Mack is a Golang wrapper for AppleScript. With Mack, you can easily trigger OS X desktop notifications and system sounds from within your Go application.

## Installation
Mack requires OS X.

`go get github.com/everdev/mack`

## Usage
Mack is ideal for local workflow optimization, OS X binary applications, or just spicing things up. For example:

### Workflow: Process notification
When executing a long-running process, trigger a notification so you can get back to development without having to check the execution status.
```
package main

import "github.com/everdev/mack"

func main() {
  mack.Say("Starting process")
  // do stuff
  mack.Notify("Complete")
}
```

### App: ToDo list
Add some cheap UI to your applications
```
package main

import "github.com/everdev/mack"

func main() {
  response, err := mack.Dialog("Enter a ToDo", "ToDo Wizard", "My new ToDo")
  if err != nil {
    panic(err)
  }

  if response.Clicked == "Cancel" {
    // handle the Cancel event
  } else {
    newToDo := response.Text
    // add ToDo to the database
    mack.Notify("Added " + newToDo + " to your calendar")  
  }
}
```

## Documentation
Currently, Mack supports the following AppleScript commands:
* Beep
* Display Alert
* Display Dialog
* Display Notification
* Say

Full documentation is available at: [godoc.org/github.com/everdev/mack](http://godoc.org/github.com/everdev/mack)

## Links
* [AppleScript Command Reference Docs](https://developer.apple.com/library/mac/documentation/AppleScript/Conceptual/AppleScriptLangGuide/reference/ASLR_cmds.html)

## License
MIT
