package main

import (
	"fmt"
	"sync"

	"time"

	"github.com/sirupsen/logrus"
)

type jobUpdate struct {
	filename string
	line     int
	row      map[string]string
	handler  HandlerFunc
	logger   logrus.FieldLogger
	wg       *sync.WaitGroup
}

func (job jobUpdate) Name() string {
	return fmt.Sprintf("update-%s:%d", job.filename, job.line)
}

func (job jobUpdate) Do() error {
	start := time.Now()
	defer job.wg.Done()

	err := job.handler(job.row)
	if err != nil {
		return err
	}

	job.logger.Debugf("Updated %s:%d in %.2fms", job.filename, job.line, time.Since(start).Seconds()*1000)

	return nil
}

func NewUpdate(filename string, line int, row map[string]string, handler HandlerFunc, logger logrus.FieldLogger, wg *sync.WaitGroup) Job {
	return &jobUpdate{
		filename: filename,
		line:     line,
		row:      row,
		handler:  handler,
		logger:   logger.WithField("component", "update"),
		wg:       wg,
	}
}
