# Concurrent Linked List in Go

## Overview

This Go program implements a concurrent linked list data structure, focusing on the insert and lookup operations. It utilizes mutex locks to ensure thread safety during concurrent access.

## Key Features

Linked List: Represents a basic linked list structure with nodes containing integer keys.
Concurrency: Implements thread-safe operations using mutex locks to prevent data races.

## Components

### Structs

Node: Represents a node in the linked list, containing a key value and a pointer to the next node.
List: Represents the linked list structure with a head node and a mutex lock for synchronization.

### Methods

Init(): Initializes the linked list by setting the head node to nil.
Insert(key int): Inserts a new node with the given key at the beginning of the list.
Lookup(key int) bool: Searches for a node with the given key in the list and returns true if found, false otherwise.
Print(): Prints the keys of all nodes in the list.
Delet(): Deletes a node with the given key from the list.

### Main Function

Initializes the linked list.
Launches goroutines to concurrently insert nodes into the list.
Waits for all goroutines to complete.
Prints the keys of all nodes in the list.
Performs lookup operations for several keys and prints the results.

## Usage

Running the Program
Initialize the List: Use the Init method to initialize the linked list.
Concurrent Insertions: Multiple goroutines insert nodes into the list concurrently using the Insert method.
Printing the List: Use the Print method to print the keys of all nodes in the list.
Performing Lookups: Use the Lookup method to search for specific keys in the list and check their presence.

## Note

This implementation focuses on concurrent insertions and lookups. Additional functionalities like deletion and traversal can be added as needed.
