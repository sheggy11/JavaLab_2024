package ru.itis.hr.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import ru.itis.hr.dto.HRResponse;
import ru.itis.hr.service.HRService;

@RestController
@RequiredArgsConstructor
public class HrController {

    private final HRService hrService;

    @GetMapping("/candidates")
    public ResponseEntity<HRResponse> getCandidates() {
        return ResponseEntity.ok(
                hrService.getCandidates());
    }

}
