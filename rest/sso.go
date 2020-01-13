package main

import (
	"context"
	"log"
	"net/http"

	pbx "github.com/abaron/chat/pbx"
	af "github.com/abaron/chat/server/adiraFinance"
	"google.golang.org/grpc"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG"))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	af.Log.Info("[ Create User ]: ")

	conn, err := grpc.Dial(config.GrpcListen, grpc.WithInsecure())
	if err != nil {
		af.Log.Error(err)
		log.Fatal("Error dialing", err)
	}

	c := pbx.NewNodeClient(conn)
	response, err := c.MessageLoop(context.Background())

	if err != nil {
		af.Log.Error(err)
		return
	}

	// hi := &pbx.ClientHi{}
	// hi.Id = "1"
	// hi.UserAgent = "Golang_Spider_Bot/3.0"
	// hi.Ver = "0.16"
	// hi.Lang = "EN"

	// msgHi := &pbx.ClientMsg_Hi{hi}
	// clientMessage := &pbx.ClientMsg{Message: msgHi}
	// err = response.Send(clientMessage)

	// if err != nil {
	// 	af.Log.Error("error sending message " + err.Error())
	// 	return
	// }

	// login := &pbx.ClientLogin{}
	// login.Id = "test1"
	// login.Scheme = "basic"
	// login.Secret = []byte("test1123")
	// clMsg := &pbx.ClientMsg_Login{login}
	// clientMessage = &pbx.ClientMsg{Message: clMsg}
	// err = response.Send(clientMessage)

	// if err != nil {
	// 	af.Log.Error("error sending message " + err.Error())
	// 	return
	// }

	// serverMsg, err := response.Recv()
	// if err != nil {
	// 	af.Log.Error("Recv1: " + err.Error())
	// 	return
	// }
	// log.Println(serverMsg)

	// serverMsg, err = response.Recv()
	// if err != nil {
	// 	af.Log.Error("Recv2: " + err.Error())
	// 	return
	// }
	// log.Println(serverMsg)

	register := &pbx.ClientAcc{
		Id: "123456",
		// UserId: "new",
		Scheme: "basic",
		Secret: []byte("dGVzdDExOnRlc3QxMTEyMw=="), // test11:test11123
		Login:  true,
		Tags:   []string{"test11", "test"},
		Cred: []*pbx.ClientCred{
			&pbx.ClientCred{
				Method: "email",
				Value:  "test11@mailinator.com",
			},
		},
		Desc: &pbx.SetDesc{
			Public: []byte("{\"fn\":\"test11\"}"),
		},
	}
	reg := &pbx.ClientMsg_Acc{register}
	clientMessage := &pbx.ClientMsg{Message: reg}
	err = response.Send(clientMessage)

	log.Println("ERR: ", err)
	serverMsg, err := response.Recv()
	if err != nil {
		af.Log.Error("Recv1: " + err.Error())
		return
	}
	log.Println(serverMsg)

	// a := pbx.AccountEvent{
	// 	Action: pbx.Crud_CREATE,
	// 	UserId: ,
	// }
}