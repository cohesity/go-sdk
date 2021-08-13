# CommonExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the ID of the External Target. | [optional] [readonly] 
**Name** | **NullableString** | Specifies the name of the External Target. | 
**PurposeType** | **NullableString** | Specifies the purpose of the External Target. | 
**Compression** | Pointer to **NullableString** | Specifies whether the type of compression of the External Target | [optional] 
**Status** | Pointer to **NullableString** | Specifies the registration status of the External Target | [optional] [readonly] 
**TenantIds** | Pointer to **[]string** | Specifies the list of tenantIds for the External Target | [optional] 
**CloudDomains** | Pointer to [**[]CloudDomain**](CloudDomain.md) | Specifies the cloud domain information. | [optional] 

## Methods

### NewCommonExternalTargetParams

`func NewCommonExternalTargetParams(name NullableString, purposeType NullableString, ) *CommonExternalTargetParams`

NewCommonExternalTargetParams instantiates a new CommonExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonExternalTargetParamsWithDefaults

`func NewCommonExternalTargetParamsWithDefaults() *CommonExternalTargetParams`

NewCommonExternalTargetParamsWithDefaults instantiates a new CommonExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonExternalTargetParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonExternalTargetParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonExternalTargetParams) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CommonExternalTargetParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CommonExternalTargetParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommonExternalTargetParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CommonExternalTargetParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonExternalTargetParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonExternalTargetParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CommonExternalTargetParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonExternalTargetParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPurposeType

`func (o *CommonExternalTargetParams) GetPurposeType() string`

GetPurposeType returns the PurposeType field if non-nil, zero value otherwise.

### GetPurposeTypeOk

`func (o *CommonExternalTargetParams) GetPurposeTypeOk() (*string, bool)`

GetPurposeTypeOk returns a tuple with the PurposeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeType

`func (o *CommonExternalTargetParams) SetPurposeType(v string)`

SetPurposeType sets PurposeType field to given value.


### SetPurposeTypeNil

`func (o *CommonExternalTargetParams) SetPurposeTypeNil(b bool)`

 SetPurposeTypeNil sets the value for PurposeType to be an explicit nil

### UnsetPurposeType
`func (o *CommonExternalTargetParams) UnsetPurposeType()`

UnsetPurposeType ensures that no value is present for PurposeType, not even an explicit nil
### GetCompression

`func (o *CommonExternalTargetParams) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *CommonExternalTargetParams) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *CommonExternalTargetParams) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *CommonExternalTargetParams) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### SetCompressionNil

`func (o *CommonExternalTargetParams) SetCompressionNil(b bool)`

 SetCompressionNil sets the value for Compression to be an explicit nil

### UnsetCompression
`func (o *CommonExternalTargetParams) UnsetCompression()`

UnsetCompression ensures that no value is present for Compression, not even an explicit nil
### GetStatus

`func (o *CommonExternalTargetParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonExternalTargetParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonExternalTargetParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommonExternalTargetParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CommonExternalTargetParams) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CommonExternalTargetParams) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTenantIds

`func (o *CommonExternalTargetParams) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *CommonExternalTargetParams) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *CommonExternalTargetParams) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *CommonExternalTargetParams) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### GetCloudDomains

`func (o *CommonExternalTargetParams) GetCloudDomains() []CloudDomain`

GetCloudDomains returns the CloudDomains field if non-nil, zero value otherwise.

### GetCloudDomainsOk

`func (o *CommonExternalTargetParams) GetCloudDomainsOk() (*[]CloudDomain, bool)`

GetCloudDomainsOk returns a tuple with the CloudDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDomains

`func (o *CommonExternalTargetParams) SetCloudDomains(v []CloudDomain)`

SetCloudDomains sets CloudDomains field to given value.

### HasCloudDomains

`func (o *CommonExternalTargetParams) HasCloudDomains() bool`

HasCloudDomains returns a boolean if a field has been set.

### SetCloudDomainsNil

`func (o *CommonExternalTargetParams) SetCloudDomainsNil(b bool)`

 SetCloudDomainsNil sets the value for CloudDomains to be an explicit nil

### UnsetCloudDomains
`func (o *CommonExternalTargetParams) UnsetCloudDomains()`

UnsetCloudDomains ensures that no value is present for CloudDomains, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


