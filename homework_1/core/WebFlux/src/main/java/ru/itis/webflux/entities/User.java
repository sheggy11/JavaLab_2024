package ru.itis.webflux.entities;

import lombok.AllArgsConstructor;
import lombok.Builder;

@AllArgsConstructor
@Builder
public class User {
    private String username;
    private String password;
    private String orientation;
}