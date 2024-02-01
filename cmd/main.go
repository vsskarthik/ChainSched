package main

import (
	"fmt"

	"github.com/vsskarthik/ChainSched/internal/types"
	"github.com/vsskarthik/ChainSched/internal/chain"

)

func main(){
	var chain1 *types.Chain = &types.Chain{};
	var testNode1 *types.EventNode = &types.EventNode{}
	testNode1.ID = 1;
	chain.Attach(chain1, testNode1);
	var testNode2 *types.EventNode = &types.EventNode{}
	testNode2.ID = 2;
	chain.Attach(chain1, testNode2);

	curr := chain1.Start;
	for curr != nil {
		fmt.Print(curr.ID, "->");
		curr = curr.NextEvent;
	}

}