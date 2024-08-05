# \MFAAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailOtp**](MFAAPI.md#CreateEmailOtp) | **Post** /email-otp | Creates a new OTP to be sent to the user email.
[**CreateTotpKey**](MFAAPI.md#CreateTotpKey) | **Post** /totp-key | Create a new TOTP secret URI and store the secret key.
[**GetMFAConfig**](MFAAPI.md#GetMFAConfig) | **Get** /mfa-config | Returns the current MFA configuration.
[**GetSupportMFAConfig**](MFAAPI.md#GetSupportMFAConfig) | **Get** /support-user/mfa | Returns the current MFA configuration.
[**SendEmailOtp**](MFAAPI.md#SendEmailOtp) | **Post** /send-email-otp | Creates a new OTP to be sent to the user email.
[**SendSupportEmailOtp**](MFAAPI.md#SendSupportEmailOtp) | **Post** /support-user/send-email-otp | Creates a new OTP to be sent to the linux support user email.
[**UpdateMFAConfig**](MFAAPI.md#UpdateMFAConfig) | **Put** /mfa-config | Stores the updated MFA configuration.
[**UpdateSupportMFAConfig**](MFAAPI.md#UpdateSupportMFAConfig) | **Patch** /support-user/mfa | Stores the updated MFA configuration.
[**VerifySupportUserTotp**](MFAAPI.md#VerifySupportUserTotp) | **Post** /support-user/verify-totp | Verify the totp code for support user.



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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewCreateEmailOtpRequestBody() // CreateEmailOtpRequestBody | Specifies the parameters to send email OTP. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MFAAPI.CreateEmailOtp(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.CreateEmailOtp``: %v\n", err)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewCreateTotpKeyRequestBody() // CreateTotpKeyRequestBody | Specifies the key id for creating the TOTP key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.CreateTotpKey(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.CreateTotpKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTotpKey`: TotpKeyInfo
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.CreateTotpKey`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.GetMFAConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.GetMFAConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMFAConfig`: MfaConfigInfo
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.GetMFAConfig`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.GetSupportMFAConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.GetSupportMFAConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportMFAConfig`: SupportMfaConfigInfo
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.GetSupportMFAConfig`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MFAAPI.SendEmailOtp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.SendEmailOtp``: %v\n", err)
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


## SendSupportEmailOtp

> SendSupportEmailOtp(ctx).Execute()

Creates a new OTP to be sent to the linux support user email.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MFAAPI.SendSupportEmailOtp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.SendSupportEmailOtp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSendSupportEmailOtpRequest struct via the builder pattern


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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewMfaConfigInfo() // MfaConfigInfo | The update request for the MFA Settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.UpdateMFAConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.UpdateMFAConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMFAConfig`: MfaConfigInfo
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.UpdateMFAConfig`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewSupportMfaConfigInfo() // SupportMfaConfigInfo | The update request for the MFA Settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.UpdateSupportMFAConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.UpdateSupportMFAConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSupportMFAConfig`: UpdateMFAResult
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.UpdateSupportMFAConfig`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewVerifyTotpRequest() // VerifyTotpRequest | Totp code to be verified.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.VerifySupportUserTotp(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.VerifySupportUserTotp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifySupportUserTotp`: VerifyTotpResult
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.VerifySupportUserTotp`: %v\n", resp)
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

