# RuntimeUnknown

Unknown allows api objects with unknown types to be passed-through. This can be used to deal with the API objects from a plug-in. Unknown objects still have functioning TypeMeta features-- kind, version, etc. metadata and field mutatation.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**content_encoding** | **str** | ContentEncoding is encoding used to encode &#39;Raw&#39; data. Unspecified means no encoding. | [default to '']
**content_type** | **str** | ContentType  is serialization method used to serialize &#39;Raw&#39;. Unspecified means ContentTypeJSON. | [default to '']
**api_version** | **str** |  | [optional] 
**kind** | **str** |  | [optional] 

## Example

```python
from tatope.models.runtime_unknown import RuntimeUnknown

# TODO update the JSON string below
json = "{}"
# create an instance of RuntimeUnknown from a JSON string
runtime_unknown_instance = RuntimeUnknown.from_json(json)
# print the JSON string representation of the object
print(RuntimeUnknown.to_json())

# convert the object into a dict
runtime_unknown_dict = runtime_unknown_instance.to_dict()
# create an instance of RuntimeUnknown from a dict
runtime_unknown_from_dict = RuntimeUnknown.from_dict(runtime_unknown_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


