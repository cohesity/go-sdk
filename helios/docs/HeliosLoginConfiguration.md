# HeliosLoginConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Specifies the consumer key configured in Salesforce | [optional] 
**PortalUrl** | Pointer to **string** | Specifies the salesforce URL to be redirected to | [optional] 
**RedirectUrl** | Pointer to **string** | Redirect URL after successful authentication | [optional] 

## Methods

### NewHeliosLoginConfiguration

`func NewHeliosLoginConfiguration() *HeliosLoginConfiguration`

NewHeliosLoginConfiguration instantiates a new HeliosLoginConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosLoginConfigurationWithDefaults

`func NewHeliosLoginConfigurationWithDefaults() *HeliosLoginConfiguration`

NewHeliosLoginConfigurationWithDefaults instantiates a new HeliosLoginConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *HeliosLoginConfiguration) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *HeliosLoginConfiguration) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *HeliosLoginConfiguration) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *HeliosLoginConfiguration) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetPortalUrl

`func (o *HeliosLoginConfiguration) GetPortalUrl() string`

GetPortalUrl returns the PortalUrl field if non-nil, zero value otherwise.

### GetPortalUrlOk

`func (o *HeliosLoginConfiguration) GetPortalUrlOk() (*string, bool)`

GetPortalUrlOk returns a tuple with the PortalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalUrl

`func (o *HeliosLoginConfiguration) SetPortalUrl(v string)`

SetPortalUrl sets PortalUrl field to given value.

### HasPortalUrl

`func (o *HeliosLoginConfiguration) HasPortalUrl() bool`

HasPortalUrl returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *HeliosLoginConfiguration) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *HeliosLoginConfiguration) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *HeliosLoginConfiguration) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *HeliosLoginConfiguration) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


