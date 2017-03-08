package entity

import (
	"time"
)

type Commit struct {
	CommitId     string
	RepositoryId string
	Ref          string

	SenderName   string
	SenderAvatar string

	CreatedAt string
}

func NewCommit(commitId, repositoryId, ref, senderName, senderAvatar string) *Commit {
	return &Commit{
		CommitId:     commitId,
		RepositoryId: repositoryId,
		Ref:          ref,
		SenderName:   senderName,
		SenderAvatar: senderAvatar,
		CreatedAt:    time.Now().Truncate(time.Second).Format(time.RFC3339),
	}
}