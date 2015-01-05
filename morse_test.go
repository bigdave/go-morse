package main

import "testing"

func TestMapLetter(t *testing.T) {
  result := mapLetter('s')
  if (result != "···") {
    t.Log("Expected '···' but got ", result)
    t.FailNow()
  }
}
