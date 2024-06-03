FROM maven:latest as build
WORKDIR /workspace/app

COPY pom-docker.xml pom.xml
COPY src src

RUN --mount=type=cache,target=/root/.m2 mvn -DskipTests=true clean package
RUN mkdir -p target/dependency && (cd target/dependency; jar -xf ../*.jar)

FROM eclipse-temurin:latest
ARG DEPENDENCY=/workspace/app/target/dependency

COPY --from=build ${DEPENDENCY}/BOOT-INF/lib /app/lib
COPY --from=build ${DEPENDENCY}/META-INF  /app/META-INF
COPY --from=build ${DEPENDENCY}/BOOT-INF/classes /app

EXPOSE 8080

ENTRYPOINT ["java", "-cp", "app:app/lib/*", "ru.itis.hr.HrServiceApplication"]