package scheduling

import (
	"go.7yes.com/j7f/proto/scheduling"
	"google.golang.org/grpc"
)

func Init(g *grpc.Server) {
	s := &PersonalCourseScheduling{}
	scheduling.RegisterPersonalCourseScheduleServerServer(g, s)
}
