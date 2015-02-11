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
        print url
        resp, content = http.request(url, "POST", payload)
        resp = {"status":               resp.get('status', ''), 
                "content-length":       resp.get('content-length', ''),
                "transfer-encoding":    resp.get('transfer-encoding', ''), 
                "server":               resp.get('server', ''), 
                "date":                 resp.get('date', ''), 
                "content-type":         resp.get('content-type', '')}
        print resp

        logging.info("End transport: " + str(resp))
