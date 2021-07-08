# OutputSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumReduceShards** | Pointer to **NullableInt32** | Number of reduce shards. | [optional] 
**OutputDir** | Pointer to **NullableString** | Name of output directory. | [optional] 
**PartitionId** | Pointer to **NullableInt64** | Partition id where output will go. | [optional] 
**ReduceOutputPrefix** | Pointer to **NullableString** | Prefix of the reduce output files. File names will be: ${reduce_output_prefix}-00000-of-00100 if num_reduce_shards&#x3D;100 This name can contain some path components. e.g. \&quot;awb_results/run1\&quot; is a valid value. output_dir is deprecated. | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | Viewbox id where the output will go. | [optional] 
**ViewName** | Pointer to **NullableString** | Name of the view where output will go. This will be filled up by yoda. | [optional] 

## Methods

### NewOutputSpec

`func NewOutputSpec() *OutputSpec`

NewOutputSpec instantiates a new OutputSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputSpecWithDefaults

`func NewOutputSpecWithDefaults() *OutputSpec`

NewOutputSpecWithDefaults instantiates a new OutputSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumReduceShards

`func (o *OutputSpec) GetNumReduceShards() int32`

GetNumReduceShards returns the NumReduceShards field if non-nil, zero value otherwise.

### GetNumReduceShardsOk

`func (o *OutputSpec) GetNumReduceShardsOk() (*int32, bool)`

GetNumReduceShardsOk returns a tuple with the NumReduceShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReduceShards

`func (o *OutputSpec) SetNumReduceShards(v int32)`

SetNumReduceShards sets NumReduceShards field to given value.

### HasNumReduceShards

`func (o *OutputSpec) HasNumReduceShards() bool`

HasNumReduceShards returns a boolean if a field has been set.

### SetNumReduceShardsNil

`func (o *OutputSpec) SetNumReduceShardsNil(b bool)`

 SetNumReduceShardsNil sets the value for NumReduceShards to be an explicit nil

### UnsetNumReduceShards
`func (o *OutputSpec) UnsetNumReduceShards()`

UnsetNumReduceShards ensures that no value is present for NumReduceShards, not even an explicit nil
### GetOutputDir

`func (o *OutputSpec) GetOutputDir() string`

GetOutputDir returns the OutputDir field if non-nil, zero value otherwise.

### GetOutputDirOk

`func (o *OutputSpec) GetOutputDirOk() (*string, bool)`

GetOutputDirOk returns a tuple with the OutputDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDir

`func (o *OutputSpec) SetOutputDir(v string)`

SetOutputDir sets OutputDir field to given value.

### HasOutputDir

`func (o *OutputSpec) HasOutputDir() bool`

HasOutputDir returns a boolean if a field has been set.

### SetOutputDirNil

`func (o *OutputSpec) SetOutputDirNil(b bool)`

 SetOutputDirNil sets the value for OutputDir to be an explicit nil

### UnsetOutputDir
`func (o *OutputSpec) UnsetOutputDir()`

UnsetOutputDir ensures that no value is present for OutputDir, not even an explicit nil
### GetPartitionId

`func (o *OutputSpec) GetPartitionId() int64`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *OutputSpec) GetPartitionIdOk() (*int64, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *OutputSpec) SetPartitionId(v int64)`

SetPartitionId sets PartitionId field to given value.

### HasPartitionId

`func (o *OutputSpec) HasPartitionId() bool`

HasPartitionId returns a boolean if a field has been set.

### SetPartitionIdNil

`func (o *OutputSpec) SetPartitionIdNil(b bool)`

 SetPartitionIdNil sets the value for PartitionId to be an explicit nil

### UnsetPartitionId
`func (o *OutputSpec) UnsetPartitionId()`

UnsetPartitionId ensures that no value is present for PartitionId, not even an explicit nil
### GetReduceOutputPrefix

`func (o *OutputSpec) GetReduceOutputPrefix() string`

GetReduceOutputPrefix returns the ReduceOutputPrefix field if non-nil, zero value otherwise.

### GetReduceOutputPrefixOk

`func (o *OutputSpec) GetReduceOutputPrefixOk() (*string, bool)`

GetReduceOutputPrefixOk returns a tuple with the ReduceOutputPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOutputPrefix

`func (o *OutputSpec) SetReduceOutputPrefix(v string)`

SetReduceOutputPrefix sets ReduceOutputPrefix field to given value.

### HasReduceOutputPrefix

`func (o *OutputSpec) HasReduceOutputPrefix() bool`

HasReduceOutputPrefix returns a boolean if a field has been set.

### SetReduceOutputPrefixNil

`func (o *OutputSpec) SetReduceOutputPrefixNil(b bool)`

 SetReduceOutputPrefixNil sets the value for ReduceOutputPrefix to be an explicit nil

### UnsetReduceOutputPrefix
`func (o *OutputSpec) UnsetReduceOutputPrefix()`

UnsetReduceOutputPrefix ensures that no value is present for ReduceOutputPrefix, not even an explicit nil
### GetViewBoxId

`func (o *OutputSpec) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *OutputSpec) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *OutputSpec) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *OutputSpec) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *OutputSpec) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *OutputSpec) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
### GetViewName

`func (o *OutputSpec) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *OutputSpec) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *OutputSpec) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *OutputSpec) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *OutputSpec) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *OutputSpec) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


