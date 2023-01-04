package main

type RequestMsg struct {
	TxID       [32]byte `json:"TxID"`
	TimeStamp  []byte   `json:"TimeStamp"`
	SequenceID int64    `json:"sequenceID"` //시퀀스아이디
}
type ResultMsg struct {
	Txid [32]byte `json:"txid"`
}

type ReplyMsg struct {
	ViewID    int64    `json:"viewID"`
	Timestamp []byte   `json:"timestamp"`
	ClientID  [32]byte `json:"clientID"`
	NodeID    string   `json:"nodeID"`
	Result    string   `json:"result"`
}

type PrePrepareMsg struct {
	ViewID     int64       `json:"viewID"`
	SequenceID int64       `json:"sequenceID"`
	Digest     string      `json:"digest"`
	RequestMsg *RequestMsg `json:"requestMsg"`
}

type VoteMsg struct {
	ViewID     int64  `json:"viewID"`
	SequenceID int64  `json:"sequenceID"`
	Digest     string `json:"digest"`
	NodeID     string `json:"nodeID"`
	MsgType    `json:"msgType"`
}

type MsgType int

const (
	PrepareMsg MsgType = iota
	CommitMsg
)
