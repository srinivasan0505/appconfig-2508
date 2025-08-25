

# Action

<p>An action defines the tasks that the extension performs during the AppConfig workflow. Each action includes an action point such as <code>ON_CREATE_HOSTED_CONFIGURATION</code>, <code>PRE_DEPLOYMENT</code>, or <code>ON_DEPLOYMENT</code>. Each action also includes a name, a URI to an Lambda function, and an Amazon Resource Name (ARN) for an Identity and Access Management assume role. You specify the name, URI, and ARN for each <i>action point</i> defined in the extension. You can specify the following actions for an extension:</p> <ul> <li> <p> <code>PRE_CREATE_HOSTED_CONFIGURATION_VERSION</code> </p> </li> <li> <p> <code>PRE_START_DEPLOYMENT</code> </p> </li> <li> <p> <code>ON_DEPLOYMENT_START</code> </p> </li> <li> <p> <code>ON_DEPLOYMENT_STEP</code> </p> </li> <li> <p> <code>ON_DEPLOYMENT_BAKING</code> </p> </li> <li> <p> <code>ON_DEPLOYMENT_COMPLETE</code> </p> </li> <li> <p> <code>ON_DEPLOYMENT_ROLLED_BACK</code> </p> </li> </ul>

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**name** | [**String**](String.md) |  |  [optional] |
|**description** | [**String**](String.md) |  |  [optional] |
|**uri** | [**String**](String.md) |  |  [optional] |
|**roleArn** | [**String**](String.md) |  |  [optional] |



