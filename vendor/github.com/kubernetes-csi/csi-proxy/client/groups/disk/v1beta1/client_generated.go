// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"net"

	"github.com/Microsoft/go-winio"
	"github.com/kubernetes-csi/csi-proxy/client"
	"github.com/kubernetes-csi/csi-proxy/client/api/disk/v1beta1"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"google.golang.org/grpc"
)

const groupName = "disk"

var version = apiversion.NewVersionOrPanic("v1beta1")

type Client struct {
	client     v1beta1.DiskClient
	connection *grpc.ClientConn
}

// NewClient returns a client to make calls to the disk API group version v1beta1.
// It's the caller's responsibility to Close the client when done.
func NewClient() (*Client, error) {
	pipePath := client.PipePath(groupName, version)

	connection, err := grpc.Dial(pipePath,
		grpc.WithContextDialer(func(context context.Context, s string) (net.Conn, error) {
			return winio.DialPipeContext(context, s)
		}),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := v1beta1.NewDiskClient(connection)
	return &Client{
		client:     client,
		connection: connection,
	}, nil
}

// Close closes the client. It must be called before the client gets GC-ed.
func (w *Client) Close() error {
	return w.connection.Close()
}

// ensures we implement all the required methods
var _ v1beta1.DiskClient = &Client{}

func (w *Client) DiskStats(context context.Context, request *v1beta1.DiskStatsRequest, opts ...grpc.CallOption) (*v1beta1.DiskStatsResponse, error) {
	return w.client.DiskStats(context, request, opts...)
}

func (w *Client) ListDiskIDs(context context.Context, request *v1beta1.ListDiskIDsRequest, opts ...grpc.CallOption) (*v1beta1.ListDiskIDsResponse, error) {
	return w.client.ListDiskIDs(context, request, opts...)
}

func (w *Client) ListDiskLocations(context context.Context, request *v1beta1.ListDiskLocationsRequest, opts ...grpc.CallOption) (*v1beta1.ListDiskLocationsResponse, error) {
	return w.client.ListDiskLocations(context, request, opts...)
}

func (w *Client) PartitionDisk(context context.Context, request *v1beta1.PartitionDiskRequest, opts ...grpc.CallOption) (*v1beta1.PartitionDiskResponse, error) {
	return w.client.PartitionDisk(context, request, opts...)
}

func (w *Client) Rescan(context context.Context, request *v1beta1.RescanRequest, opts ...grpc.CallOption) (*v1beta1.RescanResponse, error) {
	return w.client.Rescan(context, request, opts...)
}
