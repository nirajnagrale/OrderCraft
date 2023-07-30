import threading
from random import randint
import socket
import pickle
from time import sleep
from random import random
host = '127.0.0.1'
port = 59000
server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind((host, port))
server.listen()
clients = []

gs_id=0

class my_class:
    messages=""
    s_id=0

def broadcast(message):
    for client in clients:
        client.send(message)
        # sleep(randint(1,3))

w=0
def handle_client(client):
    i = 0
    while True:
        try:
            global gs_id
            data = client.recv(4096).decode('utf-8')
            obj=my_class()
            gs_id=gs_id+1
            obj.s_id=gs_id
            obj.messages=data
            message = pickle.dumps(obj)
            # global w
            # while w == 0:
            #   x = 1
            # w = 0
            broadcast(message)
            # w = 1
            
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
    
