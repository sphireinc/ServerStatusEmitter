import socket

class Transport():
    def __init__(self, payload, config, sock):
#        logger.info("Start transport")
        print "Send"

        #url = config['host'] + ((":" + str(config['port'])) if config['port'] else "")
        #response = requests.post(url, data=payload)

        payload = str(payload)
#        sock.setblocking(0)
        sock.sendto(payload, (config.get('mothership').get('host'), 
                              config.get('mothership').get('port')))

        #response = { "status_code": response.status_code,
        #             "encoding": response.encoding,
        #             "headers": response.headers,
        #             "content": response.text
        #           }

        #logger.info("End transport: " + str(response))
#        logger.info("End transport")
