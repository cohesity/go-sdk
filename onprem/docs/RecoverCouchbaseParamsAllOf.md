# RecoverCouchbaseParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshots** | [**[]RecoverCouchbaseSnapshotParams**](RecoverCouchbaseSnapshotParams.md) | Specifies the local snapshot ids of the Objects to be recovered. | 
**Suffix** | Pointer to **NullableString** | A suffix that is to be applied to all recovered objects. | [optional] 
**FilterDocumentsParams** | [**FilterDocumentsParams**](FilterDocumentsParams.md) |  | 
**AppendDocuments** | Pointer to **NullableBool** | If set to true, docuements from the bucket being recovered will be appended into the bucket at the destination. | [optional] 
**DdlOnlyRecovery** | Pointer to **NullableBool** | Set to true to recover only the bucket configurations. No documents will be recovered. | [optional] 
**OverwriteUsers** | Pointer to **NullableBool** | If set to true existing users will be replaced with users from the bucket being recovered. | [optional] 

## Methods

### NewRecoverCouchbaseParamsAllOf

`func NewRecoverCouchbaseParamsAllOf(snapshots []RecoverCouchbaseSnapshotParams, filterDocumentsParams FilterDocumentsParams, ) *RecoverCouchbaseParamsAllOf`

NewRecoverCouchbaseParamsAllOf instantiates a new RecoverCouchbaseParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverCouchbaseParamsAllOfWithDefaults

`func NewRecoverCouchbaseParamsAllOfWithDefaults() *RecoverCouchbaseParamsAllOf`

NewRecoverCouchbaseParamsAllOfWithDefaults instantiates a new RecoverCouchbaseParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshots

`func (o *RecoverCouchbaseParamsAllOf) GetSnapshots() []RecoverCouchbaseSnapshotParams`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *RecoverCouchbaseParamsAllOf) GetSnapshotsOk() (*[]RecoverCouchbaseSnapshotParams, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *RecoverCouchbaseParamsAllOf) SetSnapshots(v []RecoverCouchbaseSnapshotParams)`

SetSnapshots sets Snapshots field to given value.


### SetSnapshotsNil

`func (o *RecoverCouchbaseParamsAllOf) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *RecoverCouchbaseParamsAllOf) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil
### GetSuffix

`func (o *RecoverCouchbaseParamsAllOf) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *RecoverCouchbaseParamsAllOf) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *RecoverCouchbaseParamsAllOf) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *RecoverCouchbaseParamsAllOf) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *RecoverCouchbaseParamsAllOf) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *RecoverCouchbaseParamsAllOf) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetFilterDocumentsParams

`func (o *RecoverCouchbaseParamsAllOf) GetFilterDocumentsParams() FilterDocumentsParams`

GetFilterDocumentsParams returns the FilterDocumentsParams field if non-nil, zero value otherwise.

### GetFilterDocumentsParamsOk

`func (o *RecoverCouchbaseParamsAllOf) GetFilterDocumentsParamsOk() (*FilterDocumentsParams, bool)`

GetFilterDocumentsParamsOk returns a tuple with the FilterDocumentsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterDocumentsParams

`func (o *RecoverCouchbaseParamsAllOf) SetFilterDocumentsParams(v FilterDocumentsParams)`

SetFilterDocumentsParams sets FilterDocumentsParams field to given value.


### GetAppendDocuments

`func (o *RecoverCouchbaseParamsAllOf) GetAppendDocuments() bool`

GetAppendDocuments returns the AppendDocuments field if non-nil, zero value otherwise.

### GetAppendDocumentsOk

`func (o *RecoverCouchbaseParamsAllOf) GetAppendDocumentsOk() (*bool, bool)`

GetAppendDocumentsOk returns a tuple with the AppendDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendDocuments

`func (o *RecoverCouchbaseParamsAllOf) SetAppendDocuments(v bool)`

SetAppendDocuments sets AppendDocuments field to given value.

### HasAppendDocuments

`func (o *RecoverCouchbaseParamsAllOf) HasAppendDocuments() bool`

HasAppendDocuments returns a boolean if a field has been set.

### SetAppendDocumentsNil

`func (o *RecoverCouchbaseParamsAllOf) SetAppendDocumentsNil(b bool)`

 SetAppendDocumentsNil sets the value for AppendDocuments to be an explicit nil

### UnsetAppendDocuments
`func (o *RecoverCouchbaseParamsAllOf) UnsetAppendDocuments()`

UnsetAppendDocuments ensures that no value is present for AppendDocuments, not even an explicit nil
### GetDdlOnlyRecovery

`func (o *RecoverCouchbaseParamsAllOf) GetDdlOnlyRecovery() bool`

GetDdlOnlyRecovery returns the DdlOnlyRecovery field if non-nil, zero value otherwise.

### GetDdlOnlyRecoveryOk

`func (o *RecoverCouchbaseParamsAllOf) GetDdlOnlyRecoveryOk() (*bool, bool)`

GetDdlOnlyRecoveryOk returns a tuple with the DdlOnlyRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdlOnlyRecovery

`func (o *RecoverCouchbaseParamsAllOf) SetDdlOnlyRecovery(v bool)`

SetDdlOnlyRecovery sets DdlOnlyRecovery field to given value.

### HasDdlOnlyRecovery

`func (o *RecoverCouchbaseParamsAllOf) HasDdlOnlyRecovery() bool`

HasDdlOnlyRecovery returns a boolean if a field has been set.

### SetDdlOnlyRecoveryNil

`func (o *RecoverCouchbaseParamsAllOf) SetDdlOnlyRecoveryNil(b bool)`

 SetDdlOnlyRecoveryNil sets the value for DdlOnlyRecovery to be an explicit nil

### UnsetDdlOnlyRecovery
`func (o *RecoverCouchbaseParamsAllOf) UnsetDdlOnlyRecovery()`

UnsetDdlOnlyRecovery ensures that no value is present for DdlOnlyRecovery, not even an explicit nil
### GetOverwriteUsers

`func (o *RecoverCouchbaseParamsAllOf) GetOverwriteUsers() bool`

GetOverwriteUsers returns the OverwriteUsers field if non-nil, zero value otherwise.

### GetOverwriteUsersOk

`func (o *RecoverCouchbaseParamsAllOf) GetOverwriteUsersOk() (*bool, bool)`

GetOverwriteUsersOk returns a tuple with the OverwriteUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteUsers

`func (o *RecoverCouchbaseParamsAllOf) SetOverwriteUsers(v bool)`

SetOverwriteUsers sets OverwriteUsers field to given value.

### HasOverwriteUsers

`func (o *RecoverCouchbaseParamsAllOf) HasOverwriteUsers() bool`

HasOverwriteUsers returns a boolean if a field has been set.

### SetOverwriteUsersNil

`func (o *RecoverCouchbaseParamsAllOf) SetOverwriteUsersNil(b bool)`

 SetOverwriteUsersNil sets the value for OverwriteUsers to be an explicit nil

### UnsetOverwriteUsers
`func (o *RecoverCouchbaseParamsAllOf) UnsetOverwriteUsers()`

UnsetOverwriteUsers ensures that no value is present for OverwriteUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


