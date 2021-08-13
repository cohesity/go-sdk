# VcenterCredentialInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Specifies the username to access target entity. | 
**Password** | **string** | Specifies the password to access target entity. | 
**Endpoint** | **string** | Specifies the endpoint IPaddress, URL or hostname of the host. | 
**Description** | Pointer to **NullableString** | Specifies the description of the source being registered. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the vCenter. | [optional] 

## Methods

### NewVcenterCredentialInfo

`func NewVcenterCredentialInfo(username string, password string, endpoint string, ) *VcenterCredentialInfo`

NewVcenterCredentialInfo instantiates a new VcenterCredentialInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCredentialInfoWithDefaults

`func NewVcenterCredentialInfoWithDefaults() *VcenterCredentialInfo`

NewVcenterCredentialInfoWithDefaults instantiates a new VcenterCredentialInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *VcenterCredentialInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterCredentialInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterCredentialInfo) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *VcenterCredentialInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterCredentialInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterCredentialInfo) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEndpoint

`func (o *VcenterCredentialInfo) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *VcenterCredentialInfo) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *VcenterCredentialInfo) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetDescription

`func (o *VcenterCredentialInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterCredentialInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterCredentialInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterCredentialInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VcenterCredentialInfo) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VcenterCredentialInfo) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *VcenterCredentialInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterCredentialInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterCredentialInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterCredentialInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VcenterCredentialInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VcenterCredentialInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


