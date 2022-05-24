package main

import (
	"os"
	"strings"
	"testing"
)

func TestTaskRun(t *testing.T) {
	tmp := t.TempDir()

	mock := Task{
		Title:    "NextJs Page",
		Schedule: []Command{{
			Directory: tmp,
			Filename:  "test.html",
			Template:  "__mocks__/tpl.hbs",
			Output:    "",
		}},
		Params:   []string{"title"},
		Vars:     map[string]string{"t": "test title", "l" : "LOREM IPSUM", "u" : "some upper case ipsum", "cc" : "this should be camel case"},
		Parser:   Parser{},
		Base:     "",
	}
	
	GENIE.BASE = "."

	mock.Run()

	o, err := os.ReadFile(tmp + "/test.html")
	Check(err)
	
	oString := string(o)

	if(!strings.Contains(oString, "Test Title")){
		t.Logf("Expected: %s\n Got %s\n", "Test Title", oString)
		t.Error("To Title Case Helper failure")
	}

	lower := "lorem ipsum"
	if(!strings.Contains(oString, lower)){
		t.Logf("Expected: %s\n Got %s\n", lower, oString)
		t.Error("To Lower Helper failure")
	}

	upper := "SOME UPPER CASE IPSUM"
	if(!strings.Contains(oString, upper)){
		t.Logf("Expected: %s\n Got %s\n", upper, oString)
		t.Error("To Lower Helper failure")
	}

	cc := "ThisShouldBeCamelCase"
	if(!strings.Contains(oString, cc)){
		t.Logf("\nExpected: %s\n Got: %s\n", cc, oString)
		t.Error("To Lower Helper failure")
	}
}