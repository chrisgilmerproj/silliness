#! /usr/local/bin/python

import socket

# Settings
TCP_IP = '127.0.0.1'
TCP_PORT = 8001
BUFFER_SIZE = 1024
CONNECTIONS = 5

# Create Socket
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind((TCP_IP, TCP_PORT))
s.listen(CONNECTIONS)

# Run the server
while True:
    conn, addr = s.accept()
    print 'Connection address:', addr
    print "Received data:"
    print "\n---\n"
    while True:
        data = conn.recv(BUFFER_SIZE)
        print data
        if len(data) < BUFFER_SIZE:
            break
    print "\n---\n"
    conn.sendall('<a href="?location=home">home</a>')
    conn.close()
