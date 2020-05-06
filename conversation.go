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
    OK               bool          `json:"ok"`
    Channels         []ChannelList `json:"channels"`
    ResponseMetadata RM            `json:"response_metadata"`
    Info             map[string]I
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
    Unlinked           int               `json:"unlinked"`
    NameNormalized     string            `json:"name_normalized"`
    IsShared           bool              `json:"is_shared"`
    IsExtShared        bool              `json:"is_ext_shared"`
    IsOrgShared        bool              `json:"is_org_shared"`
    PendingShared      []string          `json:"pending_shared"`
    IsPendingExtShared bool              `json:"is_pending_ext_shared"`
    IsMember           bool              `json:"is_member"`
    IsPrivate          bool              `json:"is_private"`
    IsMpim             bool              `json:"is_mpim"`
    IsOpen             bool              `json:"is_open"`
    Topic              T                 `json:"topic"`
    Purpose            P                 `json:"purpose"`
    Priority           int               `json:"priority"`
}

type T struct {
    Value   string `json:"value"`
    Creator string `json:"creator"`
    LastSet int64  `json:"last_set"`
}

type P struct {
    Value   string `json:"value"`
    Creator string `json:"creator"`
    LastSet int64  `json:"last_set"`
}

type I struct {
    ID            string `json:"id"`
    Created       int64  `json:"created"`
    IsIm          bool   `json:"is_im"`
    IsOrShared    bool   `json:"is_org_shared"`
    User          string `json:"user"`
    IsUserDeleted bool   `json:"is_user_deleted"`
    Priority      int    `json:"priority"`
}
