package main

import (
        
        "os/exec"
)

/*ExecuteBashCommand , executes the bash os commands and return back the response*/
func ExecuteBashCommand(command string) (string, bool) {

        //commands := strings.Split(command, " ")
        //out, err := exec.Command(commands[0], commands[1:]...).Output()
        out := exec.Command("bash", "-c", command)


        if err != nil {
                #Logger.Error(stderr)
                #Logger.Error("Error in executing the command!")
                return string(stderr), false
        }
        return string(out.stdout), true

}

ExecuteBashCommand("ls")
