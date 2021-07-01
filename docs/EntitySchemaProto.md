# EntitySchemaProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributesDescriptor** | Pointer to [**EntitySchemaProtoAttributesDescriptor**](EntitySchemaProtoAttributesDescriptor.md) |  | [optional] 
**EnableRollup** | Pointer to **NullableBool** | Timeseries for an entity schema is rolled up based on this setting. Rollup is disabled by default. Rollups cannot be done for metrics with value_type other than kInt64 or kDouble. | [optional] 
**FlushIntervalSecs** | Pointer to **NullableInt32** | Defines the interval used to flush in memory stats to scribe table. During this time if the stats server is down before flushing, it could loose some of the stats. Modules can flush any critical stats via AddEntitiesStats API. But this  should be used very judiciously as it causes lot of overhead for stats. | [optional] 
**IsInternalSchema** | Pointer to **NullableBool** | Specifies if this schema should be displayed in Advanced Diagnostics of the Cohesity Dashboard. If false, the schema is displayed. | [optional] 
**LargestFlushIntervalSecs** | Pointer to **NullableInt32** | Use can change the flush interval secs via gflag and this store the largest interval seconds set. This is used to round up the timestamp to this flush interval secs during range scan. | [optional] 
**Name** | Pointer to **NullableString** | Specifies a name that uniquely identifies an entity schema such as &#39;kBridgeClusterStats&#39;. Name cannot have &#39;:&#39; as character. | [optional] 
**RollupGranularityVec** | Pointer to [**[]EntitySchemaProtoGranularity**](EntitySchemaProtoGranularity.md) |  | [optional] 
**SchemaDescriptiveName** | Pointer to **NullableString** | Specifies the name of the Schema as displayed in Advanced Diagnostics of the Cohesity Dashboard. For example for the &#39;kBridgeClusterStats&#39; Schema, the descriptive name is &#39;Cluster Physical Stats&#39;. | [optional] 
**SchemaHelpText** | Pointer to **NullableString** | Specifies an optional informational description about the schema. | [optional] 
**TimeSeriesDescriptorVec** | Pointer to [**[]EntitySchemaProtoTimeSeriesDescriptor**](EntitySchemaProtoTimeSeriesDescriptor.md) | Array of Time Series.  List of time series of data (set of data points) for metrics. | [optional] 
**TimeToLiveSecs** | Pointer to **NullableInt64** | Specifies how long the timeseries data of this schema will be stored. After expiry the entire data point(all metrics) is garbage collected. | [optional] 
**Version** | Pointer to **NullableInt64** | Specifies the version of the entity schema. | [optional] 

## Methods

### NewEntitySchemaProto

`func NewEntitySchemaProto() *EntitySchemaProto`

NewEntitySchemaProto instantiates a new EntitySchemaProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySchemaProtoWithDefaults

`func NewEntitySchemaProtoWithDefaults() *EntitySchemaProto`

NewEntitySchemaProtoWithDefaults instantiates a new EntitySchemaProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributesDescriptor

`func (o *EntitySchemaProto) GetAttributesDescriptor() EntitySchemaProtoAttributesDescriptor`

GetAttributesDescriptor returns the AttributesDescriptor field if non-nil, zero value otherwise.

### GetAttributesDescriptorOk

`func (o *EntitySchemaProto) GetAttributesDescriptorOk() (*EntitySchemaProtoAttributesDescriptor, bool)`

GetAttributesDescriptorOk returns a tuple with the AttributesDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesDescriptor

`func (o *EntitySchemaProto) SetAttributesDescriptor(v EntitySchemaProtoAttributesDescriptor)`

SetAttributesDescriptor sets AttributesDescriptor field to given value.

### HasAttributesDescriptor

`func (o *EntitySchemaProto) HasAttributesDescriptor() bool`

HasAttributesDescriptor returns a boolean if a field has been set.

### GetEnableRollup

`func (o *EntitySchemaProto) GetEnableRollup() bool`

