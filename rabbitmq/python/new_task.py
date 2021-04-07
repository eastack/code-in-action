#!/usr/bin/env python
import pika, sys

message = ''.join(sys.argv[1:]) or "Hello World!"

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()

channel.queue_declare(queue='hello')

channel.basic_publish(exchange='',
        routing_key='hello',
        body=message)

print(" [X] Sent %r" % message)

connection.close()
