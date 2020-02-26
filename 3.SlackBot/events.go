package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

func (ep *Endpoints) eventsEndpoint(w http.ResponseWriter, r *http.Request) {

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: ep.verificationToken}))

	if e != nil {
		fmt.Println(body)
		fmt.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		ep.urlVerification(w, body)
		return
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		ep.innerEvent(w, eventsAPIEvent)
		return
	}
}

func (ep *Endpoints) urlVerification(w http.ResponseWriter, body string) {
	var r *slackevents.ChallengeResponse
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text")
	w.Write([]byte(r.Challenge))

}

func (ep *Endpoints) innerEvent(w http.ResponseWriter, eventsAPIEvent slackevents.EventsAPIEvent) {

	innerEvent := eventsAPIEvent.InnerEvent

	fmt.Println(innerEvent.Data)

	switch ev := innerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:

		fmt.Println(ev.User)

		user, err := ep.client.GetUserInfo(ev.User)

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)

		ep.client.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
	}
}
