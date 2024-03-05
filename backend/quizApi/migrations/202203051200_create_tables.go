package main

func main() {
    db, err := repository.NewDatabase(&repository.Config{
    
    })
    if err != nil {
        panic("failed to connect database")
    }
    db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

    // Auto Migrate Models
    err = db.AutoMigrate(
        &models.User{},
        &models.Admin{},
        &models.Participant{},
        &models.Quiz{},
        &models.Question{},
        &models.Option{},
    )
    if err != nil {
        panic("failed to migrate database")
    }
}
