# dc.libs.metrics.golang

## Import
```
dcmetrics "github.com/deliverycenter/dc.libs.metrics.golang"
```

## Usage
The first step to use the package is to call the function `Setup`. This is where the global parameters for the logger will be set and used in every metrics call.

```
err := dcmetrics.Setup("METRICTS_GRPC_SERVER_ADDRESS", "DEVELOPMENT", "MY_PACKAGE", dcmetrics.Metrics{
	Direction:        "INCOMING",
	SourceType:       "PROVIDER",
	RootResourceType: "PRODUCT",
})
```


After the call to `Setup`, the logging functions can be used: `Debug()`, `Info()`, `Warn()`, `Error()`.

```
dcmetrics.Info("request published", dcmetrics.Metrics{
	SourceName:        "PROVIDER_NAME",
	Action:            "REQUEST_PUBLISHED",
	RootResourceType:  "PRODUCT",
	ExtStoreID:        "12345",
	IntStoreID:        "67890",
	ExtRootResourceID: "abcde",
})
```
