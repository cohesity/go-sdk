# RecoverVolumesTaskStateProtoTaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DstGuid** | Pointer to **NullableString** | Volume GUID for the Target Entity (phy host). | [optional] 
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**ProgressMonitorTaskPath** | Pointer to **NullableString** | The path relative to the root path of the restore task progress monitor. | [optional] 

## Methods

### NewRecoverVolumesTaskStateProtoTaskResult

`func NewRecoverVolumesTaskStateProtoTaskResult() *RecoverVolumesTaskStateProtoTaskResult`

NewRecoverVolumesTaskStateProtoTaskResult instantiates a new RecoverVolumesTaskStateProtoTaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVolumesTaskStateProtoTaskResultWithDefaults

`func NewRecoverVolumesTaskStateProtoTaskResultWithDefaults() *RecoverVolumesTaskStateProtoTaskResult`

NewRecoverVolumesTaskStateProtoTaskResultWithDefaults instantiates a new RecoverVolumesTaskStateProtoTaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDstGuid

`func (o *RecoverVolumesTaskStateProtoTaskResult) GetDstGuid() string`

GetDstGuid returns the DstGuid field if non-nil, zero value otherwise.

### GetDstGuidOk

`func (o *RecoverVolumesTaskStateProtoTaskResult) GetDstGuidOk() (*string, bool)`

GetDstGuidOk returns a tuple with the DstGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstGuid

`func (o *RecoverVolumesTaskStateProtoTaskResult) SetDstGuid(v string)`

SetDstGuid sets DstGuid field to given value.

### HasDstGuid

`func (o *RecoverVolumesTaskStateProtoTaskResult) HasDstGuid() bool`

HasDstGuid returns a boolean if a field has been set.

### SetDstGuidNil

`func (o *RecoverVolumesTaskStateProtoTaskResult) SetDstGuidNil(b bool)`

 SetDstGuidNil sets the value for DstGuid to be an explicit nil

### UnsetDstGuid
`func (o *RecoverVolumesTaskStateProtoTaskResult) UnsetDstGuid()`

UnsetDstGuid ensures that no value is present for DstGuid, not even an explicit nil
### GetError

`func (o *RecoverVolumesTaskStateProtoTaskResult) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RecoverVolumesTaskStateProtoTaskResult) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RecoverVolumesTaskStateProtoTaskResult) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *RecoverVolumesTaskStateProtoTaskResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetProgressMonitorTaskPath

`func (o *RecoverVolumesTaskStateProtoTaskResult) GetProgressMonitorTaskPath() string`

GetProgressMonitorTaskPath returns the ProgressMonitorTaskPath field if non-nil, zero value otherwise.

### GetProgressMonitorTaskPathOk

`func (o *RecoverVolumesTaskStateProtoTaskResult) GetProgressMonitorTaskPathOk() (*string, bool)`

GetProgressMonitorTaskPathOk returns a tuple with the ProgressMonitorTaskPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMonitorTaskPath

`func (o *RecoverVolumesTaskStateProtoTaskResult) SetProgressMonitorTaskPath(v string)`

SetProgressMonitorTaskPath sets ProgressMonitorTaskPath field to given value.

### HasProgressMonitorTaskPath

`func (o *RecoverVolumesTaskStateProtoTaskResult) HasProgressMonitorTaskPath() bool`

HasProgressMonitorTaskPath returns a boolean if a field has been set.

### SetProgressMonitorTaskPathNil

`func (o *RecoverVolumesTaskStateProtoTaskResult) SetProgressMonitorTaskPathNil(b bool)`

 SetProgressMonitorTaskPathNil sets the value for ProgressMonitorTaskPath to be an explicit nil

### UnsetProgressMonitorTaskPath
`func (o *RecoverVolumesTaskStateProtoTaskResult) UnsetProgressMonitorTaskPath()`

UnsetProgressMonitorTaskPath ensures that no value is present for ProgressMonitorTaskPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


