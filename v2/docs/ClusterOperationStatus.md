# ClusterOperationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]ClusterOperationAttribute**](ClusterOperationAttribute.md) | Specifies the attributes of the operation to provide more context for the operation. For example: Use name &#39;kPackageNameAttribute&#39; and value &#39;7.0.1-p1-2023Jul04-cc6d7c5f&#39; to indicate package being installed when operation type involves installation of package such as &#39;Upgrade&#39; or &#39;Patch&#39;. Attribute list will differ based on the type of operation.  | [optional] 
**ClusterId** | Pointer to **int64** | Specifies the ID of the cluster. | [optional] 
**ClusterIncarnationId** | Pointer to **int64** | Specifies the incarnation id of the cluster. | [optional] 
**ErrorMessage** | Pointer to **string** | Specifies the error message for the operation. | [optional] 
**Events** | Pointer to [**[]OperationEvents**](OperationEvents.md) | Specifies the list of events that took place during the operation.  | [optional] 
**FinishTimeSeconds** | Pointer to **NullableInt64** | Specifies unix epoch finish time of operation. | [optional] 
**NodesAssessmentResults** | Pointer to [**[]NodeAssessmentResults**](NodeAssessmentResults.md) | Specifies the result of running assessment on the cluster nodes. | [optional] 
**NodesOperationStatus** | Pointer to [**[]ClusterNodeOperationStatus**](ClusterNodeOperationStatus.md) | Specifies the status of operation on the cluster nodes. | [optional] 
**OperationId** | Pointer to **string** | Specifies the operation Id of cluster operation.  | [optional] 
**OperationType** | Pointer to **string** | Specifies the type of cluster operation. * &#x60;Destroy&#x60; indicates cluster destroy operation. * &#x60;Create&#x60; indicates cluster create operation. * &#x60;NodeAddition&#x60; indicates the operation to add nodes to the cluster. * &#x60;NodeRemoval&#x60; indicates a node removal operation. *  Operation types related to software update are detailed in   [UpdateClusterSoftware](#tag/Platform/operation/UpdateClusterSoftware).  | [optional] 
**Percentage** | Pointer to **int32** | Specifies an approximate completion percentage for the operation.  | [optional] 
**StartTimeSeconds** | Pointer to **NullableInt64** | Specifies unix epoch start time of operation. | [optional] 
**Status** | Pointer to **string** | Specifies the status of the operation. * &#39;Success&#39; indicates the operation is successful. * &#39;Failed&#39; indicates the operation failed due to an error. * &#39;InProgress&#39; indicates the operation is in progress.  | [optional] 
**TimeRemainingSeconds** | Pointer to **int64** | Specifies an estimated number of seconds until the operation is complete.  | [optional] 

## Methods

### NewClusterOperationStatus

`func NewClusterOperationStatus() *ClusterOperationStatus`

NewClusterOperationStatus instantiates a new ClusterOperationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOperationStatusWithDefaults

`func NewClusterOperationStatusWithDefaults() *ClusterOperationStatus`

NewClusterOperationStatusWithDefaults instantiates a new ClusterOperationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ClusterOperationStatus) GetAttributes() []ClusterOperationAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ClusterOperationStatus) GetAttributesOk() (*[]ClusterOperationAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ClusterOperationStatus) SetAttributes(v []ClusterOperationAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ClusterOperationStatus) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetClusterId

`func (o *ClusterOperationStatus) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterOperationStatus) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterOperationStatus) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ClusterOperationStatus) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterIncarnationId

`func (o *ClusterOperationStatus) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *ClusterOperationStatus) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *ClusterOperationStatus) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *ClusterOperationStatus) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ClusterOperationStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ClusterOperationStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ClusterOperationStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ClusterOperationStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetEvents

