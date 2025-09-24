package types

type Provider string

const (
	Gemini Provider = "gemini"
	// OpenAI Provider = "openai" // TODO: Add support for OpenAI in the future
)

type ServerConfig struct {
	AuthorizationKey string
	Port             int
	Provider         Provider
	Model            string
	APIKey           string
}

type Identity struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type ComposeDetails struct {
	Subject    string   `json:"subject,omitempty"`
	To         []string `json:"to,omitempty"`
	Cc         []string `json:"cc,omitempty"`
	Bcc        []string `json:"bcc,omitempty"`
	BodyPlain  string   `json:"bodyPlain,omitempty"`
	BodyHTML   string   `json:"bodyHTML,omitempty"`
	IdentityID string   `json:"identityId,omitempty"`
	IsHTML     bool     `json:"isHTML"`
}

type ComposeContext struct {
	Account Identity       `json:"account"`
	Compose ComposeDetails `json:"compose"`
}

type Payload struct {
	Prompt  string         `json:"prompt"`
	Context ComposeContext `json:"context"`
}

type SuccessResponse struct {
	Response string  `json:"response"`
	Payload  Payload `json:"payload"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
