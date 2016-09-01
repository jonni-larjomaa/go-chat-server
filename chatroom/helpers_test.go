package chatroom

import "testing"

func TestHelpers(t *testing.T) {
  strar := []string{"1","2","3","4","5"}
  intar := []int{1,2,3,4,5}

  strtoint := toIntSlice(strar)

  if strtoint[0] != intar[0] {
    t.Error("int Slice is not what is expected")
  }

  if inArray(strtoint[0], intar) != true {
    t.Error("integer value not found from intar array")
  }
}
