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
client.connect(('10.21.210.108', 59000))

p = input("enter no of client in distributed_system")
q = int(p)

m = input("enter no of message")
n = int(m)

dic = {0: {0: 'hello'}, 1: {0: 'hello'}, 2: {
    0: 'hello'}, 3: {0: 'hello'}, 4: {0: 'hello'}, 5:{0: 'hello'}}

vector_clock = []

for i in range(q):
    vector_clock.append(-1)


class my_class:
    messages = ""
    pid = 1
    lst = []


p_id = int(alias)
p=0
w=1


def client_receive():
    while True:
        try:
            s=1
            global vector_clock
            global p
            data = client.recv(4096)
            message = pickle.loads(data)
            print(f'received:{message.messages}')
            r_pid = message.pid
            t_id = message.lst[r_pid]
            if vector_clock[r_pid]+1 == t_id:
                print(f'deliver: {message.messages}')
                vector_clock[r_pid] = t_id
                t_id = t_id+1
                while s:
                        s=1
                        for key,value in dic[r_pid].items():
                            if(key == t_id):
                                print(f'deliver:{dic[r_pid][t_id]}')
                                vector_clock[r_pid] = t_id
                                t_id=t_id+1
                                s=2
                                break
                    
                        if s==2:
                            continue
                        else:
                            s=0

            else:
                dic[r_pid][t_id] = message.messages
                print("message in in buffer")


        except Exception as e: 
            print(e)
            # print(e)
            # print('Error')
            # client.close()
            # break


def client_send(threads_id):
    obj = my_class()
    obj.messages = f'process{alias}: has send message with threads id is: {threads_id} '
    obj.pid = p_id
    obj.lst = vector_clock
    obj.lst[p_id] = threads_id

    data_string = pickle.dumps(obj)
    sleep(randint(20, 60))
    global w
    while w==0:
        x=1
    client.send(data_string)
    w=1


# proc=process(target=client_receive)
# proc.start()
# proc.join()

receive_thread = threading.Thread(target=client_receive)
receive_thread.start()

# receive_thread1 = threading.Thread(target=client_receive)
# receive_thread1.start()

# receive_thread2 = threading.Thread(target=client_receive)
# receive_thread2.start()

# with ThreadPoolExecutor() as executor:
#     executor.map(client_receive, range(q))
#     print(" receiver waiting")

with ThreadPoolExecutor() as executor:
    executor.map(client_send, range(n))
    print('Waiting...')
