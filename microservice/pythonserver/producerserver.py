import pika
import uuid
import json
from fastapi import FastAPI,logger
from fastapi.middleware.cors import CORSMiddleware


app = FastAPI()
app.add_middleware(CORSMiddleware,allow_origins=["*"],allow_methods=["*"],allow_headers=["*"])


@app.get('/order')
def firstfxn():
    return {"here":"akwaaba"}


order = {
    'orderid': str(uuid.uuid4()),
    'userid': 'somusiddrid1232',
    'productname': "toy car",
    'quantity': 4,
    'price': 25.00,
    'total': 100.00,
    
}

connectionParam = pika.ConnectionParameters('localhost')

# create connection
connection = pika.BlockingConnection(connectionParam)


# create channel
channel = connection.channel()

# declare a queue
channel.queue_declare(queue='Order_Queue')

#declare exchange name of exchanfge is order
channel.exchange_declare(exchange='order',exchange_type='direct')

message = "I want two nkrumah toys"

# publish to queue
channel.basic_publish(
    exchange='',
    routing_key='Order_Queue',
    body=json.dumps(order))
print(f"Successfully published to queeu")

# close connection
connection.close()