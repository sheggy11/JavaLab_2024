package service

import (
    "context"
    "fmt"
    "job-service/internal/core"
    "job-service/proto"
)

type JobRepository interface {
    GetAll(ctx context.Context) ([]*core.Job, error)
}

type JobService struct {
    proto.JobServiceServer
    jobRepository JobRepository
}

func NewJobService(jobRepository JobRepository) *JobService {
    return &JobService{jobRepository: jobRepository}
}

func (s *JobService) GetAll(ctx context.Context, _ *proto.GetAllRequest) (*proto.GetAllResponse, error) {
    jobs, err := s.jobRepository.GetAll(ctx)

    if err != nil {
        return nil, fmt.Errorf("get all jobs in service: %w", err)
    }

    if jobs == nil {
        jobs = make([]*core.Job, 0)
    }

    jobResponses := make([]*proto.JobResponse, len(jobs))
    for i, job := range jobs {
        jobResponses[i] = &proto.JobResponse{
            Name:        job.Name,
            Experience:  job.Experience,
            Skills:      job.Skills,
            Description: job.Description,
            Company:     job.Company,
        }
    }

    return &proto.GetAllResponse{
        Jobs: jobResponses,
    }, nil
}
