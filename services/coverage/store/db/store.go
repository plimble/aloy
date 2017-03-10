package db

import (
	"database/sql"
	"fmt"
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
	_, err := s.db.Exec(`CREATE TABLE repositorys (
		id varchar(30) NOT NULL PRIMARY KEY,
		name varchar(30) NOT NULL,
		owner_name varchar(30) NOT NULL,
		source varchar(30) NOT NULL,
		homepage text,
		description text,
		created_at varchar(50)
	)`)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Store) CreateCommitsTable() {
	s.db.Exec(`CREATE TABLE commits (
		id varchar(30) NOT NULL PRIMARY KEY,
		repository_id varchar(30) NOT NULL,
		ref varchar(30) NOT NULL,
		sender_name varchar(30) NOT NULL,
		sender_avatar varchar(200),
		created_at varchar(50)
	)`)

	// s.db.Exec("CREATE INDEX commits_repositoryId_idx ON commits (repository_id)")
}

func (s *Store) CreateRepository(repository *entity.Repository) error {
	_, err := s.db.Exec(`INSERT INTO repositorys(id, name, owner_name, source, homepage, description, created_at) VALUES(?, ?, ?, ?, ?, ?, ?)`, repository.Id, repository.Name, repository.OwnerName, repository.Source, repository.HomePage, repository.Description, repository.CreatedAt)

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

func (s *Store) GetRepository(repositoryName, repositoryOwnerName, repositorySource string) (*entity.Repository, error) {
	repository := &entity.Repository{}
	rows, err := s.db.Query("SELECT * FROM repositorys WHERE name=? AND owner_name=? AND source=?", repositoryName, repositoryOwnerName, repositorySource)
	defer rows.Close()

	if err == sql.ErrNoRows {
		return nil, errors.NotFound(err.Error())
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	for rows.Next() {
		if err := rows.Scan(&repository.Id, &repository.Name, &repository.OwnerName, &repository.Source, &repository.HomePage, &repository.Description, &repository.CreatedAt); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return repository, nil
}

func (s *Store) GetAllRepositorys(limit, offset int) ([]*entity.Repository, error) {
	repositorys := []*entity.Repository{}

	rows, err := s.db.Query("SELECT * FROM repositorys ORDER BY created_at DESC LIMIT ? OFFSET ?", limit, offset)
	defer rows.Close()

	if err != nil && err != sql.ErrNoRows {
		return nil, errors.WithStack(err)
	}

	for rows.Next() {
		repository := &entity.Repository{}

		if err := rows.Scan(&repository.Id, &repository.Name, &repository.OwnerName, &repository.Source, &repository.HomePage, &repository.Description, &repository.CreatedAt); err != nil {
			return nil, errors.WithStack(err)
		}

		repositorys = append(repositorys, repository)
	}

	return repositorys, nil

}

func (s *Store) GetAllCommitsByRepository(repositoryId string, limit, offset int) ([]*entity.Commit, error) {
	commits := []*entity.Commit{}
	rows, err := s.db.Query("SELECT * FROM commits WHERE repository_id=? ORDER BY created_at DESC  LIMIT ? OFFSET ?", repositoryId, limit, offset)
	defer rows.Close()

	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return nil, errors.WithStack(err)
	}

	for rows.Next() {
		commit := &entity.Commit{}

		if err := rows.Scan(&commit.Id, &commit.RepositoryId, &commit.Ref, &commit.SenderName, &commit.SenderAvatar, &commit.CreatedAt); err != nil {
			return nil, errors.WithStack(err)
		}

		commits = append(commits, commit)
	}

	return commits, nil
}

func (s *Store) GetAllCommitsByName(name string, limit, offset int) ([]*entity.Commit, error) {
	commits := []*entity.Commit{}
	rows, err := s.db.Query("SELECT * FROM commits WHERE sender_name=? ORDER BY created_at DESC  LIMIT ? OFFSET ?", name, limit, offset)
	defer rows.Close()

	if err != nil && err != sql.ErrNoRows {
		return nil, errors.WithStack(err)
	}

	for rows.Next() {
		commit := &entity.Commit{}
		if err := rows.Scan(&commit.Id, &commit.RepositoryId, &commit.Ref, &commit.SenderName, &commit.SenderAvatar, &commit.CreatedAt); err != nil {
			return nil, errors.WithStack(err)
		}

		commits = append(commits, commit)
	}

	return commits, nil
}

func (s *Store) GetAllCommitsByRepositoryAndRef(repositoryId, ref string, limit, offset int) ([]*entity.Commit, error) {
	commits := []*entity.Commit{}
	rows, err := s.db.Query("SELECT * FROM commits WHERE repository_id=? AND ref=? ORDER BY created_at DESC LIMIT ? OFFSET ?", repositoryId, ref, limit, offset)
	defer rows.Close()

	if err != nil && err != sql.ErrNoRows {
		return nil, errors.WithStack(err)
	}

	for rows.Next() {
		commit := &entity.Commit{}
		if err := rows.Scan(&commit.Id, &commit.RepositoryId, &commit.Ref, &commit.SenderName, &commit.SenderAvatar, &commit.CreatedAt); err != nil {
			return nil, errors.WithStack(err)
		}

		commits = append(commits, commit)
	}

	return commits, nil
}
