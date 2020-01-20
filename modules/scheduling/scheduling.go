package scheduling

import (
	"go.7yes.com/go/proto/scheduling"
	"google.golang.org/grpc"
)

func Init(g *grpc.Server) {
	s := &PersonalCourseScheduling{}
	scheduling.RegisterPersonalCourseScheduleServerServer(g, s)
}
