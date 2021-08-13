# EsxiRegistrationParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinFreeDatastoreSpaceForBackupGb** | Pointer to **NullableInt64** | Specifies the minimum free space (in GB) expected to be available in the datastore where the virtual disks of the VM being backed up reside. If the space available is lower than the specified value, backup will be aborted. | [optional] 
**MaxConcurrentStreams** | Pointer to **NullableInt32** | If this value is &gt; 0 and the number of streams concurrently active on a datastore is equal to it, then any further requests to access the datastore would be denied until the number of active streams reduces. This applies for all the datastores in the specified host. | [optional] 
**DataStoreParams** | Pointer to [**[]DatastoreParams**](DatastoreParams.md) | Specifies the datastore specific params. | [optional] 

## Methods

### NewEsxiRegistrationParamsAllOf

`func NewEsxiRegistrationParamsAllOf() *EsxiRegistrationParamsAllOf`

NewEsxiRegistrationParamsAllOf instantiates a new EsxiRegistrationParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsxiRegistrationParamsAllOfWithDefaults

`func NewEsxiRegistrationParamsAllOfWithDefaults() *EsxiRegistrationParamsAllOf`

NewEsxiRegistrationParamsAllOfWithDefaults instantiates a new EsxiRegistrationParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinFreeDatastoreSpaceForBackupGb

`func (o *EsxiRegistrationParamsAllOf) GetMinFreeDatastoreSpaceForBackupGb() int64`

GetMinFreeDatastoreSpaceForBackupGb returns the MinFreeDatastoreSpaceForBackupGb field if non-nil, zero value otherwise.

### GetMinFreeDatastoreSpaceForBackupGbOk

`func (o *EsxiRegistrationParamsAllOf) GetMinFreeDatastoreSpaceForBackupGbOk() (*int64, bool)`

GetMinFreeDatastoreSpaceForBackupGbOk returns a tuple with the MinFreeDatastoreSpaceForBackupGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFreeDatastoreSpaceForBackupGb

`func (o *EsxiRegistrationParamsAllOf) SetMinFreeDatastoreSpaceForBackupGb(v int64)`

SetMinFreeDatastoreSpaceForBackupGb sets MinFreeDatastoreSpaceForBackupGb field to given value.

### HasMinFreeDatastoreSpaceForBackupGb

`func (o *EsxiRegistrationParamsAllOf) HasMinFreeDatastoreSpaceForBackupGb() bool`

HasMinFreeDatastoreSpaceForBackupGb returns a boolean if a field has been set.

### SetMinFreeDatastoreSpaceForBackupGbNil

`func (o *EsxiRegistrationParamsAllOf) SetMinFreeDatastoreSpaceForBackupGbNil(b bool)`

 SetMinFreeDatastoreSpaceForBackupGbNil sets the value for MinFreeDatastoreSpaceForBackupGb to be an explicit nil

### UnsetMinFreeDatastoreSpaceForBackupGb
`func (o *EsxiRegistrationParamsAllOf) UnsetMinFreeDatastoreSpaceForBackupGb()`

UnsetMinFreeDatastoreSpaceForBackupGb ensures that no value is present for MinFreeDatastoreSpaceForBackupGb, not even an explicit nil
### GetMaxConcurrentStreams

`func (o *EsxiRegistrationParamsAllOf) GetMaxConcurrentStreams() int32`

GetMaxConcurrentStreams returns the MaxConcurrentStreams field if non-nil, zero value otherwise.

### GetMaxConcurrentStreamsOk

`func (o *EsxiRegistrationParamsAllOf) GetMaxConcurrentStreamsOk() (*int32, bool)`

GetMaxConcurrentStreamsOk returns a tuple with the MaxConcurrentStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentStreams

`func (o *EsxiRegistrationParamsAllOf) SetMaxConcurrentStreams(v int32)`

SetMaxConcurrentStreams sets MaxConcurrentStreams field to given value.

### HasMaxConcurrentStreams

`func (o *EsxiRegistrationParamsAllOf) HasMaxConcurrentStreams() bool`

HasMaxConcurrentStreams returns a boolean if a field has been set.

### SetMaxConcurrentStreamsNil

`func (o *EsxiRegistrationParamsAllOf) SetMaxConcurrentStreamsNil(b bool)`

 SetMaxConcurrentStreamsNil sets the value for MaxConcurrentStreams to be an explicit nil

### UnsetMaxConcurrentStreams
`func (o *EsxiRegistrationParamsAllOf) UnsetMaxConcurrentStreams()`

UnsetMaxConcurrentStreams ensures that no value is present for MaxConcurrentStreams, not even an explicit nil
### GetDataStoreParams

`func (o *EsxiRegistrationParamsAllOf) GetDataStoreParams() []DatastoreParams`

GetDataStoreParams returns the DataStoreParams field if non-nil, zero value otherwise.

### GetDataStoreParamsOk

`func (o *EsxiRegistrationParamsAllOf) GetDataStoreParamsOk() (*[]DatastoreParams, bool)`

GetDataStoreParamsOk returns a tuple with the DataStoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreParams

`func (o *EsxiRegistrationParamsAllOf) SetDataStoreParams(v []DatastoreParams)`

SetDataStoreParams sets DataStoreParams field to given value.

### HasDataStoreParams

`func (o *EsxiRegistrationParamsAllOf) HasDataStoreParams() bool`

HasDataStoreParams returns a boolean if a field has been set.

### SetDataStoreParamsNil

`func (o *EsxiRegistrationParamsAllOf) SetDataStoreParamsNil(b bool)`

 SetDataStoreParamsNil sets the value for DataStoreParams to be an explicit nil

### UnsetDataStoreParams
`func (o *EsxiRegistrationParamsAllOf) UnsetDataStoreParams()`

UnsetDataStoreParams ensures that no value is present for DataStoreParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


