package main

import (
	"encoding/json"
	"fmt"

	"github.com/FedeBev/dkronc-go/pkg/client"
	"github.com/FedeBev/dkronc-go/pkg/client/jobs"
)

func main() {
	cfg := client.DefaultTransportConfig()
	cfg.WithHost("localhost:8080")
	dkronc := client.NewHTTPClientWithConfig(nil, cfg)

	// Gets Status object.
	st, err := dkronc.DefaultOperations.Status(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, _ := json.MarshalIndent(st, ">", "  ")
	fmt.Println(string(data))

	// Show a job.
	params := jobs.NewShowJobByNameParams()
	params.WithJobName("test")

	job, err := dkronc.Jobs.ShowJobByName(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, _ = json.MarshalIndent(job, ">", "  ")
	fmt.Println(string(data))
}
