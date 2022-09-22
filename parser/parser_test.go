package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadHTMLDoc(t *testing.T) {
	filename1 := "../resources/example1.html"
	f1, err := ReadHTMLDoc(filename1)
	if f1 == nil || err != nil {
		t.Fatalf("Expected no error on opening filename1: %s", filename1)
	}

	filename2 := "../resources/example2.html"
	f2, err := ReadHTMLDoc(filename2)
	if f2 == nil || err != nil {
		t.Fatalf("Expected no error on opening filename2: %s", filename2)
	}

	filename3 := "../resources/example3.html"
	f3, err := ReadHTMLDoc(filename3)
	if f3 == nil || err != nil {
		t.Fatalf("Expected no error on opening filename3: %s", filename3)
	}

	filename4 := "../resources/example4.html"
	f4, err := ReadHTMLDoc(filename4)
	if f4 == nil || err != nil {
		t.Fatalf("Expected no error on opening filename4: %s", filename4)
	}

	filename5 := "I-do-not-exist.html"
	f5, err := ReadHTMLDoc(filename5)
	if f5 != nil || err == nil {
		t.Fatalf("Test was expected to fail with error on opening filename: %s (file: %v)", filename5, f5)
	}
}

func TestParseFile1(t *testing.T) {

	filename := "../resources/example1.html"
	f1, err := ReadHTMLDoc(filename)
	if f1 == nil || err != nil {
		t.Fatalf("Expected no error on opening filename1: %s", filename)
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

	assert.ElementsMatch(t, expectedSlice, actualSlice)
}

func TestParseFile2(t *testing.T) {

	filename2 := "../resources/example2.html"
	f2, err := ReadHTMLDoc(filename2)
	if f2 == nil || err != nil {
		t.Fatalf("Expected no error on opening filename2: %s", filename2)
	}

	var link1 Link
	link1.Href = "https://www.twitter.com/joncalhoun"
	link1.Text = "Check me out on twitter"

	link2 := Link{Href: "https://github.com/gophercises", Text: "Gophercises is on Github"}

	var expectedSlice LinkList
	expectedSlice = append(expectedSlice, link1, link2)

	actualSlice := ParseFile(f2)
	if actualSlice == nil {
		t.Fatalf("Linklist should not be empty: %v", actualSlice)
	}

	assert.ElementsMatch(t, expectedSlice, actualSlice)
}

func TestParseFile3(t *testing.T) {

	filename3 := "../resources/example3.html"
	f3, err := ReadHTMLDoc(filename3)
	if f3 == nil || err != nil {
		t.Fatalf("Expected no error on opening filename2: %s", filename3)
	}

	link1 := Link{Href: "#", Text: "Login"}
	link2 := Link{Href: "/lost", Text: "Lost? Need help?"}
	link3 := Link{Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"}

	var expectedSlice LinkList
	expectedSlice = append(expectedSlice, link1, link2, link3)

	actualSlice := ParseFile(f3)
	if actualSlice == nil {
		t.Fatalf("Linklist should not be empty: %v", actualSlice)
	}

	assert.ElementsMatch(t, expectedSlice, actualSlice)
}

func TestParseFile4(t *testing.T) {

	filename4 := "../resources/example4.html"
	f4, err := ReadHTMLDoc(filename4)
	if f4 == nil || err != nil {
		t.Fatalf("Expected no error on opening filename2: %s", filename4)
	}

	link1 := Link{Href: "/dog-cat", Text: "dog cat"}

	var expectedSlice LinkList
	expectedSlice = append(expectedSlice, link1)

	actualSlice := ParseFile(f4)
	if actualSlice == nil {
		t.Fatalf("Linklist should not be empty: %v", actualSlice)
	}

	assert.ElementsMatch(t, expectedSlice, actualSlice)
}
