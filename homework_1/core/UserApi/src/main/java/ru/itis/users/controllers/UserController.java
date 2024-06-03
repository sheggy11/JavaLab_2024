package ru.itis.users.controllers;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import ru.itis.users.entities.User;

import java.util.List;

@RestController
@RequestMapping("/users")
public class UserController {

    @GetMapping()
    public List<User> getAllBooks() {

        User user1 = User.builder()
                .username("Asgat_Sheggy")
                .password("123456")
                .direction("east")
                .build();

        User user2 = User.builder()
                .username("Ivan_Ivanov")
                .password("237044")
                .direction("west")
                .build();

        User user3 = User.builder()
                .username("Denis_Biserov")
                .password("4389n33r")
                .direction("north")
                .build();

        User user4 = User.builder()
                .username("Danis_Zaripov")
                .password("jrnov7n")
                .direction("south")
                .build();

        User user5 = User.builder()
                .username("Rus_Minkh")
                .password("fhno430")
                .direction("north")
                .build();

        return  List.of(user1, user2, user3, user4, user5);
    }
}
