package coverage

type DeleteRepositoryAndCommitsReq struct {
	RepositoryId string
}

type DeleteRepositoryAndCommitsRes struct {
}

func (cs *CoverageService) DeleteRepositoryAndCommits(req *DeleteRepositoryAndCommitsReq) (*DeleteRepositoryAndCommitsRes, error) {
	if err := cs.store.DeleteRepositoryAndCommits(req.RepositoryId); err != nil {
		return nil, err
	}

	return &DeleteRepositoryAndCommitsRes{}, nil
}
