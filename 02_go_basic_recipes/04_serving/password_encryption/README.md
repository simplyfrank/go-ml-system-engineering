# Authentication with encrypted passwords

This represents a simple implementation of a custom
authentication system you can play around with. 

Functional reqs:
* User can signup with an account
* User is kept in a session after first signup
* A logout will clean the cookie for the user
* A login will allow a registered user to easily get a new session assigned
* An existing user is not allowed to register another account

### Technical implementation:
For this example i focussed on a simple layout to keep the focus on the interaction of 
the individual components along a classical authentication workflow.

1. Users are stored in a in-memory object "dbUsers" for convenience
2. Sessions are stored in an in-memory map "dbSessions" for convenience
3. Templates are stored in /templates and are extracted on init(){}
4. Handlers are simple:
    1. baseHandler: Presents session and user data when set
    2. signupHandler: Manages 'GET' and 'POST' responses with a switch for easy coding.
    3. logoutHandler: Manages logout and redirects back to "/" on success
    4. loginHandler: ...(not yet implemented)
5. All functionality is removed from the viewHandlers and moved into
separate helper functions for convenience
6. Session IDs are genereted using the google/uuid package. Here the parsing to string
and back ensures a hightened security layer, and the implemented mutex lock will allow 
the system to be unique even in concurrent designs and frequent requests.
7. 