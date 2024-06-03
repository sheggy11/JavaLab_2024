package ru.itis.webflux.clients;

import reactor.core.publisher.Flux;
import ru.itis.webflux.entities.User;

public interface Client {
    Flux<User> getUsers();
}
