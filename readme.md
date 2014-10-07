## Server Status Emitter

A server status emitter for multi-server deployments. Reports a JSON object back to 1 central server for reporting/audit/storage.


### Deployment

1. Add this as a git submodule (in .gitsubmodule file) so it is linked. 
2. Perform the following:

    $ pip -R install requirements.txt
    $ supervisord main.py


