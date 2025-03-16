package service

import (
	"context"
	"net/url"
	"strings"
	"sync/atomic"

	pb "video-balancer/proto"
)

type BalancerService struct {
	pb.UnimplementedServiceServer
	cdnHost string
	counter uint64
}

func NewBalancerService(cdnHost string) *BalancerService {
	return &BalancerService{
		cdnHost: cdnHost,
	}
}

func (s *BalancerService) Method(ctx context.Context, req *pb.VideoRequest) (*pb.VideoResponse, error) {
	count := atomic.AddUint64(&s.counter, 1)

	if count%10 == 0 {
		return &pb.VideoResponse{
			RedirectUrl: req.Video,
		}, nil
	}

	parsedUrl, err := url.Parse(req.Video)
	if err != nil {
		return nil, err
	}

	hostParts := strings.Split(parsedUrl.Host, ".")
	if len(hostParts) == 0 {
		return nil, err
	}
	originServer := hostParts[0]

	newUrl := &url.URL{
		Scheme:   parsedUrl.Scheme,
		Host:     s.cdnHost,
		Path:     "/" + originServer + parsedUrl.Path,
		RawQuery: parsedUrl.RawQuery,
		Fragment: parsedUrl.Fragment,
	}

	return &pb.VideoResponse{
		RedirectUrl: newUrl.String(),
	}, nil
}
