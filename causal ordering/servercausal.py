import threading
import socket
import pickle
from random import randint
from time import sleep
# 
host='127.0.0.1'
port = 59000
server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind((host, port))
server.listen()
clients = []


def broadcast(message, client):
    for clien in clients:
        if clien != client:
            clien.send(message)
            
            


def handle_client(client):
    i = 0
    while True:
        try:
            message = client.recv(4096)
            broadcast(message, client)
        except:
            index = clients.index(client)
            clients.remove(client)
            client.close()
            break
   
def receive():
    while True:
        print('Server is running and listening ...')
        client, address = server.accept()
        print(f'connection is established with {str(address)}')
        clients.append(client)
        thread = threading.Thread(target=handle_client, args=(client,))
        thread.start()


if __name__ == "__main__":
    receive()
    
