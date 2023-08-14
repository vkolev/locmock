# locmock - A mocking API server

This is a simple application server for mocking API services.

The idea is that you define your services and actions in a directory structure, providing
information like:
- Expected path of the request
- Expected request
- Expected response

With this information when you call the locmock service you will get
your predefined response.

## Features

* Flat-File based database
* Service per directory
  * Service meta information:
    * Authentication
* Actions per subdirectory
* YAML based configuration
* Admin API
* Logging
* Tested

