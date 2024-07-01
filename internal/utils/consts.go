package utils

type DashboardTile struct {
	Path  string
	Label string
}

var DashboardTiles []DashboardTile = []DashboardTile{
	{
		Path:  "/transactions",
		Label: "Transactions",
	},
	{
		Path:  "/accounts",
		Label: "Accounts",
	},
	// {
	// 	path: "/entries",
	// 	label: "Entries",
	// },
	{
		Path:  "/account-tags",
		Label: "Account Tags",
	},
}