`func (o *ClusterOperationStatus) GetEvents() []OperationEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ClusterOperationStatus) GetEventsOk() (*[]OperationEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ClusterOperationStatus) SetEvents(v []OperationEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ClusterOperationStatus) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetFinishTimeSeconds

`func (o *ClusterOperationStatus) GetFinishTimeSeconds() int64`

GetFinishTimeSeconds returns the FinishTimeSeconds field if non-nil, zero value otherwise.

### GetFinishTimeSecondsOk

`func (o *ClusterOperationStatus) GetFinishTimeSecondsOk() (*int64, bool)`

GetFinishTimeSecondsOk returns a tuple with the FinishTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTimeSeconds

`func (o *ClusterOperationStatus) SetFinishTimeSeconds(v int64)`

SetFinishTimeSeconds sets FinishTimeSeconds field to given value.

### HasFinishTimeSeconds

`func (o *ClusterOperationStatus) HasFinishTimeSeconds() bool`

HasFinishTimeSeconds returns a boolean if a field has been set.

### SetFinishTimeSecondsNil

`func (o *ClusterOperationStatus) SetFinishTimeSecondsNil(b bool)`

 SetFinishTimeSecondsNil sets the value for FinishTimeSeconds to be an explicit nil

### UnsetFinishTimeSeconds
`func (o *ClusterOperationStatus) UnsetFinishTimeSeconds()`

UnsetFinishTimeSeconds ensures that no value is present for FinishTimeSeconds, not even an explicit nil
### GetNodesAssessmentResults

`func (o *ClusterOperationStatus) GetNodesAssessmentResults() []NodeAssessmentResults`

GetNodesAssessmentResults returns the NodesAssessmentResults field if non-nil, zero value otherwise.

### GetNodesAssessmentResultsOk

`func (o *ClusterOperationStatus) GetNodesAssessmentResultsOk() (*[]NodeAssessmentResults, bool)`

GetNodesAssessmentResultsOk returns a tuple with the NodesAssessmentResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesAssessmentResults

`func (o *ClusterOperationStatus) SetNodesAssessmentResults(v []NodeAssessmentResults)`

SetNodesAssessmentResults sets NodesAssessmentResults field to given value.

### HasNodesAssessmentResults

`func (o *ClusterOperationStatus) HasNodesAssessmentResults() bool`

HasNodesAssessmentResults returns a boolean if a field has been set.

### GetNodesOperationStatus

`func (o *ClusterOperationStatus) GetNodesOperationStatus() []ClusterNodeOperationStatus`

GetNodesOperationStatus returns the NodesOperationStatus field if non-nil, zero value otherwise.

### GetNodesOperationStatusOk

`func (o *ClusterOperationStatus) GetNodesOperationStatusOk() (*[]ClusterNodeOperationStatus, bool)`

GetNodesOperationStatusOk returns a tuple with the NodesOperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesOperationStatus

`func (o *ClusterOperationStatus) SetNodesOperationStatus(v []ClusterNodeOperationStatus)`

SetNodesOperationStatus sets NodesOperationStatus field to given value.

### HasNodesOperationStatus

`func (o *ClusterOperationStatus) HasNodesOperationStatus() bool`

HasNodesOperationStatus returns a boolean if a field has been set.

### GetOperationId

`func (o *ClusterOperationStatus) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *ClusterOperationStatus) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *ClusterOperationStatus) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *ClusterOperationStatus) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetOperationType

`func (o *ClusterOperationStatus) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ClusterOperationStatus) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ClusterOperationStatus) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *ClusterOperationStatus) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetPercentage

`func (o *ClusterOperationStatus) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ClusterOperationStatus) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ClusterOperationStatus) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ClusterOperationStatus) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetStartTimeSeconds

`func (o *ClusterOperationStatus) GetStartTimeSeconds() int64`

GetStartTimeSeconds returns the StartTimeSeconds field if non-nil, zero value otherwise.

### GetStartTimeSecondsOk

`func (o *ClusterOperationStatus) GetStartTimeSecondsOk() (*int64, bool)`

GetStartTimeSecondsOk returns a tuple with the StartTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeSeconds

`func (o *ClusterOperationStatus) SetStartTimeSeconds(v int64)`

SetStartTimeSeconds sets StartTimeSeconds field to given value.

### HasStartTimeSeconds

`func (o *ClusterOperationStatus) HasStartTimeSeconds() bool`

HasStartTimeSeconds returns a boolean if a field has been set.

### SetStartTimeSecondsNil

`func (o *ClusterOperationStatus) SetStartTimeSecondsNil(b bool)`

 SetStartTimeSecondsNil sets the value for StartTimeSeconds to be an explicit nil

### UnsetStartTimeSeconds
`func (o *ClusterOperationStatus) UnsetStartTimeSeconds()`

UnsetStartTimeSeconds ensures that no value is present for StartTimeSeconds, not even an explicit nil
### GetStatus

`func (o *ClusterOperationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterOperationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterOperationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterOperationStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeRemainingSeconds

`func (o *ClusterOperationStatus) GetTimeRemainingSeconds() int64`

GetTimeRemainingSeconds returns the TimeRemainingSeconds field if non-nil, zero value otherwise.

### GetTimeRemainingSecondsOk

`func (o *ClusterOperationStatus) GetTimeRemainingSecondsOk() (*int64, bool)`

GetTimeRemainingSecondsOk returns a tuple with the TimeRemainingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRemainingSeconds

`func (o *ClusterOperationStatus) SetTimeRemainingSeconds(v int64)`

SetTimeRemainingSeconds sets TimeRemainingSeconds field to given value.

### HasTimeRemainingSeconds

`func (o *ClusterOperationStatus) HasTimeRemainingSeconds() bool`

HasTimeRemainingSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


