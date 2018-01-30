package rebalance

import "github.com/pborman/uuid"

// RebalanceStatus represents Rebalance Status
type RebalanceStatus uint64

const (
	// GfDefragStatusNotStarted should be set only for a volume in which rebalance process is not started
	GfDefragStatusNotStarted RebalanceStatus = iota
	// GfDefragStatusStarted should be set only for a volume that has been just started rebalance process
	GfDefragStatusStarted
	// GfDefragStatusStopped should be set only for a volume that has been just stopped rebalance process
	GfDefragStatusStopped
	// GfDefragStatusComplete should be set only for a volume that the rebalance process is completed
	GfDefragStatusComplete
	// GfDefragStatusFailed should be set only for a volume that are failed to run rebalance process
	GfDefragStatusFailed
	// GfDefragStatusLayoutFixStarted should be set only for a volume that has been just started rebalance fix-layout
	GfDefragStatusLayoutFixStarted
	// GfDefragStatusLayoutFixStopped should be set only for a volume that has been just stopped rebalance fix-layout
	GfDefragStatusLayoutFixStopped
	// GfDefragStatusLayoutFixComplete should be set only for a volume that the rebalance fix-layout is completed
	GfDefragStatusLayoutFixComplete
	// GfDefragStatusLayoutFixFailed should be set only for a volume that are failed to run rebalance fix-layout
	GfDefragStatusLayoutFixFailed
)

// RebalanceCommand represents Rebalance Commands
type RebalanceCommand uint64

const (
	// GfDefragCmdNone should be set only when given cmd is none
	GfDefragCmdNone RebalanceCommand = iota
	// GfDefragCmdStart should be set only when given cmd is rebalance start
	GfDefragCmdStart
	// GfDefragCmdStop should be set only when given cmd is rebalance stop
	GfDefragCmdStop
	// GfDefragCmdStatus should be set only when given cmd is rebalance status
	GfDefragCmdStatus
	// GfDefragCmdStartLayoutFix should be set only when given cmd is rebalance fix layout start
	GfDefragCmdStartLayoutFix
	// GfDefragCmdStartForce should be set only when given cmd is rebalance start force
	GfDefragCmdStartForce
)

// RebalanceInfo represents Rebalance details
type RebalanceInfo struct {
	Volname           string
	Status            RebalanceStatus
	Cmd               RebalanceCommand
	RebalanceID       uuid.UUID
	RebalanceFiles    uint64
	RebalanceData     uint64
	LookedupFiles     uint64
	RebalanceFailures uint64
	ElapsedTime       uint64
	SkippedFiles      uint64
	TimeLeft          uint64
	CommitHash        uint64
}

// RebalanceStartReq represents Rebalance Start Request
type RebalanceStartReq struct {
	Volname   string
	Fixlayout bool `omitempty`
	Force     bool `omitempty`
}
