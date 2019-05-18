import blockchain as bc
import json
from urllib.parse import urlparse
from uuid import uuid4

class Bot:

    def __init__(self):
        self.fields = ['data','sender','receiver','validator']

    def add_data(self,attributes):
        data, sender, receiver, validator = dict(zip(self.fields,attributes))
        bc.add_block(bc.create_block(data, sender, receiver, validator))

