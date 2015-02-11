from httplib2 import Http
import logging

class Transport():
    def __init__(self, payload, config):
        logging.basicConfig(filename="/var/log/sse.log", filemode='a',
                            format='%(asctime)s,%(msecs)d %(name)s %(levelname)s %(message)s',
                            datefmt='%H:%M:%S', level=logging.INFO)
        logging.info("Start transport")

        http = Http()
        url = config['host'] + ((":" + str(config['port'])) if config['port'] else "")
        resp, content = http.request(url, "POST", payload)
        resp = [resp.get('status', ''), resp.get('content-length', ''),
                resp.get('transfer-encoding', ''), resp.get('server', ''), 
                resp.get('date', ''), resp.get('content-type', '')]

        logging.info("End transport: " + resp)
