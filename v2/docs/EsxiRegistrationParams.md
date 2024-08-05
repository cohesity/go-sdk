# EsxiRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Specifies the password to access target entity. | 
**Username** | **string** | Specifies the username to access target entity. | 
**Description** | Pointer to **NullableString** | Specifies the description of the source being registered. | [optional] 
**Endpoint** | **string** | Specifies the endpoint IPaddress, URL or hostname of the host. | 
**DataStoreParams** | Pointer to [**[]DatastoreParams**](DatastoreParams.md) | Specifies the datastore specific params. | [optional] 
**MaxConcurrentStreams** | Pointer to **NullableInt32** | If this value is &gt; 0 and the number of streams concurrently active on a datastore is equal to it, then any further requests to access the datastore would be denied until the number of active streams reduces. This applies for all the datastores in the specified host. | [optional] 
**MinFreeDatastoreSpaceForBackupGb** | Pointer to **NullableInt64** | Specifies the minimum free space (in GB) expected to be available in the datastore where the virtual disks of the VM being backed up reside. If the space available is lower than the specified value, backup will be aborted. | [optional] 
**MinFreeDatastoreSpaceForBackupPercentage** | Pointer to **NullableInt64** | Specifies the minimum free space (in percentage) expected to be available in the datastore where the virtual disks of the VM being backed up reside. If the space available is lower than the specified value, backup will be aborted. | [optional] 

## Methods

### NewEsxiRegistrationParams

`func NewEsxiRegistrationParams(password string, username string, endpoint string, ) *EsxiRegistrationParams`

NewEsxiRegistrationParams instantiates a new EsxiRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsxiRegistrationParamsWithDefaults

`func NewEsxiRegistrationParamsWithDefaults() *EsxiRegistrationParams`

NewEsxiRegistrationParamsWithDefaults instantiates a new EsxiRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *EsxiRegistrationParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EsxiRegistrationParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EsxiRegistrationParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *EsxiRegistrationParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EsxiRegistrationParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EsxiRegistrationParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDescription

`func (o *EsxiRegistrationParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EsxiRegistrationParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EsxiRegistrationParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EsxiRegistrationParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EsxiRegistrationParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EsxiRegistrationParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndpoint

`func (o *EsxiRegistrationParams) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *EsxiRegistrationParams) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *EsxiRegistrationParams) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetDataStoreParams

`func (o *EsxiRegistrationParams) GetDataStoreParams() []DatastoreParams`

GetDataStoreParams returns the DataStoreParams field if non-nil, zero value otherwise.

### GetDataStoreParamsOk

`func (o *EsxiRegistrationParams) GetDataStoreParamsOk() (*[]DatastoreParams, bool)`

GetDataStoreParamsOk returns a tuple with the DataStoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreParams

`func (o *EsxiRegistrationParams) SetDataStoreParams(v []DatastoreParams)`

SetDataStoreParams sets DataStoreParams field to given value.

### HasDataStoreParams

`func (o *EsxiRegistrationParams) HasDataStoreParams() bool`

HasDataStoreParams returns a boolean if a field has been set.

### SetDataStoreParamsNil

`func (o *EsxiRegistrationParams) SetDataStoreParamsNil(b bool)`

 SetDataStoreParamsNil sets the value for DataStoreParams to be an explicit nil

### UnsetDataStoreParams
`func (o *EsxiRegistrationParams) UnsetDataStoreParams()`

UnsetDataStoreParams ensures that no value is present for DataStoreParams, not even an explicit nil
### GetMaxConcurrentStreams

`func (o *EsxiRegistrationParams) GetMaxConcurrentStreams() int32`

GetMaxConcurrentStreams returns the MaxConcurrentStreams field if non-nil, zero value otherwise.

### GetMaxConcurrentStreamsOk

`func (o *EsxiRegistrationParams) GetMaxConcurrentStreamsOk() (*int32, bool)`

GetMaxConcurrentStreamsOk returns a tuple with the MaxConcurrentStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentStreams

`func (o *EsxiRegistrationParams) SetMaxConcurrentStreams(v int32)`

SetMaxConcurrentStreams sets MaxConcurrentStreams field to given value.

### HasMaxConcurrentStreams

`func (o *EsxiRegistrationParams) HasMaxConcurrentStreams() bool`

HasMaxConcurrentStreams returns a boolean if a field has been set.

