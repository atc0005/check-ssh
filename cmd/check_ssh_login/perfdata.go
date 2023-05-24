// Copyright 2023 Adam Chalkley
//
// https://github.com/atc0005/check-ssh
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

// getPerfData gathers performance data metrics that we wish to report.
// func getPerfData(processes process.Processes) []nagios.PerformanceData {
//
// 	probProcs := processes.States(process.KnownProblemProcessStates())
//
// 	return []nagios.PerformanceData{
// 		// The `time` (runtime) metric is appended at plugin exit, so do not
// 		// duplicate it here.
// 		{
// 			Label: "problem_processes",
// 			Value: fmt.Sprintf("%d", len(probProcs)),
// 		},
// 		{
// 			Label: "running",
// 			Value: fmt.Sprintf("%d", processes.StateRunningCount()),
// 		},
// 		{
// 			Label: "sleeping",
// 			Value: fmt.Sprintf("%d", processes.StateSleepingCount()),
// 		},
// 		{
// 			Label: "uninterruptible_disk_sleep",
// 			Value: fmt.Sprintf("%d", processes.StateDiskSleepCount()),
// 		},
// 		{
// 			Label: "stopped",
// 			Value: fmt.Sprintf("%d", processes.StateStoppedCount()),
// 		},
// 		{
// 			Label: "zombie",
// 			Value: fmt.Sprintf("%d", processes.StateZombieCount()),
// 		},
// 		{
// 			Label: "dead",
// 			Value: fmt.Sprintf("%d", processes.StateDeadCount()),
// 		},
// 		{
// 			Label: "tracing_stop",
// 			Value: fmt.Sprintf("%d", processes.StateTracingStopCount()),
// 		},
// 		{
// 			Label: "wakekill",
// 			Value: fmt.Sprintf("%d", processes.StateWakeKillCount()),
// 		},
// 		{
// 			Label: "waking",
// 			Value: fmt.Sprintf("%d", processes.StateWakingCount()),
// 		},
// 		{
// 			Label: "idle",
// 			Value: fmt.Sprintf("%d", processes.StateIdleCount()),
// 		},
// 		{
// 			Label: "parked",
// 			Value: fmt.Sprintf("%d", processes.StateParkedCount()),
// 		},
// 	}
// }
