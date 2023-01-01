package postgres

import (
	"fmt"
	"go_practice/config"
)

func ConfigToString(psqlConfig *config.PostgresConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		psqlConfig.Host,
		psqlConfig.Port,
		psqlConfig.User,
		psqlConfig.Password,
		psqlConfig.Dbname,
		psqlConfig.Sslmode,
		psqlConfig.SearchPath,
	)	
}