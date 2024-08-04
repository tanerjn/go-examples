package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

type Album struct {
        ID int
        Artist string
        Title string
        Price float64
    }

var db *sql.DB

func main() {
    
    dsn := "username:password@tcp(127.0.0.1:3306)/recordings"

    db, err := sql.Open("mysql", dsn)
    
    if err != nil {
        fmt.Println("Error opening database:", err)
        return
    }
    defer db.Close()
    err = db.Ping()
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    fmt.Println("Successfully connected to the database!")

    rows, err := db.Query("SELECT id, title, artist, price age FROM album")
    if err != nil{
        log.Fatal(err)
    }
    defer rows.Close()
    
    var albums []Album

    for rows.Next(){
        var album Album
        err := rows.Scan(&album.ID, &album.Artist, &album.Title, &album.Price)
        if err != nil {
            log.Fatal(err)
        }
        albums = append(albums, album)
        fmt.Printf("ID: %d, artist: %s, title: %s, price: %.2f\n", album.ID, album.Artist, album.Title, album.Price)
    }

    fmt.Printf("Albums found: %v\n", albums)

    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Successfully read from the database!")

    var albumsFilteredByArtist []Album
    var albumFilteredByID Album

    //get Albums by Artist
    albumsFilteredByArtist, err = albumsByArtist("John Coltrane")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("\nAlbumsFilteredByArtist found: %v\n", albumsFilteredByArtist)

    //get Albums by ID
    albumFilteredByID, err = albumByID(2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("\nAlbumFilteredByID found: %v\n", albumFilteredByID)

    //add Album and return ID
    albID, err := addAlbum(Album{
        Title:  "The Modern Sound of Betty Carter",
        Artist: "Betty Carter",
        Price:  49.99,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("\n ID of added album: %v\n", albID)
}


// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []Album
    dsn := "username:password@tcp(127.0.0.1:3306)/recordings"
    db, err := sql.Open("mysql", dsn)
    
    fmt.Printf("before q\n")
    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    fmt.Printf("after q\n")
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    fmt.Printf("after close\n")
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("\nalbumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("\nalbumsByArtist %q: %v", name, err)
    }
    return albums, nil
}

    
    
// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
    // An album to hold data from the returned row.
    var alb Album

    dsn := "username:password@tcp(127.0.0.1:3306)/recordings"
    db, err := sql.Open("mysql", dsn)

    if err != nil {
        log.Fatalf("Error: %v", err)
        }

    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
        if err == sql.ErrNoRows {
            return alb, fmt.Errorf("\nalbumsById %d: no such album", id)
        }
        return alb, fmt.Errorf("\nalbumsById %d: %v", id, err)
    }
    return alb, nil
}

// AddNewALbum
// returning the album ID of the new entry
func addAlbum(alb Album) (int64,  error) {
    dsn := "username:password@tcp(127.0.0.1:3306)/recordings"
    db, err := sql.Open("mysql", dsn)
    result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
}



