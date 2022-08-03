package service

import (
	"io"

	"jw.lib/logx"

	pb "jw_base/proto/generated_go"
)

func (s *BaseService) StreamTest(srv pb.BaseService_StreamTestServer) error {
	for {
		//从流中获取消息
		res, err := srv.Recv()
		if err == io.EOF {
			//发送结果，并关闭
			return srv.SendAndClose(&pb.StreamTestResp{Result: "ok"})
		}
		if err != nil {
			return err
		}
		logx.Infof("-->", res.Data)
	}
}
