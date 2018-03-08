package service

import (
	"github.com/laidingqing/KeepJobs/accounts/model"
	pb "github.com/laidingqing/KeepJobs/pb"
)

func encodeAccountInfo(request *pb.Account) model.Account {
	return model.Account{
		AccountID: request.Id,
		UserName:  request.Username,
	}
}

func dencodeAccountInfo(request model.Account) *pb.Account {
	return &pb.Account{
		Id:       request.ID.Hex(),
		Username: request.UserName,
	}
}
