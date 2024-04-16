package chat_type

import "github.com/gorilla/websocket"

type ChatRoom struct {
	RoomName    string
	Connections []*websocket.Conn
	Messages    []Message
	Users       []string
}

type Message struct {
	// MessageType string `json:"type"`
	MsgID    string `json:"id"`
	Type     string `json:"type"` // text  image  roomList file
	UserName string `json:"userName"`
	Content  string `json:"content"`
	Image    string `json:"image,omitempty"` // Base64-encoded image data
	File     string `json:"file,omitempty"`  // Base64-encoded file data
	SendTime string `json:"sendTime"`
	RoomName string `json:"roomName,omitempty"`

	ChatRoomList []string `json:"chatRoomList,omitempty"`
}

type User struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}
