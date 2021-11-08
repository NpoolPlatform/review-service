// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/review-service/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/review-service/pkg/db/ent/review"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/reviewrule"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Review is the client for interacting with the Review builders.
	Review *ReviewClient
	// ReviewRule is the client for interacting with the ReviewRule builders.
	ReviewRule *ReviewRuleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Review = NewReviewClient(c.config)
	c.ReviewRule = NewReviewRuleClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Review:     NewReviewClient(cfg),
		ReviewRule: NewReviewRuleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:     cfg,
		Review:     NewReviewClient(cfg),
		ReviewRule: NewReviewRuleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Review.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Review.Use(hooks...)
	c.ReviewRule.Use(hooks...)
}

// ReviewClient is a client for the Review schema.
type ReviewClient struct {
	config
}

// NewReviewClient returns a client for the Review from the given config.
func NewReviewClient(c config) *ReviewClient {
	return &ReviewClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `review.Hooks(f(g(h())))`.
func (c *ReviewClient) Use(hooks ...Hook) {
	c.hooks.Review = append(c.hooks.Review, hooks...)
}

// Create returns a create builder for Review.
func (c *ReviewClient) Create() *ReviewCreate {
	mutation := newReviewMutation(c.config, OpCreate)
	return &ReviewCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Review entities.
func (c *ReviewClient) CreateBulk(builders ...*ReviewCreate) *ReviewCreateBulk {
	return &ReviewCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Review.
func (c *ReviewClient) Update() *ReviewUpdate {
	mutation := newReviewMutation(c.config, OpUpdate)
	return &ReviewUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReviewClient) UpdateOne(r *Review) *ReviewUpdateOne {
	mutation := newReviewMutation(c.config, OpUpdateOne, withReview(r))
	return &ReviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReviewClient) UpdateOneID(id uuid.UUID) *ReviewUpdateOne {
	mutation := newReviewMutation(c.config, OpUpdateOne, withReviewID(id))
	return &ReviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Review.
func (c *ReviewClient) Delete() *ReviewDelete {
	mutation := newReviewMutation(c.config, OpDelete)
	return &ReviewDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ReviewClient) DeleteOne(r *Review) *ReviewDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ReviewClient) DeleteOneID(id uuid.UUID) *ReviewDeleteOne {
	builder := c.Delete().Where(review.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReviewDeleteOne{builder}
}

// Query returns a query builder for Review.
func (c *ReviewClient) Query() *ReviewQuery {
	return &ReviewQuery{
		config: c.config,
	}
}

// Get returns a Review entity by its id.
func (c *ReviewClient) Get(ctx context.Context, id uuid.UUID) (*Review, error) {
	return c.Query().Where(review.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReviewClient) GetX(ctx context.Context, id uuid.UUID) *Review {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ReviewClient) Hooks() []Hook {
	return c.hooks.Review
}

// ReviewRuleClient is a client for the ReviewRule schema.
type ReviewRuleClient struct {
	config
}

// NewReviewRuleClient returns a client for the ReviewRule from the given config.
func NewReviewRuleClient(c config) *ReviewRuleClient {
	return &ReviewRuleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `reviewrule.Hooks(f(g(h())))`.
func (c *ReviewRuleClient) Use(hooks ...Hook) {
	c.hooks.ReviewRule = append(c.hooks.ReviewRule, hooks...)
}

// Create returns a create builder for ReviewRule.
func (c *ReviewRuleClient) Create() *ReviewRuleCreate {
	mutation := newReviewRuleMutation(c.config, OpCreate)
	return &ReviewRuleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ReviewRule entities.
func (c *ReviewRuleClient) CreateBulk(builders ...*ReviewRuleCreate) *ReviewRuleCreateBulk {
	return &ReviewRuleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ReviewRule.
func (c *ReviewRuleClient) Update() *ReviewRuleUpdate {
	mutation := newReviewRuleMutation(c.config, OpUpdate)
	return &ReviewRuleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReviewRuleClient) UpdateOne(rr *ReviewRule) *ReviewRuleUpdateOne {
	mutation := newReviewRuleMutation(c.config, OpUpdateOne, withReviewRule(rr))
	return &ReviewRuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReviewRuleClient) UpdateOneID(id uuid.UUID) *ReviewRuleUpdateOne {
	mutation := newReviewRuleMutation(c.config, OpUpdateOne, withReviewRuleID(id))
	return &ReviewRuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ReviewRule.
func (c *ReviewRuleClient) Delete() *ReviewRuleDelete {
	mutation := newReviewRuleMutation(c.config, OpDelete)
	return &ReviewRuleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ReviewRuleClient) DeleteOne(rr *ReviewRule) *ReviewRuleDeleteOne {
	return c.DeleteOneID(rr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ReviewRuleClient) DeleteOneID(id uuid.UUID) *ReviewRuleDeleteOne {
	builder := c.Delete().Where(reviewrule.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReviewRuleDeleteOne{builder}
}

// Query returns a query builder for ReviewRule.
func (c *ReviewRuleClient) Query() *ReviewRuleQuery {
	return &ReviewRuleQuery{
		config: c.config,
	}
}

// Get returns a ReviewRule entity by its id.
func (c *ReviewRuleClient) Get(ctx context.Context, id uuid.UUID) (*ReviewRule, error) {
	return c.Query().Where(reviewrule.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReviewRuleClient) GetX(ctx context.Context, id uuid.UUID) *ReviewRule {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ReviewRuleClient) Hooks() []Hook {
	return c.hooks.ReviewRule
}
