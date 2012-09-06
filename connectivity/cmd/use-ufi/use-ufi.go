package main

import dbg "fmt"
import "github.com/mewkiz84/algo/connectivity/ufi"

func main() {
   set := ufi.New(10)
   dbg.Println(set)

   set.Union(0, 1)
   set.Union(2, 3)
   set.Union(4, 5)
   set.Union(6, 7)
   set.Union(8, 9)

   dbg.Println(set)
}
