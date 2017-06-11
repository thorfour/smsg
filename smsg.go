package smsg

import "encoding/json"

const (
	Ephemeral = "ephemeral"
	InChannel = "in_channel"
)

type Resp struct {
	RespType    string       `json:"response_type,omitempty"`
	Text        string       `json:"text,omitempty"`
	Attachments []RespAttach `json:"attachments,omitempty"`
}

type RespAttach struct {
	Fallback    string      `json:"fallback,omitempty"`
	CallbackID  string      `json:"callback_id"`
	Color       string      `json:"color,omitempty"`
	Pretext     string      `json:"pretext,omitempty"`
	Author_name string      `json:"author_name,omitempty"`
	Author_link string      `json:"author_link,omitempty"`
	Author_icon string      `json:"author_icon,omitempty"`
	Title       string      `json:"title,omitempty"`
	Title_link  string      `json:"title_link,omitempty"`
	Text        string      `json:"text,omitempty"`
	Fields      []RespField `json:"fields,omitempty"`
	Image_url   string      `json:"image_url,omitempty"`
	Thumb_url   string      `json:"thumb_url,omitempty"`
	Footer      string      `json:"footer,omitempty"`
	Footer_icon string      `json:"footer_icon,omitempty"`
	Ts          string      `json:"ts,omitempty"`
	Actions     []Action    `json:"actions,omitempty"`
}

type RespField struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short string `json:"short,omitempty"`
}

type Action struct {
	Name    string      `json:"name,omitempty"`
	Text    string      `json:"text,omitempty"`
	Style   string      `json:"style,omitempty"`
	Type    string      `json:"type,omitempty"`
	Value   string      `json:"value,omitempty"`
	Confirm RespConfirm `json:"confirm,omitempty"`
}

type RespConfirm struct {
	Title        string `json:"title,omitempty"`
	Text         string `json:"text,omitempty"`
	Ok_text      string `json:"ok_text,omitempty"`
	Dismiss_text string `json:"dismiss_text,omitempty"`
}

func (s *Resp) String() string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	return string(b)
}

type Msg struct {
	Payload MsgPayload `json:"payload"`
}

type MsgPayload struct {
	Actions      []Action `json:"actions"`
	CallbackID   string   `json:"callback_id"`
	Team         TeamInfo `json:"team"`
	Channel      Channel  `json:"channel"`
	User         User     `json:"user"`
	ActionTS     string   `json:"action_ts"`
	MessageTS    string   `json:"message_ts"`
	AttachmentId string   `json:"attachment_id"`
	Token        string   `json:"token"`
	IsAppUnfurl  bool     `json:"is_app_unfurl"`
}

type TeamInfo struct {
	ID     string `json"id"`
	Domain string `json:"domain"`
}

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *Msg) String() string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	return string(b)
}
