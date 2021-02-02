# twitter-cli-go
## Warning
    あまりにも適当なコードがツイートを行う謎のプログラムもどき
## How to use
### First settings
    go get -u github.com/okinjp/twitter-cli-go
    export "CONSUMER_KEY" = "twitter-app-consumer-key"
    export "CONSUMER_SECRET" = "twitter-app-consumer-secret"
    export "ACCESS_TOKEN" = "access-token"
    export "ACCESS_TOKEN_SECRET" = "access-token-secret"
    
 ### Run
    go run main.go -t "I'm first tweet from cli!"
 
### OptionList
   --echo, -e     Echo
   
   --tweet, -t    tweet
   
   --at, -a       at tweet(screen name)
   
   (-t -a [to_screen_name] [tweet_texts...])
                    
   --reply, --rp  reply(tweet_id)
   
   --retweet, -r  retweet(tweet_id)
   
   --undo, -u     undo(retweet and other)
   
   --home, (or without options)         show timeline
   
   --help, -h     show help
   
   --version, -v  print the version
