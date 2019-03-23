def new_block(index, data, previous_hash):

    block = {
        'index': index,
        'timestamp': time(),
        'data': data,
        'previous_hash': previous_hash,
    }

    return block
