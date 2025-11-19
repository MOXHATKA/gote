package bot

import "gote/pkg/types"

type StateMachine struct {
	StartState *State
	States     map[int64]*State
	ResetState *State
}

type State struct {
	Name      string
	Condition string
	Parent    *State
	Children  []*State
}

func (sm *StateMachine) SetState(update *types.Update) bool {
	id := update.Message.Chat.Id
	text := update.Message.Text
	state, ok := sm.States[id]
	if !ok {
		sm.States[id] = sm.ResetState
		return false
	}

	children := state.Children
	if len(children) == 0 {
		return false
	}

	if len(children) == 1 {
		sm.States[id] = children[0]
		return true
	}

	for _, s := range children {
		if s.Condition == text {
			sm.States[id] = s
			return true
		}
	}

	return false
}

func (sm *StateMachine) GetState(id int64) *State {
	return sm.States[id]
}
