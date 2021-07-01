# ProtectionSourcesJobRunsReportElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionSource** | Pointer to [**NullableProtectionSource**](ProtectionSource.md) | Specifies the leaf Protection Source Object such as a VM. | [optional] 
**SnapshotsInfo** | Pointer to [**[]ProtectionSourceSnapshotInformation**](ProtectionSourceSnapshotInformation.md) | Array of Snapshots  Specifies the Snapshots that contain backups of the Protection Source Object. | [optional] 

## Methods

### NewProtectionSourcesJobRunsReportElement

`func NewProtectionSourcesJobRunsReportElement() *ProtectionSourcesJobRunsReportElement`

NewProtectionSourcesJobRunsReportElement instantiates a new ProtectionSourcesJobRunsReportElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionSourcesJobRunsReportElementWithDefaults

`func NewProtectionSourcesJobRunsReportElementWithDefaults() *ProtectionSourcesJobRunsReportElement`

NewProtectionSourcesJobRunsReportElementWithDefaults instantiates a new ProtectionSourcesJobRunsReportElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionSource

`func (o *ProtectionSourcesJobRunsReportElement) GetProtectionSource() ProtectionSource`

GetProtectionSource returns the ProtectionSource field if non-nil, zero value otherwise.

### GetProtectionSourceOk

`func (o *ProtectionSourcesJobRunsReportElement) GetProtectionSourceOk() (*ProtectionSource, bool)`

GetProtectionSourceOk returns a tuple with the ProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSource

`func (o *ProtectionSourcesJobRunsReportElement) SetProtectionSource(v ProtectionSource)`

SetProtectionSource sets ProtectionSource field to given value.

### HasProtectionSource

`func (o *ProtectionSourcesJobRunsReportElement) HasProtectionSource() bool`

HasProtectionSource returns a boolean if a field has been set.

### SetProtectionSourceNil

`func (o *ProtectionSourcesJobRunsReportElement) SetProtectionSourceNil(b bool)`

 SetProtectionSourceNil sets the value for ProtectionSource to be an explicit nil

### UnsetProtectionSource
`func (o *ProtectionSourcesJobRunsReportElement) UnsetProtectionSource()`

UnsetProtectionSource ensures that no value is present for ProtectionSource, not even an explicit nil
### GetSnapshotsInfo

`func (o *ProtectionSourcesJobRunsReportElement) GetSnapshotsInfo() []ProtectionSourceSnapshotInformation`

GetSnapshotsInfo returns the SnapshotsInfo field if non-nil, zero value otherwise.

### GetSnapshotsInfoOk

`func (o *ProtectionSourcesJobRunsReportElement) GetSnapshotsInfoOk() (*[]ProtectionSourceSnapshotInformation, bool)`

GetSnapshotsInfoOk returns a tuple with the SnapshotsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotsInfo

`func (o *ProtectionSourcesJobRunsReportElement) SetSnapshotsInfo(v []ProtectionSourceSnapshotInformation)`

SetSnapshotsInfo sets SnapshotsInfo field to given value.

### HasSnapshotsInfo

`func (o *ProtectionSourcesJobRunsReportElement) HasSnapshotsInfo() bool`

HasSnapshotsInfo returns a boolean if a field has been set.

### SetSnapshotsInfoNil

`func (o *ProtectionSourcesJobRunsReportElement) SetSnapshotsInfoNil(b bool)`

 SetSnapshotsInfoNil sets the value for SnapshotsInfo to be an explicit nil

### UnsetSnapshotsInfo
`func (o *ProtectionSourcesJobRunsReportElement) UnsetSnapshotsInfo()`

UnsetSnapshotsInfo ensures that no value is present for SnapshotsInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


