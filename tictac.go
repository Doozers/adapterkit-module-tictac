package tictac

import (
	"fmt"
	"time"
)

type service struct {
	UnimplementedTictacSvcServer
}

func New() TictacSvcServer {
	return &service{}
}

func (s *service) Countdown(req *CountdownReq, serv TictacSvc_CountdownServer) error {
	for i := req.Count; i > 0; i-- {
		err := serv.Send(&CountdownRes{Msg: fmt.Sprintf("%d seconds left", i)})
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}

	if req.Msg == "" {
		err := serv.Send(&CountdownRes{Msg: "Done!"})
		if err != nil {
			return err
		}
		return nil
	}

	err := serv.Send(&CountdownRes{Msg: req.Msg})
	if err != nil {
		return err
	}
	return nil
}
