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
  result = convert("1-243-867-5309")
  if (result != "·---- -····- ··--- ····- ···-- -····- ---·· -···· --··· -····- ····· ···-- ----- ----·") {
    t.Log("Expected '·---- -····- ··--- ····- ···-- -····- ---·· -···· --··· -····- ····· ···-- ----- ----·' but got ", result)
    t.FailNow()
  }
  result = convert(".,!?")
  if (result != "·-·-·- --··-- -·-·-- ··--··") {
    t.Log("Expected '·-·-·- --··-- -·-·-- ··--··' but got ", result)
    t.FailNow()
  }
  result = convert("(){}")
  if (result != "☹ ☹ ☹ ☹") {
    t.Log("Expected '☹☹☹☹' but got ", result)
    t.FailNow()
  }
}
