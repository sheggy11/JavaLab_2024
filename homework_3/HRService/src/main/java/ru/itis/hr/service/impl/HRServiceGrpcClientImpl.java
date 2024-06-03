package ru.itis.hr.service.impl;

import lombok.RequiredArgsConstructor;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.stereotype.Service;
import ru.itis.hr.dto.Candidate;
import ru.itis.hr.dto.HRResponse;
import ru.itis.hr.dto.Pair;
import ru.itis.hr.grpc.pb.cv.CVServiceGrpc;
import ru.itis.hr.grpc.pb.cv.Cv;
import ru.itis.hr.grpc.pb.job.JobOuterClass;
import ru.itis.hr.grpc.pb.job.JobServiceGrpc;
import ru.itis.hr.service.HRService;

import java.util.ArrayList;
import java.util.List;

@Service
@RequiredArgsConstructor
public class HRServiceGrpcClientImpl implements HRService {

    private final JobServiceGrpc.JobServiceBlockingStub jobService;

    private final CVServiceGrpc.CVServiceBlockingStub cvService;

    @Override
    public HRResponse getCandidates() {
        HRResponse response = new HRResponse();
        response.pairs = new ArrayList<>();

        var jobs = jobService.getJob(JobOuterClass.JobRequest.newBuilder()
                        .setPage(1)
                .build());

        for (int i = 0; i < jobs.getJobsCount(); i++) {
            List<Cv.CV> cvs = cvService.getCV(
                    Cv.CVRequest.newBuilder().
                    addAllSkills(jobs.getJobs(i).getSkillsList())
                    .build()
            ).getCvsList();

            response.pairs.add(
                    new Pair(jobs.getJobs(i).getJobDescription(),
                    convert(cvs))
            );
        }

        return response;
    }

    private List<Candidate> convert(List<Cv.CV> from) {
        List<Candidate> result = new ArrayList<>();

        for (Cv.CV cv : from) {
            result.add(new Candidate(cv.getCandidateName(),
                    cv.getCandidateAge(),
                    cv.getExperience(),
                    cv.getSkillsList(),
                    cv.getExperienceDescription()));
        }

        return result;
    }
}
