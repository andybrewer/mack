/*
** Mack: Clipboard
** Interact with the clipboard
 */

package mack

import (
	"testing"
)

func TestBuildClipboard(t *testing.T) {
	stringAssertTests := []StringAssert{
		StringAssert{
			actual:   buildClipboard(),
			expected: "the clipboard",
		},
	}

	runStringAssertTests("buildClipboard", stringAssertTests, t)
}

func TestBuildSetClipboard(t *testing.T) {
	stringAssertTests := []StringAssert{
		StringAssert{
			actual:   buildSetClipboard("text"),
			expected: "set the clipboard to \"text\"",
		},
	}

	runStringAssertTests("buildSetClipboard", stringAssertTests, t)
}

func TestSetClipboard(t *testing.T) {
	content := "testing clipboard"
	SetClipboard(content)
	result, _ := Clipboard()

	stringAssertTests := []StringAssert{
		StringAssert{
			actual:   result,
			expected: content,
		},
	}

	runStringAssertTests("buildSetClipboard", stringAssertTests, t)
}
