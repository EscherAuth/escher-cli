# Escher CLI Tool
[![Build Status](https://travis-ci.org/EscherAuth/escher-cli.svg?branch=master)](https://travis-ci.org/EscherAuth/escher-cli)

This tool is created for working with and implementing Escher Auth into web projects,
without the requirement to modify the code base.

## Requirement

The client that use escher-cli must implement the use of HTTP_PROXY like curl does for outgoing requests.

For Server, it must use the PORT env variable like at heroku in order to use the escher validation with the cli tool

you must also configure your env with the applications escher configuration such as KEY_POOL for example.

## Usage

### client

curl here is only an example application to demonstrate the usage

```shell
escher-cli curl google.com
```

### Web service

Again, you must set the port number in order to enable escher-cli validation feature. Your server also expected to use the $PORT env variable for setting it's own listening.
This is a common thing at heroku Procfile for example
```shell
export PORT="$THE_PORT_NUMBER_THAT_I_EXPECTED_TO_USE"
```

than all is left to use your application
```shell
escher-cli ./my-service-app
escher-cli bundle exec rackup -p $PORT
```
