# Go Clean Arch

## Description

This monorepo implements Clean Architecture in Go (Golang).

Rule of Clean Architecture by Uncle Bob
* Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows us to use such frameworks as tools, rather than having to cram our system into their limited constraints.
* Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
* Independent of Database. We can swap out Mongodb, for Rocksdb, Dynamodb, CouchDB, or something else. Our business rules are not bound to the database.

More [here](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)

This repo has 4 layers
* Domain (Models, entities)
* Repository
* Usecase
* Delivery (controller/delivery)

<img src="https://github.com/bxcodec/go-clean-arch/raw/master/clean-arch.png">

The original explanation about this repos structure can be read from this medium's post : [https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047](https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047).

It may be different already, but the concept is still the same in application level