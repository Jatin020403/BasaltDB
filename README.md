# BasaltDB

BasaltDB is a SQLite and LevelDB inspired key-value database. It runs from a single executable with a CLI interface and persistant storage.

# Features

  * Data is in form of Key and Value pair of String type.
  * Data is stored in different partitions.
  * Data is overwritten with according to most recent timestamp with nanosecond precision, as a future feature to allow writes from multiple nodes.
  * Basic data operations are 
    * insert val1 val2
    * delete val1
    * get val1
  * Each partition is an AVL tree with value sorted by Key.
  * Persistant stored data is the Level Order Traversal of the AVL tree.
  * Basic partition operations are
    * createPartition -p pName
    * deletePartition -p pName
    * getPartitions
  * Specify partitions with -p flag. All data operations without the -p flag are done in the "default" partition.
  * Read and write operations are of linear time complexity. 

# Limitations

  * Each read or write involves loading the entire database. Data can be split between partitions to avoid heavy operations.
  * Writes are done with strings as keys.
  * Mass writes are to be implemented and are in the TODO pipeline. 

# Getting started 

## Get source code 

```bash
git clone https://github.com/Jatin020403/BasaltDB.git
```

## Building 

Build the file and add ./bin directory to the path
```sh
make build
```

## Test

Run tests with 
```sh
make test
```

## Data Operations

### Insert Data

```sh
# Default Partition
basaltdb insert 1 69                        
basaltdb insert -k 2 -v "Inserting string data"
# test1 Partition
basaltdb -p test1 insert --key "test data" --value "inserting with text data"
basaltdb -p test1 insert 550e8400-e29b-41d4-a716-446655440000 "Inserting with UUID"
```

### Get Data
```sh
# Default Partition
basaltdb get 2 
basaltdb get -k "test data"
# test1 Partition
basaltdb -p test1 get --key 1 
```

### View Data
```sh
# Default Partition
basaltdb getTree
# test1 Partition
basaltdb -p test1 getTree
```

### Update Data
basaltdb overwrites data according to timestamp.

```sh
# test1 Partition
basaltdb -p test1 insert -k 1 -v "Updating 69 with this string"
```

### Delete Data
```sh
# Default Partition
basaltdb delete 2 
basaltdb delete "test data" 
# test1 Partition
basaltdb -p test1 delete --key 1 
```

## Partition Operations

### Create Partition
```sh
basaltdb createPartition -p test1
```

### Get All Partitions
```sh
basaltdb getPartitions
```

### Delete Partition
```sh
basaltdb deletePartition -p test1
```
