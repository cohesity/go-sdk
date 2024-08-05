# StorageDomainPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalStorageDomainId** | **NullableInt64** | Specifies the local Storage Domain id. | 
**LocalStorageDomainName** | Pointer to **NullableString** | Specifies the local Storage Domain name. | [optional] [readonly] 
**RemoteStorageDomainId** | **NullableInt64** | Specifies the remote Storage Domain id. | 
**RemoteStorageDomainName** | Pointer to **NullableString** | Specifies the remote Storage Domain name. | [optional] [readonly] 

## Methods

### NewStorageDomainPair

`func NewStorageDomainPair(localStorageDomainId NullableInt64, remoteStorageDomainId NullableInt64, ) *StorageDomainPair`

NewStorageDomainPair instantiates a new StorageDomainPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDomainPairWithDefaults

`func NewStorageDomainPairWithDefaults() *StorageDomainPair`

NewStorageDomainPairWithDefaults instantiates a new StorageDomainPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalStorageDomainId

`func (o *StorageDomainPair) GetLocalStorageDomainId() int64`

GetLocalStorageDomainId returns the LocalStorageDomainId field if non-nil, zero value otherwise.

### GetLocalStorageDomainIdOk

`func (o *StorageDomainPair) GetLocalStorageDomainIdOk() (*int64, bool)`

GetLocalStorageDomainIdOk returns a tuple with the LocalStorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorageDomainId

`func (o *StorageDomainPair) SetLocalStorageDomainId(v int64)`

SetLocalStorageDomainId sets LocalStorageDomainId field to given value.


### SetLocalStorageDomainIdNil

`func (o *StorageDomainPair) SetLocalStorageDomainIdNil(b bool)`

 SetLocalStorageDomainIdNil sets the value for LocalStorageDomainId to be an explicit nil

### UnsetLocalStorageDomainId
`func (o *StorageDomainPair) UnsetLocalStorageDomainId()`

UnsetLocalStorageDomainId ensures that no value is present for LocalStorageDomainId, not even an explicit nil
### GetLocalStorageDomainName

`func (o *StorageDomainPair) GetLocalStorageDomainName() string`

GetLocalStorageDomainName returns the LocalStorageDomainName field if non-nil, zero value otherwise.

### GetLocalStorageDomainNameOk

`func (o *StorageDomainPair) GetLocalStorageDomainNameOk() (*string, bool)`

GetLocalStorageDomainNameOk returns a tuple with the LocalStorageDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorageDomainName

`func (o *StorageDomainPair) SetLocalStorageDomainName(v string)`

SetLocalStorageDomainName sets LocalStorageDomainName field to given value.

### HasLocalStorageDomainName

`func (o *StorageDomainPair) HasLocalStorageDomainName() bool`

HasLocalStorageDomainName returns a boolean if a field has been set.

### SetLocalStorageDomainNameNil

`func (o *StorageDomainPair) SetLocalStorageDomainNameNil(b bool)`

 SetLocalStorageDomainNameNil sets the value for LocalStorageDomainName to be an explicit nil

### UnsetLocalStorageDomainName
`func (o *StorageDomainPair) UnsetLocalStorageDomainName()`

UnsetLocalStorageDomainName ensures that no value is present for LocalStorageDomainName, not even an explicit nil
### GetRemoteStorageDomainId

`func (o *StorageDomainPair) GetRemoteStorageDomainId() int64`

GetRemoteStorageDomainId returns the RemoteStorageDomainId field if non-nil, zero value otherwise.

### GetRemoteStorageDomainIdOk

`func (o *StorageDomainPair) GetRemoteStorageDomainIdOk() (*int64, bool)`

GetRemoteStorageDomainIdOk returns a tuple with the RemoteStorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteStorageDomainId

`func (o *StorageDomainPair) SetRemoteStorageDomainId(v int64)`

SetRemoteStorageDomainId sets RemoteStorageDomainId field to given value.


### SetRemoteStorageDomainIdNil

`func (o *StorageDomainPair) SetRemoteStorageDomainIdNil(b bool)`

 SetRemoteStorageDomainIdNil sets the value for RemoteStorageDomainId to be an explicit nil

### UnsetRemoteStorageDomainId
`func (o *StorageDomainPair) UnsetRemoteStorageDomainId()`

UnsetRemoteStorageDomainId ensures that no value is present for RemoteStorageDomainId, not even an explicit nil
### GetRemoteStorageDomainName

`func (o *StorageDomainPair) GetRemoteStorageDomainName() string`

GetRemoteStorageDomainName returns the RemoteStorageDomainName field if non-nil, zero value otherwise.

### GetRemoteStorageDomainNameOk

`func (o *StorageDomainPair) GetRemoteStorageDomainNameOk() (*string, bool)`

GetRemoteStorageDomainNameOk returns a tuple with the RemoteStorageDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteStorageDomainName

`func (o *StorageDomainPair) SetRemoteStorageDomainName(v string)`

SetRemoteStorageDomainName sets RemoteStorageDomainName field to given value.

### HasRemoteStorageDomainName

`func (o *StorageDomainPair) HasRemoteStorageDomainName() bool`

HasRemoteStorageDomainName returns a boolean if a field has been set.

### SetRemoteStorageDomainNameNil

`func (o *StorageDomainPair) SetRemoteStorageDomainNameNil(b bool)`

 SetRemoteStorageDomainNameNil sets the value for RemoteStorageDomainName to be an explicit nil

### UnsetRemoteStorageDomainName
`func (o *StorageDomainPair) UnsetRemoteStorageDomainName()`

UnsetRemoteStorageDomainName ensures that no value is present for RemoteStorageDomainName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


