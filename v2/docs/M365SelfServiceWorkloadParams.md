# M365SelfServiceWorkloadParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedSecurityGroups** | Pointer to [**[]M365SelfServiceSecurityGroupInfo**](M365SelfServiceSecurityGroupInfo.md) | Specifies the list of Security Groups whose members are to be allowed the Self-Service workflows. | [optional] 

## Methods

### NewM365SelfServiceWorkloadParams

`func NewM365SelfServiceWorkloadParams() *M365SelfServiceWorkloadParams`

NewM365SelfServiceWorkloadParams instantiates a new M365SelfServiceWorkloadParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewM365SelfServiceWorkloadParamsWithDefaults

`func NewM365SelfServiceWorkloadParamsWithDefaults() *M365SelfServiceWorkloadParams`

NewM365SelfServiceWorkloadParamsWithDefaults instantiates a new M365SelfServiceWorkloadParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedSecurityGroups

`func (o *M365SelfServiceWorkloadParams) GetAllowedSecurityGroups() []M365SelfServiceSecurityGroupInfo`

GetAllowedSecurityGroups returns the AllowedSecurityGroups field if non-nil, zero value otherwise.

### GetAllowedSecurityGroupsOk

`func (o *M365SelfServiceWorkloadParams) GetAllowedSecurityGroupsOk() (*[]M365SelfServiceSecurityGroupInfo, bool)`

GetAllowedSecurityGroupsOk returns a tuple with the AllowedSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSecurityGroups

`func (o *M365SelfServiceWorkloadParams) SetAllowedSecurityGroups(v []M365SelfServiceSecurityGroupInfo)`

SetAllowedSecurityGroups sets AllowedSecurityGroups field to given value.

### HasAllowedSecurityGroups

`func (o *M365SelfServiceWorkloadParams) HasAllowedSecurityGroups() bool`

HasAllowedSecurityGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


