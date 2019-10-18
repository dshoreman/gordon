package gordon

import (
	"fmt"
	irc "github.com/fluffle/goirc/client"
	"regexp"
	"strings"
)

var (
	bot  *irc.Conn
	quit chan bool
)

// CreateBot connects to IRC and returns the connection
func CreateBot(nickname, ident, realname, channel string) *irc.Conn {
	fmt.Println("Loading Gordon IRC bot...")
	bot = irc.SimpleClient(nickname, ident, realname)

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

	return bot
}

// Connect connects to IRC
func Connect(bot *irc.Conn, server string) {
	fmt.Println("Connecting to IRC...")

	if err := bot.ConnectTo(server); err != nil {
		fmt.Printf("Connection error: %s\n", err.Error())
	}
	<-quit
}

// AddTrigger adds the given trigger with a basic response
func AddTrigger(bot *irc.Conn, trigger string, response string) {
	fmt.Printf("Registering trigger '%s' with response: \"%s\"...\n", trigger, response)

	bot.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {
		if strings.HasPrefix(line.Text(), "!"+trigger) {
			msg := response
			conn.Privmsg(line.Target(), msg)
			fmt.Printf("[SEND] %s <%s> %s\n", line.Target(), conn.Me().Nick, msg)
		}
	})
}

// AddCommand handles commands matching the given trigger regex
func AddCommand(bot *irc.Conn, trigger string, handler func() string) {
	fmt.Printf("Registering command matching '%s'...\n", trigger)

	bot.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {
		if match, _ := regexp.MatchString(trigger, line.Text()); false == match {
			return
		}

		msg := handler()
		conn.Privmsg(line.Target(), msg)
		fmt.Printf("[SEND] %s <%s> %s\n", line.Target(), conn.Me().Nick, msg)
	})
}
