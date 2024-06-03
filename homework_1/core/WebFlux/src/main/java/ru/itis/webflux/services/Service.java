package ru.itis.webflux.services;

import reactor.core.publisher.Flux;
import ru.itis.webflux.entities.User;

public interface Service {
    Flux<User> getUsers();
}
