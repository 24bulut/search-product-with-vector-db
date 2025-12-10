package Structures

// EmbeddingUsage contains token usage information from OpenAI
type EmbeddingUsage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

// EmbeddingResult contains the embedding and usage information
type EmbeddingResult struct {
	Embedding []float32      `json:"embedding"`
	Model     string         `json:"model"`
	Usage     EmbeddingUsage `json:"usage"`
}
