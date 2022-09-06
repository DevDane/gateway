package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"devdane.com/gateway/graph/generated"
	"devdane.com/gateway/graph/model"
	"devdane.com/gateway/repos"
)

// Repos is the resolver for the repos field.
func (r *queryResolver) Repos(ctx context.Context, limit *int) ([]*model.Repo, error) {
	res := repos.Repos(limit)
	repos := make([]*model.Repo, len(res.Repos))

	for i, repo := range res.Repos {
		repos[i] = &model.Repo{
			ID:              int(repo.Id),
			NodeID:          repo.NodeId,
			Name:            repo.Name,
			FullName:        repo.FullName,
			Description:     repo.Description,
			HTMLURL:         repo.HtmlUrl,
			URL:             repo.Url,
			StargazersCount: int(repo.StargazersCount),
			WatchersCount:   int(repo.WatchersCount),
			Visibility:      repo.Visibility,
			Languages:       repo.Language,
			ImageURL:        repo.ImageUrl,
		}
	}

	return repos, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
