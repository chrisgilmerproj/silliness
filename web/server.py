#! /usr/local/bin/python

import os
import socket

# Settings
TCP_IP = '127.0.0.1'
TCP_PORT = 8001
BUFFER_SIZE = 1024
CONNECTIONS = 5

def parse_request(data):
    """
    The first line of the request has the
    verb, path, and request type
    """
    request = data.split('\n')[0]
    verb, path, r_type = request.split()
    return verb, path[1:]

def get_body(path):
    """
    Using the path return body information
    """
    contents = os.listdir(os.path.join(os.getcwd(), path))
    body = ""
    for item in contents:
        body += '<li><a href="{0}">{1}</a></li>'.format(item, item)
    body = "<ul>{0}</ul>".format(body)
    body = "<!DOCTYPE html><html><body>{0}</body></html>".format(body)
    return body

def main():
    # Create Socket
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.bind((TCP_IP, TCP_PORT))
    s.listen(CONNECTIONS)

    # Run the server
    while True:
        conn, addr = s.accept()
        print 'Connection address:', addr
        print "Received data"
        while True:
            data = conn.recv(BUFFER_SIZE)
            verb, path = parse_request(data)
            if len(data) < BUFFER_SIZE:
                break
        if path != 'favicon.ico':
            body = get_body(path)
            conn.sendall(body)
        conn.close()

if __name__ == "__main__":
    main()
