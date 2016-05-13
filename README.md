# Mack
Mack is a Golang wrapper for AppleScript. With Mack, you can easily trigger OS X desktop notifications and system sounds from within your Go application.

## Installation
Mack requires OS X.

`go get github.com/everdev/mack`

## Usage
Mack is ideal for local workflow optimization, OS X binary applications, or just spicing things up. For example:

### Workflow: Process notification
When executing a long-running process, trigger a notification so you can get back to development without having to check the execution status.
```go
package main

import "github.com/everdev/mack"

func main() {
  mack.Say("Starting process")
  // do stuff
  mack.Notify("Complete")
}
```

### Workflow: Open applications
Interact with any Mac application from your code, like opening a URL to a HowTo video.
```go
package main

import (
  "github.com/everdev/mack"
)

func main() {
  browsers := []string{"Some new browser", "Google Chrome", "Firefox", "Safari"}
  opened := false

  for _, browser := range browsers {
    err := mack.Tell(browser, `open location "http://youtube.com/my-howto-video"`)
    if err != nil {
      // handle error
    } else {
      // exit when we found a browser that works
      opened = true
      break
    }
  }

  if !opened {
    // alert user that a common browser could not be found
  }
}
```

### App: ToDo list
Add a cheap GUI to your applications
```go
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

### Workflow: clipboard
Manipulate the clipboard
```go
package main

import (
  "fmt"
  "github.com/everdev/mack"
)

func main() {
  // Output the content of the clipboard
  content, _ := mack.Clipboard()
  fmt.Println(content)

  // Change the content of the clipboard
  mack.SetClipboard("Hello World!")
  content, _ = mack.Clipboard()
  fmt.Println(content)
}
```

## Documentation
Currently, Mack supports the following AppleScript commands:
* Beep
* Clipboard
* Display Alert
* Display Dialog
* Display Notification
* Say
* Tell

Full documentation is available at: [godoc.org/github.com/everdev/mack](http://godoc.org/github.com/everdev/mack)

## Links
* [AppleScript Command Reference Docs](https://developer.apple.com/library/mac/documentation/AppleScript/Conceptual/AppleScriptLangGuide/reference/ASLR_cmds.html)

## Contributors
* Andy Brewer ([everdev](https://github.com/everdev))
* Hiroaki Nakamura ([hnakamur](https://github.com/hnakamur))
* Antoine Augusti ([AntoineAugusti](https://github.com/AntoineAugusti))

## License
MIT
