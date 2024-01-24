package handler

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const(
	writeWait=10*time.Second
	pongWait=60*time.Second
	pingPeriod=(pongWait*9)/10
	maxMessageSize=512
)

var(
	newLine=[]byte{'\n'}
	space=[]byte{' '}
)

type Client struct {
	Hub    *Hub
	UserId uint
	Conn   *websocket.Conn
	Send   chan *Message
}

type Message struct {
	SenderId   uint
	ReceiverId uint
	Msg        []byte
}

type Hub struct {
	Register   chan *Client
	Unregister chan *Client
	Data       chan *Message
	Clients    map[*Client]bool
}

func NewHub() *Hub {
	return &Hub{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Data:       make(chan *Message),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

		case message := <-h.Data:
			for client := range h.Clients {
				if client.UserId == message.ReceiverId {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(h.Clients, client)
					}
				}
			}
		}
	}
}

func (c *Client) WritePump(){
	ticker:=time.NewTicker(pingPeriod)
	defer func(){
		ticker.Stop()
		c.Conn.Close()
	}()
	for{
		select{
		case message,ok:=<-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok{
				return
			}
			w,err:=c.Conn.NextWriter(websocket.TextMessage)
			if err!=nil{
				return
			}
			w.Write(message.Msg)
			n:=len(c.Send)

			for i:=0;i<n;i++{
				w.Write(newLine)
				msg:= <-c.Send
				w.Write(msg.Msg)
			}
			if err:=w.Close();err!=nil{
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err:=c.Conn.WriteMessage(websocket.PingMessage,nil);err!=nil{
				return
			}
		}
	}
}

func (c *Client) ReadPump(){
	defer func(){
		c.Hub.Unregister<-c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(appData string) error {c.Conn.SetReadDeadline(time.Now().Add(pongWait));return nil})
	for{
		var msg *Message
		err:=c.Conn.ReadJSON(msg)
		if err!=nil{
			if websocket.IsUnexpectedCloseError(err,websocket.CloseGoingAway,websocket.CloseAbnormalClosure){
				log.Printf("error:%v",err)
			}
			break
		}
		c.Hub.Data<-msg		
	}
}

func PeerChatConn(c *websocket.Conn,hub *Hub){
	client:=&Client{Hub: hub,Conn: c,Send: make(chan *Message,256)}
	client.Hub.Register<-client

	go client.WritePump()
	client.ReadPump()
}