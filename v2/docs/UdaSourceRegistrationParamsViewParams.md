# UdaSourceRegistrationParamsViewParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountDir** | Pointer to **NullableString** | This field is deprecated and its value will be ignored. It was used to specify the absolute path on the host where the Cohesity view would be mounted. This is now controlled by the agent configuration and is common for all the Universal Data Adapter sources. | [optional] 

## Methods

### NewUdaSourceRegistrationParamsViewParams

`func NewUdaSourceRegistrationParamsViewParams() *UdaSourceRegistrationParamsViewParams`

NewUdaSourceRegistrationParamsViewParams instantiates a new UdaSourceRegistrationParamsViewParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdaSourceRegistrationParamsViewParamsWithDefaults

`func NewUdaSourceRegistrationParamsViewParamsWithDefaults() *UdaSourceRegistrationParamsViewParams`

NewUdaSourceRegistrationParamsViewParamsWithDefaults instantiates a new UdaSourceRegistrationParamsViewParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountDir

`func (o *UdaSourceRegistrationParamsViewParams) GetMountDir() string`

GetMountDir returns the MountDir field if non-nil, zero value otherwise.

### GetMountDirOk

`func (o *UdaSourceRegistrationParamsViewParams) GetMountDirOk() (*string, bool)`

GetMountDirOk returns a tuple with the MountDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountDir

`func (o *UdaSourceRegistrationParamsViewParams) SetMountDir(v string)`

SetMountDir sets MountDir field to given value.

### HasMountDir

`func (o *UdaSourceRegistrationParamsViewParams) HasMountDir() bool`

HasMountDir returns a boolean if a field has been set.

### SetMountDirNil

`func (o *UdaSourceRegistrationParamsViewParams) SetMountDirNil(b bool)`

 SetMountDirNil sets the value for MountDir to be an explicit nil

### UnsetMountDir
`func (o *UdaSourceRegistrationParamsViewParams) UnsetMountDir()`

UnsetMountDir ensures that no value is present for MountDir, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


