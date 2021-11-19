package models

type Metadata struct {
	Hash string
	Type string
}

var IsDB = []Metadata{
	Metadata{
		Hash: "07d14d16d21d21d00042d43d000000aa99ce74e2c6d013c745aa52b5cc042d",
		Type: "Metasploit",
	},
	Metadata{
		Hash: "22b22b09b22b22b22b22b22b22b22b352842cd5d6b0278445702035e06875c",
		Type: "Trickbot",
	},
	Metadata{
		Hash: "1dd40d40d00040d1dc1dd40d1dd40d3df2d6a0c2caaa0dc59908f0d3602943",
		Type: "AsyncRAT",
	},
	Metadata{
		Hash: "29d21b20d29d29d21c41d21b21b41d494e0df9532e75299f15ba73156cee38",
		Type: "Merlin C2",
	},
	Metadata{
		Hash: "3fd21b20d3fd3fd21c43d21b21b43d494e0df9532e75299f15ba73156cee38",
		Type: "Merlin C2 https",
	},
	// add ..
}
