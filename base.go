package dingtalk

// MsgType message type
type MsgType string

const (
	// TextMsgType text message Type
	TextMsgType MsgType = "text"
	// LinkMsgType link message type
	LinkMsgType MsgType = "link"
	// MarkdownMsgType markdown message type
	MarkdownMsgType MsgType = "markdown"
	// ActionCardMsgType actionCard message type
	ActionCardMsgType MsgType = "actionCard"
	// FeedCardMsgType feedCard message Type
	FeedCardMsgType MsgType = "feedCard"
)

// Message dingtalk message
type Message interface {
	// ToMessageMap 返回Message对象组装出来的Map对象，供后续JSON序列化
	ToMessageMap() map[string]interface{}
}

// At dingtalk at
type At struct {
	// AtMobiles 被@人的手机号
	AtMobiles []string
	// AtUserIds 被@人的用户userid
	AtUserIds []string
	// atAll 是否@所有人
	AtAll bool
}

// NewAt create At
func NewAt(isAtAll bool) *At {
	return &At{
		AtAll:     isAtAll,
		AtMobiles: []string{},
		AtUserIds: []string{},
	}
}

// SetAtMobiles set AtMobiles
func (a *At) SetAtMobiles(atMobiles ...string) *At {
	a.AtMobiles = atMobiles
	return a
}

// AddAtMobiles add AtMobiles
func (a *At) AddAtMobiles(atMobiles ...string) *At {
	a.AtMobiles = append(a.AtMobiles, atMobiles...)
	return a
}

// SetAtUserIds set AtUserIds
func (a *At) SetAtUserIds(atUserIds ...string) *At {
	a.AtUserIds = atUserIds
	return a
}

// AddAtUserIds add AtUserIds
func (a *At) AddAtUserIds(atUserIds ...string) *At {
	a.AtUserIds = append(a.AtUserIds, atUserIds...)
	return a
}

func (a *At) ToMessageMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["atMobiles"] = a.AtMobiles
	message["atUserIds"] = a.AtUserIds
	message["isAtAll"] = a.AtAll
	return message
}

// ActionCardButton ActionCardMessage btn
type ActionCardButton struct {
	// Title 按钮标题
	Title string
	// ActionUrl 点击按钮触发的URL
	ActionUrl string
}

// NewActionCardButton create ActionCardButton
func NewActionCardButton(title, actionUrl string) ActionCardButton {
	return ActionCardButton{
		title,
		actionUrl,
	}
}
func (button ActionCardButton) ToMessageMap() map[string]interface{} {
	msg := map[string]interface{}{}
	msg["title"] = button.Title
	msg["actionURL"] = button.ActionUrl
	return msg
}

// Link FeedCardMessage links
type Link struct {
	// Title 单条信息文本
	Title string
	// MessageUrl 点击单条信息到跳转链接
	MessageUrl string
	// PicUrl 单条信息后面图片的URL
	PicUrl string
}

// NewLink create Link
func NewLink(title, messageUrl, picUrl string) Link {
	return Link{
		title,
		messageUrl,
		picUrl,
	}
}

func (link Link) ToMessageMap() map[string]interface{} {
	msg := map[string]interface{}{}
	msg["title"] = link.Title
	msg["messageURL"] = link.MessageUrl
	msg["picURL"] = link.PicUrl
	return msg
}

// TextMessage dingtalk text message
type TextMessage struct {
	// Content 消息内容
	Content string
	// At @信息
	At *At
}

// NewTextMessage create TextMessage
func NewTextMessage(content string) *TextMessage {
	return &TextMessage{
		Content: content,
	}
}

func (text *TextMessage) ToMessageMap() map[string]interface{} {
	contentMessage := map[string]interface{}{}
	contentMessage["content"] = text.Content
	message := map[string]interface{}{}
	message["msgtype"] = TextMsgType
	message["text"] = contentMessage
	if text.At != nil {
		message["at"] = text.At.ToMessageMap()
	}
	return message
}

// LinkMessage link message
type LinkMessage struct {
	// Text  消息内容。如果太长只会部分展示
	Text string
	// Title 消息标题
	Title string
	// PicUrl 图片URL
	PicUrl string
	// MessageUrl 点击消息跳转的URL
	MessageUrl string
}

// NewLinkMessage create LinkMessage
func NewLinkMessage(text, title, messageUrl string) *LinkMessage {
	return &LinkMessage{
		Title:      title,
		Text:       text,
		MessageUrl: messageUrl,
	}
}

func (link *LinkMessage) ToMessageMap() map[string]interface{} {
	linkMsg := map[string]interface{}{}
	linkMsg["text"] = link.Text
	linkMsg["title"] = link.Title
	linkMsg["messageUrl"] = link.MessageUrl
	linkMsg["picUrl"] = link.PicUrl
	msg := map[string]interface{}{}
	msg["msgtype"] = LinkMsgType
	msg["link"] = linkMsg
	return msg
}

