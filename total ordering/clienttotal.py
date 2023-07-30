import threading
import socket
from random import randint
from time import sleep
from random import random
from concurrent.futures import ThreadPoolExecutor
import pickle
from multiprocessing import process
alias = input('Choose an alias >>> ')
client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client.connect(('127.0.0.1', 59000))


m = input("enter no of message")
n = int(m)

dic = {0:'hello'}


class my_class:
    messages = ""
    s_id = 0


ls_id = 0
w = 1


def client_receive():
    while True:
        try:
            s = 1
            global ls_id
            data = client.recv(4096)
            message = pickle.loads(data)
            # print(f'received: {message.messages}')
            # print(f'sequence no: {message.s_id}')
            id = message.s_id
            if id == ls_id+1:
                print(f'deliver: {message.messages} with sequence id:{message.s_id}')
                ls_id = ls_id+1
                id = id+1
                while s:
                    s = 1
                    for key in dic:
                        if key == id:
                            print(f'deliver: {dic[key]} with sequence id:{key}')
                            id = id+1
                            ls_id = ls_id+1
                            s = 2
                            break
                    if s==2:
                            continue
                    else:
                        s=0

            else:
                dic[id] = message.messages
                # print(dic[id])

        except Exception as e:
            print(e)


def client_send(threads_id):
    message=f'process{alias}: has send message with threads id is: {threads_id}'
    sleep(randint(20, 60))
    global w
    # while w == 0:
    #     x = 1
    # w = 0
    client.send(message.encode('utf-8'))
    # w = 1



receive_thread = threading.Thread(target=client_receive)
receive_thread.start()

# receive_thread1 = threading.Thread(target=client_receive)
# receive_thread1.start()


with ThreadPoolExecutor() as executor:
    executor.map(client_send, range(n))
    print('Waiting...')
