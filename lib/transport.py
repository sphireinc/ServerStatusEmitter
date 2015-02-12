import requests
import logging

class Transport():
    def __init__(self, payload, config):
        logging.basicConfig(filename="/var/log/sse.log", filemode='a',
                            format='%(asctime)s,%(msecs)d %(name)s %(levelname)s %(message)s',
                            datefmt='%H:%M:%S', level=logging.INFO)
        logging.info("Start transport")

        url = config['host'] + ((":" + str(config['port'])) if config['port'] else "")
        response = requests.post(url, data=payload)

        response = { "status_code": response.status_code,
                     "encoding": response.encoding,
                     "headers": response.headers,
                     "content": response.text
                   }

        logging.info("End transport: " + str(response))
