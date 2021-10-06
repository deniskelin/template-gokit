package rds

import "database/sql"

const (
	QueryRDSGetBillingByEnvAndID          string = `select t.billing_source from rds.get_billing_by_number($1, $2)`
	QueryRDSGetBillingEnvRouteByAccountID string = `select number_code, i_account, billing_source, env, route_label from rds.get_billing_by_number($1, $2)`
	QueryRDSSetRouteLabel                 string = `/*NO LOAD BALANCE*/ select rds.set_route_label($1, $2, $3, $4)`
	QueryRDSGetRouteLabel                 string = `select number_code, route_label from rds.get_route_label($1, $2, $3)`
)

type IClient interface {
	GetBillingByEnvAndID(accountID uint64, envID string) ([]string, error)
	GetBillingEnvRouteByAccountID(accountID uint64, envID *string) ([]BillingInfo, error)
	SetRouteLabel(accountID uint64, envID, billingSource, routeLabel string) error
	GetRouteLabel(accountID uint64, envID, billingSource string) ([]RouteLabel, error)
}

type DB interface {
	DBWriter
	DBReader
}

type DBWriter interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type DBReader interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
}

// Client is RDS-Client structure that stores config
type Client struct {
	rwdb DBWriter
	rdb  DBReader
}

func NewClient(rwdb DBWriter, rdb DBReader) IClient {
	cl := &Client{rwdb: rwdb, rdb: rdb}
	return cl
}

func (cl *Client) GetBillingByEnvAndID(accountID uint64, envID string) ([]string, error) {
	billings := make([]string, 0, 2)
	err := cl.rdb.Select(&billings, QueryRDSGetBillingByEnvAndID, accountID, envID)
	if err != nil {
		// return nil, dbpostgres.ProcessDBError(err)
		return nil, err
	}
	return billings, nil
}

func (cl *Client) GetBillingEnvRouteByAccountID(accountID uint64, envID *string) ([]BillingInfo, error) {
	billingInfo := make([]BillingInfo, 0, 2)
	err := cl.rdb.Select(&billingInfo, QueryRDSGetBillingEnvRouteByAccountID, accountID, envID)
	if err != nil {
		// return dbpostgres.ProcessDBError(err)
		return nil, err
	}
	return billingInfo, nil
}

func (cl *Client) SetRouteLabel(accountID uint64, envID, billingSource, routeLabel string) error {
	_, err := cl.rwdb.Exec(QueryRDSSetRouteLabel, accountID, envID, billingSource, routeLabel)
	if err != nil {
		// return dbpostgres.ProcessDBError(err)
		return err
	}
	return nil
}

func (cl *Client) GetRouteLabel(accountID uint64, envID, billingSource string) ([]RouteLabel, error) {
	routeLabel := make([]RouteLabel, 0, 2)
	err := cl.rdb.Select(&routeLabel, QueryRDSGetRouteLabel, accountID, billingSource, envID)
	if err != nil {
		// return GetRouteLabelResult{}, dbpostgres.ProcessDBError(err)
		return nil, err
	}
	return routeLabel, nil
}
