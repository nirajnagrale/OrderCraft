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


p = input("enter no of client in distributed_system")
q = int(p)

m = input("enter no of message")
n = int(m)

dic = {0: {0: 'hello'}, 1: {0: 'hello'}, 2: {
    0: 'hello'}, 3: {0: 'hello'}, 4: {0: 'hello'}, 5: {0: 'hello'}}

vector_clock = []

for i in range(q):
    vector_clock.append(-1)


class my_class:
    messages = ""
    pid = 1
    lst = []


p_id = int(alias)
p = 0
w = 1
z=1


def client_receive():
    while True:
        try:
            s = 1
            t = 0
            global vector_clock
            data = client.recv(4096)
            message = pickle.loads(data)
            print(f'received message{message.pid}{message.lst}')
            r_pid = message.pid
            t_id = message.lst[r_pid]
            if vector_clock[r_pid]+1 == t_id:
                for i in range(q):
                    if i != r_pid and message.lst[i] > vector_clock[i]:
                        dic[r_pid][t_id] = message.lst
                        #print("lst is greater then vector clock")
                        t = 1
                        break
                if t == 0:
                    print(f'deliver {message.lst}')
                    #print(f"before update vector clock{vector_clock}")
                    vector_clock[r_pid] = t_id
                    print(f"after update vector clock{vector_clock}")
                    #print(f'updated vecvtor clock after delivery:{vector_clock}')
                    t_id = t_id+1
                    while s:
                        a=1
                        s = 1
                        for key, value in dic[r_pid].items():
                            if (key == t_id):
                                a=2
                                messag=dic[r_pid][t_id]
                                for i in range(q):
                                    if i != r_pid and messag[i] > vector_clock[i]:
                                        s = 0
                                        break

                        if s == 1 and a==2:
                            print(f'deliver {dic[r_pid][t_id]}')
                            print(f"after update vector clock{vector_clock}")
                            vector_clock[r_pid] = t_id
                            t_id = t_id+1

                        if a==1:
                            s = 0

            else:
                dic[r_pid][t_id] = message.lst
                print("place the message is a buffer")

        except Exception as e:
            print(e)
            # print(e)
            # print('Error')
            # client.close()
            # break


def client_send(threads_id):
    global vector_clock
    
    obj = my_class()
    # obj.messages = f'process{alias}: has send message with threads id is: {} '
    obj.pid = p_id
    sleep(randint(20, 30))
    global q
    while q == 0:
        x = 1
    q=0
    vector_clock[p_id] = vector_clock[p_id]+1
    obj.lst = vector_clock
    q = 1
    
    #print(f'message send{obj.lst}')
    obj.messages = f'process{alias}: has send message with threads id is: {obj.lst[p_id]} '
    # obj.lst[p_id] = threads_id
    data_string = pickle.dumps(obj)
    
    global w
    while w == 0:
        x = 1
    w=0
    sleep(randint(2, 5))
    client.send(data_string)
    w = 1


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
