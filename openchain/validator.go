/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package openchain

import (
	"errors"
	"fmt"
	"io"

	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"

	"github.com/golang/protobuf/proto"
	"github.com/looplab/fsm"
	"github.com/op/go-logging"
	"github.com/spf13/viper"

	"github.com/openblockchain/obc-peer/openchain/consensus/pbft"
	"github.com/openblockchain/obc-peer/openchain/util"
	pb "github.com/openblockchain/obc-peer/protos"
)

var validatorLogger = logging.MustGetLogger("validator")

type Validator interface {
	Broadcast(*pb.OpenchainMessage) error
	GetHandler(stream PeerChatStream) MessageHandler
}

type SimpleValidator struct {
	validatorStreams map[string]MessageHandler
	peerStreams      map[string]MessageHandler
	leader           pb.PeerClient
	leaderHandler    MessageHandler
}

func (v *SimpleValidator) Broadcast(msg *pb.OpenchainMessage) error {
	validatorLogger.Debug("Broadcasting OpenchainMessage of type: %s", msg.Type)
	err := v.leaderHandler.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("Error broadcasting msg of type: %s", msg.Type)
	}
	return nil
}

func (v *SimpleValidator) GetHandler(stream PeerChatStream) MessageHandler {
	return NewValidatorFSM(v, "", stream)
}

func (v *SimpleValidator) chatWithLeader(peerAddress string) error {

	var errFromChat error = nil
	conn, err := NewPeerClientConnectionWithAddress(peerAddress)
	if err != nil {
		return errors.New(fmt.Sprintf("Error connecting to leader at address=%s:  %s", peerAddress, err))
	}
	serverClient := pb.NewPeerClient(conn)
	stream, err := serverClient.Chat(context.Background())
	v.leaderHandler = v.GetHandler(stream)

	if err != nil {
		return errors.New(fmt.Sprintf("Error chatting with leader at address=%s:  %s", peerAddress, err))
	} else {
		defer stream.CloseSend()
		stream.Send(&pb.OpenchainMessage{Type: pb.OpenchainMessage_DISC_HELLO})
		waitc := make(chan struct{})
		go func() {
			for {
				in, err := stream.Recv()
				if err == io.EOF {
					// read done.
					errFromChat = errors.New(fmt.Sprintf("Error sending transactions to peer address=%s, received EOF when expecting %s", peerAddress, pb.OpenchainMessage_DISC_HELLO))
					close(waitc)
					return
				}
				if err != nil {
					grpclog.Fatalf("Failed to receive a DiscoverMessage from server : %v", err)
				}
				// Call FSM.HandleMessage()
				err = v.leaderHandler.HandleMessage(in)
				if err != nil {
					validatorLogger.Error("Error handling message: %s", err)
					return
				}

				// 	if in.Type == pb.OpenchainMessage_DISC_HELLO {
				// 		peerLogger.Debug("Received %s message as expected, sending transactions...", in.Type)
				// 		payload, err := proto.Marshal(transactionsMessage)
				// 		if err != nil {
				// 			errFromChat = errors.New(fmt.Sprintf("Error marshalling transactions to peer address=%s:  %s", peerAddress, err))
				// 			close(waitc)
				// 			return
				// 		}
				// 		stream.Send(&pb.OpenchainMessage{Type: pb.OpenchainMessage_CHAIN_TRANSACTIONS, Payload: payload})
				// 		peerLogger.Debug("Transactions sent to peer address: %s", peerAddress)
				// 		close(waitc)
				// 		return
				// 	} else {
				// 		peerLogger.Debug("Got unexpected message %s, with bytes length = %d,  doing nothing", in.Type, len(in.Payload))
				// 		close(waitc)
				// 		return
				// 	}
			}
		}()
		<-waitc
		return nil
	}
}

func NewSimpleValidator() (Validator, error) {
	validator := &SimpleValidator{}
	// Only perform if NOT the leader
	if !viper.GetBool("peer.consensus.leader.enabled") {
		leaderAddress := viper.GetString("peer.consensus.leader.address")
		validatorLogger.Debug("Creating client to Peer (Leader) with address: %s", leaderAddress)
		go validator.chatWithLeader(leaderAddress)
	}
	return validator, nil
}

type ValidatorFSM struct {
	To         string
	ChatStream PeerChatStream
	FSM        *fsm.FSM
	PeerFSM    *PeerFSM
	validator  Validator
}

