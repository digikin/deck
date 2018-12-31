package dry

import (
	"github.com/hbagdi/go-kong/kong"
	"github.com/kong/deck/crud"
	"github.com/kong/deck/diff"
	"github.com/kong/deck/print"
	"github.com/kong/deck/state"
	"github.com/kong/deck/utils"
)

// UpstreamCRUD implements Actions interface
// from the github.com/kong/crud package for the Upstream entitiy of Kong.
type UpstreamCRUD struct {
	// client    *kong.Client
	// callbacks []Callback // use this to update the current in-memory state
}

func upstreamFromStuct(a diff.Event) *state.Upstream {
	upstream, ok := a.Obj.(*state.Upstream)
	if !ok {
		panic("unexpected type, expected *state.upstream")
	}

	return upstream
}

// Create creates a fake Upstream.
// The arg should be of type diff.Event, containing the upstream to be created,
// else the function will panic.
// It returns a the created *state.Upstream.
func (s *UpstreamCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	upstream := upstreamFromStuct(event)

	print.CreatePrintln("creating upstream", *upstream.Name)
	upstream.ID = kong.String(utils.UUID())
	return upstream, nil
}

// Delete deletes a fake Upstream.
// The arg should be of type diff.Event, containing the upstream to be deleted,
// else the function will panic.
// It returns a the deleted *state.Upstream.
func (s *UpstreamCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	upstream := upstreamFromStuct(event)

	print.DeletePrintln("deleting upstream", *upstream.Name)
	return upstream, nil
}

// Update updates a fake Upstream.
// The arg should be of type diff.Event, containing the upstream to be updated,
// else the function will panic.
// It returns a the updated *state.Upstream.
func (s *UpstreamCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	upstream := upstreamFromStuct(event)
	oldUpstreamObj, ok := event.OldObj.(*state.Upstream)
	if !ok {
		panic("unexpected type, expected *state.upstream")
	}
	oldUpstream := oldUpstreamObj.DeepCopy()
	// TODO remove this hack
	oldUpstream.CreatedAt = nil
	diff := getDiff(oldUpstream, &upstream.Upstream)
	print.UpdatePrintln("updating upstream", *upstream.Name)
	print.UpdatePrintf("%s", diff)
	return upstream, nil
}
