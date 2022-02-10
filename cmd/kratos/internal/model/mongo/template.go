package mongo

import (
	"html/template"
	"bytes"
)

var repoTemplate = `
{{- /* delete empty line */ -}}
package data

import (
	"context"
	"gitlab.intra.knownsec.com/sysarch/gous/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *Database) Create{{.ModelName}}(ctx context.Context, entity *model.{{.ModelName}}) error {
	if _, err := d.DB.Collection("servers").
		InsertOne(ctx, entity); err != nil {
		logger.Errorf("failed to create entity: %v", err)
		return err
	}
	return nil
}
`

// Service is a proto service.
type Repo struct {
	ModelName     string
	Fields     []string
}

func (s *Repo) execute() ([]byte, error) {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("mongo").Parse(repoTemplate)
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
