package randomizer

import (
	"container/list"
	"math/rand"
	"sort"
	"strings"
)

type multiRoute struct {
	ri     *routeInfo
	local  map[*node]bool
	checks map[*node]*node
}

func randomMultiCheck(src *rand.Rand, mr *multiRoute) (*node, *node) {
	i, slots := 0, make([]*node, 0, len(mr.checks))
	for slot, item := range mr.checks {
		// TODO: this ring check could be done more precisely
		if !mr.local[slot] && !strings.Contains(item.name, "ring") &&
			!strings.HasSuffix(item.name, "small key") &&
			!strings.HasSuffix(item.name, "boss key") &&
			!strings.HasSuffix(item.name, "dungeon map") &&
			!strings.HasSuffix(item.name, "compass") {
			slots = append(slots, slot)
			i++
		}
	}

	sort.Slice(slots, func(i, j int) bool {
		return slots[i].name < slots[j].name
	})

	slot := slots[src.Intn(len(slots))]
	return slot, mr.checks[slot]
}

func shuffleMultiworld(
	ris []*routeInfo, roms []*romState, verbose bool, logf logFunc) {
	mrs := make([]*multiRoute, len(ris))
	src := ris[0].src

	for i, ri := range ris {
		// mark graph nodes as belonging to each player
		for _, v := range ri.graph {
			v.player = i + 1
		}

		// accumulate checks in a sane way
		mrs[i] = &multiRoute{
			ri:     ri,
			checks: getChecks(ri.usedItems, ri.usedSlots),
			local:  make(map[*node]bool, ri.usedItems.Len()),
		}
		for slot := range mrs[i].checks {
			mrs[i].local[slot] = roms[i].itemSlots[slot.name].localOnly
		}
	}

	// swap some random items ???
	swaps := 0
	for i := 0; i < 10000*len(mrs); i++ {
		slot1, item1 := randomMultiCheck(src, mrs[src.Intn(len(mrs))])
		slot2, item2 := randomMultiCheck(src, mrs[src.Intn(len(mrs))])

		// haha
		if slot1 == slot2 || slot1.player == slot2.player {
			continue
		}

		if verbose {
			logf("swapping {p%d %s <- p%d %s} with {p%d %s <- p%d %s}",
				slot1.player, slot1.name, item1.player, item1.name,
				slot2.player, slot2.name, item2.player, item2.name)
		}

		// swap parents
		item1.removeParent(slot1)
		item2.removeParent(slot2)
		item1.addParent(slot2)
		item2.addParent(slot1)

		// test whether seeds are still beatable w/ item placement
		success := true
		for i, ri := range ris {
			ri.graph.reset()
			ri.graph["start"].explore()
			if !ri.graph["done"].reached {
				item1.removeParent(slot2)
				item2.removeParent(slot1)
				item1.addParent(slot1)
				item2.addParent(slot2)
				success = false
				if verbose {
					logf("player %d route no longer viable", i)
				}
				break
			}
		}

		// update check maps
		if success {
			mrs[slot1.player-1].checks[slot1] = item2
			mrs[slot2.player-1].checks[slot2] = item1
			swaps++
			println(swaps, "swaps")
		}
	}

	if verbose {
		logf("made %d successful swaps", swaps)
	}

	// reconstruct used item and slot lists
	for i, ri := range ris {
		ri.usedItems, ri.usedSlots = list.New(), list.New()
		for slot, item := range mrs[i].checks {
			ri.usedItems.PushBack(item)
			ri.usedSlots.PushBack(slot)
			roms[i].itemSlots[slot.name].player = byte(item.player)
		}
	}
}