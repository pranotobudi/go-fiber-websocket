// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	socketio "github.com/googollee/go-socket.io"
// )

// func main() {
// 	// app := fiber.New()
// 	server := socketio.NewServer(nil)

// 	server.OnConnect("/", func(s socketio.Conn) error {
// 		s.Join("chat")
// 		s.SetContext("")
// 		fmt.Println("connected:", s.ID())
// 		return nil
// 	})

// 	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
// 		fmt.Println("notice:", msg)
// 		s.Emit("reply", "have "+msg)
// 	})

// 	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
// 		s.SetContext(msg)
// 		return "recv " + msg
// 	})

// 	server.OnEvent("/", "bye", func(s socketio.Conn) string {
// 		last := s.Context().(string)
// 		s.Emit("bye", last)
// 		s.Close()
// 		return last
// 	})

// 	server.OnError("/", func(s socketio.Conn, e error) {
// 		fmt.Println("meet error:", e)
// 	})

// 	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
// 		fmt.Println("closed", reason)
// 	})

// 	go server.Serve()
// 	defer server.Close()

// 	http.Handle("/socket.io/", server)
// 	http.Handle("/", http.FileServer(http.Dir("./asset")))
// 	log.Println("Serving at localhost:8000...")
// 	log.Fatal(http.ListenAndServe(":8000", nil))

// }
package main

import (
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func main() {
	// This handler demonstrates how to safely accept cross origin WebSockets
	// from the origin example.com.
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{"example.com"},
		})
		if err != nil {
			log.Println(err)
			return
		}
		c.Close(websocket.StatusNormalClosure, "cross origin WebSocket accepted")
	})
	log.Println("Serving at localhost:8080...")
	// 	log.Fatal(http.ListenAndServe(":8000", nil))

	err := http.ListenAndServe("localhost:8080", fn)
	log.Fatal(err)
}
