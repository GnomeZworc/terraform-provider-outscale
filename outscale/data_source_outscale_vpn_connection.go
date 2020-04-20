package outscale

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/antihax/optional"
	oscgo "github.com/marinsalinas/osc-sdk-go"
	"github.com/spf13/cast"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceOutscaleVPNConnection() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOutscaleVPNConnectionRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"vpn_connection_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_gateway_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_gateway_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"static_routes_only": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"client_gateway_configuration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"routes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_ip_range": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tags": tagsOAPIListSchemaComputed(),
			"request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOutscaleVPNConnectionRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	filters, filtersOk := d.GetOk("filter")
	vpnConnectionID, vpnConnectionOk := d.GetOk("vpn_connection_id")

	if !filtersOk && !vpnConnectionOk {
		return fmt.Errorf("One of filters, or vpn_connection_id must be assigned")
	}

	params := oscgo.ReadVpnConnectionsRequest{}

	if vpnConnectionOk {
		params.Filters = &oscgo.FiltersVpnConnection{
			VpnConnectionIds: &[]string{vpnConnectionID.(string)},
		}
	}

	if filtersOk {
		params.Filters = buildOutscaleDataSourceVPNConnectionFilters(filters.(*schema.Set))
	}

	var resp oscgo.ReadVpnConnectionsResponse
	var err error
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		resp, _, err = conn.VpnConnectionApi.ReadVpnConnections(context.Background(), &oscgo.ReadVpnConnectionsOpts{
			ReadVpnConnectionsRequest: optional.NewInterface(params),
		})
		if err != nil {
			if strings.Contains(fmt.Sprint(err), "RequestLimitExceeded:") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	if len(resp.GetVpnConnections()) == 0 {
		return fmt.Errorf("Unable to find Client Gateway")
	}

	if len(resp.GetVpnConnections()) > 1 {
		return fmt.Errorf("multiple results returned, please use a more specific criteria in your query")
	}

	vpnConnection := resp.GetVpnConnections()[0]

	if err := d.Set("client_gateway_id", vpnConnection.GetClientGatewayId()); err != nil {
		return err
	}
	if err := d.Set("virtual_gateway_id", vpnConnection.GetVirtualGatewayId()); err != nil {
		return err
	}
	if err := d.Set("connection_type", vpnConnection.GetConnectionType()); err != nil {
		return err
	}
	if err := d.Set("static_routes_only", vpnConnection.GetStaticRoutesOnly()); err != nil {
		return err
	}
	if err := d.Set("client_gateway_configuration", vpnConnection.GetClientGatewayId()); err != nil {
		return err
	}
	if err := d.Set("vpn_connection_id", vpnConnection.GetVpnConnectionId()); err != nil {
		return err
	}
	if err := d.Set("state", vpnConnection.GetState()); err != nil {
		return err
	}
	if err := d.Set("routes", flattenVPNConnection(vpnConnection.GetRoutes())); err != nil {
		return err
	}
	if err := d.Set("tags", tagsOSCAPIToMap(vpnConnection.GetTags())); err != nil {
		return err
	}
	if err := d.Set("request_id", resp.ResponseContext.GetRequestId()); err != nil {
		return err
	}

	d.SetId(vpnConnection.GetVpnConnectionId())

	return nil
}

func buildOutscaleDataSourceVPNConnectionFilters(set *schema.Set) *oscgo.FiltersVpnConnection {
	var filters oscgo.FiltersVpnConnection
	for _, v := range set.List() {
		m := v.(map[string]interface{})
		var filterValues []string
		for _, e := range m["values"].([]interface{}) {
			filterValues = append(filterValues, e.(string))
		}

		var filteBgpAsnsValues []int64
		for _, e := range m["values"].([]interface{}) {
			filteBgpAsnsValues = append(filteBgpAsnsValues, cast.ToInt64(e))
		}

		switch name := m["name"].(string); name {
		case "vpn_connection_ids":
			filters.SetVpnConnectionIds(filterValues)
		case "virtual_gateway_ids":
			filters.SetVirtualGatewayIds(filterValues)
		case "client_gateway_ids":
			filters.SetVpnConnectionIds(filterValues)
		case "connection_types":
			filters.SetConnectionTypes(filterValues)
		case "route_destination_ip_rRanges":
			filters.SetRouteDestinationIpRanges(filterValues)
		case "states":
			filters.SetStates(filterValues)
		case "tag_keys":
			filters.SetStaticRoutesOnly(cast.ToBool(filterValues[0]))
		case "bgp_asns":
			filters.SetBgpAsns(filteBgpAsnsValues)
		default:
			log.Printf("[Debug] Unknown Filter Name: %s.", name)
		}
	}
	return &filters
}
