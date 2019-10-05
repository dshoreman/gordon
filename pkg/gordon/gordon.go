package gordon

import (
	"fmt"
	irc "github.com/fluffle/goirc/client"
	"strings"
)

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
