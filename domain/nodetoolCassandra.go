package domain

import (
	"bytes"
	"github.com/shshang/gorilla-mux-cassandra/errs"
	"log"
	"os/exec"
	"regexp"
)

//Cassandra is the adaptor of Nodetool
type NodetoolCassandra struct {
	command string
	args    []string
}

func (c NodetoolCassandra) RetrieveNodetoolStatus() ([]NodetoolStatus, *errs.AppError) {
	cmd := exec.Command(c.command, c.args...)
	var output bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		// this is the log for the developer to debug api server
		log.Printf("command %s with args %s returned error: %s", c.command, c.args, stderr.String())
		// this is the error message returned to the end user
		return nil, errs.NewInternalServerError("unexpected cassandra command error")
	}

	getFullName := func(s string) string {
		status, ok := map[string]string{
			"U": "up",
			"D": "down",
			"N": "normal",
			"L": "leaving",
			"J": "joining",
			"M": "moving",
			"S": "stopped",
		}[string(s)]

		if !ok {
			status = s
		}
		return status
	}

	nodeTexts := regexp.MustCompile(`(?m)^.*(([0-9a-fA-F]+-){4}([0-9a-fA-F]+)).*$`).FindAllString(output.String(), -1)
	var nodeStatus []NodetoolStatus
	for _, nodeText := range nodeTexts {
		comps := regexp.MustCompile(`[[:space:]]+`).Split(nodeText, -1)
		nodeStatus = append(nodeStatus,
			NodetoolStatus{
				Status:  getFullName(string(comps[0][0])),
				State:   getFullName(string(comps[0][1])),
				Address: comps[1],
				Load:    comps[2] + "GB",
				Tokens:  comps[4],
				Owns:    comps[5],
				HostId:  comps[len(comps)-2],
				Rack:    comps[len(comps)-1],
			})
	}

	return nodeStatus, nil
}

// NewNodetool is the helper function
func NewNodetool() NodetoolCassandra {
	command := "oc"
	args := []string{"-n", "sysdigcloud", "exec", "sysdigcloud-cassandra", "--", "nodetool", "status"}

	return NodetoolCassandra{command: command, args: args}
}
