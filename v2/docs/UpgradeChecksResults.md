# UpgradeChecksResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinishTimeSecs** | Pointer to **NullableInt64** | Specifies unix epoch finish time of checks(in seconds). | [optional] 
**NodeResults** | Pointer to [**[]UpgradeCheckNodeResult**](UpgradeCheckNodeResult.md) | The healthcheck result for node | [optional] 
**RequestType** | Pointer to **string** | type of checks(preupgrade/postupgrade) | [optional] 
**ResultStatus** | Pointer to **string** | final result (running/pass/fail) of run | [optional] 
**StartTimeSecs** | Pointer to **NullableInt64** | Specifies unix epoch start time of checks(in seconds). | [optional] 
**TestRunInstanceId** | Pointer to **string** | Specifies test run instance of upgrade checks | [optional] 

## Methods

### NewUpgradeChecksResults

`func NewUpgradeChecksResults() *UpgradeChecksResults`

NewUpgradeChecksResults instantiates a new UpgradeChecksResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeChecksResultsWithDefaults

`func NewUpgradeChecksResultsWithDefaults() *UpgradeChecksResults`

NewUpgradeChecksResultsWithDefaults instantiates a new UpgradeChecksResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinishTimeSecs

`func (o *UpgradeChecksResults) GetFinishTimeSecs() int64`

GetFinishTimeSecs returns the FinishTimeSecs field if non-nil, zero value otherwise.

### GetFinishTimeSecsOk

`func (o *UpgradeChecksResults) GetFinishTimeSecsOk() (*int64, bool)`

GetFinishTimeSecsOk returns a tuple with the FinishTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTimeSecs

`func (o *UpgradeChecksResults) SetFinishTimeSecs(v int64)`

SetFinishTimeSecs sets FinishTimeSecs field to given value.

### HasFinishTimeSecs

`func (o *UpgradeChecksResults) HasFinishTimeSecs() bool`

HasFinishTimeSecs returns a boolean if a field has been set.

### SetFinishTimeSecsNil

`func (o *UpgradeChecksResults) SetFinishTimeSecsNil(b bool)`

 SetFinishTimeSecsNil sets the value for FinishTimeSecs to be an explicit nil

### UnsetFinishTimeSecs
`func (o *UpgradeChecksResults) UnsetFinishTimeSecs()`

UnsetFinishTimeSecs ensures that no value is present for FinishTimeSecs, not even an explicit nil
### GetNodeResults

`func (o *UpgradeChecksResults) GetNodeResults() []UpgradeCheckNodeResult`

GetNodeResults returns the NodeResults field if non-nil, zero value otherwise.

### GetNodeResultsOk

`func (o *UpgradeChecksResults) GetNodeResultsOk() (*[]UpgradeCheckNodeResult, bool)`

GetNodeResultsOk returns a tuple with the NodeResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeResults

`func (o *UpgradeChecksResults) SetNodeResults(v []UpgradeCheckNodeResult)`

SetNodeResults sets NodeResults field to given value.

### HasNodeResults

`func (o *UpgradeChecksResults) HasNodeResults() bool`

HasNodeResults returns a boolean if a field has been set.

### GetRequestType

`func (o *UpgradeChecksResults) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *UpgradeChecksResults) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *UpgradeChecksResults) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *UpgradeChecksResults) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetResultStatus

`func (o *UpgradeChecksResults) GetResultStatus() string`

GetResultStatus returns the ResultStatus field if non-nil, zero value otherwise.

### GetResultStatusOk

`func (o *UpgradeChecksResults) GetResultStatusOk() (*string, bool)`

GetResultStatusOk returns a tuple with the ResultStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultStatus

`func (o *UpgradeChecksResults) SetResultStatus(v string)`

SetResultStatus sets ResultStatus field to given value.

### HasResultStatus

`func (o *UpgradeChecksResults) HasResultStatus() bool`

HasResultStatus returns a boolean if a field has been set.

### GetStartTimeSecs

`func (o *UpgradeChecksResults) GetStartTimeSecs() int64`

GetStartTimeSecs returns the StartTimeSecs field if non-nil, zero value otherwise.

### GetStartTimeSecsOk

`func (o *UpgradeChecksResults) GetStartTimeSecsOk() (*int64, bool)`

GetStartTimeSecsOk returns a tuple with the StartTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeSecs

`func (o *UpgradeChecksResults) SetStartTimeSecs(v int64)`

SetStartTimeSecs sets StartTimeSecs field to given value.

### HasStartTimeSecs

`func (o *UpgradeChecksResults) HasStartTimeSecs() bool`

HasStartTimeSecs returns a boolean if a field has been set.

### SetStartTimeSecsNil

`func (o *UpgradeChecksResults) SetStartTimeSecsNil(b bool)`

 SetStartTimeSecsNil sets the value for StartTimeSecs to be an explicit nil

### UnsetStartTimeSecs
`func (o *UpgradeChecksResults) UnsetStartTimeSecs()`

UnsetStartTimeSecs ensures that no value is present for StartTimeSecs, not even an explicit nil
### GetTestRunInstanceId

`func (o *UpgradeChecksResults) GetTestRunInstanceId() string`

GetTestRunInstanceId returns the TestRunInstanceId field if non-nil, zero value otherwise.

### GetTestRunInstanceIdOk

`func (o *UpgradeChecksResults) GetTestRunInstanceIdOk() (*string, bool)`

GetTestRunInstanceIdOk returns a tuple with the TestRunInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunInstanceId

`func (o *UpgradeChecksResults) SetTestRunInstanceId(v string)`

SetTestRunInstanceId sets TestRunInstanceId field to given value.

### HasTestRunInstanceId

`func (o *UpgradeChecksResults) HasTestRunInstanceId() bool`

HasTestRunInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


