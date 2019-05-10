package rom

var agesFixedMutables = map[string]Mutable{
	// first and second time portals (near maku tree) are always active
	"first portal active": MutableString(Addr{0x10, 0x7d4e},
		"\x20\x0e", "\x20\x00"),
	"second portal active": MutableString(Addr{0x10, 0x7d57},
		"\x38\x05", "\x38\x00"),

	// allow access to nayru's house check from start
	"move impa": MutableByte(Addr{0x09, 0x6567}, 0xd0, 0x00),

	// prevent stairs disappearing in event where maku tree is attacked by
	// moblins, preventing softlock if player gets there with seed satchel and
	// no sword or something stupid
	"maku tree stairs": MutableString(Addr{0x15, 0x6bf3},
		"\x84\x05\x00", "\xc4\x15\xc3"),

	// change harp interaction to allow sub ID
	"create harp with sub ID": MutableString(Addr{0x0b, 0x6825},
		"\xcd\xef\x3a\xc0\x36\x60\x2c\x36\x11",
		"\xc5\x01\x00\x11\xcd\xd4\x27\xc1\xc0"),

	// never spawn hide and seek event in fairies' woods. apparently you're
	// frozen if you enter on an animal?
	"don't spawn fairies": MutableByte(Addr{0x0a, 0x52bf}, 0xc2, 0xc3),

	// make guy in front of d2 go away if you have bombs
	"d2 guy flag check": MutableString(Addr{0x09, 0x5242},
		"\x3e\x0b\xcd\xf3\x31\xc2", "\x3e\x03\xcd\x48\x17\xda"),
	// and center him on a tile so you can't get stuck in a currents loop
	"d2 guy position": MutableByte(Addr{0x12, 0x611c}, 0x4e, 0x48),

	// clear rubble from rolling ridge base present without d4 essence
	"clear rubble": MutableByte(Addr{0x04, 0x6a44}, 0xc8, 0x00),
	// open rolling ridge present tunnel without completing d5
	"open tunnel": MutableByte(Addr{0x04, 0x6a35}, 0xc8, 0x00),

	// cut off the end of deku forest soldier's text so it makes sense when
	// giving item
	"soldier text end": MutableByte(Addr{0x23, 0x6656}, 0x01, 0x00),
	// and position the "you may go now" text correctly on screen
	"soldier text position": MutableByte(Addr{0x23, 0x65d8}, 0x22, 0x00),
	// and remove the usual soldier event (taken to palace etc)
	"remove soldier event": MutableByte(Addr{0x12, 0x58f5}, 0xcd, 0xc9),

	// skip essence checks for the following events:
	"rafton essence check":    MutableByte(Addr{0x0a, 0x4d7a}, 0x20, 0x18),
	"dimitri essence check 1": MutableByte(Addr{0x09, 0x5816}, 0x13, 0x00),
	"dimitri essence check 2": MutableByte(Addr{0x0a, 0x4bb3}, 0xc8, 0x00),
	"open palace": MutableString(Addr{0x09, 0x51f8},
		"\x3e\x40", "\xaf\xc9"),

	// moosh should always appear in the graveyard
	"moosh essence checks": MutableStrings([]Addr{{0x0a, 0x5dd5},
		{0x0a, 0x5943}, {0x0a, 0x4b85}}, "\xcb\x4f", "\xf6\x01"),
	"moosh rope checks": MutableStrings([]Addr{{0x05, 0x78b8}, {0x0a, 0x4b92},
		{0x0a, 0x4ba3}}, "\xcd\x48\x17", "\xaf\xaf\xaf"),
	"moosh cheval checks": MutableStrings([]Addr{{0x0a, 0x5ddc},
		{0x0a, 0x594b}, {0x0a, 0x4b8c}}, "\xcb\x77", "\xf6\x01"),

	// ricky shouldn't leave after talking to tingle
	"end tingle script": MutableString(Addr{0x0c, 0x7e2a},
		"\x91\x03\xd1\x02", "\xbe\x7d\xfe\xba"),
	// and check fake treasure ID 13 (slingshot) instead of island chart
	"tingle fake ID": MutableByte(Addr{0x0c, 0x7e00}, 0x54, 0x13),
	// ignore satchel level when talking to tingle for second item
	"tingle satchel check": MutableByte(Addr{0x0b, 0x75c5}, 0x3d, 0xaf),
	// dig up item on south shore regardless of ricky state
	"south shore ricky check 1": MutableByte(Addr{0x04, 0x6b77}, 0x0a, 0x00),
	"south shore ricky check 2": MutableByte(Addr{0x04, 0x6b7b}, 0x06, 0x00),
	"south shore ricky check 3": MutableByte(Addr{0x0a, 0x5e2f}, 0x12, 0x00),
	"south shore ricky check 4": MutableByte(Addr{0x0a, 0x5e33}, 0x0e, 0x00),
	// and check fake treasure ID 08 (magnet gloves) instead of ricky's gloves
	"south shore fake ID": MutableStrings([]Addr{{0x04, 0x6b7d},
		{0x0a, 0x5e35}}, "\x48", "\x08"),

	// don't refill seeds when getting item from tingle
	"tingle seed refill": MutableString(Addr{0x0c, 0x7e7d},
		"\xe0\x0c\x18", "\xc4\x80\x7e"),

	// remove storm event that washes link up on crescent island without raft,
	// and the event where tokays steal link's items
	"remove storm event": MutableByte(Addr{0x0b, 0x52e3}, 0xc2, 0xc3),
	"remove tokay event": MutableStrings([]Addr{{0x09, 0x5756}, {0x09, 0x5731},
		{0x0a, 0x4fe1}}, "\xc2", "\xc3"),
	"remove tokay items": MutableString(Addr{0x09, 0x57a5},
		"\xcb\x77", "\x3c\x3c"),
	"tokay trading hut": MutableStrings([]Addr{{0x0a, 0x623a}, {0x0a, 0x62d7}},
		"\xcd\xf3\x31", "\xb7\xb7\xb7"),
	// don't have an item in the chicken hut
	"tokay bomb hut": MutableString(Addr{0x12, 0x638f},
		"\xf2\x6b\x0a\x28", "\xf3\x57\x41\xff"),

	// sell 150 rupee item from lynna city shop from the start
	"shop flute flag check": MutableString(Addr{0x09, 0x4333},
		"\x28\x04", "\x00\x00"),
	// check for fake treasure ID 07 (rod) so that non-unique items can be sold
	"shop fake ID": MutableStrings([]Addr{{0x09, 0x4328}, {0x09, 0x42a5}},
		"\x0e", "\x07"),

	// remove flute item from shooting gallery prizes
	"shooting gallery script": MutableString(Addr{0x15, 0x51d8},
		"\xdf\x0e", "\xdf\x02"),
	// prevent bridge-building foreman from setting flag 22 so that
	// animal/flute event doesn't happen in fairies' woods
	"bridge foreman script": MutableString(Addr{0x15, 0x75bf},
		"\xb6\x22", "\xb6\xa2"),

	// skip normal boomerang check in target carts, since EOB code handles it
	"skip target carts boomerang check": MutableString(Addr{0x15, 0x66ae},
		"\x20\x0b", "\x18\x0b"),
	// and remove "boomerang" from random prizes
	"target carts prize table": MutableString(Addr{0x15, 0x66e5},
		"\x04\x04\x04", "\x03\x03\x03"),
	// and don't give boomerang as a shooting gallery prize
	"no goron gallery boomerang": MutableString(Addr{0x15, 0x52b6},
		"\xdf\x06\xc3\x52", "\xc4\xc3\x52\x00"),

	// trade lava juice without mermaid key
	"trade lava juice without key": MutableString(Addr{0x15, 0x6879},
		"\x30\x07", "\x30\x00"),

	// stop d6 boss key chest from setting past boss key flag
	"stop d6 boss key chest": MutableString(Addr{0x10, 0x793c},
		"\xc3\x0e\x02", "\xc9\x00\x00"),

	// fix pickup text for harp tunes
	"tune of echoes text": MutableString(Addr{0x1e, 0x4c3e}, "\x49",
		"\x02\x06"+ // You got the
			"\x09\x01Tune\x04\xceE\x05\x0d\x04\x91"+ // Tune of Echoes!
			"Play\x04\x0f\x01"+ // Play it to
			"awaken \x04\xa8\x04\x5a"+ // awaken sleeping
			"\x09\x03Time Portals\x09\x00!\x00"), // Time Portals!
	"tune of currents text": MutableString(Addr{0x1d, 0x7e48}, "\x59",
		"\x02\x06"+ // You got the
			"\x09\x01Tune\x04\xce\x01"+ // Tune of
			"Currents\x05\x95Play\x01"+ // Currents! Play
			"it\x04\x57\x05\x5b\x03\x50"+ // it to move from
			"\x02\x81 \x02\x64\x01"+ // the past to the
			"\x03\x2e!\x00"), // the present!
	"tune of ages text": MutableString(Addr{0x1d, 0x7e8e}, "\x59",
		"\x02\x06"+ // You got the
			"\x09\x01Tune \x03\x31\x04\x91"+ // Tune of Ages!
			"Play\x04\x0f\x04\xdf"+ // Play it to move
			"freely \x02\x77\x01"+ // freely through
			"\x04\xdd!\x00"), // time!

	// make tokay who gives iron shield always give the same item, and in a
	// format compatible with lookupItemSpriteAddr.
	"give hidden tokay item": MutableString(Addr{0x15, 0x5b35},
		"\x06\x01\x0e\x01\xfa\xaf\xc6\xfe\x02",
		"\x01\x01\x01\x78\x41\x4f\x37\x00\x00"),

	// buy tokay trader's shield if you have scent seeds but not satchel
	"tokay trader satchel check": MutableString(Addr{0x0a, 0x629c},
		"\x30\x16", "\x30\x00"),

	// game has zora scale palette in item gfx wrong for some reason
	"fix zora scale palette": MutableByte(Addr{0x3f, 0x67d0}, 0x13, 0x43),

	// check fake ID 1e (fool's ore) for symmetry city brother's item
	"brother fake ID": MutableStrings([]Addr{{0x15, 0x77f0}, {0x15, 0x78f6}},
		"\x4c", "\x1e"),
	// and don't change the brothers' state if the tuni nut has been placed
	"brother ignore flag": MutableString(Addr{0x15, 0x78e5},
		"\xb5\x29", "\xb0\x02"),
	// and change wives' state based on whether brother's item was obtained
	"wife ignore flag": MutableString(Addr{0x15, 0x784f},
		"\xb5\x29", "\xdf\x1e"),

	// check fake ID 10 (nothing) for king zora's item
	"king zora fake ID": MutableByte(Addr{0x0b, 0x548a}, 0x46, 0x10),

	// check fake ID 12 (nothing) for first goron dance
	"check dance 1 fake ID": MutableStrings([]Addr{{0x0c, 0x67e0},
		{0x0c, 0x685a}, {0x0c, 0x6983}}, "\x5b", "\x12"),
	// check fake ID 14 (nothing) for goron dance with letter of introduction
	"check dance 2 fake ID": MutableStrings([]Addr{{0x0c, 0x67d8},
		{0x0c, 0x6852}, {0x0c, 0x697b}}, "\x44", "\x14"),

	// skip essence checks for goron elder event
	"skip goron elder essence checks": MutableStrings(
		[]Addr{{0x0c, 0x6b1d}, {0x0c, 0x6b83}, {0x15, 0x735d}},
		"\xc7\xdb\xcd\x80", "\xc7\xdb\xcd\x00"),

	// add railing to ricky nuun screen and move worker off the "roof"
	"ricky nuun railing": MutableString(Addr{0x23, 0x718e},
		"\x69\x07\x07\x6a", "\x72\x50\x50\x73"),
	"move nuun worker": MutableString(Addr{0x12, 0x5a9e},
		"\x28\x50", "\x68\x48"),

	// text for special crescent island present portal
	"portal sign text": MutableString(Addr{0x23, 0x583f}, "\x0c\x20\x02\x18",
		"\x0c\x00C\x04\x23s only.\x01"+ // Currents only.
			" -\x04\x56Management\x00"), // -The Management

	// skip essence check for comedian
	"comedian essence check": MutableString(Addr{0x15, 0x6261},
		"\x38\x02", "\x38\x00"),

	// change conditions for rafton 2's script based on whether the player has
	// the magic oar, not on essences.
	"rafton script check": MutableString(Addr{0x15, 0x6b42},
		"\xc7\xdb\xcd\x80", "\xcb\xc0\xc6\x09"),

	// vanilla bug: compass doesn't show D6 boss key chest.
	"fix d6 compass": MutableByte(Addr{0x01, 0x4eea}, 0x14, 0x34),
}

