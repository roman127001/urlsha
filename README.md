# URL Shortener (straightforward version)


#### Description

A simple URL shortener service written in Go. Future improvements are planned to enhance its 
functionality and robustness.

Encode use random string of length 6 to encode the origin url to short url.

Length of the random string can be changed by changing the value of `ShortUrlLength` in `config.go` file.

[![asciicast](https://asciinema.org/a/6NBaGoAcchgyzEU87bwb5kOAm.svg)](https://asciinema.org/a/6NBaGoAcchgyzEU87bwb5kOAm)

### **Questions**

1. Describe your solution. What tradeoffs did you make while designing it, and why? 
2. If this were a real project, how would you improve it further? 
3. What is the math & logic behind “short url” if there is any?


> 1. Describe your solution. What tradeoffs did you make while designing it, and why? 

1. Solution Description and Tradeoffs:
    - Solution Overview: The solution comprises a simple URL shortener service written in Go. 
    It uses a straightforward approach to generate short aliases for long URLs and provides functionality 
    to map those aliases back to the original URLs. Encode use random string of length 6 
    to encode the origin url to short url.
      
    - Design Tradeoffs: Several design decisions were made while developing this solution:
      - Random String Generation: To generate short URLs, the system utilizes a random string of a fixed length (6 characters by default). This approach provides a balance between shortness and uniqueness of URLs. However, shorter strings might lead to increased collision probability, while longer strings could result in less user-friendly URLs.
        - in case with md5/sha256 hash, it will be more unique, but it will be longer and not user-friendly.
      - Use pregenerated random strings: we can use pregenerated random string and store it in database and response faster.
      - In-memory Data Store: The system currently employs an in-memory data store to map short URLs to their corresponding long URLs. 
      While this approach ensures simplicity and fast lookup times, it might not be suitable for high-traffic or distributed 
      environments due to scalability limitations and potential data loss on server restarts.
      - HTTP Server Implementation: The HTTP server is implemented using the standard Go net/http package. 
      While this provides a lightweight and efficient solution, it lacks features such as built-in middleware 
      for authentication, logging, and error handling, which are essential for production-grade applications.
      - Error Handling: The system includes basic error handling to manage invalid requests and edge cases. 
      However, more robust error handling mechanisms, such as middleware or custom error responses, 
      could enhance the service's reliability and user experience. 
      Logging throw by UDP to avoid network tradeoff (as another way to logging)?
      - In this application I do not use expiration time for short url, but it can be added in future improvements.
      - Use gRPC for communication between services, to avoid network tradeoff (for big data/payload).
      - Also, We must add health check and version endpoint to check the health of the service.

> 2. If this were a real project, how would you improve it further?

2. Potential Improvements for a Real Project:
   - **Authentication**: Implement authentication mechanisms to ensure that only authorized users can encode 
   or decode URLs. This could involve integrating with OAuth providers like Google or implementing a custom 
   authentication service.
   - **Middleware**: Incorporate logging and error handling middleware to improve observability and reliability. 
   This includes logging request and response details for analysis and handling errors gracefully to prevent 
   service disruptions.
   - **Data Storage**: Replace the in-memory data store with a more robust and scalable solution 
   like Redis or a relational database (e.g., PostgreSQL). This would enable persistence of URL mappings 
   across server restarts and support for distributed deployments.
   - **Expiration URLs**: Introduce functionality to automatically expire short URLs after a certain 
   period to enhance security and prevent outdated links. This could involve implementing a background 
   task scheduler to periodically clean up expired entries in the data store.
   - **Enhanced Metadata**: Allow users to associate additional metadata (e.g., author information, description) 
   with each URL to provide context and improve organization.
   - **Code Quality Tools**: Integrate linters and code quality tools into the project's development workflow 
   to enforce coding standards and improve maintainability.
   - **Continuous Integration/Deployment**: Set up CI/CD pipelines using tools
   like GitHub Actions or GitLab CI/CD to automate testing, builds, and deployments. 
   This ensures consistent and reliable delivery of updates to production environments.
   - **ENVs tools and credentials**: Environment variables for configuration, to avoid hardcoding of configuration. 
   Manage secrets and credentials securely using tools like AWS Secrets Manager or HashiCorp Vault.
   - **Documentation**: Enhance the documentation with detailed guides, API references, and examples to use and modify the service effectively. 
   Include service exploitation, maintenance and troubleshooting guides!


> 3. What is the math & logic behind “short url” if there is any?

3. Math & Logic Behind Short URLs:
- The short URLs are generated using a random string of characters chosen from a predefined set 
(uppercase letters, lowercase letters, and digits). The length of the string determines the total 
number of possible combinations, which in turn affects the likelihood of collisions 
(two different long URLs mapping to the same short URL).

- In this implementation, a random string of length 6 is used by default. 
With a character set of 62 (26 lowercase letters + 26 uppercase letters + 10 digits), 
there are 62^6 possible combinations, resulting in a large address space to accommodate many unique short URLs. 

- While the approach of using random strings provides simplicity and avoids sequential patterns, 
more sophisticated techniques such as hashing algorithms (e.g., MD5, SHA) 
or base conversion methods can also be employed to generate short URLs with certain properties 
(e.g., cryptographic security, deterministic mapping).



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


#### Makefile commands

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
