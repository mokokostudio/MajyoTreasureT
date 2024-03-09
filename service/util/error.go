package util

import (
	"errors"

	"google.golang.org/grpc/status"
)

func GRPCErrIs(err, targetErr error) bool {
	if targetErr == nil {
		return err == nil
	}
	errStatus, ok := status.FromError(err)
	if !ok {
		return errors.Is(err, targetErr)
	}
	return errStatus.Message() == targetErr.Error()
}
