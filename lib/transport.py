import socket

class Transport():
    def __init__(self, payload, config, logger):
#        logger.info("Start transport")
        print "Send"

        #url = config['host'] + ((":" + str(config['port'])) if config['port'] else "")
        #response = requests.post(url, data=payload)

        payload = str(payload)
        sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        print sock, config
#        sock.setblocking(0)
        sock.sendto(payload, (config['host'], config['port']))

        #response = { "status_code": response.status_code,
        #             "encoding": response.encoding,
        #             "headers": response.headers,
        #             "content": response.text
        #           }

        #logger.info("End transport: " + str(response))
#        logger.info("End transport")
