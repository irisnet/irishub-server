package vo

import "github.com/irisnet/irishub-server/models/document"

type CandidateListReqVO struct {
	Address              string   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Page                 uint64   `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
	PerPage              uint64   `protobuf:"varint,3,opt,name=perPage" json:"perPage,omitempty"`
	Sort                 string   `protobuf:"bytes,4,opt,name=sort" json:"sort,omitempty"`
	Q                    string   `protobuf:"bytes,5,opt,name=q" json:"q,omitempty"`
}

type CandidateListResVO struct {
	Candidates []document.Candidate
}

