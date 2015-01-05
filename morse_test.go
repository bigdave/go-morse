package main

import "testing"

func TestMapLetter(t *testing.T) {
  result := mapLetter('s')
  if (result != "···") {
    t.Log("Expected '···' but got ", result)
    t.FailNow()
  }
}

func TestConvert(t *testing.T) {
  result := convert("The quick brown fox jumped over the lazy dog")
  if (result != "- ···· ·        --·- ··- ·· -·-· -·-        -··· ·-· --- ·-- -·        ··-· --- -··-        ·--- ··- -- ·--· · -··        --- ···- · ·-·        - ···· ·        ·-·· ·- --·· -·--        -·· --- --·") {
    t.Log("Expected '- ···· ·        --·- ··- ·· -·-· -·-        -··· ·-· --- ·-- -·        ··-· --- -··-        ·--- ··- -- ·--· · -··        --- ···- · ·-·        - ···· ·        ·-·· ·- --·· -·--        -·· --- --·' but got ", result)
    t.FailNow()
  }
}
