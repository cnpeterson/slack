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
    ID                 string            `json:"id"`
    Name               string            `json:"name"`
    IsChannel          bool              `json:"is_channel"`
    IsGroup            bool              `json:"is_group"`
    IsIm               bool              `json:"is_im"`
    Created            int64             `json:"created"`
    Creator            string            `json:"creator"`
    IsArchived         bool              `json:"is_archived"`
    IsGeneral          bool              `json:"is_general"`
    Unlinked           string            `json:"unlinked"`
    NameNormalized     string            `json:"name_normalized"`
    IsShared           bool              `json:"is_shared"`
    IsExtShared        bool              `json:"is_ext_shared"`
    IsOrgShared        bool              `json:"is_org_shared"`
    PendingShared      string            `json:"pending_shared"`
    IsPendingExtShared bool              `json:"is_pending_ext_shared"`
    IsMember           bool              `json:"is_member"`
    IsPrivate          bool              `json:"is_private"`
    IsMpim             bool              `json:"is_mpim"`
    Topic              map[string]string `json:"topic"`
    Purpose            map[string]string `json:"purpose"`
    PreviousNames      []string          `json:"previous_names"`
    NumMembers         int               `json:"num_members"`
}
