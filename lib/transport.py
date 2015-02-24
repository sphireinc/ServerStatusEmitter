from json import dumps as json_dumps

class Transport():
    def __init__(self, payload, config, sock):
        payload = str(json_dumps(payload))
        print payload
        sock.setblocking(0)
        sock.sendto(payload, (config.get('mothership').get('host'), 
                              config.get('mothership').get('port')))
