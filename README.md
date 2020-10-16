# dc.libs.metrics.golang

Golang implementation for DeliveryCenter's structured logging format.

> This package is part of a family of libraries that implement DeliveryCenter's metrics pattern  for different languages. 
Check also our [Elixir](https://github.com/deliverycenter/dc.libs.metrics.elixir), 
>[Node](https://github.com/deliverycenter/dc.libs.metrics.node) and 
>[Ruby](https://github.com/deliverycenter/dc.libs.metrics.ruby) versions.

## Import
```
dcmetrics "github.com/deliverycenter/dc.libs.metrics.golang"
```

## Usage
The first step to use the package is to call the function `Setup`. This is where the global parameters for the logger will be set and used in every metrics call.

```go
err := dcmetrics.Setup("GOOGLE_PROJECT_ID", "PUBSUB_TOPIC_NAME", "DEVELOPMENT", "MY_PACKAGE", dcmetrics.Metrics{
	Direction:        "INCOMING",
	SourceType:       "PROVIDER",
	RootResourceType: "PRODUCT",
})
```


After the call to `Setup`, the logging functions can be used: `Debug()`, `Info()`, `Warn()`, `Error()`.

```go
dcmetrics.Info("request published", dcmetrics.Metrics{
	SourceName:        "PROVIDER_NAME",
	Action:            "REQUEST_PUBLISHED",
	RootResourceType:  "PRODUCT",
	ExtStoreID:        "12345",
	IntStoreID:        "67890",
	ExtRootResourceID: "abcde",
})
```

### Disable logs

If you need to disable the package behavior (for testing, for example), you can use:

```go
dcmetrics.Disable()
``` 