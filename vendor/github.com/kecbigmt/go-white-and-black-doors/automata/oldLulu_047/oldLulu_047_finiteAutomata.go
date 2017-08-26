package oldLulu_047
import (
  "github.com/kecbigmt/go-white-and-black-doors/automata/common"
)

type room interface{
  openDoor(i int, v uint8) (interface{}, error)
}

type entrance struct{}
func (r *entrance) openDoor(i int, v uint8) (interface{}, error){
  switch v{
  case uint8(0):
    return &roomA{}, nil
  case uint8(1):
    return &entrance{}, nil
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v, nil}
  }
}

type roomA struct{}
func (r *roomA) openDoor(i int, v uint8) (interface{}, error){
  switch v{
  case uint8(0):
    return &roomA{}, nil
  case uint8(1):
    return &exit{}, nil
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v, nil}
  }
}

type exit struct{}
func (r *exit) openDoor(i int, v uint8) (interface{}, error){
  switch v{
  case uint8(0):
    return &roomA{}, nil
  case uint8(1):
    return &entrance{}, nil
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v, nil}
  }
}

func Validate(b []byte) error{
  var state room
  state = &entrance{}
  for i,v := range b{
    result, err := state.openDoor(i, uint8(v))
    if err != nil{
      return err
    }
    switch tmp := result.(type){
    case *entrance:
      state = tmp
    case *roomA:
      state = tmp
    case *exit:
      state = tmp
    }
  }
  if _, ok:=state.(*exit);ok{
    return nil
  }else{
    return common.ValidationError{"failed to reach exit", 101, state, len(b)-1, uint8(b[len(b)-1]), nil}
  }
}

var TestPatterns [][]byte = [][]byte{
  {1,1,1,1,0,1},
  {1,1,1,0,0,1},
  {1,1,0,1,0,1},
  {1,1,0,0,0,1},
  {1,0,1,1,0,1},
  {1,0,1,0,0,1},
  {1,0,0,1,0,1},
  {1,0,0,0,0,1},
  {0,1,1,1,0,1},
  {0,1,1,0,0,1},
  {0,1,1,0,0,0},
}
