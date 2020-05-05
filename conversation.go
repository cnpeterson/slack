package slack

const (
    ConvHistoryUrl = "https://slack.com/api/conversations.history"
    ConvListUrl = "https://slack.com/api/conversations.list"
)

type ConvHistoryRequest struct {
    Token     string   `json:"token"`
    Channel   string   `json:"channel"`
    Cursor    *string  `json:"cursor,omitempty"`
    Inclusive *bool    `json:"inclusive,omitempty"`
    Latest    *float64 `json:"latest,omitempty"`
    Limit     *int     `json:"limit,omitempty"`
    Oldest    *float64 `json:"oldest,omitempty"`
}

type ConvHistoryResponse struct {
    Ok               bool                 `json:"ok"`
    Messages         []ConvHistoryMessage `json:"messages"`
    HasMore          bool                 `json:"has_more"`
    PinCount         int                  `json:"pin_count"`
    ResponseMetadata RM                   `json:"response_metadata"`
}

type ConvHistoryMessage struct {
    Type string `json:"type"`
    User string `json:"user"`
    Text string `json:"text"`
    Ts   string `json:"ts"`
}

type RM struct {
    NextCursor string `json:"next_cursor"`
}

type ConvList struct {
    OK      bool `json:"ok"`
    Channels []ChannelList `json:"channels"`
    ResponseMetadata RM
}

type ChannelList struct {
    ID                 string
    Name               string
    IsChannel          bool
    IsGroup            bool
    IsIm               bool
    Created            int64
    Creator            string
    IsArchived         bool
    IsGeneral          bool
    Unlinked           string
    NameNormalized     string
    IsShared           bool
    IsExtShared        bool
    IsOrgShared        bool
    PendingShared      string
    IsPendingExtShared bool
    IsMember           bool
    IsPrivate          bool
    IsMpim             bool
    Topic              map[string]string
    Purpose            string
    PreviousNames      []string
    NumMembers         int
}
