package coverage

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"strings"

	"github.com/plimble/errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Store interface {
	GetCoverateResult(repo, commitID string) (*CoverateResult, error)
	GetCommits(repo, ref, last string, limit int) ([]*Commit, error)
	SaveCommit(commit *Commit) error
	SaveCoverateResult(result *CoverateResult) error
}

type leveldbstore struct {
	*leveldb.DB
}

func NewLeveldbStore(path string) (*leveldbstore, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &leveldbstore{db}, nil
}

func (db *leveldbstore) GetCoverateResult(repo, commitID string) (*CoverateResult, error) {
	key := []byte(fmt.Sprintf("%s:%s", repo, commitID))
	data, err := db.Get(key, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	b := bytes.NewBuffer(data)
	dec := gob.NewDecoder(b)
	result := &CoverateResult{}
	err = dec.Decode(result)

	return result, errors.WithStack(err)
}

func (db *leveldbstore) GetCommits(repo, ref, last string, limit int) ([]*Commit, error) {
	var key []byte
	var err error
	var rag *util.Range

	if last != "" {
		idx := strings.LastIndex(last, ":")
		key = []byte(last[:idx])
		limit := fmt.Sprint(last[:idx], ":", "\u02ad")
		rag = &util.Range{Start: []byte(last), Limit: []byte(limit)}
	} else if ref != "" {
		key = []byte(fmt.Sprintf("%s:%s", repo, ref))
		rag = util.BytesPrefix(key)
	} else {
		key = []byte(repo)
		rag = util.BytesPrefix(key)
	}

	var b *bytes.Buffer
	var dec *gob.Decoder
	var commit *Commit
	count := 0
	commits := make([]*Commit, 0)

	iter := db.NewIterator(rag, nil)
	for iter.Next() {
		if count == limit {
			break
		}
		if last != "" {
			last = ""
			continue
		}

		b = bytes.NewBuffer(iter.Value())
		dec = gob.NewDecoder(b)

		commit = &Commit{}
		if err = dec.Decode(commit); err != nil {
			return commits, errors.WithStack(err)
		}

		commits = append(commits, commit)
		count++
	}
	iter.Release()

	return commits, errors.WithStack(iter.Error())
}

func (db *leveldbstore) SaveCommit(commit *Commit) error {
	var err error

	repcer := strings.NewReplacer("9", "0", "8", "1", "7", "2", "6", "3", "5", "4", "4", "5", "3", "6", "2", "7", "1", "8", "0", "9")
	strTime := repcer.Replace(fmt.Sprintf("%d", commit.Timestamp))

	key := []byte(fmt.Sprintf("%s:%s:%s", commit.Repo, commit.Ref, strTime))

	b := bytes.NewBuffer(nil)

	enc := gob.NewEncoder(b)
	if err = enc.Encode(commit); err != nil {
		return errors.WithStack(err)
	}

	err = db.Put(key, b.Bytes(), nil)

	return errors.WithStack(err)
}

func (db *leveldbstore) SaveCoverateResult(result *CoverateResult) error {
	var err error
	key := []byte(fmt.Sprintf("%s:%s", result.Repo, result.CommitID))

	b := bytes.NewBuffer(nil)

	enc := gob.NewEncoder(b)
	if err = enc.Encode(result); err != nil {
		return errors.WithStack(err)
	}

	err = db.Put(key, b.Bytes(), nil)

	return errors.WithStack(err)
}
