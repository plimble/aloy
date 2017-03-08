package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/plimble/aloy/services/coverage/entity"
	"github.com/plimble/errors"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	s := &Store{db}

	s.CreateRepositorysTable()
	s.CreateCommitsTable()

	return s
}

func (s *Store) CreateRepositorysTable() {
	s.db.Exec(`CREATE TABLE repositorys (
		id varchar(30) NOT NULL,
		name varchar(30) NOT NULL,
		owner_name varchar(30) NOT NULL,
		source varchar(30) NOT NULL,
		homepage text,
		description text,
		created_at timestamp without time zone,
		CONSTRAINT repositorys_id_pk PRIMARY KEY (id)
	)`)
}

func (s *Store) CreateCommitsTable() {
	s.db.Exec(`CREATE TABLE commits (
		id varchar(30) NOT NULL,
		repository_id varchar(30) NOT NULL,
		ref varchar(30) NOT NULL,
		sender_name varchar(30) NOT NULL,
		sender_avatar varchar(200),
		created_at timestamp without time zone,
		CONSTRAINT commits_id_pk PRIMARY KEY (id)
	)`)

	s.db.Exec("CREATE INDEX commits_repositoryId_idx ON commits (repository_id)")
}

func (s *Store) CreateRepository(repository *entity.Repository) error {
	_, err := s.db.Exec(`INSERT INTO repositorys(id, name, owner_name, soruce, homepage, description, created_at) VALUES(?, ?, ?, ?, ?, ?, ?)`, repository.Id, repository.Name, repository.OwnerName, repository.Source, repository.HomePage, repository.Description, repository.CreatedAt)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *Store) CreateCommit(commit *entity.Commit) error {
	_, err := s.db.Exec(`INSERT INTO commits(id, repository_id, ref, sender_name, sender_avatar, created_at) VALUES(?, ?, ?, ?, ?, ?)`, commit.Id, commit.RepositoryId, commit.Ref, commit.SenderName, commit.SenderAvatar, commit.CreatedAt)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *Store) DeleteRepositoryAndCommits(repositoryId string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM commits WHERE repository_id=?`, repositoryId)
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}

	_, err = tx.Exec(`DELETE FROM repositorys WHERE id=?`, repositoryId)
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}

	return nil
}

func (s *Store) DeleteCommit(commitId string) error {
	_, err := s.db.Exec(`DELETE FROM commits WHERE id=?`, commitId)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *Store) GetRepositoryByName(repositoryName string) (*entity.Repository, error) {
	var repository *entity.Repository
	err := s.db.QueryRow("SELECT * FROM repositorys WHERE name=?", repositoryName).Scan(&repository)

	if err == sql.ErrNoRows {
		return nil, errors.NotFound(err.Error())
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return repository, nil
}

func (s *Store) GetAllRepositorys(limit, offset int) ([]*entity.Repository, error) {
	var repositorys []*entity.Repository
	rows, err := s.db.Query("SELECT * FROM repositorys LIMIT ? OFFSET ? ORDER BY created_at DESC", limit, offset)
	defer rows.Close()

	if err != nil && err != sql.ErrNoRows {
		return nil, errors.WithStack(err)
	}

	for rows.Next() {
		var repository *entity.Repository
		if err := rows.Scan(&repository); err != nil {
			return nil, errors.WithStack(err)
		}

		repositorys = append(repositorys, repository)
	}

	return repositorys, nil

}

func (s *Store) GetAllCommitsByRepository(repositoryId string, limit, offset int) ([]*entity.Commit, error) {
	var commits []*entity.Commit
	rows, err := s.db.Query("SELECT * FROM commits WHERE repository_id=? LIMIT ? OFFSET ? ORDER BY created_at DESC", repositoryId, limit, offset)
	defer rows.Close()

	if err != nil && err != sql.ErrNoRows {
		return nil, errors.WithStack(err)
	}

	for rows.Next() {
		var commit *entity.Commit
		if err := rows.Scan(&commit); err != nil {
			return nil, errors.WithStack(err)
		}

		commits = append(commits, commit)
	}

	return commits, nil
}

func (s *Store) GetAllCommitsByName(name string, limit, offset int) ([]*entity.Commit, error) {
	var commits []*entity.Commit
	rows, err := s.db.Query("SELECT * FROM commits WHERE sender_name=? LIMIT ? OFFSET ? ORDER BY created_at DESC", name, limit, offset)
	defer rows.Close()

	if err != nil && err != sql.ErrNoRows {
		return nil, errors.WithStack(err)
	}

	for rows.Next() {
		var commit *entity.Commit
		if err := rows.Scan(&commit); err != nil {
			return nil, errors.WithStack(err)
		}

		commits = append(commits, commit)
	}

	return commits, nil
}

func (s *Store) GetAllCommitsByRepositoryAndRef(repositoryId, ref string, limit, offset int) ([]*entity.Commit, error) {
	var commits []*entity.Commit
	rows, err := s.db.Query("SELECT * FROM commits WHERE repository_id=? AND ref=? LIMIT ? OFFSET ? ORDER BY created_at DESC", repositoryId, ref, limit, offset)
	defer rows.Close()

	if err != nil && err != sql.ErrNoRows {
		return nil, errors.WithStack(err)
	}

	for rows.Next() {
		var commit *entity.Commit
		if err := rows.Scan(&commit); err != nil {
			return nil, errors.WithStack(err)
		}

		commits = append(commits, commit)
	}

	return commits, nil
}
