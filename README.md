# BasaltDB

BasaltDB is a SQLite and LevelDB inspired key-value database. It runs serverless from a single executable via CLI interface. The main features is user-defined data partitioning system. 

# Features

  * Data is in form of Key and Value pair of String type.
  * This is stored as an AVL tree in a sorted by Key.
  * Key is hashed using [MurmurHash3](https://github.com/aappleby/smhasher/wiki/MurmurHash3) (64bit). Murmur3 is chosen for it's fast speed and even spread of keys. 
  * Basic data operations are 
    * insert val1 val2
    * delete val1
    * get val1
  * A patition comprises of multiple parts. Each part is a file containing Level Order Traversal of AVL tree.
  * The configuration of partition is defined in the config.yaml file. It comprises of: 
    * number of parts
    * map of part id and its location
  * Basic partition operations are:
    * createPartition -p pName
    * deletePartition -p pName
    * rebalancePartition -p pName
  * Specify partitions with -p flag. All data operations without the -p flag are done in the "default" partition.
  * Rebalance Partition allows the user to redistribute data in the in a new configuration. This is to be mentioned in a new_config.yaml file.

# Limitations

  * Each read or write involves loading the entire database. Data can be split between partitions to avoid heavy operations.
  * Writes are done with strings as keys.
  * Mass writes are to be implemented and are in the TODO pipeline. 

# Getting started 

## Installation

### Get source code 

```bash
git clone https://github.com/Jatin020403/BasaltDB.git
```

### Build

```sh
make build
```
Build the file and add ./bin directory to the path

### Test

```sh
make test
```

## Partition Operations

### Create Partiton

To create a partition, first initialise the partition. Change the config file as required. Finally create the partition.

```sh
# Initialise Partition
basaltdb init -p test1

# Adjust the config as needed 

# Create Partition 
basaltdb createPartition -p test1
```

### Create Default Partition

To create default partition. This initialises and creates the the partition in the default configuration.

```sh
basaltdb initDefault
```

### Rebalance Partition

In the directory of the partition, create a new file called new_config.yaml. Then run the rebalance command.

```sh
# Create new_conf.yaml file with new conf
basaltdb rebalancePartition -p test1
```

### Delete Partition
```sh
basaltdb deletePartition -p test1
```

## Data Operations

### Insert Data

```sh
# Default Partition
basaltdb insert -k 2 -v "Inserting string data"
basaltdb insert 1 69                        
# test1 Partition
basaltdb -p test1 insert --key "test data" --value "inserting with text data"
basaltdb -p test1 insert 550e8400-e29b-41d4-a716-446655440000 "Inserting with UUID"
```

### Get Data
```sh
# Default Partition
basaltdb get -k "test data"
basaltdb get 2 
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
