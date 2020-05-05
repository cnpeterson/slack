package slack

import (
    "net/http"
)

type CommandRequest struct {
    Token        string `json:"token"`
    Command      string `json:"command"`
    Text         string `json:"text"`
    ResponseUrl  string `json:"response_url"`
    TriggerId    string `json:"trigger_id"`
    UserId       string `json:"user_id"`
    UserName     string `json:"user_name"`
    TeamId       string `json:"team_id"`
    EnterpriseId string `json:"enterprise_id"`
    ChannelId    string `json:"channel_id"`
}

type Message struct {
    Text string `json:"text"`
}

func RequestParse (r *http.Request) (s CommandRequest, err error) {
    if err = r.ParseForm(); err != nil {
        return s, err
    }
    s.Token = r.PostForm.Get("token")
    s.Command = r.PostForm.Get("command")
    s.Text = r.PostForm.Get("text")
    s.ResponseUrl = r.PostForm.Get("response_url")
    s.TriggerId = r.PostForm.Get("trigger_id")
    s.UserId = r.PostForm.Get("user_id")
    s.UserName = r.PostForm.Get("user_name")
    s.TeamId = r.PostForm.Get("team_id")
    s.EnterpriseId = r.PostForm.Get("enterprise_id")
    s.ChannelId = r.PostForm.Get("channel_id")
    return s, err
}

