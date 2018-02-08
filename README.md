# authentication-service
This service handles authenticating users that login or sign up to the application

docker build -t authentication-service .

docker run --publish 8080:8080 --name test --rm authentication-service
