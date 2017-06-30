package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "strconv"
	_ "time"

	"github.com/googollee/go-socket.io"
)

type elem struct {
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	Id string  `json:"id"`
}

func add_elem(current elem, all_elem *[]elem) {
	*all_elem = append(*all_elem, current)
}

func convert_into_struct(curr string) elem {
	var tmp elem
	json.Unmarshal([]byte(curr), &tmp)
	return tmp
}

func get_id(curr string) string {
	var tmp elem
	json.Unmarshal([]byte(curr), &tmp)
	return tmp.Id
}

func remove_elem(id string, all_elem *[]elem) {

	for i, curr := range *all_elem {
		if curr.Id == id {
			*all_elem = append((*all_elem)[:i], (*all_elem)[i+1:]...)
		}
	}
}

func check_id(pos string, all_elem *[]elem) {
	tmp := convert_into_struct(pos)
	if tmp.Id == "" {
		return
	}

	for i, curr := range *all_elem {
		if curr.Id == tmp.Id {
			(*all_elem)[i].X = tmp.X
			(*all_elem)[i].Y = tmp.Y
			return
		}
	}

	add_elem(convert_into_struct(pos), all_elem)
}

func convert_into_json(all_elem []elem) string {

	var tmp string
	var count int

	tmp += "["
	count = len(all_elem)
	for i, curr := range all_elem {
		tmp += "{" + "\"x\":" + fmt.Sprintf("%.2f", curr.X) + ",\"y\":" + fmt.Sprintf("%.2f", curr.Y) + ",\"id\":" + "\"" + curr.Id + "\"" + "}"
		if i != count-1 {
			tmp += ","
		}
	}
	tmp += "]"
	fmt.Println(tmp)
	return tmp
}

func main() {

	//var all_elem []elem

	//ticker := time.NewTicker(time.Millisecond * 30)
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}
	/*
		server.On("connection", func(so socketio.Socket) {
			var id_string string
			so.On("position", func(pos string) {
				id_string = get_id(pos)
				check_id(pos, &all_elem)
				//so.Emit("position", convert_into_json(all_elem))
				//		for t := range ticker.C {
				//			fmt.Println("tick at", t)

				so.Emit("position", convert_into_json(all_elem))
				//		}

			})

			so.On("disconnection", func(data string) {
				remove_elem(id_string, &all_elem)
			})

		})

		server.On("error", func(so socketio.Socket, err error) {
			log.Println("error:", err)
		})
	*/

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Fatal(http.ListenAndServe(":5000", nil))
}
