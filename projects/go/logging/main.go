package main

import (
   "github.com/ca-mmis/ref-logging/projects/go/logging/logcli"
   "os"
)


func main() {

   if len(os.Args) == 0 {
      logcli.Info();
   }
   if len(os.Args) > 1 {
      logcli.Init()
      logcli.Info()
      os.Exit(logcli.Exec(os.Args))
   }

}

func Status() string {
   return "OKAY"
}

