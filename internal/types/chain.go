package types


type Chain struct{
	ID int;
	Start *EventNode;
}

type EventNode struct{
	Event *Event;
	ID int;
	NextEvent *EventNode;
}