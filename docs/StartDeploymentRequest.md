

# StartDeploymentRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**deploymentStrategyId** | **String** | The deployment strategy ID. |  |
|**configurationProfileId** | **String** | The configuration profile ID. |  |
|**configurationVersion** | **String** | The configuration version to deploy. If deploying an AppConfig hosted configuration version, you can specify either the version number or version label. |  |
|**description** | **String** | A description of the deployment. |  [optional] |
|**tags** | **Map&lt;String, String&gt;** | Metadata to assign to the deployment. Tags help organize and categorize your AppConfig resources. Each tag consists of a key and an optional value, both of which you define. |  [optional] |
|**kmsKeyIdentifier** | **String** | The KMS key identifier (key ID, key alias, or key ARN). AppConfig uses this ID to encrypt the configuration data using a customer managed key.  |  [optional] |



