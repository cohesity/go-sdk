# ApplicationServerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseCopyInfoList** | Pointer to [**[]ExchangeDatabaseCopyInfo**](ExchangeDatabaseCopyInfo.md) | Specifies the list of all the copies of the Exchange databases(that are part of DAG) that are present on this Exchange Node. | [optional] 
**DatabaseInfoList** | Pointer to [**[]ExchangeDatabaseInfo**](ExchangeDatabaseInfo.md) | Specifies the list of all the databases available on the standalone Exchange server node. This is populated for the Standlone Exchange Servers. | [optional] 
**Fqdn** | Pointer to **NullableString** | Specifies the fully qualified domain name of the Exchange Server. | [optional] 
**Guid** | Pointer to **NullableString** | Specifies the Guid of the Exchange Application Server. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the display name of the Exchange Application Server. | [optional] 
**TotalSizeBytes** | Pointer to **NullableInt64** | Specifies the total size of all Exchange database copies in all the Exchange Application Servers that are part of the DAG. | [optional] 

## Methods

### NewApplicationServerInfo

`func NewApplicationServerInfo() *ApplicationServerInfo`

NewApplicationServerInfo instantiates a new ApplicationServerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationServerInfoWithDefaults

`func NewApplicationServerInfoWithDefaults() *ApplicationServerInfo`

NewApplicationServerInfoWithDefaults instantiates a new ApplicationServerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseCopyInfoList

`func (o *ApplicationServerInfo) GetDatabaseCopyInfoList() []ExchangeDatabaseCopyInfo`

GetDatabaseCopyInfoList returns the DatabaseCopyInfoList field if non-nil, zero value otherwise.

### GetDatabaseCopyInfoListOk

`func (o *ApplicationServerInfo) GetDatabaseCopyInfoListOk() (*[]ExchangeDatabaseCopyInfo, bool)`

GetDatabaseCopyInfoListOk returns a tuple with the DatabaseCopyInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseCopyInfoList

`func (o *ApplicationServerInfo) SetDatabaseCopyInfoList(v []ExchangeDatabaseCopyInfo)`

SetDatabaseCopyInfoList sets DatabaseCopyInfoList field to given value.

### HasDatabaseCopyInfoList

`func (o *ApplicationServerInfo) HasDatabaseCopyInfoList() bool`

HasDatabaseCopyInfoList returns a boolean if a field has been set.

### SetDatabaseCopyInfoListNil

`func (o *ApplicationServerInfo) SetDatabaseCopyInfoListNil(b bool)`

 SetDatabaseCopyInfoListNil sets the value for DatabaseCopyInfoList to be an explicit nil

### UnsetDatabaseCopyInfoList
`func (o *ApplicationServerInfo) UnsetDatabaseCopyInfoList()`

UnsetDatabaseCopyInfoList ensures that no value is present for DatabaseCopyInfoList, not even an explicit nil
### GetDatabaseInfoList

`func (o *ApplicationServerInfo) GetDatabaseInfoList() []ExchangeDatabaseInfo`

GetDatabaseInfoList returns the DatabaseInfoList field if non-nil, zero value otherwise.

### GetDatabaseInfoListOk

`func (o *ApplicationServerInfo) GetDatabaseInfoListOk() (*[]ExchangeDatabaseInfo, bool)`

GetDatabaseInfoListOk returns a tuple with the DatabaseInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseInfoList

`func (o *ApplicationServerInfo) SetDatabaseInfoList(v []ExchangeDatabaseInfo)`

SetDatabaseInfoList sets DatabaseInfoList field to given value.

### HasDatabaseInfoList

`func (o *ApplicationServerInfo) HasDatabaseInfoList() bool`

HasDatabaseInfoList returns a boolean if a field has been set.

### SetDatabaseInfoListNil

`func (o *ApplicationServerInfo) SetDatabaseInfoListNil(b bool)`

 SetDatabaseInfoListNil sets the value for DatabaseInfoList to be an explicit nil

### UnsetDatabaseInfoList
`func (o *ApplicationServerInfo) UnsetDatabaseInfoList()`

UnsetDatabaseInfoList ensures that no value is present for DatabaseInfoList, not even an explicit nil
### GetFqdn

`func (o *ApplicationServerInfo) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ApplicationServerInfo) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ApplicationServerInfo) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ApplicationServerInfo) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *ApplicationServerInfo) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *ApplicationServerInfo) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetGuid

`func (o *ApplicationServerInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ApplicationServerInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ApplicationServerInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ApplicationServerInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *ApplicationServerInfo) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *ApplicationServerInfo) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetName

`func (o *ApplicationServerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationServerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationServerInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationServerInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApplicationServerInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApplicationServerInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTotalSizeBytes

`func (o *ApplicationServerInfo) GetTotalSizeBytes() int64`

GetTotalSizeBytes returns the TotalSizeBytes field if non-nil, zero value otherwise.

### GetTotalSizeBytesOk

`func (o *ApplicationServerInfo) GetTotalSizeBytesOk() (*int64, bool)`

GetTotalSizeBytesOk returns a tuple with the TotalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeBytes

`func (o *ApplicationServerInfo) SetTotalSizeBytes(v int64)`

SetTotalSizeBytes sets TotalSizeBytes field to given value.

### HasTotalSizeBytes

`func (o *ApplicationServerInfo) HasTotalSizeBytes() bool`

HasTotalSizeBytes returns a boolean if a field has been set.

### SetTotalSizeBytesNil

`func (o *ApplicationServerInfo) SetTotalSizeBytesNil(b bool)`

 SetTotalSizeBytesNil sets the value for TotalSizeBytes to be an explicit nil

### UnsetTotalSizeBytes
`func (o *ApplicationServerInfo) UnsetTotalSizeBytes()`

UnsetTotalSizeBytes ensures that no value is present for TotalSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


