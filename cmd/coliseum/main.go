package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// if this process was spawned as a node, run node logic and skip the launcher.
	if len(os.Args) > 1 && os.Args[1] == "node" {
		fmt.Println("this is a node - do not proceed to run launch cmd")
		return
	}

	createColiseum()
}

// program launcher function
// builds the starting nodes in the cluster (coliseum)
// TODO: 
// - improve doc comments and explain in more detail what this function process 
// - implement Flag pkg to parse the args from the cmd
// - add ability to run the node building cmd from the terminal as an option, rather than the default of having the program do it
// - tie this together with node logic function that the created node processes will run 
// - create http server for each node, using the port from the cmd line args
func createColiseum() {
	coliseum, err := os.Executable() // return the path to this executable
	if err != nil {
		log.Fatalf("ERROR: executable path not found:\n%v", err)
	}
  
	for i := 1; i <= 3; i++ {
		// create command that builds a node, taking node subcommand, node-id, and port as args
		cmd := exec.Command(
			coliseum,
			"node",
			fmt.Sprintf("--node-id=node-%d", i),
			fmt.Sprintf("--port=%d", 8000+i),
		)

		cmd.Stdout = os.Stdout	
		cmd.Stderr = os.Stderr	
		if err := cmd.Start(); err != nil {
			log.Fatalln(err)
		}
		log.Printf("Started node %d with PID %d", i, cmd.Process.Pid)
	}

}
