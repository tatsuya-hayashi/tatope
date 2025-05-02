# V1InternalEvent

InternalEvent makes watch.Event versioned

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**object** | [**K8sIoApimachineryPkgRuntimeObject**](K8sIoApimachineryPkgRuntimeObject.md) |  | 
**type** | **str** |  | [default to '']

## Example

```python
from tatope.models.v1_internal_event import V1InternalEvent

# TODO update the JSON string below
json = "{}"
# create an instance of V1InternalEvent from a JSON string
v1_internal_event_instance = V1InternalEvent.from_json(json)
# print the JSON string representation of the object
print(V1InternalEvent.to_json())

# convert the object into a dict
v1_internal_event_dict = v1_internal_event_instance.to_dict()
# create an instance of V1InternalEvent from a dict
v1_internal_event_from_dict = V1InternalEvent.from_dict(v1_internal_event_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


