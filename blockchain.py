import hashlib
import json
from time import time
from urllib.parse import urlparse
from uuid import uuid4

blockchain_file = "./.blockchain"

def create_block(data):
    last_block = get_prev_block()
    block = {
        'index': last_block['index']+1,
        'timestamp': time(),
        'data': data,
        'previous_hash': hash_block(last_block),
    }
    return block

def add_block(block):
    file = open(blockchain_file, "a")
    file.write(json.dumps(block))
    file.write("\n")
    file.close()

def read_blockchain():
    file = open(blockchain_file, "r")
    content = file.read()
    print(content)
    file.close()


def get_prev_block():
    file = open(blockchain_file, "r")
    content = file.readlines()
    json_block = json.loads(content[-1].strip())
    file.close()
    return json_block

def read_prev_block():
    file = open(blockchain_file, "r")
    content = file.readlines()
    print(content[-1])
    file.close()

def hash_block(block):
    block_string = json.dumps(block, sort_keys=True).encode()
    return hashlib.sha256(block_string).hexdigest()
