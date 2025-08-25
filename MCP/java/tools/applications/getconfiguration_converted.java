/**
 * MCP Server function for No description available
 */

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.*;
import java.util.function.Function;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;

class Get_Applications_Application_Environments_Environment_Configurations_Configuration_Client_IdMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getGet_Applications_Application_Environments_Environment_Configurations_Configuration_Client_IdHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("X-Amz-Content-Sha256")) {
            queryParams.add("X-Amz-Content-Sha256=" + args.get("X-Amz-Content-Sha256"));
        }
        if (args.containsKey("X-Amz-Date")) {
            queryParams.add("X-Amz-Date=" + args.get("X-Amz-Date"));
        }
        if (args.containsKey("X-Amz-Algorithm")) {
            queryParams.add("X-Amz-Algorithm=" + args.get("X-Amz-Algorithm"));
        }
        if (args.containsKey("X-Amz-Credential")) {
            queryParams.add("X-Amz-Credential=" + args.get("X-Amz-Credential"));
        }
        if (args.containsKey("X-Amz-Security-Token")) {
            queryParams.add("X-Amz-Security-Token=" + args.get("X-Amz-Security-Token"));
        }
        if (args.containsKey("X-Amz-Signature")) {
            queryParams.add("X-Amz-Signature=" + args.get("X-Amz-Signature"));
        }
        if (args.containsKey("X-Amz-SignedHeaders")) {
            queryParams.add("X-Amz-SignedHeaders=" + args.get("X-Amz-SignedHeaders"));
        }
        if (args.containsKey("Application")) {
            queryParams.add("Application=" + args.get("Application"));
        }
        if (args.containsKey("Environment")) {
            queryParams.add("Environment=" + args.get("Environment"));
        }
        if (args.containsKey("Configuration")) {
            queryParams.add("Configuration=" + args.get("Configuration"));
        }
        if (args.containsKey("client_id")) {
            queryParams.add("client_id=" + args.get("client_id"));
        }
        if (args.containsKey("client_configuration_version")) {
            queryParams.add("client_configuration_version=" + args.get("client_configuration_version"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/get_applications_application_environments_environment_configurations_configuration_client_id" + queryString;
                
                HttpClient client = HttpClient.newHttpClient();
                HttpRequest httpRequest = HttpRequest.newBuilder()
                    .uri(URI.create(url))
                    .header("Authorization", "Bearer " + config.getApiKey())
                    .header("Accept", "application/json")
                    .GET()
                    .build();
                
                HttpResponse<String> response = client.send(httpRequest, HttpResponse.BodyHandlers.ofString());
                
                if (response.statusCode() >= 400) {
                    return new MCPServer.MCPToolResult("API error: " + response.body(), true);
                }
                
                // Pretty print JSON
                ObjectMapper mapper = new ObjectMapper();
                JsonNode jsonNode = mapper.readTree(response.body());
                String prettyJson = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(jsonNode);
                
                return new MCPServer.MCPToolResult(prettyJson);
                
            } catch (IOException | InterruptedException e) {
                return new MCPServer.MCPToolResult("Request failed: " + e.getMessage(), true);
            } catch (Exception e) {
                return new MCPServer.MCPToolResult("Unexpected error: " + e.getMessage(), true);
            }
        };
    }
    
    public static MCPServer.Tool createGet_Applications_Application_Environments_Environment_Configurations_Configuration_Client_IdTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> X-Amz-Content-Sha256Property = new HashMap<>();
        X-Amz-Content-Sha256Property.put("type", "string");
        X-Amz-Content-Sha256Property.put("required", false);
        X-Amz-Content-Sha256Property.put("description", "");
        properties.put("X-Amz-Content-Sha256", X-Amz-Content-Sha256Property);
        Map<String, Object> X-Amz-DateProperty = new HashMap<>();
        X-Amz-DateProperty.put("type", "string");
        X-Amz-DateProperty.put("required", false);
        X-Amz-DateProperty.put("description", "");
        properties.put("X-Amz-Date", X-Amz-DateProperty);
        Map<String, Object> X-Amz-AlgorithmProperty = new HashMap<>();
        X-Amz-AlgorithmProperty.put("type", "string");
        X-Amz-AlgorithmProperty.put("required", false);
        X-Amz-AlgorithmProperty.put("description", "");
        properties.put("X-Amz-Algorithm", X-Amz-AlgorithmProperty);
        Map<String, Object> X-Amz-CredentialProperty = new HashMap<>();
        X-Amz-CredentialProperty.put("type", "string");
        X-Amz-CredentialProperty.put("required", false);
        X-Amz-CredentialProperty.put("description", "");
        properties.put("X-Amz-Credential", X-Amz-CredentialProperty);
        Map<String, Object> X-Amz-Security-TokenProperty = new HashMap<>();
        X-Amz-Security-TokenProperty.put("type", "string");
        X-Amz-Security-TokenProperty.put("required", false);
        X-Amz-Security-TokenProperty.put("description", "");
        properties.put("X-Amz-Security-Token", X-Amz-Security-TokenProperty);
        Map<String, Object> X-Amz-SignatureProperty = new HashMap<>();
        X-Amz-SignatureProperty.put("type", "string");
        X-Amz-SignatureProperty.put("required", false);
        X-Amz-SignatureProperty.put("description", "");
        properties.put("X-Amz-Signature", X-Amz-SignatureProperty);
        Map<String, Object> X-Amz-SignedHeadersProperty = new HashMap<>();
        X-Amz-SignedHeadersProperty.put("type", "string");
        X-Amz-SignedHeadersProperty.put("required", false);
        X-Amz-SignedHeadersProperty.put("description", "");
        properties.put("X-Amz-SignedHeaders", X-Amz-SignedHeadersProperty);
        Map<String, Object> ApplicationProperty = new HashMap<>();
        ApplicationProperty.put("type", "string");
        ApplicationProperty.put("required", true);
        ApplicationProperty.put("description", "The application to get. Specify either the application name or the application ID.");
        properties.put("Application", ApplicationProperty);
        Map<String, Object> EnvironmentProperty = new HashMap<>();
        EnvironmentProperty.put("type", "string");
        EnvironmentProperty.put("required", true);
        EnvironmentProperty.put("description", "The environment to get. Specify either the environment name or the environment ID.");
        properties.put("Environment", EnvironmentProperty);
        Map<String, Object> ConfigurationProperty = new HashMap<>();
        ConfigurationProperty.put("type", "string");
        ConfigurationProperty.put("required", true);
        ConfigurationProperty.put("description", "The configuration to get. Specify either the configuration name or the configuration ID.");
        properties.put("Configuration", ConfigurationProperty);
        Map<String, Object> client_idProperty = new HashMap<>();
        client_idProperty.put("type", "string");
        client_idProperty.put("required", true);
        client_idProperty.put("description", "The clientId parameter in the following command is a unique, user-specified ID to identify the client for the configuration. This ID enables AppConfig to deploy the configuration in intervals, as defined in the deployment strategy.");
        properties.put("client_id", client_idProperty);
        Map<String, Object> client_configuration_versionProperty = new HashMap<>();
        client_configuration_versionProperty.put("type", "string");
        client_configuration_versionProperty.put("required", false);
        client_configuration_versionProperty.put("description", "<p>The configuration version returned in the most recent <code>GetConfiguration</code> response.</p> <important> <p>AppConfig uses the value of the <code>ClientConfigurationVersion</code> parameter to identify the configuration version on your clients. If you don’t send <code>ClientConfigurationVersion</code> with each call to <code>GetConfiguration</code>, your clients receive the current configuration. You are charged each time your clients receive a configuration.</p> <p>To avoid excess charges, we recommend you use the <a href=\"https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/StartConfigurationSession.html\">StartConfigurationSession</a> and <a href=\"https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/GetLatestConfiguration.html\">GetLatestConfiguration</a> APIs, which track the client configuration version on your behalf. If you choose to continue using <code>GetConfiguration</code>, we recommend that you include the <code>ClientConfigurationVersion</code> value with every call to <code>GetConfiguration</code>. The value to use for <code>ClientConfigurationVersion</code> comes from the <code>ConfigurationVersion</code> attribute returned by <code>GetConfiguration</code> when there is new or updated data, and should be saved for subsequent calls to <code>GetConfiguration</code>.</p> </important> <p>For more information about working with configurations, see <a href=\"http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration.html\">Retrieving the Configuration</a> in the <i>AppConfig User Guide</i>.</p>");
        properties.put("client_configuration_version", client_configuration_versionProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "get_applications_application_environments_environment_configurations_configuration_client_id",
            "No description available",
            parameters
        );
        
        return new MCPServer.Tool(definition, getGet_Applications_Application_Environments_Environment_Configurations_Configuration_Client_IdHandler(config));
    }
    
}