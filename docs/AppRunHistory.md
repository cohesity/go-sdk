# AppRunHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppInfo** | Pointer to [**MapReduceInfo**](MapReduceInfo.md) |  | [optional] 
**MrInstances** | Pointer to [**[]MapReduceInstanceWrapper**](MapReduceInstanceWrapper.md) | InstancesWrapper is the slice containing the information about the map reduce application instances. | [optional] 

## Methods

### NewAppRunHistory

`func NewAppRunHistory() *AppRunHistory`

NewAppRunHistory instantiates a new AppRunHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRunHistoryWithDefaults

`func NewAppRunHistoryWithDefaults() *AppRunHistory`

NewAppRunHistoryWithDefaults instantiates a new AppRunHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppInfo

`func (o *AppRunHistory) GetAppInfo() MapReduceInfo`

GetAppInfo returns the AppInfo field if non-nil, zero value otherwise.

### GetAppInfoOk

`func (o *AppRunHistory) GetAppInfoOk() (*MapReduceInfo, bool)`

GetAppInfoOk returns a tuple with the AppInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInfo

`func (o *AppRunHistory) SetAppInfo(v MapReduceInfo)`

SetAppInfo sets AppInfo field to given value.

### HasAppInfo

`func (o *AppRunHistory) HasAppInfo() bool`

HasAppInfo returns a boolean if a field has been set.

### GetMrInstances

`func (o *AppRunHistory) GetMrInstances() []MapReduceInstanceWrapper`

GetMrInstances returns the MrInstances field if non-nil, zero value otherwise.

### GetMrInstancesOk

`func (o *AppRunHistory) GetMrInstancesOk() (*[]MapReduceInstanceWrapper, bool)`

GetMrInstancesOk returns a tuple with the MrInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrInstances

`func (o *AppRunHistory) SetMrInstances(v []MapReduceInstanceWrapper)`

SetMrInstances sets MrInstances field to given value.

### HasMrInstances

`func (o *AppRunHistory) HasMrInstances() bool`

HasMrInstances returns a boolean if a field has been set.

### SetMrInstancesNil

`func (o *AppRunHistory) SetMrInstancesNil(b bool)`

 SetMrInstancesNil sets the value for MrInstances to be an explicit nil

### UnsetMrInstances
`func (o *AppRunHistory) UnsetMrInstances()`

UnsetMrInstances ensures that no value is present for MrInstances, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


