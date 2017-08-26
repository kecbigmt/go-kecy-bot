package common
import "fmt"

type ValidationError struct{
  Msg string
  Code uint8
  State interface{}
  Index int
  Input uint8
  Stack []byte
}

func (e ValidationError) Error() string{
  return fmt.Sprintf("%v(%v) state=%T, index=%v, input=%v, stack=%v", e.Msg, e.Code, e.State, e.Index, e.Input, e.Stack)
}
