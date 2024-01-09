package security

// func (u *models.NewUser) HashPassword() error {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = string(hashedPassword)
// 	return nil
// }

// // CheckPasswordHash compares the provided password with the hashed password in the database.
// func (u *models.NewUser) CheckPasswordHash(password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
// }
