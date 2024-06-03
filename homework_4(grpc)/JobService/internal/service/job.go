package service

import (
	"context"
	"job-service/internal/core"
	"job-service/proto"
)

type JobRepository interface {
	GetAllJob(ctx context.Context, page int64) ([]*core.Job, error)
}

type JobService struct {
	proto.JobServiceServer
	jobRepository JobRepository
}

func NewJobService(jobRepository JobRepository) *JobService {
	return &JobService{
		jobRepository: jobRepository,
	}
}

func (service *JobService) GetAllJob(ctx context.Context, request *proto.JobRequest) (*proto.JobResponse, error) {
	jobs, err := service.jobRepository.GetAllJob(ctx, request.Page)

	if err != nil {
		return nil, err
	}

	res := convert(jobs)

	return &proto.JobResponse{Jobs: res}, nil
}

func convert(jobs []*core.Job) []*proto.Job {
	res := make([]*proto.Job, 0)
	for i := 0; i < len(jobs); i++ {
		job := jobs[i]
		res = append(res, &proto.Job{
			Revenue: uint32(job.Revenue),
			Skills:  job.Skills,
		})
	}
	return res
}
