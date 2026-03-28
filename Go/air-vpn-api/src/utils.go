package main

import (
	"encoding/json"
	"slices"

	"github.com/Flori991/ProgrammingLearning/types"
)

func mergeResponsesIntoSummaries(sessions []types.Session, status []types.ServerStatus) types.SessionSummaries {
	var sessionSummaries []types.SessionSummary
	for _, session := range sessions {

		index := slices.IndexFunc(status, func(s types.ServerStatus) bool { return s.ServerName == session.ServerName })
		matchingServerStatus := status[index]

		sessionSummaries = append(sessionSummaries, types.SessionSummary{
			DeviceName:         session.DeviceName,
			DeviceDescription:  session.DeviceDescription,
			ExitIpv4:           session.ExitIpv4,
			ServerName:         session.ServerName,
			ServerCountry:      session.ServerCountry,
			BandwidthUsed:      matchingServerStatus.BandwidthUsed,
			BandwidthMax:       matchingServerStatus.BandwidthMax,
			BytesRead:          session.BytesRead,
			BytesWrite:         session.BytesWrite,
			ConnectedSinceDate: session.ConnectedSinceDate,
			ConnectedSinceUnix: session.ConnectedSinceUnix,
		})
	}
	return types.SessionSummaries{
		Sessions: sessionSummaries,
	}
}

func safeJsonParse[T any](body []byte, target T) (T, error) {
	if err := json.Unmarshal(body, &target); err != nil {
		return target, err
	}
	return target, nil
}
