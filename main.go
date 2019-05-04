package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "tweet-cli-go"
	app.Usage = "tweet-cli-go"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		args := context.Args()
		if context.Bool("echo") {
			fmt.Println(context.Args().Get(0))
		} else if len(args) > 0 {
			if context.Bool("retweet") {
				var tweetID int64
				tweetID, err := strconv.ParseInt(args.Get(0), 10, 64)
				if err != nil {
					fmt.Println(err)
					return err
				}
				api := getTwitterApi()
				tweet := anaconda.Tweet{}
				if !context.Bool("undo") {
					tweet, err = api.Retweet(tweetID, false)
					if err != nil {
						fmt.Println(err)
						return err
					}
				} else {
					tweet, err = api.UnRetweet(tweetID, false)
					if err != nil {
						fmt.Println(err)
						return err
					}
				}
				fmt.Println(tweet.Id)
				fmt.Println(tweet.FullText)
				fmt.Println(tweet.CreatedAt)
				fmt.Println("Is retweet")
				fmt.Println(tweet.Retweeted)
				return nil
			}
			if context.Bool("tweet") {
				var text string
				v := url.Values{}
				api := getTwitterApi()
				for index := 0; index < len(args); index++ {
					text += args.Get(index) + " "
				}
				if context.Bool("at") {
					text = strings.Replace(text, args.Get(0), "@"+args.Get(0), 1)
				}
				if context.Bool("reply") {
					tweetID, err := strconv.ParseInt(args.Get(0), 10, 64)
					if err != nil {
						fmt.Println(err)
						return err
					}
					baseTweet, err := api.GetTweet(tweetID, url.Values{})
					if err != nil {
						fmt.Println(err)
						return err
					}
					v.Add("in_reply_to_status_id", args.Get(0))
					text = strings.Replace(text, args.Get(0), "", 1)
					text = strings.Replace(text, text, "@"+baseTweet.User.ScreenName+" "+text, 1)
				}

				tweet, err := api.PostTweet(text, url.Values{})
				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println(tweet.Id)
				fmt.Println(tweet.FullText)
				fmt.Println(tweet.CreatedAt)
				return nil
			}

		}
		if context.Bool("home") || true {
			api := getTwitterApi()
			tweets, err := api.GetHomeTimeline(url.Values{})
			if err != nil {
				fmt.Println(err)
				return err
			}
			for _, tweet := range tweets {
				fmt.Println("----")
				fmt.Println(tweet.User.Name + " /@" + tweet.User.ScreenName)
				fmt.Println(tweet.Id)
				fmt.Println(tweet.FullText)
				fmt.Println(tweet.CreatedAt)
				fmt.Println("RT :" + fmt.Sprint(tweet.RetweetCount) + " Fav :" + fmt.Sprint(tweet.FavoriteCount))
				if nil != tweet.QuotedStatus {
					fmt.Println("--QT--")
					fmt.Println("    " + tweet.QuotedStatus.User.Name + " /@" + tweet.QuotedStatus.User.ScreenName)
					fmt.Println("    " + tweet.QuotedStatus.FullText)
					fmt.Println("--QT--")
				}
				fmt.Println("----")
			}
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "echo, e",
			Usage: "Echo",
		},
		cli.BoolFlag{
			Name:  "tweet, t",
			Usage: "tweet",
		},
		cli.BoolFlag{
			Name: "at, a",
			Usage: "at tweet(screen name) \n		(-t -a [to_screen_name] [tweet_texts...])",
		},
		cli.BoolFlag{
			Name:  "reply, rp",
			Usage: "reply(tweet_id)",
		},
		cli.BoolFlag{
			Name:  "retweet, r",
			Usage: "retweet(tweet_id)",
		},
		cli.BoolFlag{
			Name:  "undo, u",
			Usage: "undo(retweet and other)",
		},
		cli.BoolFlag{
			Name:  "home",
			Usage: "show timeline",
		},
	}

	app.Run(os.Args)
}

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}