GetEnableRollup returns the EnableRollup field if non-nil, zero value otherwise.

### GetEnableRollupOk

`func (o *EntitySchemaProto) GetEnableRollupOk() (*bool, bool)`

GetEnableRollupOk returns a tuple with the EnableRollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRollup

`func (o *EntitySchemaProto) SetEnableRollup(v bool)`

SetEnableRollup sets EnableRollup field to given value.

### HasEnableRollup

`func (o *EntitySchemaProto) HasEnableRollup() bool`

HasEnableRollup returns a boolean if a field has been set.

### SetEnableRollupNil

`func (o *EntitySchemaProto) SetEnableRollupNil(b bool)`

 SetEnableRollupNil sets the value for EnableRollup to be an explicit nil

### UnsetEnableRollup
`func (o *EntitySchemaProto) UnsetEnableRollup()`

UnsetEnableRollup ensures that no value is present for EnableRollup, not even an explicit nil
### GetFlushIntervalSecs

`func (o *EntitySchemaProto) GetFlushIntervalSecs() int32`

GetFlushIntervalSecs returns the FlushIntervalSecs field if non-nil, zero value otherwise.

### GetFlushIntervalSecsOk

`func (o *EntitySchemaProto) GetFlushIntervalSecsOk() (*int32, bool)`

GetFlushIntervalSecsOk returns a tuple with the FlushIntervalSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushIntervalSecs

`func (o *EntitySchemaProto) SetFlushIntervalSecs(v int32)`

SetFlushIntervalSecs sets FlushIntervalSecs field to given value.

### HasFlushIntervalSecs

`func (o *EntitySchemaProto) HasFlushIntervalSecs() bool`

HasFlushIntervalSecs returns a boolean if a field has been set.

### SetFlushIntervalSecsNil

`func (o *EntitySchemaProto) SetFlushIntervalSecsNil(b bool)`

 SetFlushIntervalSecsNil sets the value for FlushIntervalSecs to be an explicit nil

### UnsetFlushIntervalSecs
`func (o *EntitySchemaProto) UnsetFlushIntervalSecs()`

UnsetFlushIntervalSecs ensures that no value is present for FlushIntervalSecs, not even an explicit nil
### GetIsInternalSchema

`func (o *EntitySchemaProto) GetIsInternalSchema() bool`

GetIsInternalSchema returns the IsInternalSchema field if non-nil, zero value otherwise.

### GetIsInternalSchemaOk

`func (o *EntitySchemaProto) GetIsInternalSchemaOk() (*bool, bool)`

GetIsInternalSchemaOk returns a tuple with the IsInternalSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternalSchema

`func (o *EntitySchemaProto) SetIsInternalSchema(v bool)`

SetIsInternalSchema sets IsInternalSchema field to given value.

### HasIsInternalSchema

`func (o *EntitySchemaProto) HasIsInternalSchema() bool`

HasIsInternalSchema returns a boolean if a field has been set.

### SetIsInternalSchemaNil

`func (o *EntitySchemaProto) SetIsInternalSchemaNil(b bool)`

 SetIsInternalSchemaNil sets the value for IsInternalSchema to be an explicit nil

### UnsetIsInternalSchema
`func (o *EntitySchemaProto) UnsetIsInternalSchema()`

UnsetIsInternalSchema ensures that no value is present for IsInternalSchema, not even an explicit nil
### GetLargestFlushIntervalSecs

`func (o *EntitySchemaProto) GetLargestFlushIntervalSecs() int32`

GetLargestFlushIntervalSecs returns the LargestFlushIntervalSecs field if non-nil, zero value otherwise.

### GetLargestFlushIntervalSecsOk

`func (o *EntitySchemaProto) GetLargestFlushIntervalSecsOk() (*int32, bool)`

GetLargestFlushIntervalSecsOk returns a tuple with the LargestFlushIntervalSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestFlushIntervalSecs

`func (o *EntitySchemaProto) SetLargestFlushIntervalSecs(v int32)`

