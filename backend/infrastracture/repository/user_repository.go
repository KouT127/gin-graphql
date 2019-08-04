package repository


//type UserRepository struct {
//	db *gorm.DB
//}
//
//func NewUserRepository(db *gorm.DB) *UserRepository {
//	return &UserRepository{
//		db: db,
//	}
//}
//func (ur *UserRepository) FindAll(p *form.Pagination) ([]*model.User, error) {
//	user := model.User{}
//	u := ur.db.Model(&user).Related(&user.Tasks, "UserRefer").Order("-updated_at")
//	u = p.Paging(u)
//	rows, err := u.Rows()
//	defer rows.Close()
//	users, err := ur.getPointerList(rows)
//	if err != nil {
//		return users, nil
//	}
//	return users, nil
//}
//
//func (ur *UserRepository) Create(frm *form.UserForm) (*model.User, error) {
//	tx := ur.db.Begin()
//	defer func() {
//		if r := recover(); r != nil {
//			tx.Rollback()
//		}
//	}()
//
//	u := model.User{
//		Name:     frm.Name,
//		BirthDay: "",
//		Gender:   frm.Gender,
//		PhotoURL: "",
//		Active:   true,
//	}
//	if err := tx.Create(&u).Error; err != nil {
//		tx.Rollback()
//		return &u, err
//	}
//	task := model.Task{
//		UserRefer:   u.ID,
//		Title:       "test",
//		Description: "testd",
//	}
//	if err := tx.Create(&task).Error; err != nil {
//		tx.Rollback()
//		return &u, err
//	}
//	return &u, tx.Commit().Error
//}
//
//func (ur *UserRepository) getPointerList(rows *sql.Rows) ([]*model.User, error) {
//	var list []*model.User
//	for rows.Next() {
//		mem := &model.User{}
//		err := ur.db.ScanRows(rows, &mem)
//		if err != nil {
//			fmt.Print(err.Error())
//			return list, err
//		}
//		list = append(list, mem)
//	}
//	return list, nil
//}
//
//func (ur *UserRepository) GetUserMaxPage(limit int) int {
//	var cnt int
//	ur.db.Model(&[]model.User{}).Count(&cnt)
//	return int(math.Ceil(float64(cnt) / float64(limit)))
//}
