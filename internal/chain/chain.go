package chain

import (
	"github.com/vsskarthik/ChainSched/internal/types"
)

func Attach(chain *types.Chain, next *types.EventNode){
	if(chain.Start == nil){
		chain.Start = next;
		return;
	}

	var curr *types.EventNode = chain.Start;
	for curr.NextEvent!=nil{
		curr = curr.NextEvent;
	}
	curr.NextEvent = next;
}