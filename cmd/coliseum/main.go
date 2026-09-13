package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"flag"

	"github.com/dannyjimenez98/cluster-coliseum/internal/node"
)

func main() {
	// if this process was spawned as a node, run node logic and skip the launcher.
	if len(os.Args) > 1 && os.Args[1] == "node" {
		startNode()
		return
	}

	createColiseum()
}


func startNode() {
// 1. parse args from cmd
// 2. create node
// 3. start http server
	nodeFlagSet := flag.NewFlagSet("node", flag.ExitOnError)

	nodeID := nodeFlagSet.String("node-id", "", "Node ID (required)")
	port :=  nodeFlagSet.Int("port", 0, "Port number for the node server")

	// Skip executable and "node" subcommand, then parse node flags
	_ = nodeFlagSet.Parse(os.Args[2:])

	if *nodeID == "" || *port == 0 {
		log.Fatal("ERROR: required node flags are missing or invalid")
	}

	n := node.NewNode(*nodeID)
	fmt.Printf("port %v --> node: %#v\n",*port, n)
}


// program launcher function
// builds the starting nodes in the cluster (coliseum)
// TODO: 
// - improve doc comments and explain in more detail what this function process 
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
			"--node-id", fmt.Sprintf("%d", i),
			"--port", fmt.Sprintf("%d", 8000+i),
		)

		cmd.Stdout = os.Stdout	
		cmd.Stderr = os.Stderr	
		if err := cmd.Start(); err != nil {
			log.Fatalln(err)
		}
		log.Printf("Started node %d with PID %d", i, cmd.Process.Pid)
	}

}
