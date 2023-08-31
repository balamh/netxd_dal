package netxd_dal_interfaces

import "github.com/balamh/netxd_dal/netxd_dal_models"

// import "netxd_dal/netxd_dal_models"

type INetxdCustomers interface {
	CreateCustomer(*netxd_dal_models.NetxdCustomer) (*netxd_dal_models.CustomerResponse, error)
}
