package kvstore

import (
	log "github.com/sirupsen/logrus"
)

type Store interface {
	Put(key string, value string) error
	Get(key string) (string, error)
	Del(key string) error
}

type KvStore struct {
}

func (k KvStore) Put(key string, value string) error {
	return nil
}

func (k KvStore) Get(key string) (string, error) {
	return "", nil
}

func (k KvStore) Del(key string) error {
	return nil
}

func NewKvStore() *KvStore {
	log.Info("Creating new Kv Store.")
	return &KvStore{}
}
