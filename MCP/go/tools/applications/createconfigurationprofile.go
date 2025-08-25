package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-appconfig/mcp-server/config"
	"github.com/amazon-appconfig/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreateconfigurationprofileHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		ApplicationIdVal, ok := args["ApplicationId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: ApplicationId"), nil
		}
		ApplicationId, ok := ApplicationIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: ApplicationId"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CreateConfigurationProfilerequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/applications/%s/configurationprofiles", cfg.BaseURL, ApplicationId)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Content-Sha256"]; ok {
			req.Header.Set("X-Amz-Content-Sha256", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Date"]; ok {
			req.Header.Set("X-Amz-Date", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Algorithm"]; ok {
			req.Header.Set("X-Amz-Algorithm", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Credential"]; ok {
			req.Header.Set("X-Amz-Credential", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Security-Token"]; ok {
			req.Header.Set("X-Amz-Security-Token", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Signature"]; ok {
			req.Header.Set("X-Amz-Signature", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-SignedHeaders"]; ok {
			req.Header.Set("X-Amz-SignedHeaders", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.ConfigurationProfile
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCreateconfigurationprofileTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_applications_ApplicationId_configurationprofiles",
		mcp.WithDescription("<p>Creates a configuration profile, which is information that enables AppConfig to access the configuration source. Valid configuration sources include the following:</p> <ul> <li> <p>Configuration data in YAML, JSON, and other formats stored in the AppConfig hosted configuration store</p> </li> <li> <p>Configuration data stored as objects in an Amazon Simple Storage Service (Amazon S3) bucket</p> </li> <li> <p>Pipelines stored in CodePipeline</p> </li> <li> <p>Secrets stored in Secrets Manager</p> </li> <li> <p>Standard and secure string parameters stored in Amazon Web Services Systems Manager Parameter Store</p> </li> <li> <p>Configuration data in SSM documents stored in the Systems Manager document store</p> </li> </ul> <p>A configuration profile includes the following information:</p> <ul> <li> <p>The URI location of the configuration data.</p> </li> <li> <p>The Identity and Access Management (IAM) role that provides access to the configuration data.</p> </li> <li> <p>A validator for the configuration data. Available validators include either a JSON Schema or an Amazon Web Services Lambda function.</p> </li> </ul> <p>For more information, see <a href="http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile.html">Create a Configuration and a Configuration Profile</a> in the <i>AppConfig User Guide</i>.</p>"),
		mcp.WithString("X-Amz-Content-Sha256", mcp.Description("")),
		mcp.WithString("X-Amz-Date", mcp.Description("")),
		mcp.WithString("X-Amz-Algorithm", mcp.Description("")),
		mcp.WithString("X-Amz-Credential", mcp.Description("")),
		mcp.WithString("X-Amz-Security-Token", mcp.Description("")),
		mcp.WithString("X-Amz-Signature", mcp.Description("")),
		mcp.WithString("X-Amz-SignedHeaders", mcp.Description("")),
		mcp.WithString("ApplicationId", mcp.Required(), mcp.Description("The application ID.")),
		mcp.WithString("LocationUri", mcp.Required(), mcp.Description("Input parameter: <p>A URI to locate the configuration. You can specify the following:</p> <ul> <li> <p>For the AppConfig hosted configuration store and for feature flags, specify <code>hosted</code>.</p> </li> <li> <p>For an Amazon Web Services Systems Manager Parameter Store parameter, specify either the parameter name in the format <code>ssm-parameter://&lt;parameter name&gt;</code> or the ARN.</p> </li> <li> <p>For an Secrets Manager secret, specify the URI in the following format: <code>secrets-manager</code>://&lt;secret name&gt;.</p> </li> <li> <p>For an Amazon S3 object, specify the URI in the following format: <code>s3://&lt;bucket&gt;/&lt;objectKey&gt; </code>. Here is an example: <code>s3://my-bucket/my-app/us-east-1/my-config.json</code> </p> </li> <li> <p>For an SSM document, specify either the document name in the format <code>ssm-document://&lt;document name&gt;</code> or the Amazon Resource Name (ARN).</p> </li> </ul>")),
		mcp.WithString("Name", mcp.Required(), mcp.Description("Input parameter: A name for the configuration profile.")),
		mcp.WithString("RetrievalRoleArn", mcp.Description("Input parameter: <p>The ARN of an IAM role with permission to access the configuration at the specified <code>LocationUri</code>.</p> <important> <p>A retrieval role ARN is not required for configurations stored in the AppConfig hosted configuration store. It is required for all other sources that store your configuration. </p> </important>")),
		mcp.WithObject("Tags", mcp.Description("Input parameter: Metadata to assign to the configuration profile. Tags help organize and categorize your AppConfig resources. Each tag consists of a key and an optional value, both of which you define.")),
		mcp.WithString("Type", mcp.Description("Input parameter: <p>The type of configurations contained in the profile. AppConfig supports <code>feature flags</code> and <code>freeform</code> configurations. We recommend you create feature flag configurations to enable or disable new features and freeform configurations to distribute configurations to an application. When calling this API, enter one of the following values for <code>Type</code>:</p> <p> <code>AWS.AppConfig.FeatureFlags</code> </p> <p> <code>AWS.Freeform</code> </p>")),
		mcp.WithArray("Validators", mcp.Description("Input parameter: A list of methods for validating the configuration.")),
		mcp.WithString("Description", mcp.Description("Input parameter: A description of the configuration profile.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateconfigurationprofileHandler(cfg),
	}
}
