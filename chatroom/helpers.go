package chatroom

import (
  "strconv"
)

func toIntSlice(strs []string) []int {
    list := make([]int, 0)

    for _, str := range strs {
      num, _ := strconv.ParseInt(str, 10, 0)
      list = append(list, int(num))
    }

    return list
}

func inArray(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
