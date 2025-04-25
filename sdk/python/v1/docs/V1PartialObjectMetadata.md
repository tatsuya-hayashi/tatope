# V1PartialObjectMetadata

PartialObjectMetadata is a generic representation of any object with ObjectMeta. It allows clients to get access to a particular ObjectMeta schema without knowing the details of the version.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_version** | **str** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**kind** | **str** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**metadata** | [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 

## Example

```python
from tatope.models.v1_partial_object_metadata import V1PartialObjectMetadata

# TODO update the JSON string below
json = "{}"
# create an instance of V1PartialObjectMetadata from a JSON string
v1_partial_object_metadata_instance = V1PartialObjectMetadata.from_json(json)
# print the JSON string representation of the object
print(V1PartialObjectMetadata.to_json())

# convert the object into a dict
v1_partial_object_metadata_dict = v1_partial_object_metadata_instance.to_dict()
# create an instance of V1PartialObjectMetadata from a dict
v1_partial_object_metadata_from_dict = V1PartialObjectMetadata.from_dict(v1_partial_object_metadata_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


