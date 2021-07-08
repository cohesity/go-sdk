# BulkInstallAppTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **NullableString** | Job id of the task. | [optional] 
**NumMachinesFailed** | Pointer to **NullableInt32** | Number of machines on which task is started. | [optional] 
**NumMachinesPassed** | Pointer to **NullableInt32** | Number of machines on which task is started. | [optional] 
**NumMachinesTotal** | Pointer to **NullableInt32** | Number of machines on which task is started. | [optional] 
**RegisteringApp** | Pointer to **NullableString** | Application being registered. This param is used to indicate the app for which the job is created. &#39;oracle&#39; indicates that the job was created for oracle app. &#39;msSql&#39; indicates that the job was created for msSql app. &#39;physical&#39; indicates that the job was created for physical machine. | [optional] 
**State** | Pointer to **NullableString** | Current state of the task. This param is used to indicate the state of the job created by the bulk install app. &#39;started&#39; indicates that the job has been started by the user. &#39;completed&#39; indicates that the job has completed. | [optional] 

## Methods

### NewBulkInstallAppTaskInfo

`func NewBulkInstallAppTaskInfo() *BulkInstallAppTaskInfo`

NewBulkInstallAppTaskInfo instantiates a new BulkInstallAppTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkInstallAppTaskInfoWithDefaults

`func NewBulkInstallAppTaskInfoWithDefaults() *BulkInstallAppTaskInfo`

NewBulkInstallAppTaskInfoWithDefaults instantiates a new BulkInstallAppTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *BulkInstallAppTaskInfo) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BulkInstallAppTaskInfo) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BulkInstallAppTaskInfo) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *BulkInstallAppTaskInfo) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *BulkInstallAppTaskInfo) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *BulkInstallAppTaskInfo) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetNumMachinesFailed

`func (o *BulkInstallAppTaskInfo) GetNumMachinesFailed() int32`

GetNumMachinesFailed returns the NumMachinesFailed field if non-nil, zero value otherwise.

### GetNumMachinesFailedOk

`func (o *BulkInstallAppTaskInfo) GetNumMachinesFailedOk() (*int32, bool)`

GetNumMachinesFailedOk returns a tuple with the NumMachinesFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachinesFailed

`func (o *BulkInstallAppTaskInfo) SetNumMachinesFailed(v int32)`

SetNumMachinesFailed sets NumMachinesFailed field to given value.

### HasNumMachinesFailed

`func (o *BulkInstallAppTaskInfo) HasNumMachinesFailed() bool`

HasNumMachinesFailed returns a boolean if a field has been set.

### SetNumMachinesFailedNil

`func (o *BulkInstallAppTaskInfo) SetNumMachinesFailedNil(b bool)`

 SetNumMachinesFailedNil sets the value for NumMachinesFailed to be an explicit nil

### UnsetNumMachinesFailed
`func (o *BulkInstallAppTaskInfo) UnsetNumMachinesFailed()`

UnsetNumMachinesFailed ensures that no value is present for NumMachinesFailed, not even an explicit nil
### GetNumMachinesPassed

`func (o *BulkInstallAppTaskInfo) GetNumMachinesPassed() int32`

GetNumMachinesPassed returns the NumMachinesPassed field if non-nil, zero value otherwise.

### GetNumMachinesPassedOk

`func (o *BulkInstallAppTaskInfo) GetNumMachinesPassedOk() (*int32, bool)`

GetNumMachinesPassedOk returns a tuple with the NumMachinesPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachinesPassed

`func (o *BulkInstallAppTaskInfo) SetNumMachinesPassed(v int32)`

SetNumMachinesPassed sets NumMachinesPassed field to given value.

### HasNumMachinesPassed

`func (o *BulkInstallAppTaskInfo) HasNumMachinesPassed() bool`

HasNumMachinesPassed returns a boolean if a field has been set.

### SetNumMachinesPassedNil

`func (o *BulkInstallAppTaskInfo) SetNumMachinesPassedNil(b bool)`

 SetNumMachinesPassedNil sets the value for NumMachinesPassed to be an explicit nil

### UnsetNumMachinesPassed
`func (o *BulkInstallAppTaskInfo) UnsetNumMachinesPassed()`

UnsetNumMachinesPassed ensures that no value is present for NumMachinesPassed, not even an explicit nil
### GetNumMachinesTotal

`func (o *BulkInstallAppTaskInfo) GetNumMachinesTotal() int32`

GetNumMachinesTotal returns the NumMachinesTotal field if non-nil, zero value otherwise.

### GetNumMachinesTotalOk

`func (o *BulkInstallAppTaskInfo) GetNumMachinesTotalOk() (*int32, bool)`

GetNumMachinesTotalOk returns a tuple with the NumMachinesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMachinesTotal

`func (o *BulkInstallAppTaskInfo) SetNumMachinesTotal(v int32)`

SetNumMachinesTotal sets NumMachinesTotal field to given value.

### HasNumMachinesTotal

`func (o *BulkInstallAppTaskInfo) HasNumMachinesTotal() bool`

HasNumMachinesTotal returns a boolean if a field has been set.

### SetNumMachinesTotalNil

`func (o *BulkInstallAppTaskInfo) SetNumMachinesTotalNil(b bool)`

 SetNumMachinesTotalNil sets the value for NumMachinesTotal to be an explicit nil

### UnsetNumMachinesTotal
`func (o *BulkInstallAppTaskInfo) UnsetNumMachinesTotal()`

UnsetNumMachinesTotal ensures that no value is present for NumMachinesTotal, not even an explicit nil
### GetRegisteringApp

`func (o *BulkInstallAppTaskInfo) GetRegisteringApp() string`

GetRegisteringApp returns the RegisteringApp field if non-nil, zero value otherwise.

### GetRegisteringAppOk

`func (o *BulkInstallAppTaskInfo) GetRegisteringAppOk() (*string, bool)`

GetRegisteringAppOk returns a tuple with the RegisteringApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteringApp

`func (o *BulkInstallAppTaskInfo) SetRegisteringApp(v string)`

SetRegisteringApp sets RegisteringApp field to given value.

### HasRegisteringApp

`func (o *BulkInstallAppTaskInfo) HasRegisteringApp() bool`

HasRegisteringApp returns a boolean if a field has been set.

### SetRegisteringAppNil

`func (o *BulkInstallAppTaskInfo) SetRegisteringAppNil(b bool)`

 SetRegisteringAppNil sets the value for RegisteringApp to be an explicit nil

### UnsetRegisteringApp
`func (o *BulkInstallAppTaskInfo) UnsetRegisteringApp()`

UnsetRegisteringApp ensures that no value is present for RegisteringApp, not even an explicit nil
### GetState

`func (o *BulkInstallAppTaskInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BulkInstallAppTaskInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BulkInstallAppTaskInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BulkInstallAppTaskInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *BulkInstallAppTaskInfo) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *BulkInstallAppTaskInfo) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


