## Project structure

```
├── LICENSE
├── README.md
├── config.go
├── go.mod
├── go.sum
├── clientlib
│   ├── lib.go
│   └── lib_test.go
├── cmd // ->
│   ├── _you_app_client
│   │   └── main.go
│   └── _you_app_server
│       └── main.go
├── internal
│   └── auth
│       ├── auth.go
│       └── auth_test.go
└── serverlib
    └── lib.go
|
|
|
|
|
|


```

# /cmd

- Main applications for this project
- Link đường dẫn cho từng file chính nên đúng với mẫu sau (e.g., /cmd/myapp).
- Không nên cho nhiều code trong file application. Nếu như đoạn mã sẽ được tái sử dụng cho nhiều project khác thì nên để code ở file /pkg. Ngược lại nếu bạn không muốn đoạn mã đó bị người khác sử dụng thì hãy để nó ở thư mục /internal.

# /internal

- This is private application and library code.
- Đây sẽ là nơi chứa những đoạn mã mà bạn không muốn bất kỳ ai truy cập. Lưu ý rằng bạn không bị giới hạn thư mục internal ở cấp cao, bạn có thể tạo nhiều hơn 1 thư mục internal.
- Những package nằm bên trong thư mục chỉ có thể được truy cập mới các code trong cũng cấp hoặc nhỏ hơn trong thư mục. Ví dụ

```
/myproject
  /internal
    /mypackage
  /app

```

Với cây thư mục như trên mypackage nằm bên trong thư mục internal. Package /mypackage sẽ chỉ import được trong phạm vi các thư mục nằm trong /myproject.

- Ngoài ra bạn có thể tuỳ biến cho cây thư mục của bạn bằng cách chia phần code muốn share và không muốn share. File main chính của bạn có thể nằm trong thư mục /internal/app (ví dụ: /internal/app/myapp) và code sẽ được ứng dụng đó chia sẻ trong thư mục /internal/pkg (ví dụ: /internal/pkg/myprivlib).

# /pkg

- This is library code that's ok to use by external applications (eg., /pkg/my_public_lib)
- Các ứng dụng bên ngoài sẽ truy cập được vào thư mục này nên hay suy nghĩ cẩn thận khi tạo file code ở đây.
- Lưu ý rằng với thư mục internal sẽ đảm bảo độ riêng tư cho tệp tin hơn bởi

# /vendor

- This is application dependencies (quản lý toàn bộ hoặc các thư viện bạn hay dùng, yêu thích )
- Sử dụng lệnh để tạo thư mục /vendor

```
go mod vendor
```

# Service Application Directories

# /api

- OpenAPI/Swagger specs, JSON schema files, protocol definition files.

# /web

- Chứa các components web, static web assets, server side template

# Common Applications Directories

/configs
Configuration file templates or default configs.

Put your confd or consul-template template files here.

/init
System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs.

/scripts
Scripts to perform various build, install, analysis, etc operations.

These scripts keep the root level Makefile small and simple (e.g., https://github.com/hashicorp/terraform/blob/main/Makefile).

See the /scripts directory for examples.

/build
Packaging and Continuous Integration.

Put your cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the /build/package directory.

Put your CI (travis, circle, drone) configurations and scripts in the /build/ci directory. Note that some of the CI tools (e.g., Travis CI) are very picky about the location of their config files. Try putting the config files in the /build/ci directory linking them to the location where the CI tools expect them (when possible).

/deployments
IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, terraform). Note that in some repos (especially apps deployed with kubernetes) this directory is called /deploy.

/test
Additional external test apps and test data. Feel free to structure the /test directory anyway you want. For bigger projects it makes sense to have a data subdirectory. For example, you can have /test/data or /test/testdata if you need Go to ignore what's in that directory. Note that Go will also ignore directories or files that begin with "." or "\_", so you have more flexibility in terms of how you name your test data directory.

See the /test directory for examples.

# Other Directories

/docs
Design and user documents (in addition to your godoc generated documentation).

See the /docs directory for examples.

/tools
Supporting tools for this project. Note that these tools can import code from the /pkg and /internal directories.

See the /tools directory for examples.

/examples
Examples for your applications and/or public libraries.

See the /examples directory for examples.

/third_party
External helper tools, forked code and other 3rd party utilities (e.g., Swagger UI).

/githooks
Git hooks.

/assets
Other assets to go along with your repository (images, logos, etc).

/website
This is the place to put your project's website data if you are not using GitHub pages.

See the /website directory for examples.

