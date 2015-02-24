class Transport():
    def __init__(self, payload, config, sock):
        payload = str(payload)
        sock.setblocking(0)
        sock.sendto(payload, (config.get('mothership').get('host'), 
                              config.get('mothership').get('port')))
