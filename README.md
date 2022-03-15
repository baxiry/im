#                                           *mersal* *مرسال* 
##### Instant messaging server, easy to use as a server or as a library

   
how to use ?

create a websocket on browser and send message as json :

to subscribe in a topic (channel) send :
```json
{
   "event":"subscribe",
   "channel":"my-channel-id"
}
```
event must be : ```subscribe```, ```unsubscribe```, ```message```,

Later we will add events:  ```reseive```, and```seen```,
We will also work to achieve quality service ```qos``` later .

to send message to channel/topic:
```json
{
   "event" : "message",
   "channel" : "my-channel-123",
   "data" : "hi frends"
}
```

then all client subscribe with "my-channel-123" will be receive "hi frinds" message

