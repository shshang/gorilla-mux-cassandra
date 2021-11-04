package domain

import (
	"log"
	"os/exec"
	"regexp"
)

type NodetoolStatus struct {
	Status  string `json:"status"`
	State   string `json:"state"`
	Address string `json:"address"`
	Load    string `json:"load"`
	Tokens  string `json:"tokens"`
	Owns    string `json:"owns"`
	HostId  string `json:"hostid"`
	Rack    string `json:"rack"`
}

//Nodetool is the primary port on server side
type Nodetool interface {
	RetrieveNodetoolStatus() ([]NodetoolStatus, error)
}

//Cassandra is the adaptor of Nodetool
type Cassandra struct {
	command string
	args    []string
}

func (c Cassandra) RetrieveNodetoolStatus() ([]NodetoolStatus, error) {
	cmd := exec.Command(c.command, c.args...)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return nil, err
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

	nodeTexts := regexp.MustCompile(`(?m)^.*(([0-9a-fA-F]+-){4}([0-9a-fA-F]+)).*$`).FindAllString(string(output), -1)
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

// NewCassandra is the helper function
func NewNodetool() Cassandra {
	command := "oc"
	args := []string{"-n", "sysdigcloud", "exec", "sysdigcloud-cassandra-0", "--", "nodetool", "status"}

	return Cassandra{command: command, args: args}
}
