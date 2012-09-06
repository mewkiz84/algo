package main

import "github.com/mewkiz84/algo/connectivity/homebrew"
import "runtime/pprof"
import "os"
import "log"

func main() {
   f, err := os.Create("a.prof")
   if err != nil {
      log.Fatal(err)
   }
   pprof.StartCPUProfile(f)

   n := 134217728
   set := homebrew.New(n)
   for x := 1; x < n; x *= 2 {
      for j := 0; j < n; j += x * 2 {
         set.Union(j, j+x)
      }
   }
   pprof.StopCPUProfile()
}
