# gowitter
golang twitter bot

### how to
```
rs@serv:~/gowitter$ l
gowitter-netgo-linux64*  run.sh* .tweets

nohup run.sh &
```

### run.sh content
```
#!/bin/bash -e

export CONSUMER_KEY=<from twitter dev> ;\
export CONSUMER_KEY_SECRET=<from twitter dev> ;\
export ACCESS_TOKEN=<from twitter dev> ;\
export ACCESS_TOKEN_SECRET=<from twitter dev> ;\
export TWEETS_JSON=./.tweets ;\
export INTERVAL_MIN=45 ;\
./gowitter-netgo-linux64 ;\
```

### .tweets content
```
{
  "tweets": ["a tweet", "another tweet"]
}
```
