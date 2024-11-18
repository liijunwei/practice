### overview

practice project from "Let's Go Further"

and after it's done, I plan to make this project a "helper" api backend for many other purposes

### notes

- router example: https://github.com/benhoyt/go-routing/blob/master/stdlib/route.go
- api authentication approaches
    - basic authentication
    - stateful token authentication
        - server manage, can revoke from backend
    - stateless token authentication
        - client side encode, e.g. JWT token
        - downside: can't easily be revoked once issued
    - api key authentication
    - OAuth 2.0/ OpenID Connect
- **it's important to note that API keys themselves should only ever be communicated to users over a secure channel**, and you should treat them with the same level of care that you would a user's password

### TODOs

- TODO practice cursor based pagination
- TODO try another email package
    - less dependency
    - https://github.com/wneessen/go-mail/wiki/Simple-Mailer-Example
- TODO use https://github.com/google/wire
- TODO visualizing and analyzing metrics(maybe telegraf+influxdb, and probably uber jaeger)
