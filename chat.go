package slack

const (
    ChatPostMessageUrl = "https://slack.com/api/chat.postMessage"
)

type ChatPostMessageRequest struct {
    Token          string              `json:"token"`
    Channel        string              `json:"channel"`
    Text           string              `json:"text"`
    AsUser         string              `json:"as_user"`
    Attachments    []map[string]string `json:"attachments"`
    Blocks         string              `json:"blocks"`
    IconEmoji      string              `json:"icon_emjoi"`
    IconUrl        string              `json:"icon_url"`
    LinkNames      string              `json:"link_names"`
    Mrkdwn         string              `json:"mrkdwn"`
    Parse          string              `json:"parse"`
    ReplyBroadcast bool                `json:"reply_broadcast"`
    ThreadTs       string              `json:"thread_ts"`
    UnfurlLinks    bool                `json:"unfurl_links"`
    UnfurlMedia    float64             `json:"unfurl_media"`
    Username       int                 `json:"username"`
}


type ChatPostMessageResponse struct {
    Ok               bool              `json:"ok"`
    Channel          string            `json:"channel"`
    Timestamp        string            `json:"ts"`
    Message          ChatPostMessage   `json:"message"`
}

type ChatPostMessage struct {
    User        string            `json:"user"`
    Text        string            `json:"text"`
    Username    string            `json:"username"`
    BotId       string            `json:"bot_id"`
    Attachments AttachmentsType   `json:"attachments,omitempty"`
    Type        string            `json:"type"`
    Subtype     string            `json:"subtype"`
    TS          string            `json:"ts"`
    Team        string            `json:"team"`
    BotProfile  BP                `json:"bot_profile"`
}

type AttachmentsType struct {
    Text     string `json:"text"`
    Id       int    `json:"id"`
    Fallback string `json:"fallback"`
}

type BP struct {
    ID       string `json:"id"`
    Deleted  bool   `json:"deleted"`
    Name     string `json:"name"`
    Updated  int64  `json:"updated"`
    AppId    string `json:"app_id"`
    Icons    Icon   `json:"icons"`
    TeamId   string `json:"team_id"`
}

type Icon struct {
    Image36 string `json:"image_36"`
    Image48 string `json:"image_48"`
    Image72 string `json:"image_72"`
}

