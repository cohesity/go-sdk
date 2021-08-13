# RecoverCassandraParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshots** | [**[]RecoverCassandraSnapshotParams**](RecoverCassandraSnapshotParams.md) | Specifies the local snapshot ids and other details of the Objects to be recovered. | 
**Suffix** | Pointer to **NullableString** | A suffix that is to be applied to all recovered objects. | [optional] 
**SelectedDataCenters** | Pointer to **[]string** | Selected Data centers for this cluster. | [optional] 
**StagingDirectoryList** | Pointer to **[]string** | Specifies the directory on the primary to copy the files which are to be uploaded using destination sstableloader. | [optional] 

## Methods

### NewRecoverCassandraParamsAllOf

`func NewRecoverCassandraParamsAllOf(snapshots []RecoverCassandraSnapshotParams, ) *RecoverCassandraParamsAllOf`

NewRecoverCassandraParamsAllOf instantiates a new RecoverCassandraParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverCassandraParamsAllOfWithDefaults

`func NewRecoverCassandraParamsAllOfWithDefaults() *RecoverCassandraParamsAllOf`

NewRecoverCassandraParamsAllOfWithDefaults instantiates a new RecoverCassandraParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshots

`func (o *RecoverCassandraParamsAllOf) GetSnapshots() []RecoverCassandraSnapshotParams`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *RecoverCassandraParamsAllOf) GetSnapshotsOk() (*[]RecoverCassandraSnapshotParams, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *RecoverCassandraParamsAllOf) SetSnapshots(v []RecoverCassandraSnapshotParams)`

SetSnapshots sets Snapshots field to given value.


### SetSnapshotsNil

`func (o *RecoverCassandraParamsAllOf) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *RecoverCassandraParamsAllOf) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil
### GetSuffix

`func (o *RecoverCassandraParamsAllOf) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *RecoverCassandraParamsAllOf) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *RecoverCassandraParamsAllOf) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *RecoverCassandraParamsAllOf) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *RecoverCassandraParamsAllOf) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *RecoverCassandraParamsAllOf) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetSelectedDataCenters

`func (o *RecoverCassandraParamsAllOf) GetSelectedDataCenters() []string`

GetSelectedDataCenters returns the SelectedDataCenters field if non-nil, zero value otherwise.

### GetSelectedDataCentersOk

`func (o *RecoverCassandraParamsAllOf) GetSelectedDataCentersOk() (*[]string, bool)`

GetSelectedDataCentersOk returns a tuple with the SelectedDataCenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDataCenters

`func (o *RecoverCassandraParamsAllOf) SetSelectedDataCenters(v []string)`

SetSelectedDataCenters sets SelectedDataCenters field to given value.

### HasSelectedDataCenters

`func (o *RecoverCassandraParamsAllOf) HasSelectedDataCenters() bool`

HasSelectedDataCenters returns a boolean if a field has been set.

### GetStagingDirectoryList

`func (o *RecoverCassandraParamsAllOf) GetStagingDirectoryList() []string`

GetStagingDirectoryList returns the StagingDirectoryList field if non-nil, zero value otherwise.

### GetStagingDirectoryListOk

`func (o *RecoverCassandraParamsAllOf) GetStagingDirectoryListOk() (*[]string, bool)`

GetStagingDirectoryListOk returns a tuple with the StagingDirectoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingDirectoryList

`func (o *RecoverCassandraParamsAllOf) SetStagingDirectoryList(v []string)`

SetStagingDirectoryList sets StagingDirectoryList field to given value.

### HasStagingDirectoryList

`func (o *RecoverCassandraParamsAllOf) HasStagingDirectoryList() bool`

HasStagingDirectoryList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


