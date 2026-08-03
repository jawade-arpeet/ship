package repository

import "ship/internal/client"

type Repository struct{}

func New(client *client.Client) *Repository {
	return &Repository{}
}
