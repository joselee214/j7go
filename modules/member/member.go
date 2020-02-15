package member

import (
	"github.com/joselee214/j7f/components/errors"
	"j7go/proto/member"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	memberService "j7go/services/member"
	"j7go/utils"
)

func Init(g *grpc.Server) {
	s := &memberServer{}
	member.RegisterMemberServerServer(g, s)
}

type memberServer struct{}

//新增会员
func (s *memberServer) AddMember(server member.MemberServer_AddMemberServer) error {
	for {
		request, err := server.Recv()
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("recv", zap.String("grpc", err.Error()))
			return err
		}

		memberId, err := memberService.AddMember(server.Context(), request.MemberInfo)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("member", zap.String("add_member", err.Error()))
		}

		imagesResponse := &member.AddMemberResponse{
			Status:   errors.GetResHeader(err),
			MemberId: memberId,
		}

		err = server.Send(imagesResponse)
		if err != nil {
			utils.GetTraceLog(server.Context()).Error("send", zap.String("grpc", err.Error()))
			return err
		}
	}
}

func (s *memberServer) EditMember(server member.MemberServer_EditMemberServer) error {
	var err error
	return err
}
