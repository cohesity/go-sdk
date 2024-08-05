# VcdConnectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Specifies the password to access target entity. | 
**Username** | **string** | Specifies the username to access target entity. | 
**Host** | **NullableString** | IP or hostname of the source. | 

## Methods

### NewVcdConnectionParams

`func NewVcdConnectionParams(password string, username string, host NullableString, ) *VcdConnectionParams`

NewVcdConnectionParams instantiates a new VcdConnectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcdConnectionParamsWithDefaults

`func NewVcdConnectionParamsWithDefaults() *VcdConnectionParams`

NewVcdConnectionParamsWithDefaults instantiates a new VcdConnectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *VcdConnectionParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcdConnectionParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcdConnectionParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *VcdConnectionParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcdConnectionParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcdConnectionParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetHost

`func (o *VcdConnectionParams) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VcdConnectionParams) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VcdConnectionParams) SetHost(v string)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *VcdConnectionParams) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *VcdConnectionParams) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


