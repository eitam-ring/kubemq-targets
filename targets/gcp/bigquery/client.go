package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"encoding/json"
	"fmt"
	"github.com/kubemq-hub/kubemq-target-connectors/config"
	"github.com/kubemq-hub/kubemq-target-connectors/pkg/logger"
	"github.com/kubemq-hub/kubemq-target-connectors/targets"
	"github.com/kubemq-hub/kubemq-target-connectors/types"
	"google.golang.org/api/iterator"
)

type Client struct {
	name   string
	opts   options
	client *bigquery.Client
	log    *logger.Logger
	target targets.Target
}

func New() *Client {
	return &Client{}

}
func (c *Client) Name() string {
	return c.name
}

func (c *Client) Init(ctx context.Context, cfg config.Metadata) error {
	c.name = cfg.Name
	c.log = logger.NewLogger(cfg.Name)
	var err error
	c.opts, err = parseOptions(cfg)
	if err != nil {
		return err
	}

	Client, err := bigquery.NewClient(ctx, c.opts.projectID)
	if err != nil {
		return err
	}
	c.client = Client
	return nil
}

func (c *Client) Do(ctx context.Context, req *types.Request) (*types.Response, error) {
	meta, err := parseMetadata(req.Metadata)
	if err != nil {
		return nil, err
	}
	switch meta.method {
	case "query":
		return c.query(ctx, meta)
	case "create_table":
		return c.createTable(ctx, meta, req.Data)
	case "get_table_info":
		return c.getTableInfo(ctx, meta)
	case "get_data_sets":
		return c.getDataSets(ctx)
	default:
		return nil, fmt.Errorf(getValidMethodTypes())
	}
}

func (c *Client) getTableInfo(ctx context.Context, meta metadata) (*types.Response, error) {
	m, err := c.client.Dataset(meta.datasetID).Table(meta.tableName).Metadata(ctx)
	if err != nil {
		if err != nil {
			return types.NewResponse().
				SetMetadataKeyValue("error", "true").
				SetMetadataKeyValue("message", err.Error()), nil
		}
	}
	b, err := json.Marshal(m)
	if err != nil {
		return types.NewResponse().
			SetMetadataKeyValue("error", "true").
			SetMetadataKeyValue("message", err.Error()), nil
	}
	return types.NewResponse().
			SetMetadataKeyValue("result", "ok").
			SetData(b),
		nil
}

func (c *Client) getDataSets(ctx context.Context) (*types.Response, error) {
	i := c.client.Datasets(ctx)
	s, err := c.getDataSetsFromIterator(i)
	if err != nil {
		if err != nil {
			return types.NewResponse().
				SetMetadataKeyValue("error", "true").
				SetMetadataKeyValue("message", err.Error()), nil
		}
	}
	b, err := json.Marshal(s)
	if err != nil {
		return types.NewResponse().
			SetMetadataKeyValue("error", "true").
			SetMetadataKeyValue("message", err.Error()), nil
	}
	if len(s) == 0 {
		return types.NewResponse().
			SetMetadataKeyValue("error", "true").
			SetMetadataKeyValue("message", "no datasets found"), nil
	}
	return types.NewResponse().
			SetMetadataKeyValue("result", "ok").
			SetData(b),
		nil
}

func (c *Client) query(ctx context.Context, meta metadata) (*types.Response, error) {
	query := c.client.Query(meta.query)
	i, err := query.Read(ctx)
	if err != nil {
		return types.NewResponse().
			SetMetadataKeyValue("error", "true").
			SetMetadataKeyValue("message", err.Error()), nil
	}
	rows, err := c.getRowsFromIterator(i)
	if err != nil {
		return types.NewResponse().
			SetMetadataKeyValue("error", "true").
			SetMetadataKeyValue("message", err.Error()), nil
	}
	if len(rows) == 0 {
		return types.NewResponse().
			SetMetadataKeyValue("error", "true").
			SetMetadataKeyValue("message", "no rows found for this query"), nil
	}
	b, err := json.Marshal(rows)
	if err != nil {
		return types.NewResponse().
			SetMetadataKeyValue("error", "true").
			SetMetadataKeyValue("message", err.Error()), nil
	}
	return types.NewResponse().
			SetMetadataKeyValue("result", "ok").
			SetData(b),
		nil
}

func (c *Client) createTable(ctx context.Context, meta metadata, body []byte) (*types.Response, error) {
	metaData := &bigquery.TableMetadata{}

	err := json.Unmarshal(body, &metaData)
	if err != nil {
		return types.NewResponse().
			SetMetadataKeyValue("error", "true").
			SetMetadataKeyValue("message", err.Error()), nil
	}
	tableRef := c.client.Dataset(meta.datasetID).Table(meta.tableName)
	err = tableRef.Create(ctx, metaData)
	if err != nil {
		return types.NewResponse().
			SetMetadataKeyValue("error", "true").
			SetMetadataKeyValue("message", err.Error()), nil
	}
	return types.NewResponse().
			SetMetadataKeyValue("result", "ok"),
		nil
}

func (c *Client) getRowsFromIterator(iter *bigquery.RowIterator) ([]map[string]bigquery.Value, error) {
	var rows []map[string]bigquery.Value
	for {
		row := make(map[string]bigquery.Value)
		err := iter.Next(&row)
		if err == iterator.Done {
			return rows, nil
		}
		if err != nil {
			return rows, fmt.Errorf("error iterating through results: %v", err)
		}
		rows = append(rows, row)
	}
}

func (c *Client) getDataSetsFromIterator(iter *bigquery.DatasetIterator) ([]*bigquery.Dataset, error) {
	var datasets []*bigquery.Dataset
	for {
		dataset, err := iter.Next()
		if err == iterator.Done {
			return datasets, nil
		}
		if err != nil {
			return datasets, fmt.Errorf("error iterating through results: %v", err)
		}
		datasets = append(datasets, dataset)
	}
}

func (c *Client) CloseClient() error {
	return c.client.Close()
}