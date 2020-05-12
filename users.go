package slack

const (
    UsersInfoUrl = "https://slack.com/api/users.info"
)

type UsersInfoRequest struct {
    Token         string `json:"token"`
    User          string `json:"user"`
    IncludeLocale bool   `json:"include_locale,omitempty"`
}

type UsersInfoResponse struct {
    Ok   bool `json:"ok"`
    User user `json:"user"`
}

type user struct {
    Id                string  `json:"id"`
    TeamId            string  `json:"team_id"`
    Name              string  `json:"name"`
    Deleted           bool    `json:"deleted"`
    Color             string  `json:"color"`
    RealName          string  `json:"real_name"`
    Tz                string  `json:"tz"`
    TzLabel           string  `json:"tz_label"`
    TzOffset          int64   `json:"tz_offset"`
    Profile           profile `json:"profile"`
    IsAdmin           bool    `json:"is_admin"`
    IsOwner           bool    `json:"is_owner"`
    IsPrimaryOwner    bool    `json:"is_primary_owner"`
    IsRestricted      bool    `json:"is_restricted"`
    IsUltraRestricted bool    `json:"is_ultra_restricted"`
    IsBot             bool    `json:"is_bot"`
    Updated           int64   `json:"updated"`
    IsAppUser         bool    `json:"is_app_user"`
    Has2fa            bool    `json:"has_2fa"`
}

type profile struct {
    AvatarHash            string `json:"avatar_hash"`
    StatusText            string `json:"status_text"`
    StatusEmoji           string `json:"status_emoji"`
    RealName              string `json:"real_name"`
    DisplayName           string `json:"display_name"`
    RealNameNormalized    string `json:"real_name_normalized"`
    DisplayNameNormalized string `json:"display_name_normalized"`
    Email                 string `json:"email"`
    ImageOriginal         string `json:"image_original"`
    Image24               string `json:"image_24"`
    Image32               string `json:"image_32"`
    Image48               string `json:"image_48"`
    Image72               string `json:"image_72"`
    Image192              string `json:"image_192"`
    Image512              string `json:"image_512"`
    Team                  string `json:"team"`
}
