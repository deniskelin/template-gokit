package rds

type BillingInfo struct {
	NumberCode    uint64  `db:"number_code"`
	IAccount      string  `db:"i_account"`
	BillingSource string  `db:"billing_source"`
	EnvID         string  `db:"env"`
	RouteLabel    *string `db:"route_label"`
}

// todo fix this struct for correct data!
type RouteLabel struct {
	AccountID  uint64  `db:"number_code"`
	RouteLabel *string `db:"route_label"`
	EnvID      string  `db:"env"`
}
