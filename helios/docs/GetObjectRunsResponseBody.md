# GetObjectRunsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionRuns** | Pointer to [**[]ObjectProtectionRunSummary**](ObjectProtectionRunSummary.md) | Specifies the protection runs of the given object. | [optional] 
**PaginationInfo** | Pointer to [**PaginationInfo**](PaginationInfo.md) |  | [optional] 

## Methods

### NewGetObjectRunsResponseBody

`func NewGetObjectRunsResponseBody() *GetObjectRunsResponseBody`

NewGetObjectRunsResponseBody instantiates a new GetObjectRunsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectRunsResponseBodyWithDefaults

`func NewGetObjectRunsResponseBodyWithDefaults() *GetObjectRunsResponseBody`

NewGetObjectRunsResponseBodyWithDefaults instantiates a new GetObjectRunsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionRuns

`func (o *GetObjectRunsResponseBody) GetProtectionRuns() []ObjectProtectionRunSummary`

GetProtectionRuns returns the ProtectionRuns field if non-nil, zero value otherwise.

### GetProtectionRunsOk

`func (o *GetObjectRunsResponseBody) GetProtectionRunsOk() (*[]ObjectProtectionRunSummary, bool)`

GetProtectionRunsOk returns a tuple with the ProtectionRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionRuns

`func (o *GetObjectRunsResponseBody) SetProtectionRuns(v []ObjectProtectionRunSummary)`

SetProtectionRuns sets ProtectionRuns field to given value.

### HasProtectionRuns

`func (o *GetObjectRunsResponseBody) HasProtectionRuns() bool`

HasProtectionRuns returns a boolean if a field has been set.

### GetPaginationInfo

`func (o *GetObjectRunsResponseBody) GetPaginationInfo() PaginationInfo`

GetPaginationInfo returns the PaginationInfo field if non-nil, zero value otherwise.

### GetPaginationInfoOk

`func (o *GetObjectRunsResponseBody) GetPaginationInfoOk() (*PaginationInfo, bool)`

GetPaginationInfoOk returns a tuple with the PaginationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationInfo

`func (o *GetObjectRunsResponseBody) SetPaginationInfo(v PaginationInfo)`

SetPaginationInfo sets PaginationInfo field to given value.

### HasPaginationInfo

`func (o *GetObjectRunsResponseBody) HasPaginationInfo() bool`

HasPaginationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


