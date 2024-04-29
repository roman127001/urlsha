# URL Shortener (straightforward version)


#### Description

A simple URL shortener service written in Go. Future improvements are planned to enhance its 
functionality and robustness.

Encode use random string of length 6 to encode the origin url to short url.

Length of the random string can be changed by changing the value of `ShortUrlLength` in `config.go` file.


#### TODO List (Future improvements)


- **Authentication**: Consider adding authentication to secure access to the URL shortener service, 
ensuring that only authorized users can encode or decode URLs.

- **Logging Middleware**: Implement logging middleware to track and analyze requests, providing insights 
into the service's usage patterns and potential issues.

- **Error Handling Middleware**: Add error handling middleware to gracefully manage errors and exceptions, 
improving the service's reliability and user experience.

- **Expiration URLs**: Explore the implementation of expiration URLs to automatically expire short URLs after 
a specified period, enhancing security and preventing outdated links.

- **Author Metadata**: Integrate the capability to associate an author with each URL, providing additional 
context and attribution for shared links.

- **Description Metadata**: Enable the inclusion of a description for each URL, allowing users to provide 
brief descriptions or summaries for better organization and understanding.

- **External Storage**: Investigate using external storage solutions like Redis to store and manage URL mappings, 
potentially improving scalability and performance.

- **User Identifier Prefix**: Evaluate the option of using user names or IDs as prefixes for short URLs, 
facilitating URL management and tracking for individual users.

- **Add linter**: Add linter to the project to enforce code quality and consistency, 
ensuring adherence to Project's and Team's standards.

- **CI/CD**: Implement continuous integration and continuous deployment pipelines to automate testing. 
Think about using GitHub Actions or GitLab CI/CD. 
In perfect case building process must be completed in specific machine to avoid affecting local or production machine and environment.

- **Load Testing**: Conduct load testing to evaluate the service's performance under various traffic conditions.

These enhancements aim to fortify the service's security, reliability, and usability, ensuring its 
effectiveness in URL management and redirection.



### Endpoints


To encode origin (long) url to short url, send a POST request to `/encode` endpoint with the following JSON payload:
```json
{
  "url": "https://www.google.com/"
}
```

To decode short url to origin url, send a GET request to `/decode/{short_url}` endpoint.


## Example of usage

### 1. Run the server
```shell
make run
```

#### 2. Encode a URL
```shell
$ curl -s -d '{"url":"https://www.google.com/"}' -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/encode | jq                                              (492ms) 2024-04-27 14:43:31 +0100
{
  "short_url": "http://127.0.0.1:8080/decode/lT3rYV",
  "error": ""
}
```

#### 3. Decode a URL
```shell

$ curl -s "http://127.0.0.1:8080/decode/lT3rYV" | jq                                                                                                                                2024-04-27 14:44:06 +0100
{
  "origin_url": "https://www.google.com/",
  "error": ""
}
```


Makefile commands

Help - is default command for make (run: `make help` or `make`)
```shell
$ make help
Usage:
  make <target>

  help                  Available commands

Local
  run                   Run the application
  test                  Run the tests

Docker
  run-docker            Run the application in a docker container
  stop-docker           Stop the application in a docker container

```
