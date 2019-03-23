import blockchain as bc
import json
from urllib.parse import urlparse
from uuid import uuid4

class Bot:

    def __init__(self, fields=['data']):
        self.fields = fields

    def add_data(self,attributes):
        data=dict(zip(self.fields,attributes))
        bc.add_block(bc.create_block(data))

