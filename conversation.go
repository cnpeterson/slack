package slack

const (
    ConvHistoryUrl = "https://slack.com/api/conversations.history"
    ConvListUrl = "https://slack.com/api/conversations.list"
)

type SlackCovoHistoryRequest struct {
    Token     string   `json:"token"`
    Channel   string   `json:"channel"`
    Cursor    *string  `json:"cursor,omitempty"`
    Inclusive *bool    `json:"inclusive,omitempty"`
    Latest    *float64 `json:"latest,omitempty"`
    Limit     *int     `json:"limit,omitempty"`
    Oldest    *float64 `json:"oldest,omitempty"`
}

type SlackConvoHistoryResponse struct {
    Ok               bool                      `json:"ok"`
    Messages         []SlackConvHistoryMessage `json:"messages"`
    HasMore          bool                      `json:"has_more"`
    PinCount         int                       `json:"pin_count"`
    ResponseMetadata RM                        `json:"response_metadata"`
}

type SlackConvHistoryMessage struct {
    Type string `json:"type"`
    User string `json:"user"`
    Text string `json:"text"`
    Ts   string `json:"ts"`
}

type RM struct {
    NextCursor string `json:"next_cursor"`
}
