
# *im*  

##### Instant messaging server, easy to use as a server or as a library

### install :
```go get -u github.com/baxiry/im```

### how to use ?

create a websocket on browser and send message as json :

to subscribe to a topic (channel) send :
```json
{
   "event":"subscribe",
   "channel":"my-channel-id"
}
```
```event``` must be : ```subscribe```, ```unsubscribe```, ```message```,

Later we will add events:  ```reseive```, and```seen```,

to send message to channel/topic:
```json
{
   "event" : "message",
   "channel" : "my-channel-123",
   "data" : "hi frends"
}
```
then all client subscribe with "my-channel-123" will be receive "hi frinds" message.

### How do I create a private connection? between two clients?
Just create a channel and share it between two clients only, this is how you create a private connection

### Project status
- [x] websocket
- [x] Pubsub 
- [ ] cache message
- [ ] 100% Unit testing


### License
Use this library with whatever license you prefer.
