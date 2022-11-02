/* Alta3 Research | RZFeeser
   Executing system commands  */

package main

import (
    "fmt"
    "log"
    "os/exec"
)

func main() {

    // prepares a "cmd" struct
//    cmd := exec.Command("ls")
    out, err := exec.Command("nslookup", "google.com").Output()

//    test, err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(out))
}
