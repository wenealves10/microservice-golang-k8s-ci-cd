# Domain-Driven Design in Golang

### In this article we have covered the basics of Domain-Driven Design, in short.

Entities — Mutable Identifiable Structs.
Value Objects — Immutable Unidentifiable Structs.
Aggregates — Combined set of Entities and Value objects, stored in Repositories.
Repository— A implementation of storing aggregates or other information
Factory— A constructor to create complex objects and make creating new instance easier for the developers of other domains
Service — A collection of repositories and sub-services that builds together the business flow
