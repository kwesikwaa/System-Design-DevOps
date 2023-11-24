import pika
import json

connectionParam = pika.ConnectionParameters('localhost')

# create connection
connection = pika.BlockingConnection(connectionParam)


# create channel
channel = connection.channel()

def onReceived(ch,method,properties,body):
    x= json.loads(body)
    print('----DONE-----')
    ch.basic_ack(delivery_tag=method.delivery_tag)


# declare a queue
# this is idempotent so rabbimq will know if it 
# has been declared twice and ignore
channel.queue_declare(queue='Order_Queue')

channel.queue_bind(
    exchange='order',
    queue='Order_Queue',
    routing_key="Order_Queue"
)

# consume
channel.basic_consume(queue='Order_Queue',auto_ack=True,on_message_callback=onReceived)


print('Start Consuming')

channel.start_consuming()