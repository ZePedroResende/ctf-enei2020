package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/nlopes/slack"
)

func (ep *Endpoints) processSlash(command slack.SlashCommand, w http.ResponseWriter) {
	userid := command.UserID
	text := command.Text
	fmt.Println(userid)
	fmt.Println(text)
	re := regexp.MustCompile("[A-Za-z0-9]{9}")

	userid2 := re.FindStringSubmatch(text)

	if len(userid2) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(userid2[0], userid)

	user1, _ := ep.client.GetUserInfo(userid)
	user2, _ := ep.client.GetUserInfo(userid2[0])

	_, _, channelID, err := ep.client.OpenIMChannel("ULP7BM2UW")

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	inform := user1.Profile.RealName + "," + user2.Profile.RealName + "," + user1.Profile.Email + "," + user2.Profile.Email

	result := "Inscrito no ctf com " + user1.Profile.RealName + " e " + user2.Profile.RealName + " Devido a situacao de redes do forum o ctf vai correr num ambiente de rede local na qual a password do ap sera a derivacao desta primeira flag. Assim que esta rede local estiver montada irei avisar toda a gente que provou chegar a esta primeira flag.  A primeira pista e https://secret-brushlands-52678.herokuapp.com/audio.  GL HF"

	ep.client.PostMessage(channelID, slack.MsgOptionText(inform, false))

	_, _, channelIDUser1, _ := ep.client.OpenIMChannel(userid)
	_, _, channelIDUser2, _ := ep.client.OpenIMChannel(userid2[0])

	ep.client.PostMessage(channelIDUser1, slack.MsgOptionText(result, false))
	ep.client.PostMessage(channelIDUser2, slack.MsgOptionText(result, false))

	w.Write([]byte(result))
}

func (ep *Endpoints) processIncomingRequest(r *http.Request, w http.ResponseWriter) (slashCommand slack.SlashCommand, errs error) {
	verifier, err := slack.NewSecretsVerifier(r.Header, ep.verificationToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body = ioutil.NopCloser(io.TeeReader(r.Body, &verifier))
	slashCommand, err = slack.SlashCommandParse(r)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return slashCommand, err
	}

	return slashCommand, nil
}

// The HTTP request handler
func (ep *Endpoints) ctfEndpoint(w http.ResponseWriter, r *http.Request) {
	slashCommand, err := ep.processIncomingRequest(r, w)
	if err != nil {
		return
	}

	// See which slash command the message contains
	switch slashCommand.Command {
	case "/ctf":
		ep.processSlash(slashCommand, w)

	default:
		// Unknown command
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
