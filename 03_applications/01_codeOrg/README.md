# Code organization sample project

This project defines a basic MVC based template code layout following
common best practices for code reuse across multiple packages. 

### It addresses the following assumptions:
1. It is a web based project 
2. It will work with multiple database systems for caching, storage, analytics
3. It is written using a REST API structure that needs versioning and backwards compatability
4. It will be written as a single codebase for ease of maintenance by a single developer
5. The frontend will be written independently by another contractor, so we can focus on data
engineering tasks and workflows. The REST API will act as our binding contract
6. Error handling and logging will be a main source for future maintainability as we are planning
to deliver this project to a customer for ongoing usage. We might be called in for bug fixes, and 
ongoing development of new features. 

### It will use the following architecture:
1. We are running a base server from the net/http package
2. We are serving a versioned rest API that is backwards compatible
3. We are extracting all handlers as controller files into their own module
4. We are defining our data structures in their own module and provide a
creator method to manage initialization of the values according to their types, without
overloading the functional API to the individual datatypes
5. We are following google practices passing context to each handler, accepting timeout
settings from the user as max-timeout for our long-running processes
6. We are implementing a management API to the system that allow us to debug the state
of the running system in production.


### Functional requirements
- We need to run a suite of automated testcases for both integration with
external datasources, as well as databases, as well as the analytical micro-service interface
- Automatic regression tests need to run on each package as it is updated

#### Technical Requirements
- Analytical processes that have long-running state will need to be designed for fail safety which
includes continuous backup and backup management of their state with automatic recovery from failure
 
#### Extensibility
- External data sources will change their API and need continuous maintenance
- Data contracts between our system and external data sources will need to be easily updatable
- 

