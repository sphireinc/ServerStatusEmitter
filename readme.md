## Server Status Emitter

A server status emitter for multi-server deployments. Reports a JSON object back to 1 central server for 
reporting/audit/storage.

### Requirements

The SSE requires a few packages, these are available from the distro package manager, or via pip:

1. psutil (pip)
2. python2.7 (pkg)
3. python-dev (pkg)
4. python-pip (pkg)
5. supervisor (pkg)

### Deployment

```
#!bash

$ ./setup.sh
```

If after that the program has not yet started, feel free to do:

```
#!bash

$ supervisorctl 
> stop SSE_Python
> start SSE_Python
CTRL+C
```
