# DataTieringAnalysisGroupRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the ID of the data tiering analysis group run. | [optional] 
**Objects** | Pointer to [**[]DataTieringObjectInfo**](DataTieringObjectInfo.md) | Specifies the objects details analyzed during data tiering analysis group run. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of analysis run in Unix epoch Timestamp(in microseconds). | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the end time of analysis run in Unix epoch Timestamp(in microseconds). | [optional] 
**AnalysisInfo** | Pointer to [**DataTieringAnalysisInfo**](DataTieringAnalysisInfo.md) |  | [optional] 

## Methods

### NewDataTieringAnalysisGroupRun

`func NewDataTieringAnalysisGroupRun() *DataTieringAnalysisGroupRun`

NewDataTieringAnalysisGroupRun instantiates a new DataTieringAnalysisGroupRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringAnalysisGroupRunWithDefaults

`func NewDataTieringAnalysisGroupRunWithDefaults() *DataTieringAnalysisGroupRun`

NewDataTieringAnalysisGroupRunWithDefaults instantiates a new DataTieringAnalysisGroupRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataTieringAnalysisGroupRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataTieringAnalysisGroupRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataTieringAnalysisGroupRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataTieringAnalysisGroupRun) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DataTieringAnalysisGroupRun) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DataTieringAnalysisGroupRun) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjects

`func (o *DataTieringAnalysisGroupRun) GetObjects() []DataTieringObjectInfo`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *DataTieringAnalysisGroupRun) GetObjectsOk() (*[]DataTieringObjectInfo, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *DataTieringAnalysisGroupRun) SetObjects(v []DataTieringObjectInfo)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *DataTieringAnalysisGroupRun) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *DataTieringAnalysisGroupRun) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *DataTieringAnalysisGroupRun) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetStartTimeUsecs

`func (o *DataTieringAnalysisGroupRun) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *DataTieringAnalysisGroupRun) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *DataTieringAnalysisGroupRun) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *DataTieringAnalysisGroupRun) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *DataTieringAnalysisGroupRun) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *DataTieringAnalysisGroupRun) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetEndTimeUsecs

`func (o *DataTieringAnalysisGroupRun) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *DataTieringAnalysisGroupRun) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *DataTieringAnalysisGroupRun) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *DataTieringAnalysisGroupRun) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *DataTieringAnalysisGroupRun) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *DataTieringAnalysisGroupRun) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetAnalysisInfo

`func (o *DataTieringAnalysisGroupRun) GetAnalysisInfo() DataTieringAnalysisInfo`

GetAnalysisInfo returns the AnalysisInfo field if non-nil, zero value otherwise.

### GetAnalysisInfoOk

`func (o *DataTieringAnalysisGroupRun) GetAnalysisInfoOk() (*DataTieringAnalysisInfo, bool)`

GetAnalysisInfoOk returns a tuple with the AnalysisInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisInfo

`func (o *DataTieringAnalysisGroupRun) SetAnalysisInfo(v DataTieringAnalysisInfo)`

SetAnalysisInfo sets AnalysisInfo field to given value.

### HasAnalysisInfo

`func (o *DataTieringAnalysisGroupRun) HasAnalysisInfo() bool`

HasAnalysisInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


