package controller

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/shimanekb/project1a/store"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

const (
	GET_COMMAND       string = "get"
	PUT_COMMAND       string = "put"
	DEL_COMMAND       string = "del"
	FIRST_LINE_RECORD string = "type"
)

type Command struct {
	Type  string
	Key   string
	Value string
}

func ReadCsvCommands(file_path string) {
	csv_file, err := os.Open(file_path)

	log.Infof("Opening csv file %s", file_path)

	if err != nil {
		log.Fatalln("FATAL: Could not open csv file.", err)
	}

	reader := csv.NewReader(csv_file)
	kvStore := kvstore.NewKvStore()

	log.Infoln("Reading in csv records.")
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if record[0] == FIRST_LINE_RECORD {
			log.Infoln("First line detected, skipping.")
			continue
		}
		command := Command{record[0], record[1], record[3]}
		cmd_err := ProcessCommand(command, kvStore)
		if cmd_err != nil {
			log.Errorln(cmd_err)
		}
	}
}

func ProcessCommand(command Command, storage kvstore.Store) error {
	switch {
	case GET_COMMAND == command.Type:
		log.Infof("Get command given for key: %s, value: %s", command.Key,
			command.Value)
		value, err := storage.Get(command.Key)
		if err == nil {
			log.Infof("Get command successful found value: %s, for key: %s",
				value, command.Key)
		}

		return err
	case PUT_COMMAND == command.Type:
		log.Infof("Put command given for key: %s, value: %s", command.Key,
			command.Value)
		return storage.Put(command.Key, command.Value)
	case DEL_COMMAND == command.Type:
		log.Infof("Del command given for key: %s, value: %s", command.Key,
			command.Value)
		return storage.Del(command.Key)
	}

	return errors.New(fmt.Sprintf("Invalid command given: %s", command))
}
