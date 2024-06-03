package ru.itis.webflux.services;

import lombok.AllArgsConstructor;
import org.springframework.stereotype.Component;
import reactor.core.publisher.Flux;
import reactor.core.scheduler.Schedulers;
import ru.itis.webflux.clients.Client;
import ru.itis.webflux.entities.User;

import java.util.List;

@Component
@AllArgsConstructor
public class UserServiceImpl implements Service {

    private final List<Client> clients;

    @Override
    public Flux<User> getUsers() {
        List<Flux<User>> fluxes = clients.stream().map(this::getBooks).toList();
        return Flux.merge((fluxes));
    }

    private Flux<User> getBooks(Client client) {
        return client.getUsers()
                .subscribeOn(Schedulers.boundedElastic());
    }
}
