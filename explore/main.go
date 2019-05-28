package main

import (
	"fmt"
	"os"
)

const version = "0.0.1"
const title = `

___________              .__                        
\_   _____/__  _________ |  |   ___________   ____  
 |    __)_\  \/  /\____ \|  |  /  _ \_  __ \_/ __ \ 
 |        \>    < |  |_> >  |_(  <_> )  | \/\  ___/ 
/_______  /__/\_ \|   __/|____/\____/|__|    \___  >
        \/      \/|__|                           \/

`

const directions = `
This will be an explanation of how this works


Have fun exploring!`

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
