package service

import (
	"CVService/internal/core"
	"CVService/proto"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CVRepository interface {
	GetByActivityType(ctx context.Context, activityType string) (*core.CV, error)
	GetAllCV(ctx context.Context) ([]*core.CV, error)
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

func (service *CVService) GetCV(ctx context.Context, request *proto.CVRequest) (response *proto.CVResponse, err error) {
	cv, err := service.cvRepository.GetByActivityType(ctx, request.GetActivityType())

	if err != nil {
		return nil, err
	}

	if cv == nil {
		cv = &core.CV{
			ID:           primitive.ObjectID{},
			Name:         "",
			ActivityType: "",
			Summary:      "",
		}
	}

	return &proto.CVResponse{Name: cv.Name, ActivityType: cv.ActivityType, Summary: cv.Summary}, nil
}

func (service *CVService) GetAll(ctx context.Context, request *proto.EmptyRequest) (response *proto.GetAllResponse, err error) {
	cvs, err := service.cvRepository.GetAllCV(ctx)

	if err != nil {
		return nil, err
	}

	if cvs == nil {
		var emptyResponse []*proto.CVResponse

		return &proto.GetAllResponse{Cvs: emptyResponse}, nil
	}

	cvResponseList := make([]*proto.CVResponse, len(cvs))
	for i, cv := range cvs {
		cvResponseList[i] = mapCVToCVResponse(cv)
	}

	return &proto.GetAllResponse{Cvs: cvResponseList}, nil
}

func mapCVToCVResponse(cv *core.CV) *proto.CVResponse {
	return &proto.CVResponse{
		Name:         cv.Name,
		ActivityType: cv.ActivityType,
		Summary:      cv.Summary,
	}
}
