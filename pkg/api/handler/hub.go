package handler

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newLine = []byte{'\n'}
	space   = []byte{' '}
)

type Peers struct{
	ListLock sync.Mutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}
func (p *Peers) AddTrack(t *webrtc.TrackRemote)*webrtc.TrackLocalStaticRTP{
	p.ListLock.Lock()
	defer func ()  {
		p.ListLock.Unlock()
		//p.SignalPeerconnection()
	}()
	trackLocal,err:=webrtc.NewTrackLocalStaticRTP(t.Codec().RTPCodecCapability,t.ID(),t.StreamID())
		if err!=nil{
			log.Println(err.Error())
			return nil
		}
		p.TrackLocals[t.ID()]=trackLocal
		return trackLocal

}

func (p *Peers) RemoveTrack(t *webrtc.TrackLocalStaticRTP){
	p.ListLock.Lock()
	defer func ()  {
		p.ListLock.Unlock()
		//p.SignalPeerconnection()
	}()
	delete(p.TrackLocals,t.ID())
}
// func(p *Peers) SignalPeerconnection(){
// 	p.ListLock.Lock()
// 	defer func ()  {
// 		p.ListLock.Unlock()
// 		p.DispatchKeyFrame()
// 	}()
// 	attemptSync:=func ()(tryAgain bool)  {
// 		for i:=range p.Connections{
// 			if p.Connections[i].PeerConnection.ConnectionState()==webrtc.PeerConnectionStateClosed{
// 				p.Connections=append(p.Connections[:i],p.Connections[i+1:]... )
// 				log.Println("a",p.Connections)
// 				return true
// 			}
// 			existingSenders:=map[string]bool{}
// 			for _,sender:=range p.Connections[i].PeerConnection.GetSenders(){
// 				if sender.Track()==nil{
// 					continue
// 				}	
// 				existingSenders[sender.Track().ID()]=true
// 				if _,ok:=p.TrackLocals[sender.Track().ID()];!ok{
// 					if err:=p.Connections[i].PeerConnection.RemoveTrack(sender);err!=nil{
// 						return true
// 					}
// 				}
// 			}
// 			for _,receiver:=range p.Connections[i].PeerConnection.GetReceivers(){
// 				if receiver.Track()==nil{
// 					continue
// 				}
// 				existingSenders[receiver.Track().ID()]	=true
// 			}

// 			for trackId:=range p.TrackLocals{
// 				if _,ok:=existingSenders[trackId];!ok{
// 					if _,err:=p.Connections[i].PeerConnection.AddTrack(p.TrackLocals[trackId]);err!=nil{
// 						return true
// 					}
// 				}	
// 			}
// 		}
// 	}
// }
func (p *Peers) DispatchKeyFrame(){

}

type ThreadSafeWriter struct {
	Conn  *websocket.Conn
	Mutex sync.Mutex
}

type PeerConnectionState struct{
	PeerConnection *webrtc.PeerConnection
	WebSocket *ThreadSafeWriter
}

func (t *ThreadSafeWriter) WriteJson(v interface{})error{
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	return t.Conn.WriteJSON(v)
}

type Client struct {
	Hub    *Hub
	UserId uint
	Conn   *websocket.Conn
	Send   chan *Message
}

type WebsocketMessage struct {
	Event string
	Data  string
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

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				return
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message.Msg)
			n := len(c.Send)

			for i := 0; i < n; i++ {
				w.Write(newLine)
				msg := <-c.Send
				w.Write(msg.Msg)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(appData string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var msg *Message
		err := c.Conn.ReadJSON(msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error:%v", err)
			}
			break
		}
		c.Hub.Data <- msg
	}
}

func PeerChatConn(c *websocket.Conn, hub *Hub) {
	client := &Client{Hub: hub, Conn: c, Send: make(chan *Message, 256)}
	client.Hub.Register <- client

	go client.WritePump()
	client.ReadPump()
}