func NewValidatorFSM(parent Validator, to string, peerChatStream PeerChatStream) *ValidatorFSM {
	v := &ValidatorFSM{
		To:         to,
		ChatStream: peerChatStream,
		validator:  parent,
	}

	v.FSM = fsm.NewFSM(
		"created",
		fsm.Events{
			{Name: pb.OpenchainMessage_DISC_HELLO.String(), Src: []string{"created"}, Dst: "established"},
			{Name: pb.OpenchainMessage_CHAIN_TRANSACTIONS.String(), Src: []string{"established"}, Dst: "established"},
		},
		fsm.Callbacks{
			"enter_state":                                               func(e *fsm.Event) { v.enterState(e) },
			"before_" + pb.OpenchainMessage_DISC_HELLO.String():         func(e *fsm.Event) { v.beforeHello(e) },
			"before_" + pb.OpenchainMessage_CHAIN_TRANSACTIONS.String(): func(e *fsm.Event) { v.beforeChainTransactions(e) },
		},
	)

	return v
}

func (v *ValidatorFSM) enterState(e *fsm.Event) {
	validatorLogger.Debug("The Validators's bi-directional stream to %s is %s, from event %s\n", v.To, e.Dst, e.Event)
}

func (v *ValidatorFSM) beforeHello(e *fsm.Event) {
	validatorLogger.Debug("Sending back %s", pb.OpenchainMessage_DISC_HELLO.String())
	if err := v.ChatStream.Send(&pb.OpenchainMessage{Type: pb.OpenchainMessage_DISC_HELLO}); err != nil {
		e.Cancel(err)
	}
}

func (v *ValidatorFSM) beforeChainTransactions(e *fsm.Event) {
	validatorLogger.Debug("Sending broadcast to all validators upon receipt of %s", pb.OpenchainMessage_DISC_HELLO.String())
	if _, ok := e.Args[0].(*pb.OpenchainMessage); !ok {

	}
	msg := e.Args[0].(*pb.OpenchainMessage)

	//
	//proto.Marshal()
	uuid, err := util.GenerateUUID()
	if err != nil {
		e.Cancel(fmt.Errorf("Error generating UUID: %s", err))
		return
	}
	request := &pbft.Request{Id: uuid, Payload: msg.Payload}
	data, err := proto.Marshal(request)
	if err != nil {
		e.Cancel(fmt.Errorf("Error marshalling Request: %s", err))
		return
	}
	newMsg := &pb.OpenchainMessage{Type: pb.OpenchainMessage_CONSENSUS, Payload: data}
	validatorLogger.Debug("Getting ready to create CONSENSUS from this message type : %s", msg.Type)
	v.validator.Broadcast(newMsg)
	// if err := v.ChatStream.Send(&pb.OpenchainMessage{Type: pb.OpenchainMessage_DISC_HELLO}); err != nil {
	// 	e.Cancel(err)
	// }
}

func (v *ValidatorFSM) when(stateToCheck string) bool {
	return v.FSM.Is(stateToCheck)
}

func (v *ValidatorFSM) HandleMessage(msg *pb.OpenchainMessage) error {
	validatorLogger.Debug("Handling OpenchainMessage of type: %s ", msg.Type)
	if v.FSM.Cannot(msg.Type.String()) {
		return fmt.Errorf("Validator FSM cannot handle message (%s) with payload size (%d) while in state: %s", msg.Type.String(), len(msg.Payload), v.FSM.Current())
	}
	err := v.FSM.Event(msg.Type.String(), msg)
	if err != nil {
		if _, ok := err.(*fsm.NoTransitionError); !ok {
			// Only allow NoTransitionError's, all others are considered true error.
			return fmt.Errorf("Peer FSM failed while handling message (%s): current state: %s, error: %s", msg.Type.String(), v.FSM.Current(), err)
			//t.Error("expected only 'NoTransitionError'")
		}
	}
	return nil
}

func (v *ValidatorFSM) SendMessage(msg *pb.OpenchainMessage) error {
	validatorLogger.Debug("Sending message to stream of type: %s ", msg.Type)
	err := v.ChatStream.Send(msg)
	if err != nil {
		return fmt.Errorf("Error Sending message through ChatStream: %s", err)
	}
	return nil
}
