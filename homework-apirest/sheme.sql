CREATE TABLE books(
    book_id    INT PRIMARY KEY AUTO_INCREMENT,
    name       VARCHAR(50),
    author     VARCHAR(50),
    status     BOOLEAN DEFAULT true)
    
CREATE TABLE  loans(
    loan_id    INT PRIMARY KEY AUTO_INCREMENT,
    book_id    INT NOT NULL, 
    user_id    INT NOT NULL,
    loan_date    DATE,
    return_date    DATE,
    FOREIGN KEY (book_id) REFERENCES books(book_id)
)
