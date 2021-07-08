# SourceForPrincipalParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionSourceIds** | Pointer to **[]int64** | Array of Protection Source Ids.  For the specified principal, grant access permissions to the Protection Sources listed in this array. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the SID of the principal to grant access permissions to. | [optional] 
**ViewNames** | Pointer to **[]string** | Array of View names.  For the specified principal, grant access permissions to the Views names listed in this array. | [optional] 

## Methods

### NewSourceForPrincipalParam

`func NewSourceForPrincipalParam() *SourceForPrincipalParam`

NewSourceForPrincipalParam instantiates a new SourceForPrincipalParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceForPrincipalParamWithDefaults

`func NewSourceForPrincipalParamWithDefaults() *SourceForPrincipalParam`

NewSourceForPrincipalParamWithDefaults instantiates a new SourceForPrincipalParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionSourceIds

`func (o *SourceForPrincipalParam) GetProtectionSourceIds() []int64`

GetProtectionSourceIds returns the ProtectionSourceIds field if non-nil, zero value otherwise.

### GetProtectionSourceIdsOk

`func (o *SourceForPrincipalParam) GetProtectionSourceIdsOk() (*[]int64, bool)`

GetProtectionSourceIdsOk returns a tuple with the ProtectionSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceIds

`func (o *SourceForPrincipalParam) SetProtectionSourceIds(v []int64)`

SetProtectionSourceIds sets ProtectionSourceIds field to given value.

### HasProtectionSourceIds

`func (o *SourceForPrincipalParam) HasProtectionSourceIds() bool`

HasProtectionSourceIds returns a boolean if a field has been set.

### SetProtectionSourceIdsNil

`func (o *SourceForPrincipalParam) SetProtectionSourceIdsNil(b bool)`

 SetProtectionSourceIdsNil sets the value for ProtectionSourceIds to be an explicit nil

### UnsetProtectionSourceIds
`func (o *SourceForPrincipalParam) UnsetProtectionSourceIds()`

UnsetProtectionSourceIds ensures that no value is present for ProtectionSourceIds, not even an explicit nil
### GetSid

`func (o *SourceForPrincipalParam) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SourceForPrincipalParam) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SourceForPrincipalParam) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SourceForPrincipalParam) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SourceForPrincipalParam) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SourceForPrincipalParam) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetViewNames

`func (o *SourceForPrincipalParam) GetViewNames() []string`

GetViewNames returns the ViewNames field if non-nil, zero value otherwise.

### GetViewNamesOk

`func (o *SourceForPrincipalParam) GetViewNamesOk() (*[]string, bool)`

GetViewNamesOk returns a tuple with the ViewNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewNames

`func (o *SourceForPrincipalParam) SetViewNames(v []string)`

SetViewNames sets ViewNames field to given value.

### HasViewNames

`func (o *SourceForPrincipalParam) HasViewNames() bool`

HasViewNames returns a boolean if a field has been set.

### SetViewNamesNil

`func (o *SourceForPrincipalParam) SetViewNamesNil(b bool)`

 SetViewNamesNil sets the value for ViewNames to be an explicit nil

### UnsetViewNames
`func (o *SourceForPrincipalParam) UnsetViewNames()`

UnsetViewNames ensures that no value is present for ViewNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


