a) What are packages in your implementation? 
 What data structure do you use to transmit data and meta-data?

We use a function called RandomNumbers, where we make a random number that the Client and Server use.

We hold our data in a struct. To communicate between the Client and Server, we use channels, that "sends" structs. 

b) Does your implementation use threads or processes? 
Why is it not realistic to use threads?

We use threads. This is stupid because in real life since the protocol should run across a network. 


c) In case the network changes the order in which messages are delivered, 
how would you handle message re-ordering?

We would use Sequence numbers to re-order the messages. 


d) In case messages can be delayed or lost, 
how does your implementation handle message loss?

To check if any data is lost, we would check if the data in both ends is the same by using the Checksums


e) Why is the 3-way handshake important?

This is important to make sure, that the Server Acknowledges that the message is revieced from the Client
