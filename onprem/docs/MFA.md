# \MFA

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailOtp**](MFA.md#CreateEmailOtp) | **Post** /email-otp | Creates a new OTP to be sent to the user email.
[**CreateTotpKey**](MFA.md#CreateTotpKey) | **Post** /totp-key | Create a new TOTP secret URI and store the secret key.
[**GetMFAConfig**](MFA.md#GetMFAConfig) | **Get** /mfa-config | Returns the current MFA configuration.
[**GetSupportMFAConfig**](MFA.md#GetSupportMFAConfig) | **Get** /support-user/mfa | Returns the current MFA configuration.
[**SendEmailOtp**](MFA.md#SendEmailOtp) | **Post** /send-email-otp | Creates a new OTP to be sent to the user email.
[**UpdateMFAConfig**](MFA.md#UpdateMFAConfig) | **Put** /mfa-config | Stores the updated MFA configuration.
[**UpdateSupportMFAConfig**](MFA.md#UpdateSupportMFAConfig) | **Patch** /support-user/mfa | Stores the updated MFA configuration.
[**VerifySupportUserTotp**](MFA.md#VerifySupportUserTotp) | **Post** /support-user/verify-totp | Verify the totp code for support user.



## CreateEmailOtp

> CreateEmailOtp(ctx).Body(body).Execute()

Creates a new OTP to be sent to the user email.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewCreateEmailOtpRequestBody() // CreateEmailOtpRequestBody | Specifies the parameters to send email OTP. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiCreateEmailOtpRequest{
        Body: &body
    }

    resp, r, err := api_client.MFA.CreateEmailOtp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFA.CreateEmailOtp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailOtpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateEmailOtpRequestBody**](CreateEmailOtpRequestBody.md) | Specifies the parameters to send email OTP. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTotpKey

> TotpKeyInfo CreateTotpKey(ctx).Body(body).Execute()

Create a new TOTP secret URI and store the secret key.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewCreateTotpKeyRequestBody() // CreateTotpKeyRequestBody | Specifies the key id for creating the TOTP key.

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiCreateTotpKeyRequest{
        Body: &body
    }

    resp, r, err := api_client.MFA.CreateTotpKey(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFA.CreateTotpKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTotpKey`: TotpKeyInfo
    fmt.Fprintf(os.Stdout, "Response from `MFA.CreateTotpKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTotpKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateTotpKeyRequestBody**](CreateTotpKeyRequestBody.md) | Specifies the key id for creating the TOTP key. | 

### Return type

[**TotpKeyInfo**](TotpKeyInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMFAConfig

> MfaConfigInfo GetMFAConfig(ctx).Execute()

Returns the current MFA configuration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiGetMFAConfigRequest{
    }

    resp, r, err := api_client.MFA.GetMFAConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFA.GetMFAConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMFAConfig`: MfaConfigInfo
    fmt.Fprintf(os.Stdout, "Response from `MFA.GetMFAConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMFAConfigRequest struct via the builder pattern


### Return type

[**MfaConfigInfo**](MfaConfigInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportMFAConfig

> SupportMfaConfigInfo GetSupportMFAConfig(ctx).Execute()

Returns the current MFA configuration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiGetSupportMFAConfigRequest{
    }

    resp, r, err := api_client.MFA.GetSupportMFAConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFA.GetSupportMFAConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportMFAConfig`: SupportMfaConfigInfo
    fmt.Fprintf(os.Stdout, "Response from `MFA.GetSupportMFAConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportMFAConfigRequest struct via the builder pattern


### Return type

[**SupportMfaConfigInfo**](SupportMfaConfigInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEmailOtp

> SendEmailOtp(ctx).Execute()

Creates a new OTP to be sent to the user email.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiSendEmailOtpRequest{
    }

    resp, r, err := api_client.MFA.SendEmailOtp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFA.SendEmailOtp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailOtpRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMFAConfig

> MfaConfigInfo UpdateMFAConfig(ctx).Body(body).Execute()

Stores the updated MFA configuration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewMfaConfigInfo() // MfaConfigInfo | The update request for the MFA Settings

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiUpdateMFAConfigRequest{
        Body: &body
    }

    resp, r, err := api_client.MFA.UpdateMFAConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFA.UpdateMFAConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMFAConfig`: MfaConfigInfo
    fmt.Fprintf(os.Stdout, "Response from `MFA.UpdateMFAConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMFAConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MfaConfigInfo**](MfaConfigInfo.md) | The update request for the MFA Settings | 

### Return type

[**MfaConfigInfo**](MfaConfigInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSupportMFAConfig

> UpdateMFAResult UpdateSupportMFAConfig(ctx).Body(body).Execute()

Stores the updated MFA configuration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewSupportMfaConfigInfo() // SupportMfaConfigInfo | The update request for the MFA Settings

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiUpdateSupportMFAConfigRequest{
        Body: &body
    }

    resp, r, err := api_client.MFA.UpdateSupportMFAConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFA.UpdateSupportMFAConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSupportMFAConfig`: UpdateMFAResult
    fmt.Fprintf(os.Stdout, "Response from `MFA.UpdateSupportMFAConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSupportMFAConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SupportMfaConfigInfo**](SupportMfaConfigInfo.md) | The update request for the MFA Settings | 

### Return type

[**UpdateMFAResult**](UpdateMFAResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifySupportUserTotp

> VerifyTotpResult VerifySupportUserTotp(ctx).Body(body).Execute()

Verify the totp code for support user.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewVerifyTotpRequest() // VerifyTotpRequest | Totp code to be verified.

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiVerifySupportUserTotpRequest{
        Body: &body
    }

    resp, r, err := api_client.MFA.VerifySupportUserTotp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFA.VerifySupportUserTotp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifySupportUserTotp`: VerifyTotpResult
    fmt.Fprintf(os.Stdout, "Response from `MFA.VerifySupportUserTotp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifySupportUserTotpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VerifyTotpRequest**](VerifyTotpRequest.md) | Totp code to be verified. | 

### Return type

[**VerifyTotpResult**](VerifyTotpResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