Commands Executed:
cd /Users/kelvin/server-go/cmd/server: Changes the current directory to the server directory inside cmd, which is part of the server-go project.
go mod init server: Initializes a new module named server in the current directory, creating a new go.mod file. This is typically done to track the dependencies of this particular part of your application.
go mod tidy: (Suggested by the output) This command is recommended to add missing module requirements and remove unnecessary ones. It wasn't executed but is a good practice to keep the module tidy and up-to-date.
go run cmd/server/main.go: Attempts to run the main.go file located in the cmd/server directory from the project's root. However, this command fails with the error stat cmd/server/main.go: no such file or directory, indicating that the path provided doesn't point to an existing main.go file.

The error message "could not import github.com/canhcutcon/server-go-postgres/pkg/gee (current file is not included in a workspace module)" indicates a couple of potential issues with your Go environment and project setup. Let's address these:

Module Not Found in GOPATH or GO111MODULE:

Ensure that your Go environment is correctly set up with GOPATH. If you are using Go modules (which is likely, since there's a go.mod file), ensure that the GO111MODULE environment variable is set to on.
Run go mod tidy in your project root to ensure all dependencies are correctly fetched and recorded in your go.mod file.
Workspace Configuration:

The error suggests that the file server.go is not recognized as part of a Go module. This can happen if the IDE or text editor you are using is not correctly configured to recognize your project's structure.
Ensure that your IDE or editor is opened at the root directory of your Go module (where the go.mod file is located).
Incorrect Import Path:

Verify that the import path github.com/canhcutcon/server-go-postgres/pkg/gee is correct. If this is a local package within your project, it should be referenced relative to the module name declared in your go.mod file.
If gee is a local package, its import path should typically be something like module-name/pkg/gee, where module-name is the module name declared in your go.mod file.
To further assist, I can examine the go.mod file and the structure of the gee package in your project. This will help in determining if the issue is due to an incorrect import path or a misconfigured module. Shall we proceed with this examination?

which frame work and structure should choose when build a client-server application with golang?

#### SOCIAL MEDIA APPLICATION

1. Viper: The-Twelve-factor-app to :

- Setting default for applications
- Reading from JSON, TOML, YAML, HCL, .env files, and Java properties config files
- Reading from environment variables — it has a precedence order in which it reads from the configuration file:
  an explicit call to Set
  flag
  .env files
  config
  key/value store
  default

Using viper:

```
go get github.com/spf13/viper
```

## **REST API best practices**

- Exchange of data via JSON(sử dụng JSON để chuyển đổi dữ liệu): When we use REST API architecture we should use JSON format for sending data
  because:
  other format like YAML, XML lack some common procedures

- Nesting on endpoints(endpoint lồng nhau): The resultant endpoint will become https://www.example.com/posts/postId/comments. This is a good practice through which we can avoid mirroring database structure in our endpoints, keeping the data safe from attackers.

- Using nouns insted of verd

Comparisons of Different API Architecture Styles

#REST API

REST API can be useful in the following case:

Should use:

- Flexibility: REST also provides flexibility due to its lesser burden on the server side than GraphQL
- As a wrapper: We can use REST interface can also be used around the gRPC channel internally to route the request to other APIs
- Rest is comprehend: Rest is the most popular framework in the industry, and it will be very easy for developer to be able to use API

Avoid:

- In the event-driven system
- In big data processing

## GraphQL

Benefit:

- When there is a client application gathering data from multiple data sources. GraphQL aggregates data from multiple data sources and sends a consolidated response to the client. For example, in the case of a Stripe payment gateway, the data is retrieved from various endpoints, including customers, invoices, charges, and so on.
- When many client applications share one data source, but their view is different. GraphQL allows the applications to access the shared data where they can use it in a way that makes sense. Via GraphQL, applications can ask for the specific fields they want to present to the user instead of requesting all the fields.

Should use:

- When application is requesting data from multiple sources
- When multiple applications share a single database

Avoid:

- Server - to - server communication.

## gRPC

The gRPC API framework combines all the performance improvement and capabilities of HTTP/2 as a single package. It can be utilized in several cases that require high throughput backends to communicate with limited CPU and memory devices. In connection with this, gRPC is a framework with built-in features required to run a system and efficient serialization and deserialization capabilities.

Should use:

- When building low latency, highly scalable distributed system
- To create a backend system with a large number of microservices

Avoid:

- If the developer's or consumer’s language is not supported by the framework
- In applications if they are calling a limited number of back-end services


