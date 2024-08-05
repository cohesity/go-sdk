# VcdRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Specifies the password to access target entity. | 
**Username** | **string** | Specifies the username to access target entity. | 
**Description** | Pointer to **NullableString** | Specifies the description of the source being registered. | [optional] 
**Endpoint** | **string** | Specifies the endpoint IPaddress, URL or hostname of the host. | 
**VcenterCredentialInfoList** | [**[]VcenterCredentialInfo**](VcenterCredentialInfo.md) | Specifies the credentials information for all the vcenters in vcloud director. | 

## Methods

### NewVcdRegistrationParams

`func NewVcdRegistrationParams(password string, username string, endpoint string, vcenterCredentialInfoList []VcenterCredentialInfo, ) *VcdRegistrationParams`

NewVcdRegistrationParams instantiates a new VcdRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcdRegistrationParamsWithDefaults

`func NewVcdRegistrationParamsWithDefaults() *VcdRegistrationParams`

NewVcdRegistrationParamsWithDefaults instantiates a new VcdRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *VcdRegistrationParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcdRegistrationParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcdRegistrationParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *VcdRegistrationParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcdRegistrationParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcdRegistrationParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDescription

`func (o *VcdRegistrationParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcdRegistrationParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcdRegistrationParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcdRegistrationParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VcdRegistrationParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VcdRegistrationParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndpoint

`func (o *VcdRegistrationParams) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *VcdRegistrationParams) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *VcdRegistrationParams) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetVcenterCredentialInfoList

`func (o *VcdRegistrationParams) GetVcenterCredentialInfoList() []VcenterCredentialInfo`

GetVcenterCredentialInfoList returns the VcenterCredentialInfoList field if non-nil, zero value otherwise.

### GetVcenterCredentialInfoListOk

`func (o *VcdRegistrationParams) GetVcenterCredentialInfoListOk() (*[]VcenterCredentialInfo, bool)`

GetVcenterCredentialInfoListOk returns a tuple with the VcenterCredentialInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterCredentialInfoList

`func (o *VcdRegistrationParams) SetVcenterCredentialInfoList(v []VcenterCredentialInfo)`

SetVcenterCredentialInfoList sets VcenterCredentialInfoList field to given value.


### SetVcenterCredentialInfoListNil

`func (o *VcdRegistrationParams) SetVcenterCredentialInfoListNil(b bool)`

 SetVcenterCredentialInfoListNil sets the value for VcenterCredentialInfoList to be an explicit nil

### UnsetVcenterCredentialInfoList
`func (o *VcdRegistrationParams) UnsetVcenterCredentialInfoList()`

UnsetVcenterCredentialInfoList ensures that no value is present for VcenterCredentialInfoList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