### SetMaxConcurrentStreamsNil

`func (o *EsxiRegistrationParams) SetMaxConcurrentStreamsNil(b bool)`

 SetMaxConcurrentStreamsNil sets the value for MaxConcurrentStreams to be an explicit nil

### UnsetMaxConcurrentStreams
`func (o *EsxiRegistrationParams) UnsetMaxConcurrentStreams()`

UnsetMaxConcurrentStreams ensures that no value is present for MaxConcurrentStreams, not even an explicit nil
### GetMinFreeDatastoreSpaceForBackupGb

`func (o *EsxiRegistrationParams) GetMinFreeDatastoreSpaceForBackupGb() int64`

GetMinFreeDatastoreSpaceForBackupGb returns the MinFreeDatastoreSpaceForBackupGb field if non-nil, zero value otherwise.

### GetMinFreeDatastoreSpaceForBackupGbOk

`func (o *EsxiRegistrationParams) GetMinFreeDatastoreSpaceForBackupGbOk() (*int64, bool)`

GetMinFreeDatastoreSpaceForBackupGbOk returns a tuple with the MinFreeDatastoreSpaceForBackupGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFreeDatastoreSpaceForBackupGb

`func (o *EsxiRegistrationParams) SetMinFreeDatastoreSpaceForBackupGb(v int64)`

SetMinFreeDatastoreSpaceForBackupGb sets MinFreeDatastoreSpaceForBackupGb field to given value.

### HasMinFreeDatastoreSpaceForBackupGb

`func (o *EsxiRegistrationParams) HasMinFreeDatastoreSpaceForBackupGb() bool`

HasMinFreeDatastoreSpaceForBackupGb returns a boolean if a field has been set.

### SetMinFreeDatastoreSpaceForBackupGbNil

`func (o *EsxiRegistrationParams) SetMinFreeDatastoreSpaceForBackupGbNil(b bool)`

 SetMinFreeDatastoreSpaceForBackupGbNil sets the value for MinFreeDatastoreSpaceForBackupGb to be an explicit nil

### UnsetMinFreeDatastoreSpaceForBackupGb
`func (o *EsxiRegistrationParams) UnsetMinFreeDatastoreSpaceForBackupGb()`

UnsetMinFreeDatastoreSpaceForBackupGb ensures that no value is present for MinFreeDatastoreSpaceForBackupGb, not even an explicit nil
### GetMinFreeDatastoreSpaceForBackupPercentage

`func (o *EsxiRegistrationParams) GetMinFreeDatastoreSpaceForBackupPercentage() int64`

GetMinFreeDatastoreSpaceForBackupPercentage returns the MinFreeDatastoreSpaceForBackupPercentage field if non-nil, zero value otherwise.

### GetMinFreeDatastoreSpaceForBackupPercentageOk

`func (o *EsxiRegistrationParams) GetMinFreeDatastoreSpaceForBackupPercentageOk() (*int64, bool)`

GetMinFreeDatastoreSpaceForBackupPercentageOk returns a tuple with the MinFreeDatastoreSpaceForBackupPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFreeDatastoreSpaceForBackupPercentage

`func (o *EsxiRegistrationParams) SetMinFreeDatastoreSpaceForBackupPercentage(v int64)`

SetMinFreeDatastoreSpaceForBackupPercentage sets MinFreeDatastoreSpaceForBackupPercentage field to given value.

### HasMinFreeDatastoreSpaceForBackupPercentage

`func (o *EsxiRegistrationParams) HasMinFreeDatastoreSpaceForBackupPercentage() bool`

HasMinFreeDatastoreSpaceForBackupPercentage returns a boolean if a field has been set.

### SetMinFreeDatastoreSpaceForBackupPercentageNil

`func (o *EsxiRegistrationParams) SetMinFreeDatastoreSpaceForBackupPercentageNil(b bool)`

 SetMinFreeDatastoreSpaceForBackupPercentageNil sets the value for MinFreeDatastoreSpaceForBackupPercentage to be an explicit nil

### UnsetMinFreeDatastoreSpaceForBackupPercentage
`func (o *EsxiRegistrationParams) UnsetMinFreeDatastoreSpaceForBackupPercentage()`

UnsetMinFreeDatastoreSpaceForBackupPercentage ensures that no value is present for MinFreeDatastoreSpaceForBackupPercentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


