# NodeAssessmentResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the ID of the node. | [optional] 
**Ip** | Pointer to **NullableString** | Specifies the IP address of the node. | [optional] 
**Results** | Pointer to [**[]AssessmentTestResult**](AssessmentTestResult.md) | Specifies the test results for node. | [optional] 
**Status** | Pointer to **string** | Specifies the test run status for node. | [optional] 

## Methods

### NewNodeAssessmentResults

`func NewNodeAssessmentResults() *NodeAssessmentResults`

NewNodeAssessmentResults instantiates a new NodeAssessmentResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeAssessmentResultsWithDefaults

`func NewNodeAssessmentResultsWithDefaults() *NodeAssessmentResults`

NewNodeAssessmentResultsWithDefaults instantiates a new NodeAssessmentResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeAssessmentResults) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeAssessmentResults) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeAssessmentResults) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NodeAssessmentResults) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NodeAssessmentResults) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NodeAssessmentResults) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIp

`func (o *NodeAssessmentResults) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NodeAssessmentResults) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NodeAssessmentResults) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NodeAssessmentResults) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *NodeAssessmentResults) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *NodeAssessmentResults) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetResults

`func (o *NodeAssessmentResults) GetResults() []AssessmentTestResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *NodeAssessmentResults) GetResultsOk() (*[]AssessmentTestResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *NodeAssessmentResults) SetResults(v []AssessmentTestResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *NodeAssessmentResults) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStatus

`func (o *NodeAssessmentResults) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeAssessmentResults) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeAssessmentResults) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeAssessmentResults) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


