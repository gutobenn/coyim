package gui

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"os"

	. "gopkg.in/check.v1"
)

type UIReaderSuite struct{}

var _ = Suite(&UIReaderSuite{})

const testFile string = `
<interface>
  <object class="GtkWindow" id="conversation">
    <property name="default-height">500</property>
    <property name="default-width">400</property>

      <child>
	<object class="GtkVBox">
	</object>
      </child>

  </object>
</interface>
`

func writeTestFile(name, content string) {
	desc, _ := os.Create(name)
	desc.WriteString(content)
}

func removeFile(name string) {
	os.Remove(name)
}

type TestWindow struct{}

func (tw TestWindow) getDefinition() string {
	return testFile
}

func (s *UIReaderSuite) Test_parseUI_useXMLIfExists(c *C) {
	gtk.Init(nil)
	removeFile("definitions/TestWindow.xml")
	writeTestFile("definitions/TestWindow.xml", testFile)
	ui := "TestWindow"

	builder, parseErr := parseUI(ui)
	if parseErr != nil {
		fmt.Errorf("\nFailed!\n%s", parseErr.Error())
		c.Fail()
	}

	win, _ := builder.GetObject("conversation")
	w, h := win.(*gtk.Window).GetSize()
	c.Assert(h, Equals, 500)
	c.Assert(w, Equals, 400)
}

func (s *UIReaderSuite) Test_parseUI_useGoFileIfXMLDoesntExists(c *C) {
	gtk.Init(nil)
	removeFile("definitions/TestWindow.xml")
	//writeTestFile("definitions/TestWindow.xml", testFile)
	ui := "TestWindow"

	builder, parseErr := parseUI(ui)
	if parseErr != nil {
		fmt.Errorf("\nFailed!\n%s", parseErr.Error())
		c.Fail()
	}

	win, _ := builder.GetObject("conversation")
	w, h := win.(*gtk.Window).GetSize()
	c.Assert(h, Equals, 500)
	c.Assert(w, Equals, 400)
}

func (s *UIReaderSuite) Test_parseUI_shouldReturnErrorWhenDefinitionDoesntExist(c *C) {
	removeFile("definitions/nonexistent")
	ui := "nonexistent"

	_, parseErr := parseUI(ui)

	expected := "There's no definition for nonexistent"
	c.Assert(parseErr.Error(), Equals, expected)
}
