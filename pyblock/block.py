from time import time

def new_block(index, sender, receiver, validator, data, previous_hash):

    block = {
        'index': index,
        'timestamp': time(),
        'sender': sender,
        'receiver': receiver,
        'validator': validator,
        'data': data,
        'previous_hash': previous_hash,
    }

    return block
