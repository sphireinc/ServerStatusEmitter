import requests
import logging

class Transport():
    def __init__(self, payload, config, logger):
        logger.info("Start transport")

        url = config['host'] + ((":" + str(config['port'])) if config['port'] else "")
        response = requests.post(url, data=payload)

        response = { "status_code": response.status_code,
                     "encoding": response.encoding,
                     "headers": response.headers,
                     "content": response.text
                   }

        logger.info("End transport: " + str(response))
