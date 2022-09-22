package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadHTMLDoc(t *testing.T) {
	filename1 := "../resources/example1.html"
	f1, err := ReadHTMLDoc(filename1)
	if f1 == nil || err != nil {
		t.Fatalf("Expected no error on opening filename: %s", filename1)
	}

	filename2 := "../resources/example2.html"
	f2, err := ReadHTMLDoc(filename2)
	if f2 == nil || err != nil {
		t.Fatalf("Expected no error on opening filename: %s", filename2)
	}

	filename3 := "../resources/example3.html"
	f3, err := ReadHTMLDoc(filename3)
	if f3 == nil || err != nil {
		t.Fatalf("Expected no error on opening filename: %s", filename3)
	}

	filename4 := "../resources/example4.html"
	f4, err := ReadHTMLDoc(filename4)
	if f4 == nil || err != nil {
		t.Fatalf("Expected no error on opening filename: %s", filename4)
	}

	filename5 := "I-do-not-exist.html"
	f5, err := ReadHTMLDoc(filename5)
	if f5 != nil || err == nil {
		t.Fatalf("Test should fail with error on opening filename: %s (file: %v)", filename5, f5)
	}
}

func TestParseFile(t *testing.T) {

	filename1 := "../resources/example1.html"
	f1, err := ReadHTMLDoc(filename1)
	if f1 == nil || err != nil {
		t.Fatalf("Expected no error on opening filename: %s", filename1)
	}

	var linkExample1 Link
	linkExample1.Href = "/other-page"
	linkExample1.Text = "A link to another page"

	var expectedSlice LinkList
	expectedSlice = append(expectedSlice, linkExample1)

	actualSlice := ParseFile(f1)

	if actualSlice == nil {
		t.Fatalf("Linklist should not be empty: %v", actualSlice)
	}

	t.Log(actualSlice)
	t.Log(expectedSlice)

	assert.ElementsMatch(t, expectedSlice, actualSlice)

}
