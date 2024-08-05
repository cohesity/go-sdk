# AwsUseSTSParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthBlob** | Pointer to **NullableString** | Credential blob to use when interacting with credential endpoint. | [optional] 
**AuthEndpoint** | **NullableString** | Specifies the credential endpoint to use to generate the security token. | 

## Methods

### NewAwsUseSTSParams

`func NewAwsUseSTSParams(authEndpoint NullableString, ) *AwsUseSTSParams`

NewAwsUseSTSParams instantiates a new AwsUseSTSParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsUseSTSParamsWithDefaults

`func NewAwsUseSTSParamsWithDefaults() *AwsUseSTSParams`

NewAwsUseSTSParamsWithDefaults instantiates a new AwsUseSTSParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthBlob

`func (o *AwsUseSTSParams) GetAuthBlob() string`

GetAuthBlob returns the AuthBlob field if non-nil, zero value otherwise.

### GetAuthBlobOk

`func (o *AwsUseSTSParams) GetAuthBlobOk() (*string, bool)`

GetAuthBlobOk returns a tuple with the AuthBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthBlob

`func (o *AwsUseSTSParams) SetAuthBlob(v string)`

SetAuthBlob sets AuthBlob field to given value.

### HasAuthBlob

`func (o *AwsUseSTSParams) HasAuthBlob() bool`

HasAuthBlob returns a boolean if a field has been set.

### SetAuthBlobNil

`func (o *AwsUseSTSParams) SetAuthBlobNil(b bool)`

 SetAuthBlobNil sets the value for AuthBlob to be an explicit nil

### UnsetAuthBlob
`func (o *AwsUseSTSParams) UnsetAuthBlob()`

UnsetAuthBlob ensures that no value is present for AuthBlob, not even an explicit nil
### GetAuthEndpoint

`func (o *AwsUseSTSParams) GetAuthEndpoint() string`

GetAuthEndpoint returns the AuthEndpoint field if non-nil, zero value otherwise.

### GetAuthEndpointOk

`func (o *AwsUseSTSParams) GetAuthEndpointOk() (*string, bool)`

GetAuthEndpointOk returns a tuple with the AuthEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthEndpoint

`func (o *AwsUseSTSParams) SetAuthEndpoint(v string)`

SetAuthEndpoint sets AuthEndpoint field to given value.


### SetAuthEndpointNil

`func (o *AwsUseSTSParams) SetAuthEndpointNil(b bool)`

 SetAuthEndpointNil sets the value for AuthEndpoint to be an explicit nil

### UnsetAuthEndpoint
`func (o *AwsUseSTSParams) UnsetAuthEndpoint()`

UnsetAuthEndpoint ensures that no value is present for AuthEndpoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


