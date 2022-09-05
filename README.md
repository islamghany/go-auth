# go-auth

an overview of different techniques to implement authentication and authorization in go web application

### Authentication is about confirming who a user is, whereas Authorization is about checking whether that user is premitted to something.

## Authentication

#### different approaches to API Authentication:

- Basic authentication
- Stateful token authentication
- Stateless token authentication
- API key authentication
- OAuth2.0 / OpenID Client

---

## Basic Authentication

in this approach, the client include an Authorization header with every request containing their credentials, in the format username:password and base-64 encoded.
for example to authenticatde `smith@gmail.com:pa55word`

```
    Authorization: Basic c21pdGhAZ21haWwuY29tOnBhNTV3b3Jk
```

then you can access the credentials from this header using Go's `Request.BasicAuth()` and verify that they are correct before continue

### pros

- Simple for clients.
- Is supported out-of-the-box by most programming languages, web browsers, and tools such as curl and wget.
- It’s often useful in the scenario where your API doesn’t have ‘real’ user accounts.

### cons

- It sends the credentials encoded but not encrypted. This would be completely insecure unless the exchange was over a secure connection (HTTPS/TLS).
- For APIs with ‘real’ user accounts and — in particular — hashed passwords, it’s not such a great fit. Comparing the password provided by a client against a (slow) hashed password is a deliberately costly operation, and when using HTTP basic authentication you need to do that check for every request. That will create a lot of extra work for your API server and add significant latency to responses.

---

## Token Authentication (AKA bearer token authentication)

1. The client sends a request for the server with his credentials(username or email, password) _the login step_.

2. then the server verify that the credentials are correct then generates a _bearer token_ which
   represents the user, and sends it back to the user, the token expires after specified time, after that the user will resubmit his credentials to get a new token.
3. for subsequent requests to the server the client will include the token in the Authorization header like this

```
    Authorization: Bearer <token>
```

4. when the server receive the request it examines the token, checks if it hasn't expired, the token value determine who the user is.

### pros

- for APIs where user's pasword are hashed, this approach is better than the basic authentication, because the in the basic auth the hashed password must be evaluated in every request that causes leak in the performance.

### cons

- the downside is that managing the token can be complicated for the clients, they will need to implement the necessary logic for token caching, momonitoring and managing token expiry, and periodically generating new tokens

<i>
    <strong>we can break down token authentication further into two sub-type: Statful and Statless token authentication.</strong>
</i>
<br />

## Stateful Token Authentication

In stateful token auhtencation approach, the value of the token is high-entorpy cryptographically-secure random string, this token _or fast hash of it_ is stored server-side in a database alongside with userID and and an expiry time for the token.

When a user sends the token with the subsequent requests the server lookup in database for that token, check that it hasn't expired, and retrieve the corresponding userID to find out who is the request coming from.

### pros

- API maintanes control ovwe the token, it very trival to get the token from the database and deleting them or make mark then expired.

- it also simple and robust, the security is provided by the token being _unguessable_ which is why it's important to use high-entopy cryptographiclly-secure random string.

### cons

- it will need to a databse lookup, a high number could be database lookup negative, but in muse cases you will need to databse lookup to check the user's activation status or retrieve additional information.
