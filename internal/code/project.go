package code

var ExampleProject = []byte(`enum Role {
  admin = "admin"
  user = "user"
}

model User {
  id      string  @id @default(uuid())
  name    string
  surname string
  role    Role
  posts   Post[]
}

model Post {
  id        int     @id @default(autoincrement())
  title     string
  content   string  @nullable
  public    bool    @default(false)
  author    User    @relation(field: authorId, reference: id)
  authorId  int
}
`)
