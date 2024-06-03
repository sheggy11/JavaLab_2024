package service

import (
	"context"
	"cv-service/internal/core"
	"cv-service/proto"
)

type CVRepository interface {
	GetBySkills(ctx context.Context, skills []string) ([]*core.CV, error)
}

type CVService struct {
	proto.CVServiceServer
	cvRepository CVRepository
}

func NewCVService(cvRepository CVRepository) *CVService {
	return &CVService{
		cvRepository: cvRepository,
	}
}

func (service *CVService) GetAllCV(ctx context.Context, request *proto.CVRequest) (response *proto.CVResponse, err error) {
	cvs, err := service.cvRepository.GetBySkills(ctx, request.Skills)

	if err != nil {
		return nil, err
	}

	res := convert(cvs)

	return &proto.CVResponse{Cvs: res}, nil
}

func convert(cvs []*core.CV) []*proto.CV {
	res := make([]*proto.CV, 0)
	for i := 0; i < len(cvs); i++ {
		cv := cvs[i]
		res = append(res, &proto.CV{
			Name:   cv.CandidateName,
			Age:    uint32(cv.CandidateAge),
			Skills: cv.Skills,
		})
	}
	return res
}
