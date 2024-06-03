package ru.itis.users.entities;

import lombok.AllArgsConstructor;
import lombok.Builder;

@AllArgsConstructor
@Builder
public class User {
    private String username;
    private String password;
    private String orientation;
}
