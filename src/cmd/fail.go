package main

import (
  "os"
  "fmt"
  "flag"
)

func main() {
  
  cmdline     := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
  fCode       := cmdline.Int      ("e",      1,           "The exit code to return after printing the message.")
  fSuppressNL := cmdline.Bool     ("n",      false,       "Suppress the trailing newline character.")
  cmdline.Parse(os.Args[1:])
  
  for i, e := range cmdline.Args() {
    if i > 0 { fmt.Fprint(os.Stderr, " ") }
    fmt.Fprint(os.Stderr, e)
  }
  if !*fSuppressNL {
    fmt.Fprintln(os.Stderr)
  }
  
  os.Exit(*fCode)
}
