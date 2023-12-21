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
