# DirQuotaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**DirQuotaConfig**](DirQuotaConfig.md) |  | [optional] 
**Quotas** | Pointer to [**[]DirQuotaPolicy**](DirQuotaPolicy.md) | Specifies the list of directory quota policies applied on the view. | [optional] 

## Methods

### NewDirQuotaInfo

`func NewDirQuotaInfo() *DirQuotaInfo`

NewDirQuotaInfo instantiates a new DirQuotaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirQuotaInfoWithDefaults

`func NewDirQuotaInfoWithDefaults() *DirQuotaInfo`

NewDirQuotaInfoWithDefaults instantiates a new DirQuotaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *DirQuotaInfo) GetConfig() DirQuotaConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DirQuotaInfo) GetConfigOk() (*DirQuotaConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DirQuotaInfo) SetConfig(v DirQuotaConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DirQuotaInfo) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetQuotas

`func (o *DirQuotaInfo) GetQuotas() []DirQuotaPolicy`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *DirQuotaInfo) GetQuotasOk() (*[]DirQuotaPolicy, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *DirQuotaInfo) SetQuotas(v []DirQuotaPolicy)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *DirQuotaInfo) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.

### SetQuotasNil

`func (o *DirQuotaInfo) SetQuotasNil(b bool)`

 SetQuotasNil sets the value for Quotas to be an explicit nil

### UnsetQuotas
`func (o *DirQuotaInfo) UnsetQuotas()`

UnsetQuotas ensures that no value is present for Quotas, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


