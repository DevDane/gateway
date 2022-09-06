package repos

import (
	"context"

	"devdane.com/gateway/internal/gsvc"
	pb "devdane.com/gateway/repos/proto"
)

func Repos(limit *int) *pb.RepoResponse {
	client := pb.NewReposServiceClient(gsvc.CreateRepoClient("localhost:4000"))

	var lim int64 = -1
	if limit != nil {
		lim = int64(*limit)
	}
	res, err := client.GetRepos(context.Background(), &pb.RepoRequest{
		Limit: &lim,
	})
	if err != nil {
		panic(err)
	}

	return res
}