var agesVarMutables = map[string]Mutable{
	// seed tree types
	"symmetry city tree sub ID": MutableByte(Addr{0x12, 0x59a1}, 0x35, 0x35),
	"south lynna present tree sub ID": MutableByte(Addr{0x12, 0x5ca4},
		0x06, 0x06),
	"crescent island tree sub ID": MutableByte(Addr{0x12, 0x59b8}, 0x17, 0x17),
	"zora village present tree sub ID": MutableByte(Addr{0x12, 0x59bf},
		0x38, 0x38),
	"rolling ridge west tree sub ID": MutableByte(Addr{0x12, 0x5e4d},
		0x29, 0x29),
	"ambi's palace tree sub ID": MutableByte(Addr{0x12, 0x5e5b}, 0x1a, 0x1a),
	"rolling ridge east tree sub ID": MutableByte(Addr{0x12, 0x5f46},
		0x4b, 0x4b),
	"south lynna past tree sub ID": MutableByte(Addr{0x12, 0x5e62}, 0x0c, 0x0c),
	"deku forest tree sub ID":      MutableByte(Addr{0x12, 0x6101}, 0x4d, 0x4d),
	"zora village past tree sub ID": MutableByte(Addr{0x12, 0x5e6f},
		0x3e, 0x3e),

	// first satchel should give the seeds on the south lynna tree.
	"satchel initial seeds": MutableByte(Addr{0x3f, 0x453b}, 0x20, 0x20),

	// set default satchel and shooter selection based on south lynna tree.
	// overwrites unimportant bytes in file initialization.
	"satchel initial selection": MutableWord(Addr{0x07, 0x418e}, 0x0700, 0xc400),
	"shooter initial selection": MutableWord(Addr{0x07, 0x4190}, 0x0e00, 0xc500),

	// map pop-up icons for seed trees
	"crescent island tree map icon": MutableByte(
		Addr{0x02, 0x6d05}, 0x16, 0x16),
	"symmetry city tree map icon": MutableByte(
		Addr{0x02, 0x6d08}, 0x18, 0x18),
	"south lynna tree map icon": MutableStrings(
		[]Addr{{0x02, 0x6d0b}, {0x02, 0x6d29}}, "\x15", "\x15"),
	"zora village tree map icon": MutableStrings(
		[]Addr{{0x02, 0x6d0e}, {0x02, 0x6d2f}}, "\x18", "\x18"),
	"rolling ridge west tree map icon": MutableByte(
		Addr{0x02, 0x6d20}, 0x17, 0x17),
	"ambi's palace tree map icon": MutableByte(
		Addr{0x02, 0x6d23}, 0x16, 0x16),
	"rolling ridge east tree map icon": MutableByte(
		Addr{0x02, 0x6d26}, 0x19, 0x19),
	"deku forest tree map icon": MutableByte(
		Addr{0x02, 0x6d2c}, 0x19, 0x19),

	// 33 for ricky, 23 for dimitri, 13 for moosh
	"flute palette": MutableByte(Addr{0x3f, 0x6746}, 0x03, 0x03),
	// 0b for ricky, 0c for dimitri, 0d for moosh
	"animal region": MutableByte(Addr{0x03, 0x7fff}, 0x00, 0x0d),
}
