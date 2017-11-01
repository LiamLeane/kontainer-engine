package cmd

import (
	generic "github.com/rancher/kontainer-engine/driver"
)

// runRPCDriver runs the rpc server and returns
func runRPCDriver(driverName string) (*generic.GrpcClient, string, error) {
	// addrChan is the channel to receive the server listen address
	addrChan := make(chan string)
	runDriver(driverName, addrChan)

	addr := <- addrChan
	rpcClient, err := generic.NewClient(driverName, addr)
	if err != nil {
		return nil, "", err
	}
	return rpcClient, addr, nil
}