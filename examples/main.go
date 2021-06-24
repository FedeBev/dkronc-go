package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FedeBev/dkronc-go/v3/client"
)

func main() {
	c, err := client.NewClientWithResponses("http://localhost:8080/v1")
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()

	st, err := c.StatusWithResponse(ctx)
	if err != nil {
		fmt.Println(err)
	}

	data, _ := json.MarshalIndent(st.JSON200, ">", "  ")
	fmt.Println(string(data))

	job, err := c.ShowJobByNameWithResponse(ctx, "test")
	if err != nil {
		fmt.Println(err)
	}

	data, _ = json.MarshalIndent(job.JSON200, ">", "  ")
	fmt.Println(string(data))

	meta := &map[string]interface{}{
		"key": "value1",
	}

	getJobsParams := &client.GetJobsParams{
		Metadata: meta,
	}
	r, err := c.GetJobsWithResponse(ctx, getJobsParams)
	if err != nil {
		fmt.Println(err)
	}
	data, _ = json.MarshalIndent(r.JSON200, ">", "  ")
	fmt.Println(string(data))

}
