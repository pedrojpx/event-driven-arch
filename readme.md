# Walletcore & Balancesapp

This simple project consists of a core application, called `Walletcore`, capable of creating clients and accounts, and perform transactions between accounts throuh its Rest API.

Every transaction by Walletcore produces two events: one indicating the transaction that was created and one that states the new balances for the accounts involved.

`Balancesapp` consumes the event stating the current balances, updates its own database and provides an endpoint for requesting this information.

`kafka` is used as the event bus and each application has is own database and could be deployed as separate microservices.

___
### This project exemplifies the usage of three main concepts:

- Hexagonal Architecture
- Event Oriented Architecture using kafka as event bus
- Containerization of an application involving multiple services such that both applications, their auxiliary services and their configuration can be deployed with a single `docker compose up -d`
