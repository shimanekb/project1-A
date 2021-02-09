package kvstore

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

const (
	STORAGE_DIR  string = "storage"
	STORAGE_FILE string = "kv_store.csv"
)

type Store interface {
	Put(key string, value string) error
	Get(key string) (string, error)
	Del(key string) error
}

type KvStore struct {
	Storage_File *os.File
}

func (k KvStore) Put(key string, value string) error {
	return WritePut(k.Storage_File, key, value)
}

func (k KvStore) Get(key string) (string, error) {
	return "", nil
}

func (k KvStore) Del(key string) error {
	return nil
}

func NewKvStore() *KvStore {
	log.Info("Creating new Kv Store.")

	log.Info("Creating storage directory if does not exist.")
	newpath := filepath.Join(".", STORAGE_DIR)
	err := os.MkdirAll(newpath, os.ModePerm)

	if err != nil {
		log.Fatalf("Cannot create directory for storage at %s", STORAGE_DIR)
	}
	log.Info("Created storage directory.")

	log.Info("Creating storage file if not exist.")
	newpath = filepath.Join(newpath, STORAGE_FILE)
	store_file, file_err := os.OpenFile(newpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if file_err != nil {
		log.Fatalf("Cannot create file for %s", newpath)
	}
	log.Info("Created storage file.")

	return &KvStore{store_file}
}

func WritePut(store_file *os.File, key string, value string) error {
	_, err := store_file.WriteString(fmt.Sprintf("%s,%s\n", key, value))
	return err
}
