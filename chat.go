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
    Messages         []ChatPostMessage `json:"message"`
}

type ChatPostMessage struct {
    Text        string            `json:"text"`
    Username    string            `json:"username"`
    BotId       string            `json:"bot_id"`
    Attachments []AttachmentsType `json:"attachments"`
    Type        string            `json:"type"`
    Subtype     string            `json:"subtype"`
    TS          string            `json:"ts"`
}

type AttachmentsType struct {
    Text     string `json:"text"`
    Id       int    `json:"id"`
    Fallback string `json:"fallback"`
}
