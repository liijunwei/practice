package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// practice project from https://youtu.be/Sphme0BqJiY

// commands
//
// /nick {nickname}   - set nickname
// /join {room_name}  - join chat room
// /rooms             - list room names
// /msg {msg}         - broadcast msg to everyone in the room
// /quit              - quit chat room

// TODO how to test this app automatically/efficiently?
func main() {
	s := newServer()
	go s.run()

	address := ":8989"
	listener, err := net.Listen("tcp", address)
	boom(err)

	defer listener.Close()
	log.Println("server started", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v", err)
			continue
		}

		go s.newClient(conn)
	}
}

func assert(ok bool) {
	if !ok {
		panic("contract violation")
	}
}

func boom(err error) {
	if err != nil {
		panic(err)
	}
}

type commandID int8

const (
	CMD_NICK commandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type command struct {
	id     commandID
	client *client
	args   []string
}

type client struct {
	conn     net.Conn
	nickname string
	room     *room // TODO: many to many as another practice
	commands chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Printf("client disconnected: %s", c.conn.RemoteAddr().String())
				return
			}

			log.Printf("failed to read from client: %v", err)
			return
		}

		log.Println("raw client message:", strings.TrimSuffix(msg, "\n"))
		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		assert(len(args) >= 1)

		clientCommand := command{
			client: c,
			args:   args,
		}

		cmd := strings.TrimSpace(args[0])
		switch cmd {
		case "/nick":
			clientCommand.id = CMD_NICK
			c.commands <- clientCommand
		case "/join":
			clientCommand.id = CMD_JOIN
			c.commands <- clientCommand
		case "/rooms":
			clientCommand.id = CMD_ROOMS
			c.commands <- clientCommand
		case "/msg":
			clientCommand.id = CMD_MSG
			c.commands <- clientCommand
		case "/quit":
			clientCommand.id = CMD_QUIT
			c.commands <- clientCommand
		default:
			c.sendErr(fmt.Errorf("unknown command %s", cmd))
		}
	}
}

func (c *client) sendErr(err error) {
	c.conn.Write([]byte("ERROR: " + err.Error() + "\n"))
}

func (c *client) sendMsg(sender, msg string) {
	var buf bytes.Buffer

	buf.WriteString("> ")
	buf.WriteString(sender)
	buf.WriteString(": ")
	buf.WriteString(msg)
	buf.WriteString("\n")

	c.conn.Write(buf.Bytes())
}

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(tp string, sender *client, msg string) {
	assert(tp == "system" || tp == "user")

	sendername := "system"
	if tp == "user" {
		sendername = sender.nickname
	}

	for addr, member := range r.members {
		if addr != sender.conn.RemoteAddr() {
			member.sendMsg(sendername, msg)
		}
	}
}

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Println("new client connected:", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nickname: "visitor",
		commands: s.commands,
	}

	c.readInput()
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listrooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		default:
			assert(false) // should not be here
		}
	}
}

func (s *server) nick(c *client, args []string) {
	assert(len(args) == 2)
	c.nickname = args[1]
	c.sendMsg("system", "hi, I will call you "+c.nickname+" from now")
}

func (s *server) join(c *client, args []string) {
	if len(args) != 2 {
		c.sendMsg("system", "room name is required. usage: /join {room_name}")
		return
	}

	assert(len(args) == 2)
	assert(args[0] == "/join")
	roomname := args[1]

	r, ok := s.rooms[roomname]
	if !ok {
		r = &room{
			name:    roomname,
			members: make(map[net.Addr]*client),
		}

		s.rooms[roomname] = r
	}

	r.members[c.conn.RemoteAddr()] = c
	s.quitCurrentRoom(c)
	c.room = r
	r.broadcast("system", c, c.nickname+" has joined the room")
	c.sendMsg("system", "welcome to "+r.name)
}

func (s *server) listrooms(c *client, args []string) {
	assert(len(args) == 1)
	assert(args[0] == "/rooms")

	rooms := make([]string, 0, len(s.rooms))

	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.sendMsg("system", fmt.Sprintf("available rooms are:\n%s", strings.Join(rooms, "\n")))
}

func (s *server) msg(c *client, args []string) {
	assert(len(args) >= 2)
	assert(args[0] == "/msg")

	if c.room == nil {
		c.sendErr(fmt.Errorf("you must join the room first"))
		return
	}

	msg := strings.Join(args[1:], " ")
	c.room.broadcast("user", c, msg)
}

func (s *server) quit(c *client, args []string) {
	log.Printf("client has disconnected: %s", c.conn.RemoteAddr().String())
	s.quitCurrentRoom(c)
	c.sendMsg("system", "see you.")
	c.conn.Close()
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast("system", c, c.nickname+" has left the room")
	}
}
