package coverage

type DeleteCommitReq struct {
	CommitId string
}

type DeleteCommitRes struct {
}

func (cs *CoverageService) DeleteCommit(req *DeleteCommitReq) (*DeleteCommitRes, error) {
	if err := cs.store.DeleteCommit(req.CommitId); err != nil {
		return nil, err
	}

	return &DeleteCommitRes{}, nil
}
