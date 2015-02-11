## Server Status Emitter

A server status emitter for multi-server deployments. Reports a JSON object back to 1 central server for reporting/audit/storage.

### Deployment

```
#!bash

$ apt-get install python2.7
$ apt-get install python-pip
$ pip -R install requirements.txt
$ supervisord main.py
```
