package main

import (
	"fmt"
	irc "github.com/fluffle/goirc/client"
	flag "github.com/ogier/pflag"
	"os"
	"strings"
)

const version = "0.0.0"

var (
	bot      *irc.Conn
	channel  string
	server   string
	ident    string
	nickname string
	realname string
	quit     chan bool
)

func main() {
	fmt.Println("Loading Gordon IRC bot...")
	bot = irc.SimpleClient(nickname, ident, realname)

	registerCoreHandlers()
	registerCommands()

	fmt.Println("Connecting to IRC...")
	if err := bot.ConnectTo(server); err != nil {
		fmt.Printf("Connection error: %s\n", err.Error())
	}

	<-quit
}

func registerCoreHandlers() {
	bot.HandleFunc(irc.CONNECTED, func(conn *irc.Conn, line *irc.Line) {
		fmt.Println("Gordon's alive!")

		conn.Join(channel)
	})

	bot.HandleFunc(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) {
		quit <- true
	})

	bot.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {
		fmt.Printf("[RECV] %s <%s> %s\n", line.Target(), line.Nick, line.Text())
	})
}

func registerCommands() {
	bot.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {
		if strings.ToLower(line.Text()) != "ping" {
			return
		}

		conn.Privmsg(line.Target(), "PONG!")
		fmt.Printf("[SEND] %s <%s> %s\n", line.Target(), conn.Me().Nick, "PONG!")
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
}
