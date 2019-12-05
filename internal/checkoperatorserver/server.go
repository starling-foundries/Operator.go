package operator

import (
	"context"

	pb "github.com/starling-foundries/Operator.go/rpc/checkoperator"

	"github.com/twitchtv/twirp"
)

// Server implements the Operator service
type Server struct{}

//filler consts to simulate a basic transaction
const LARRY_NONCE uint64 = 0
const LARRY_ADDR = ""

// sends the signed transaction from the client to the operator.
func (s *Server) SendCheck(ctx context.Context, check *pb.CheckTransaction) (*pb.QueueID, error) {

	//Check the sender is authorized
	if check.Senderpubkey.String() != LARRY_ADDR {
		return nil, twirp.InvalidArgumentError("ID:", "ID provided is not larry")
	}
	//Check the nonce is valid
	if check.Nonce != LARRY_NONCE {
		return nil, twirp.InvalidArgumentError("Nonce:", "Nonce is incorrect")
	}
	//validate transaction content
	if ValidateCheck(check) {
		return nil, twirp.InvalidArgumentError("Check:", "Check is invalid")
	}

	//attempt to execute transaction on-chain.

}

// retrieves the transaction information from the perspective of the operator
func (s *Server) Status(ctx context.Context, reciept *pb.QueueID) (*pb.StatusResp, error) {
	//Check the sender is authorized
	if !(reciept.QueueID > 0) {
		return nil, twirp.InvalidArgumentError("ID:", "ID must be provided")
	}

	//TODO:lookup the reciept in the database
	// var resp StatusResp

	//do more processing
	return nil, nil
}

// Gets the status of the last transaction sent by this user to this operator
func (s *Server) GetLastTransaction(context.Context, *pb.Empty) (*pb.StatusResp, error) {
	return nil, nil
}

func (s *Server) OperatorTarget(context.Context, *pb.Empty) (*pb.OperatorTargetResp, error) {
	return nil, nil
}

func (s *Server) GetMinFee(context.Context, *pb.Empty) (*pb.FeeResp, error) { return nil, nil }

func ValidateCheck(check *pb.CheckTransaction) bool {
	return true
}
