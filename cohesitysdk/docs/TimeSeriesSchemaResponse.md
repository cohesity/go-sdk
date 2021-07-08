# TimeSeriesSchemaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaInfoList** | Pointer to [**[]SchemaInfo**](SchemaInfo.md) | Specifies the list of the schema info for an entity. | [optional] 

## Methods

### NewTimeSeriesSchemaResponse

`func NewTimeSeriesSchemaResponse() *TimeSeriesSchemaResponse`

NewTimeSeriesSchemaResponse instantiates a new TimeSeriesSchemaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesSchemaResponseWithDefaults

`func NewTimeSeriesSchemaResponseWithDefaults() *TimeSeriesSchemaResponse`

NewTimeSeriesSchemaResponseWithDefaults instantiates a new TimeSeriesSchemaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaInfoList

`func (o *TimeSeriesSchemaResponse) GetSchemaInfoList() []SchemaInfo`

GetSchemaInfoList returns the SchemaInfoList field if non-nil, zero value otherwise.

### GetSchemaInfoListOk

`func (o *TimeSeriesSchemaResponse) GetSchemaInfoListOk() (*[]SchemaInfo, bool)`

GetSchemaInfoListOk returns a tuple with the SchemaInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaInfoList

`func (o *TimeSeriesSchemaResponse) SetSchemaInfoList(v []SchemaInfo)`

SetSchemaInfoList sets SchemaInfoList field to given value.

### HasSchemaInfoList

`func (o *TimeSeriesSchemaResponse) HasSchemaInfoList() bool`

HasSchemaInfoList returns a boolean if a field has been set.

### SetSchemaInfoListNil

`func (o *TimeSeriesSchemaResponse) SetSchemaInfoListNil(b bool)`

 SetSchemaInfoListNil sets the value for SchemaInfoList to be an explicit nil

### UnsetSchemaInfoList
`func (o *TimeSeriesSchemaResponse) UnsetSchemaInfoList()`

UnsetSchemaInfoList ensures that no value is present for SchemaInfoList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


