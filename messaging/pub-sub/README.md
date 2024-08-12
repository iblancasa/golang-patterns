# Pub/sub
The Pub/Sub pattern is a way of communicating where there are two main roles: publishers and subscribers.

* Publishers: These are the entities that send out messages or information. They don't need to know who will receive the messages or how many people are listening. They just focus on sending out their messages.
* Subscribers: These are the entities that are interested in receiving messages. They "subscribe" to receive any message that a publisher sends out.

How it Works:
* Publishing: Imagine a person who wants to share news, like a town crier. This person stands in the middle of the village and announces news out loud. In the Pub/Sub pattern, this person is the publisher.
* Subscribing: The villagers who are interested in the news gather around to listen. These villagers are the subscribers. Whenever the town crier has something to announce, the subscribers who are present hear the news.
* Receiving Messages: In this pattern, subscribers receive messages as soon as the publisher sends them. If a subscriber is not present (or not listening), they miss the message.
