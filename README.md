# dkron-go: a simple go client for dkron

**dkron-go** is a very minimal and simple client for [dkron v3](https://github.com/distribworks/dkron), generated from dkron official swagger file

**NOTE**: At the moment the swagger file contains a little patch to [expose `next` job field](https://github.com/distribworks/dkron/issues/990)

## Usage example

```go

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


```
