package gota

import (
	"os"
	"testing"
)

func TestExtendPaths(t *testing.T) {
	p := []string{"a", "b"}
	newP := extendPaths("a", p)
	if newP[0] != "a/a" || newP[1] != "a/b" {
		t.Error("Expected 'a/a a/b', got ", newP[0], newP[1])
	}
}

func TestCreatePage(t *testing.T) {
	tmplSet := []string{"testlayout.html", "testcontent.html"}
	CreatePage("testpage.html", tmplSet)
	fileInfo, err := os.Stat("testpage.html")
	if err != nil {
		check("os.Stat: ", err)
	}
	if fileInfo == nil {
		t.Error("Expected 'testpage.html' file info, got ", fileInfo)
	}
	os.Remove("testpage.html")
}
