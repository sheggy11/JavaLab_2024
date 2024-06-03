package ru.itis.webflux.clients;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Component;
import org.springframework.web.reactive.function.client.WebClient;
import reactor.core.publisher.Flux;
import ru.itis.webflux.entities.User;

import java.util.Arrays;

@Component
public class UserClient implements Client {
    private final WebClient client;

    @Override
    public Flux<User> getUsers(){
        return client.get()
                .accept(MediaType.APPLICATION_JSON)
                .exchangeToFlux(clientResponse -> clientResponse.bodyToFlux(User[].class))
                .flatMapIterable(Arrays::asList);
    }

    public UserClient(@Value("${user.api.url}") String url) {
        client = WebClient.builder()
                .baseUrl(url)
                .build();
    }

}
