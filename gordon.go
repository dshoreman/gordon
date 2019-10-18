package main

import (
	"fmt"
	"github.com/dshoreman/gordon/pkg/gordon"
	irc "github.com/fluffle/goirc/client"
	flag "github.com/ogier/pflag"
	"math/rand"
	"os"
	"time"
)

const version = "0.0.0"

var (
	bot      *irc.Conn
	channel  string
	server   string
	ident    string
	nickname string
	realname string
)

var squirrels = []string{
	"http://28.media.tumblr.com/tumblr_lybw63nzPp1r5bvcto1_500.jpg",
	"http://i.imgur.com/DPVM1.png",
	"http://d2f8dzk2mhcqts.cloudfront.net/0772_PEW_Roundup/09_Squirrel.jpg",
	"http://www.cybersalt.org/images/funnypictures/s/supersquirrel.jpg",
	"http://www.zmescience.com/wp-content/uploads/2010/09/squirrel.jpg",
	"http://1.bp.blogspot.com/_v0neUj-VDa4/TFBEbqFQcII/AAAAAAAAFBU/E8kPNmF1h1E/s640/squirrelbacca-thumb.jpg",
}

func main() {
	bot = gordon.CreateBot(nickname, ident, realname, channel)

	registerCommands()
	registerWatchers()

	gordon.Connect(bot, server)
}

func registerCommands() {
	gordon.AddTrigger(bot, "dataja", "Don't ask to ask, just ask!")
	gordon.AddTrigger(bot, "ping", "Pong!")
}

func registerWatchers() {
	gordon.AddCommand(bot, `ship\s*it`, func() string {
		return squirrels[rand.Intn(len(squirrels))]
	})
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n\n")
		flag.PrintDefaults()
	}

	flag.StringVarP(&server, "host", "h", "irc.freenode.net", "Which IRC network to connect to.")
	flag.StringVarP(&channel, "join", "j", "#gordon", "Set a channel to autojoin.")
	flag.StringVarP(&ident, "ident", "i", "gordon", "Set the bot's IRC nickname.")
	flag.StringVarP(&nickname, "nick", "n", "Gordon", "Set the bot's IRC nickname.")
	flag.StringVarP(&realname, "realname", "r", "Gordon", "Set the bot's real name.")
	showVersionInfo := flag.BoolP("version", "V", false, "Print version info and quit.")
	flag.Parse()

	if *showVersionInfo {
		fmt.Println("Gordon " + version)
		os.Exit(0)
	}

	rand.Seed(time.Now().Unix())
}
