package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func setIfEmpty(id *string) {
	if *id == "" { *id = uuid.NewString() }
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error)       { setIfEmpty(&u.ID); return }
func (a *About) BeforeCreate(tx *gorm.DB) (err error)      { setIfEmpty(&a.ID); return }
func (c *Category) BeforeCreate(tx *gorm.DB) (err error)   { setIfEmpty(&c.ID); return }
func (t *Tribe) BeforeCreate(tx *gorm.DB) (err error)      { setIfEmpty(&t.ID); return }
func (r *Region) BeforeCreate(tx *gorm.DB) (err error)     { setIfEmpty(&r.ID); return }
func (c *Content) BeforeCreate(tx *gorm.DB) (err error)    { setIfEmpty(&c.ID); return }