// MarkdownMessage markdown message
type MarkdownMessage struct {
	// Title  首屏会话透出的展示内容
	Title string
	// Text markdown格式的消息
	Text string
	// At @信息
	At *At
}

// NewMarkdownMessage create MarkdownMessage
func NewMarkdownMessage(title, text string) *MarkdownMessage {
	return &MarkdownMessage{
		Title: title,
		Text:  text,
	}
}
func (markdown *MarkdownMessage) ToMessageMap() map[string]interface{} {
	markdownMsg := map[string]interface{}{}
	markdownMsg["title"] = markdown.Title
	markdownMsg["text"] = markdown.Text
	msg := map[string]interface{}{}
	msg["msgtype"] = MarkdownMsgType
	msg["markdown"] = markdownMsg
	if markdown.At != nil {
		msg["at"] = markdown.At.ToMessageMap()
	}
	return msg
}

// SingleActionCardMessage 整体跳转ActionCard类型
type SingleActionCardMessage struct {
	// Title 首屏会话透出的展示内容
	Title string
	// Text markdown格式的消息
	Text string
	// SingleTitle 单个按钮的标题
	SingleTitle string
	// SingleUrl 点击消息跳转的URL
	SingleUrl string
	// BtnOrientation 按钮排布方式
	// true: 按钮横向排列
	// false: 按钮竖直排列
	BtnOrientation bool
}

// NewSingleActionCardMessage create SingleActionCardMessage
func NewSingleActionCardMessage(title, text, singleTitle, singleUrl string) *SingleActionCardMessage {
	return &SingleActionCardMessage{
		Title:       title,
		Text:        text,
		SingleTitle: singleTitle,
		SingleUrl:   singleUrl,
	}
}

func (single *SingleActionCardMessage) ToMessageMap() map[string]interface{} {
	var strBtnOrientation = "0"
	if single.BtnOrientation {
		strBtnOrientation = "1"
	}

	singleMsg := map[string]interface{}{}
	singleMsg["title"] = single.Title
	singleMsg["text"] = single.Text
	singleMsg["singleTitle"] = single.SingleTitle
	singleMsg["singleURL"] = single.SingleUrl
	singleMsg["btnOrientation"] = strBtnOrientation
	msg := map[string]interface{}{}
	msg["msgtype"] = ActionCardMsgType
	msg["actionCard"] = singleMsg
	return msg
}

// ActionCardMessage 独立跳转ActionCard类型
type ActionCardMessage struct {
	// Title 首屏会话透出的展示内容
	Title string
	// Text markdown格式的消息
	Text string
	// Buttons 按钮
	Buttons []ActionCardButton
	// BtnOrientation 按钮排布方式
	// true: 按钮横向排列
	// false: 按钮竖直排列
	BtnOrientation bool
}

// NewActionCardMessage create ActionCardMessage
func NewActionCardMessage(title, text string, btns ...ActionCardButton) *ActionCardMessage {
	return &ActionCardMessage{
		Title:   title,
		Text:    text,
		Buttons: btns,
	}
}

// AddButtons add ActionCardMessage Buttons
func (actionCard *ActionCardMessage) AddButtons(btns ...ActionCardButton) *ActionCardMessage {
	actionCard.Buttons = append(actionCard.Buttons, btns...)
	return actionCard
}

func (actionCard *ActionCardMessage) ToMessageMap() map[string]interface{} {
	var btnsMsg []map[string]interface{}
	for _, btn := range actionCard.Buttons {
		btnsMsg = append(btnsMsg, btn.ToMessageMap())
	}
	var strBtnOrientation = "0"
	if actionCard.BtnOrientation {
		strBtnOrientation = "1"
	}
	actionCardMsg := map[string]interface{}{}
	actionCardMsg["title"] = actionCard.Title
	actionCardMsg["text"] = actionCard.Text
	actionCardMsg["btns"] = btnsMsg
	actionCardMsg["btnOrientation"] = strBtnOrientation

	msg := map[string]interface{}{}
	msg["msgtype"] = ActionCardMsgType
	msg["actionCard"] = actionCardMsg
	return msg
}

// FeedCardMessage FeedCard message
type FeedCardMessage struct {
	// Links Link list
	Links []Link
}

// NewFeedCardMessage create FeedCardMessage
func NewFeedCardMessage(link ...Link) *FeedCardMessage {
	return &FeedCardMessage{
		link,
	}
}

// AddLinks add FeedCardMessage Links
func (feedCard *FeedCardMessage) AddLinks(link ...Link) *FeedCardMessage {
	feedCard.Links = append(feedCard.Links, link...)
	return feedCard
}
func (feedCard *FeedCardMessage) ToMessageMap() map[string]interface{} {
	var links []map[string]interface{}
	for _, link := range feedCard.Links {
		links = append(links, link.ToMessageMap())
	}
	linkMsg := map[string]interface{}{}
	linkMsg["links"] = links
	msg := map[string]interface{}{}
	msg["msgtype"] = FeedCardMsgType
	msg["feedCard"] = linkMsg
	return msg
}
