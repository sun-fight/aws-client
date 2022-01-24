package model

type Grid struct {
}

type User struct {
	UserID                    int64  `json:"user_id"`
	Nickname                  string `json:"nickname"`
	PlunderAttackCount        int32  `json:"plunder_attack_count"`
	PlunderCounterattackCount int32  `json:"plunder_counterattack_count"`
	CreatedAt                 int64  `json:"created_at"`
	DeletedAt                 int64  `json:"deleted_at"`
	LastLoginAt               int64  `json:"last_login_at"`
	Maps                      map[string]map[int32]Grid
	// Roles
	// Items
}
