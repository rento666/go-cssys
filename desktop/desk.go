package desktop

import (
	"os"
	"os/exec"
	"os/signal"
)

func abc() {

	edgePath := os.Getenv("ProgramFiles(x86)") + "/Microsoft/Edge/Application/msedge.exe"
	cmd := exec.Command(edgePath, "--app=http://127.0.0.1:8080/")
	cmd.Start()

	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	select {
	case <-chSignal:
		cmd.Process.Kill()
	}

}
