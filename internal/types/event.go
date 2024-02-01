package types

import(
	"time"
)

type EventCallback func();
  
type Event struct{
	Timestamp *time.Time;
	Callback EventCallback;
}


  