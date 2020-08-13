package main

import (
	"log"
	"os"
	"time"

	obsws "github.com/christopher-dG/go-obs-websocket"
)

func main() {
	// Get the args, excluding the name of the program
	args := os.Args[1:]

	log.Println("Primary scene is: ", args[0])
	// Connect a client.
	c := obsws.Client{Host: "192.168.13.98", Port: 4444}

	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}
	defer c.Disconnect()

	// Send and receive a request asynchronously.
	req := obsws.NewGetStreamingStatusRequest()
	if err := req.Send(c); err != nil {
		log.Fatal(err)
	}
	// This will block until the response comes (potentially forever).
	resp, err := req.Receive()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("streaming:", resp.Streaming)

	// Set the amount of time we can wait for a response.
	obsws.SetReceiveTimeout(time.Second * 2)

	log.Println("Setting transition to fade")

	transReq := obsws.NewSetCurrentTransitionRequest("Fade")
	transResp, err := transReq.SendReceive(c)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Transition response: ", transResp.Status_)

	log.Println("Setting transition timer to 1500ms")
	transTimeReq := obsws.NewSetTransitionDurationRequest(1500)
	transTimeResp, err := transTimeReq.SendReceive(c)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Transition response: ", transTimeResp.Status_)

	log.Println("Setting scene to \"reconnecting\"")

	sceneChangeReq := obsws.NewSetCurrentSceneRequest("Reconnecting")
	sceneChangeResp, err := sceneChangeReq.SendReceive(c)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Scene change response: ", sceneChangeResp.Status_)

	time.Sleep(time.Second * 5)

	log.Println("Setting scene to \"reconnecting\"")

	sceneChangeReq2 := obsws.NewSetCurrentSceneRequest(args[0])
	sceneChangeResp2, err := sceneChangeReq2.SendReceive(c)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Scene change response: ", sceneChangeResp2.Status_)
}
