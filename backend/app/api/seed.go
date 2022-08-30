package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/internal/repo"
	"github.com/hay-kot/content/backend/internal/types"
	"github.com/hay-kot/content/backend/pkgs/hasher"
	"github.com/hay-kot/content/backend/pkgs/logger"
)

const (
	DefaultGroup    = "Default"
	DefaultName     = "Admin"
	DefaultEmail    = "admin@admin.com"
	DefaultPassword = "admin"
)

// EnsureAdministrator ensures that there is at least one superuser in the database
// if one isn't found a default is generate using the default credentials
func (a *app) EnsureAdministrator() {
	superusers, err := a.repos.Users.GetSuperusers(context.Background())

	if err != nil {
		a.logger.Fatal(err, nil)
	}
	if len(superusers) > 0 {
		return
	}

	pw, _ := hasher.HashPassword(DefaultPassword)
	newSuperUser := types.UserCreate{
		Name:        DefaultName,
		Email:       DefaultEmail,
		IsSuperuser: true,
		Password:    pw,
	}

	a.logger.Info("creating default superuser", logger.Props{
		"name":  newSuperUser.Name,
		"email": newSuperUser.Email,
	})

	_, err = a.repos.Users.Create(context.Background(), newSuperUser)

	if err != nil {
		a.logger.Fatal(err, nil)
	}

}

func (a *app) SeedDatabase(repos *repo.AllRepos) {
	if !a.conf.Seed.Enabled {
		return
	}

	group, err := repos.Groups.Create(context.Background(), DefaultGroup)
	if err != nil {
		a.logger.Fatal(err, nil)
	}

	for _, user := range a.conf.Seed.Users {

		// Check if User Exists
		usr, _ := repos.Users.GetOneEmail(context.Background(), user.Email)

		if usr.ID != uuid.Nil {
			a.logger.Info("seed user already exists", logger.Props{
				"user": user.Name,
			})
			continue
		}

		hashedPw, err := hasher.HashPassword(user.Password)

		if err != nil {
			a.logger.Error(err, logger.Props{
				"details": "failed to hash password",
				"user":    user.Name,
			})
		}

		_, err = repos.Users.Create(context.Background(), types.UserCreate{
			Name:        user.Name,
			Email:       user.Email,
			IsSuperuser: user.IsSuperuser,
			Password:    hashedPw,
			GroupID:     group.ID,
		})

		if err != nil {
			a.logger.Error(err, logger.Props{
				"details": "failed to create seed user",
				"name":    user.Name,
			})
		}

		a.logger.Info("creating seed user", logger.Props{
			"name": user.Name,
		})
	}
}
