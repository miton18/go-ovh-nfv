# go-ovh-nvf

## configuration

```go
import(
    "context"

    ovhnfv"github.com/miton18/go-ovh-nfv"
    "github.com/ovh/go-ovh/ovh"
)

client, err = ovh.NewClient(
	ovh.OvhEU,
	os.Getenv("APP_KEY"),
	os.Getenv("APP_SECRET"),
	os.Getenv("CONSUMER_KEY"),
)
if err != nil {
	log.Fatal(err.Error())
}

nfv := ovhnfv.New(client)
```

## Usage

```go
ctx := context.Background()
loadbalancers, err := nfv.Loadbalancer(tenant).List(ctx)
if err != nil {
    t.Error("cannot list lbs", err)
}
fmt.Println(fmt.Sprintf("Loadbalancers %+v", lbs))
```
