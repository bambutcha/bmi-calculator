package fsm

type BotState int

const (
    WaitingForHeight BotState = iota
    WaitingForWeight
)

type UserState struct {
    State  BotState
    Height float64
}

func NewUserState() *UserState {
    return &UserState{
        State: WaitingForHeight,
    }
}