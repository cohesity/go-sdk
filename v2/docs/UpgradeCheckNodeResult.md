# UpgradeCheckNodeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeIp** | Pointer to **string** | The node ip | [optional] 
**NodeTestResults** | Pointer to [**[]UpgradeCheckTestResult**](UpgradeCheckTestResult.md) | The healthcheck test results for node | [optional] 
**NodeTestStatus** | Pointer to **string** | The healthcheck run status for node | [optional] 

## Methods

### NewUpgradeCheckNodeResult

`func NewUpgradeCheckNodeResult() *UpgradeCheckNodeResult`

NewUpgradeCheckNodeResult instantiates a new UpgradeCheckNodeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeCheckNodeResultWithDefaults

`func NewUpgradeCheckNodeResultWithDefaults() *UpgradeCheckNodeResult`

NewUpgradeCheckNodeResultWithDefaults instantiates a new UpgradeCheckNodeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeIp

`func (o *UpgradeCheckNodeResult) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *UpgradeCheckNodeResult) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *UpgradeCheckNodeResult) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *UpgradeCheckNodeResult) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### GetNodeTestResults

`func (o *UpgradeCheckNodeResult) GetNodeTestResults() []UpgradeCheckTestResult`

GetNodeTestResults returns the NodeTestResults field if non-nil, zero value otherwise.

### GetNodeTestResultsOk

`func (o *UpgradeCheckNodeResult) GetNodeTestResultsOk() (*[]UpgradeCheckTestResult, bool)`

GetNodeTestResultsOk returns a tuple with the NodeTestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeTestResults

`func (o *UpgradeCheckNodeResult) SetNodeTestResults(v []UpgradeCheckTestResult)`

SetNodeTestResults sets NodeTestResults field to given value.

### HasNodeTestResults

`func (o *UpgradeCheckNodeResult) HasNodeTestResults() bool`

HasNodeTestResults returns a boolean if a field has been set.

### GetNodeTestStatus

`func (o *UpgradeCheckNodeResult) GetNodeTestStatus() string`

GetNodeTestStatus returns the NodeTestStatus field if non-nil, zero value otherwise.

### GetNodeTestStatusOk

`func (o *UpgradeCheckNodeResult) GetNodeTestStatusOk() (*string, bool)`

GetNodeTestStatusOk returns a tuple with the NodeTestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeTestStatus

`func (o *UpgradeCheckNodeResult) SetNodeTestStatus(v string)`

SetNodeTestStatus sets NodeTestStatus field to given value.

### HasNodeTestStatus

`func (o *UpgradeCheckNodeResult) HasNodeTestStatus() bool`

HasNodeTestStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


