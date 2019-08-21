# gcputil/vm

This utility exposes a simple functions shut down host VM. It looks up the current project ID as well as
host VM name and zone in which that VM executes and calls the compute API to shut it down.

This is helpful when container finished running in VM and you want to shutdown the host VM to avoid paying for idle time.

## Import

```shell
import "github.com/mchmarny/gcputil/vm"
```

## Usage

```shell
ctx := context.Background()
vm.ShutdownHostVM(ctx, "my-long-running-job")
```
