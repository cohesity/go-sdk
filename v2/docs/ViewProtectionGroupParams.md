# ViewProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternallyTriggeredJobParams** | Pointer to [**ExternallyTriggeredJobParams**](ExternallyTriggeredJobParams.md) |  | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**Objects** | [**[]ViewProtectionGroupObjectParams**](ViewProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | 
**ReplicationParams** | Pointer to [**NullableViewProtectionGroupReplicationParams**](ViewProtectionGroupReplicationParams.md) |  | [optional] 

## Methods

### NewViewProtectionGroupParams

`func NewViewProtectionGroupParams(objects []ViewProtectionGroupObjectParams, ) *ViewProtectionGroupParams`

NewViewProtectionGroupParams instantiates a new ViewProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProtectionGroupParamsWithDefaults

`func NewViewProtectionGroupParamsWithDefaults() *ViewProtectionGroupParams`

NewViewProtectionGroupParamsWithDefaults instantiates a new ViewProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternallyTriggeredJobParams

`func (o *ViewProtectionGroupParams) GetExternallyTriggeredJobParams() ExternallyTriggeredJobParams`

GetExternallyTriggeredJobParams returns the ExternallyTriggeredJobParams field if non-nil, zero value otherwise.

### GetExternallyTriggeredJobParamsOk

`func (o *ViewProtectionGroupParams) GetExternallyTriggeredJobParamsOk() (*ExternallyTriggeredJobParams, bool)`

GetExternallyTriggeredJobParamsOk returns a tuple with the ExternallyTriggeredJobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyTriggeredJobParams

`func (o *ViewProtectionGroupParams) SetExternallyTriggeredJobParams(v ExternallyTriggeredJobParams)`

SetExternallyTriggeredJobParams sets ExternallyTriggeredJobParams field to given value.

### HasExternallyTriggeredJobParams

`func (o *ViewProtectionGroupParams) HasExternallyTriggeredJobParams() bool`

HasExternallyTriggeredJobParams returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *ViewProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *ViewProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *ViewProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *ViewProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetObjects

`func (o *ViewProtectionGroupParams) GetObjects() []ViewProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ViewProtectionGroupParams) GetObjectsOk() (*[]ViewProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ViewProtectionGroupParams) SetObjects(v []ViewProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### GetReplicationParams

`func (o *ViewProtectionGroupParams) GetReplicationParams() ViewProtectionGroupReplicationParams`

GetReplicationParams returns the ReplicationParams field if non-nil, zero value otherwise.

### GetReplicationParamsOk

`func (o *ViewProtectionGroupParams) GetReplicationParamsOk() (*ViewProtectionGroupReplicationParams, bool)`

GetReplicationParamsOk returns a tuple with the ReplicationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationParams

`func (o *ViewProtectionGroupParams) SetReplicationParams(v ViewProtectionGroupReplicationParams)`

SetReplicationParams sets ReplicationParams field to given value.

### HasReplicationParams

`func (o *ViewProtectionGroupParams) HasReplicationParams() bool`

HasReplicationParams returns a boolean if a field has been set.

### SetReplicationParamsNil

`func (o *ViewProtectionGroupParams) SetReplicationParamsNil(b bool)`

 SetReplicationParamsNil sets the value for ReplicationParams to be an explicit nil

### UnsetReplicationParams
`func (o *ViewProtectionGroupParams) UnsetReplicationParams()`

UnsetReplicationParams ensures that no value is present for ReplicationParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


