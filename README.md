# BasaltDB

BasaltDB is a SQLite and LevelDB inspired key-value database. It runs from a single executable with a CLI interface and persistant storage.

# Features

  * Keys and values are of string type.
  * Data is stored in AVL tree sorted by Key.
  * Basic operations are 
    * insert val1 val2
    * delete val1
    * get val1
  * Data is overwritten with according to most recent timestamp with nanosecond precision, as a future feature of allowing multiple access.
  * Persistant stored data is the Level Traversal of the AVL tree.
  * Read and write operations are of linear time complexity. 

# Limitations

  * Each read or write involves loading the entire database.
  * Mass writes are to be implemented and are in the TODO pipeline. 
  * Support for partitions is to be added.

# Getting started 

## Get source code 

```bash
git clone https://github.com/Jatin020403/BasaltDB.git
```

## Building 

```go
go install
```

## Example Code

### Insert Data

```sh
BasaltDB insert 1 69
BasaltDB insert 2 "Inserting string data"
BasaltDB insert "test data" "inserting with text data"
BasaltDB insert 550e8400-e29b-41d4-a716-446655440000 "Inserting with UUID"
```

### Get Data
```sh
BasaltDB get 2 
BasaltDB get "test data" 
```

### View Data
```sh
BasaltDB getAll
```

### Update Data
BasaltDB overwrites data according to timestamp.

```sh
BasaltDB insert 1 69
BasaltDB insert 1 "Updating 69 with this string"
```

### Delete Data
```sh
BasaltDB delete 2 
BasaltDB delete "test data" 
```
