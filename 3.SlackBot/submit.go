package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

func checkSubstrings(str string, subs ...string) bool {

	isCompleteMatch := true

	fmt.Printf("String: \"%s\", Substrings: %s\n", str, subs)

	for _, sub := range subs {
		if !strings.Contains(str, sub) {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch
}

func (ep *Endpoints) processSubmit(command slack.SlashCommand, w http.ResponseWriter) {
	m := map[string]int{
		"0b103ea6c189b0de5225855b31f5df73": 0,
		"a2e5f5bd750845b90016819e0111b4de": 1,
		"f197d7becb7b982a6e549b1c86afee7d": 2,
		"9d9eff5a54ebf6f14b07774abce00139": 3,
		"c045ddc4e04697fadcc12d414c105554": 4,
		"9d617b02bca0dc654f8ccdda1c714573": 5,
		"eaca7fac76f767601f0958a62448e970": 6,
		"f120c1ea8d7e6201a1210ee555cb58ab": 7,
		"003f44066473359772dfb31ca4d149ed": 8,
		"e6054d83c7c045fcd391c4be1f16e7a6": 9,
		"333a8a0de38e4e659c3fa307b1cfda24": 10,
	}

	userid := command.UserID
	text := command.Text

	words := strings.Fields(text)

	re := regexp.MustCompile("^[a-f0-9]{32}$")

	if "0b103ea6c189b0de5225855b31f5df73" == words[0] {
		w.Write([]byte("Ou tas a mandar a flag de exemplo ou encontraste uma decoy. Da proxima corre melhor"))
		return
	}

	if !re.MatchString(words[0]) {
		w.Write([]byte("Not a valid md5 hash"))
		return
	}
	user1, _ := ep.client.GetUserInfo(userid)

	_, _, channelID, err := ep.client.OpenIMChannel("ULP7BM2UW")

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	value, ok := m[words[0]]
	if ok {
		inform := user1.Profile.RealName + "," + user1.Profile.Email + "," + words[0] + "," + string(value)
		result := "Flag correta"

		ep.client.PostMessage(channelID, slack.MsgOptionText(inform, false))

		_, _, channelIDUser1, _ := ep.client.OpenIMChannel(userid)

		ep.client.PostMessage(channelIDUser1, slack.MsgOptionText(result, false))

		w.Write([]byte(result))
	} else {
		inform := user1.Profile.RealName + "," + user1.Profile.Email + " falhou: " + strings.Join(words, " ")
		result := "Flag errada"

		ep.client.PostMessage(channelID, slack.MsgOptionText(inform, false))

		_, _, channelIDUser1, _ := ep.client.OpenIMChannel(userid)

		ep.client.PostMessage(channelIDUser1, slack.MsgOptionText(result, false))
		log.Println(inform)
		log.Println(result)
		w.Write([]byte(result))
	}

}

func (ep *Endpoints) submitEndpoint(w http.ResponseWriter, r *http.Request) {
	slashCommand, err := ep.processIncomingRequest(r, w)
	if err != nil {
		return
	}

	// See which slash command the message contains
	switch slashCommand.Command {
	case "/submit":
		ep.processSubmit(slashCommand, w)

	default:
		// Unknown command
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
