use serde::{Deserialize, Serialize};
use std::collections::HashMap;

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct OpenApi {
    pub openapi: String,
    pub info: Info,
    pub paths: HashMap<String, PathItem>,
    pub components: Option<Components>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Info {
    pub title: String,
    pub version: String,
    pub description: Option<String>,
    pub terms_of_service: Option<String>,
    pub contact: Option<Contact>,
    pub license: Option<License>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Contact {
    pub name: Option<String>,
    pub url: Option<String>,
    pub email: Option<String>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct License {
    pub name: String,
    pub url: Option<String>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct PathItem {
    pub ref_path: Option<String>,
    pub summary: Option<String>,
    pub description: Option<String>,
    pub get: Option<Operation>,
    pub put: Option<Operation>,
    pub post: Option<Operation>,
    pub delete: Option<Operation>,
    pub options: Option<Operation>,
    pub head: Option<Operation>,
    pub patch: Option<Operation>,
    pub trace: Option<Operation>,
    pub servers: Option<Vec<Server>>,
    pub parameters: Option<Vec<Parameter>>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Operation {
    pub summary: Option<String>,
    pub description: Option<String>,
    pub operation_id: Option<String>,
    pub tags: Option<Vec<String>>,
    pub parameters: Option<Vec<Parameter>>,
    pub request_body: Option<RequestBody>,
    pub responses: Option<Responses>,
    pub security: Option<Vec<SecurityRequirement>>,
    pub servers: Option<Vec<Server>>,
    pub deprecated: Option<bool>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Parameter {
    pub name: String,
    pub in_field: Option<String>,
    pub description: Option<String>,
    pub required: Option<bool>,
    pub deprecated: Option<bool>,
    pub allow_empty_value: Option<bool>,
    pub style: Option<String>,
    pub explode: Option<bool>,
    pub allow_reserved: Option<bool>,
    pub schema: Option<Schema>,
    pub example: Option<serde_json::Value>,
    pub examples: Option<HashMap<String, Example>>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Example {
    pub summary: Option<String>,
    pub description: Option<String>,
    pub value: serde_json::Value,
    pub external_value: Option<String>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct RequestBody {
    pub description: Option<String>,
    pub content: HashMap<String, MediaType>,
    pub required: Option<bool>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Responses {
    pub default: Option<Response>,
    pub status: Option<HashMap<String, Response>>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Response {
    pub description: Option<String>,
    pub headers: Option<HashMap<String, Header>>,
    pub content: Option<HashMap<String, MediaType>>,
    pub links: Option<HashMap<String, Link>>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Header {
    pub description: Option<String>,
    pub required: Option<bool>,
    pub deprecated: Option<bool>,
    pub allow_empty_value: Option<bool>,
    pub style: Option<String>,
    pub explode: Option<bool>,
    pub allow_reserved: Option<bool>,
    pub schema: Option<Schema>,
    pub example: Option<serde_json::Value>,
    pub examples: Option<HashMap<String, Example>>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct MediaType {
    pub schema: Option<Schema>,
    pub example: Option<serde_json::Value>,
    pub examples: Option<HashMap<String, Example>>,
    pub encoding: Option<HashMap<String, Encoding>>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Encoding {
    pub content_type: Option<String>,
    pub headers: Option<HashMap<String, Header>>,
    pub style: Option<String>,
    pub explode: Option<bool>,
    pub allow_reserved: Option<bool>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Components {
    pub schemas: Option<HashMap<String, Schema>>,
    pub responses: Option<HashMap<String, Response>>,
    pub parameters: Option<HashMap<String, Parameter>>,
    pub examples: Option<HashMap<String, Example>>,
    pub request_bodies: Option<HashMap<String, RequestBody>>,
    pub headers: Option<HashMap<String, Header>>,
    pub security_schemes: Option<HashMap<String, SecurityScheme>>,
    pub links: Option<HashMap<String, Link>>,
    pub callbacks: Option<HashMap<String, Callback>>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct SecurityScheme {
    pub r#type: String,
    pub description: Option<String>,
    pub name: Option<String>,
    pub in_field: Option<String>,
    pub scheme: Option<String>,
    pub bearer_format: Option<String>,
    pub flows: Option<OAuthFlows>,
    pub open_id_connect_url: Option<String>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct OAuthFlows {
    pub implicit: Option<OAuthFlow>,
    pub password: Option<OAuthFlow>,
    pub client_credentials: Option<OAuthFlow>,
    pub authorization_code: Option<OAuthFlow>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct OAuthFlow {
    pub authorization_url: Option<String>,
    pub token_url: Option<String>,
    pub refresh_url: Option<String>,
    pub scopes: Option<HashMap<String, String>>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct SecurityRequirement {
    pub schemes: HashMap<String, Vec<String>>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Server {
    pub url: String,
    pub description: Option<String>,
    pub variables: Option<HashMap<String, ServerVariable>>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct ServerVariable {
    pub r#enum: Option<Vec<String>>,
    pub default: String,
    pub description: Option<String>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Link {
    pub operation_ref: Option<String>,
    pub operation_id: Option<String>,
    pub parameters: Option<HashMap<String, serde_json::Value>>,
    pub request_body: Option<serde_json::Value>,
    pub description: Option<String>,
    pub server: Option<Server>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Callback {
    pub expression: String,
    pub callback: HashMap<String, PathItem>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Schema {
    pub title: Option<String>,
    pub multiple_of: Option<f64>,
    pub maximum: Option<f64>,
    pub exclusive_maximum: Option<bool>,
    pub minimum: Option<f64>,
    pub exclusive_minimum: Option<bool>,
    pub r#type: Option<String>,
    pub items: Option<Box<Schema>>,
    pub max_items: Option<usize>,
    pub min_items: Option<usize>,
    pub unique_items: Option<bool>,
    pub max_properties: Option<usize>,
    pub min_properties: Option<usize>,
    pub required: Option<Vec<String>>,
    pub properties: Option<HashMap<String, Schema>>,
    pub all_of: Option<Vec<Box<Schema>>>,
    pub any_of: Option<Vec<Box<Schema>>>,
    pub one_of: Option<Vec<Box<Schema>>>,
    pub not: Option<Box<Schema>>,
    pub additional_properties: Option<Box<Schema>>,
    pub description: Option<String>,
    pub format: Option<String>,
    pub default: Option<serde_json::Value>,
    pub nullable: Option<bool>,
    pub discriminator: Option<Discriminator>,
    pub r#enum: Option<Vec<serde_json::Value>>,
    pub deprecated: Option<bool>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}

#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct Discriminator {
    pub property_name: String,
    pub mapping: Option<HashMap<String, String>>,
    #[serde(flatten)]
    pub extensions: HashMap<String, serde_json::Value>,
}


// my own implementation
#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct RequestTree {
    pub handlers: HashMap<String, Vec<RequestHandler>>,
}

impl RequestTree {
    #[must_use]
    pub fn new() -> Self {
        Self {
            handlers: HashMap::new(),
        }
    }
}

impl Default for RequestTree {
    fn default() -> Self {
        Self::new()
    }
}


#[derive(Clone,Debug, Serialize, Deserialize)]
pub struct RequestHandler {
    pub method: Option<String>,
    pub path: Option<String>,
    pub parameters: Option<Vec<Parameter>>,
    pub request_body: Option<RequestBody>,
    pub responses: Option<Responses>,
    pub summary: Option<String>,
    pub description: Option<String>,
    pub tags: Option<Vec<String>>,
}

impl From<&Operation> for RequestHandler {
    fn from(operation: &Operation) -> Self {
        Self {
            method: None,
            path: None,
            parameters: operation.parameters.clone(),
            request_body: operation.request_body.clone(),
            responses: operation.responses.clone(),
            summary: operation.summary.clone(),
            description: operation.description.clone(),
            tags: operation.tags.clone(),
        }
    }
}