SetLargestFlushIntervalSecs sets LargestFlushIntervalSecs field to given value.

### HasLargestFlushIntervalSecs

`func (o *EntitySchemaProto) HasLargestFlushIntervalSecs() bool`

HasLargestFlushIntervalSecs returns a boolean if a field has been set.

### SetLargestFlushIntervalSecsNil

`func (o *EntitySchemaProto) SetLargestFlushIntervalSecsNil(b bool)`

 SetLargestFlushIntervalSecsNil sets the value for LargestFlushIntervalSecs to be an explicit nil

### UnsetLargestFlushIntervalSecs
`func (o *EntitySchemaProto) UnsetLargestFlushIntervalSecs()`

UnsetLargestFlushIntervalSecs ensures that no value is present for LargestFlushIntervalSecs, not even an explicit nil
### GetName

`func (o *EntitySchemaProto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitySchemaProto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitySchemaProto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntitySchemaProto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EntitySchemaProto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EntitySchemaProto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRollupGranularityVec

`func (o *EntitySchemaProto) GetRollupGranularityVec() []EntitySchemaProtoGranularity`

GetRollupGranularityVec returns the RollupGranularityVec field if non-nil, zero value otherwise.

### GetRollupGranularityVecOk

`func (o *EntitySchemaProto) GetRollupGranularityVecOk() (*[]EntitySchemaProtoGranularity, bool)`

GetRollupGranularityVecOk returns a tuple with the RollupGranularityVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupGranularityVec

`func (o *EntitySchemaProto) SetRollupGranularityVec(v []EntitySchemaProtoGranularity)`

SetRollupGranularityVec sets RollupGranularityVec field to given value.

### HasRollupGranularityVec

`func (o *EntitySchemaProto) HasRollupGranularityVec() bool`

HasRollupGranularityVec returns a boolean if a field has been set.

### SetRollupGranularityVecNil

`func (o *EntitySchemaProto) SetRollupGranularityVecNil(b bool)`

 SetRollupGranularityVecNil sets the value for RollupGranularityVec to be an explicit nil

### UnsetRollupGranularityVec
`func (o *EntitySchemaProto) UnsetRollupGranularityVec()`

UnsetRollupGranularityVec ensures that no value is present for RollupGranularityVec, not even an explicit nil
### GetSchemaDescriptiveName

`func (o *EntitySchemaProto) GetSchemaDescriptiveName() string`

GetSchemaDescriptiveName returns the SchemaDescriptiveName field if non-nil, zero value otherwise.

### GetSchemaDescriptiveNameOk

`func (o *EntitySchemaProto) GetSchemaDescriptiveNameOk() (*string, bool)`

GetSchemaDescriptiveNameOk returns a tuple with the SchemaDescriptiveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaDescriptiveName

`func (o *EntitySchemaProto) SetSchemaDescriptiveName(v string)`

SetSchemaDescriptiveName sets SchemaDescriptiveName field to given value.

### HasSchemaDescriptiveName

`func (o *EntitySchemaProto) HasSchemaDescriptiveName() bool`

HasSchemaDescriptiveName returns a boolean if a field has been set.

### SetSchemaDescriptiveNameNil

`func (o *EntitySchemaProto) SetSchemaDescriptiveNameNil(b bool)`

 SetSchemaDescriptiveNameNil sets the value for SchemaDescriptiveName to be an explicit nil

### UnsetSchemaDescriptiveName
`func (o *EntitySchemaProto) UnsetSchemaDescriptiveName()`

UnsetSchemaDescriptiveName ensures that no value is present for SchemaDescriptiveName, not even an explicit nil
### GetSchemaHelpText

`func (o *EntitySchemaProto) GetSchemaHelpText() string`

GetSchemaHelpText returns the SchemaHelpText field if non-nil, zero value otherwise.

### GetSchemaHelpTextOk

`func (o *EntitySchemaProto) GetSchemaHelpTextOk() (*string, bool)`

GetSchemaHelpTextOk returns a tuple with the SchemaHelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaHelpText

`func (o *EntitySchemaProto) SetSchemaHelpText(v string)`

SetSchemaHelpText sets SchemaHelpText field to given value.

### HasSchemaHelpText

`func (o *EntitySchemaProto) HasSchemaHelpText() bool`

HasSchemaHelpText returns a boolean if a field has been set.

### SetSchemaHelpTextNil

`func (o *EntitySchemaProto) SetSchemaHelpTextNil(b bool)`

 SetSchemaHelpTextNil sets the value for SchemaHelpText to be an explicit nil

### UnsetSchemaHelpText
`func (o *EntitySchemaProto) UnsetSchemaHelpText()`

UnsetSchemaHelpText ensures that no value is present for SchemaHelpText, not even an explicit nil
### GetTimeSeriesDescriptorVec

`func (o *EntitySchemaProto) GetTimeSeriesDescriptorVec() []EntitySchemaProtoTimeSeriesDescriptor`

GetTimeSeriesDescriptorVec returns the TimeSeriesDescriptorVec field if non-nil, zero value otherwise.

### GetTimeSeriesDescriptorVecOk

`func (o *EntitySchemaProto) GetTimeSeriesDescriptorVecOk() (*[]EntitySchemaProtoTimeSeriesDescriptor, bool)`

GetTimeSeriesDescriptorVecOk returns a tuple with the TimeSeriesDescriptorVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSeriesDescriptorVec

`func (o *EntitySchemaProto) SetTimeSeriesDescriptorVec(v []EntitySchemaProtoTimeSeriesDescriptor)`

SetTimeSeriesDescriptorVec sets TimeSeriesDescriptorVec field to given value.

### HasTimeSeriesDescriptorVec

`func (o *EntitySchemaProto) HasTimeSeriesDescriptorVec() bool`

HasTimeSeriesDescriptorVec returns a boolean if a field has been set.

### SetTimeSeriesDescriptorVecNil

`func (o *EntitySchemaProto) SetTimeSeriesDescriptorVecNil(b bool)`

 SetTimeSeriesDescriptorVecNil sets the value for TimeSeriesDescriptorVec to be an explicit nil

### UnsetTimeSeriesDescriptorVec
`func (o *EntitySchemaProto) UnsetTimeSeriesDescriptorVec()`

UnsetTimeSeriesDescriptorVec ensures that no value is present for TimeSeriesDescriptorVec, not even an explicit nil
### GetTimeToLiveSecs

`func (o *EntitySchemaProto) GetTimeToLiveSecs() int64`

GetTimeToLiveSecs returns the TimeToLiveSecs field if non-nil, zero value otherwise.

### GetTimeToLiveSecsOk

`func (o *EntitySchemaProto) GetTimeToLiveSecsOk() (*int64, bool)`

GetTimeToLiveSecsOk returns a tuple with the TimeToLiveSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToLiveSecs

`func (o *EntitySchemaProto) SetTimeToLiveSecs(v int64)`

SetTimeToLiveSecs sets TimeToLiveSecs field to given value.

### HasTimeToLiveSecs

`func (o *EntitySchemaProto) HasTimeToLiveSecs() bool`

HasTimeToLiveSecs returns a boolean if a field has been set.

### SetTimeToLiveSecsNil

`func (o *EntitySchemaProto) SetTimeToLiveSecsNil(b bool)`

 SetTimeToLiveSecsNil sets the value for TimeToLiveSecs to be an explicit nil

### UnsetTimeToLiveSecs
`func (o *EntitySchemaProto) UnsetTimeToLiveSecs()`

UnsetTimeToLiveSecs ensures that no value is present for TimeToLiveSecs, not even an explicit nil
### GetVersion

`func (o *EntitySchemaProto) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EntitySchemaProto) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EntitySchemaProto) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EntitySchemaProto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *EntitySchemaProto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *EntitySchemaProto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


