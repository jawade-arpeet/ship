package client

import (
	"ship/internal/config"

	"github.com/resend/resend-go/v3"
)

type ResendClient struct {
	client *resend.Client
}

func newResendClient(cfg *config.ResendConfig) *ResendClient {
	client := resend.NewClient(cfg.APIKey)
	return &ResendClient{client: client}
}
