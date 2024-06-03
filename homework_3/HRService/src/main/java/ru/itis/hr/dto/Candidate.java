package ru.itis.hr.dto;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class Candidate {
    public String name;
    public int age;
    public int experience;
    public List<String> skills;
    public String experienceDescription;
}
