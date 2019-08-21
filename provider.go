package main

import (
	"compress/gzip"
	"context"
	"encoding/csv"
	"io"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
)

func provide(ctx context.Context, p *publisher, bucketName, objectURI string) (count int64, err error) {

	c, e := storage.NewClient(ctx)
	if e != nil {
		return 0, errors.Wrap(e, "Failed to create client")
	}
	defer c.Close()

	bucket := c.Bucket(bucketName)
	obj := bucket.Object(objectURI).ReadCompressed(true)
	rdr, e := obj.NewReader(ctx)
	if e != nil {
		return 0, errors.Wrap(e, "Error creating object reader")
	}
	defer rdr.Close()

	gzr, e := gzip.NewReader(rdr)
	if e != nil {
		return 0, errors.Wrap(e, "Error creating gzip reader")
	}
	defer gzr.Close()

	csvr := csv.NewReader(gzr)
	var recCount int64

	for {
		values, e := csvr.Read()

		if e == io.EOF {
			break
		}

		if e != nil {
			return 0, errors.Wrap(e, "Error reading CSV line")
		}

		// Pringint purely for demo logging effect
		line := strings.Join(values, "|")
		recCount++
		logger.Printf("Line[%d] %s", recCount, line)

		p.publish(ctx, []byte(line))

	}

	return recCount, nil

}
