package pkg

import (
	"bytes"
	"fmt"
	"os/exec"
)

//shell命令
 func RunCmd(cmdstring string) (string, error) {
	 var out bytes.Buffer
	 var stderr bytes.Buffer
	 cmd := exec.Command("/bin/sh", "-c", cmdstring)
	 cmd.Stdout = &out
	 cmd.Stderr = &stderr
	 err := cmd.Run()
	 if err != nil {
	 	return fmt.Sprintf("%s",stderr.String()),err
	 }
	   return fmt.Sprintf("Result:%v",out.String()),nil
 }
